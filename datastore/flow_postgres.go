package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/environment"
	"github.com/OpenSlides/openslides-go/oslog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate  sh -c "go run genfields/main.go > field_def.go"

var (
	envPostgresHost         = environment.NewVariable("DATABASE_HOST", "localhost", "Postgres Host.")
	envPostgresPort         = environment.NewVariable("DATABASE_PORT", "5432", "Postgres Post.")
	envPostgresDatabase     = environment.NewVariable("DATABASE_NAME", "openslides", "Postgres User.")
	envPostgresUser         = environment.NewVariable("DATABASE_USER", "openslides", "Postgres Database.")
	envPostgresPasswordFile = environment.NewVariable("DATABASE_PASSWORD_FILE", "/run/secrets/postgres_password", "Postgres Password.")

	envPostgresNotifyHost         = environment.NewVariable("DATABASE_NOTIFY_HOST", "DATABASE_HOST", "Postgres Host for notify.")
	envPostgresNotifyPort         = environment.NewVariable("DATABASE_NOTIFY_PORT", "DATABASE_PORT", "Postgres Port for notify.")
	envPostgresNotifyDatabase     = environment.NewVariable("DATABASE_NOTIFY_NAME", "DATABASE_NAME", "Postgres Database for notify.")
	envPostgresNotifyUser         = environment.NewVariable("DATABASE_NOTIFY_USER", "DATABASE_USER", "Postgres User for notify.")
	envPostgresNotifyPasswordFile = environment.NewVariable("DATABASE_NOTIFY_PASSWORD_FILE", "DATABASE_PASSWORD_FILE", "Postgres Password for notify.")
)

// FlowPostgres uses postgres to get the connections.
type FlowPostgres struct {
	Pool         *pgxpool.Pool
	notifyConfig *pgx.ConnConfig
	enums        map[uint32]struct{}
	enumArray    map[uint32]struct{}
}

// NewFlowPostgres initializes a SourcePostgres.
func NewFlowPostgres(lookup environment.Environmenter) (*FlowPostgres, error) {
	addr, err := postgresDSN(lookup)
	if err != nil {
		return nil, fmt.Errorf("reading postgres dns: %w", err)
	}

	config, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %w", err)
	}

	notifyAddr, err := postgresDSNNotify(lookup)
	if err != nil {
		return nil, fmt.Errorf("reading postgres dns for notify: %w", err)
	}

	notifyConf, err := pgx.ParseConfig(notifyAddr)
	if err != nil {
		return nil, fmt.Errorf("generate config for notify: %w", err)
	}

	flow := FlowPostgres{
		Pool:         pool,
		notifyConfig: notifyConf,
	}

	if err := flow.updateEnums(ctx); err != nil {
		return nil, err
	}

	return &flow, nil
}

func (p *FlowPostgres) updateEnums(ctx context.Context) error {
	c, err := p.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer c.Release()

	sql := `SELECT oid, typarray FROM pg_type WHERE typtype = 'e';`
	rows, err := c.Conn().Query(ctx, sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	p.enums = map[uint32]struct{}{}
	p.enumArray = map[uint32]struct{}{}
	for rows.Next() {
		var oid uint32
		var typarray uint32
		if err := rows.Scan(&oid, &typarray); err != nil {
			return err
		}

		p.enums[oid] = struct{}{}
		p.enumArray[typarray] = struct{}{}
	}

	return nil
}

// Close closes the connection pool.
func (p *FlowPostgres) Close() {
	p.Pool.Close()
}

// Get fetches the keys from postgres.
func (p *FlowPostgres) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("acquiring connection: %w", err)
	}
	defer conn.Release()

	return p.getWithConn(ctx, conn.Conn(), keys...)
}

func (p *FlowPostgres) getWithConn(ctx context.Context, conn *pgx.Conn, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	collectionIDs := make(map[string][]int)
	collectionFields := make(map[string][]string)

	for _, key := range keys {
		collection := key.Collection()

		collectionIDs[collection] = append(collectionIDs[collection], key.ID())

		if field := key.Field(); field != "id" {
			collectionFields[collection] = append(collectionFields[collection], field)
		}
	}

	for collection := range collectionIDs {
		slices.Sort(collectionIDs[collection])
	}

	for collection := range collectionFields {
		slices.Sort(collectionFields[collection])
		collectionFields[collection] = slices.Compact(collectionFields[collection])
	}

	keyValues := make(map[dskey.Key][]byte, len(keys))
	for collection, ids := range collectionIDs {
		fields := []string{"id"}
		fields = append(fields, collectionFields[collection]...)

		// TODO: Remove me, if the new vote service and projector service are merged
		switch collection {
		case "poll":
			fields = slices.DeleteFunc(fields, func(s string) bool {
				return s == "live_votes"
			})
		case "projection":
			fields = slices.DeleteFunc(fields, func(s string) bool {
				return s == "content"
			})
		}

		sql := fmt.Sprintf(
			`SELECT %s FROM "%s" WHERE id = ANY ($1) `,
			strings.Join(fields, ","),
			collection,
		)

		const batchSize = 5000
		for i := 0; i < len(ids); i += batchSize {
			end := min(i+batchSize, len(ids))
			batch := ids[i:end]

			rows, err := conn.Query(ctx, sql, batch)
			if err != nil {
				return nil, fmt.Errorf("sending query `%s`: %w", sql, err)
			}

			fieldDescription := rows.FieldDescriptions()
			if fieldDescription[0].Name != "id" {
				return nil, fmt.Errorf("invalid row for collection %s, expect first value to be the id. Got %s", collection, fieldDescription[0].Name)
			}

			err = forEachRow(rows, func(row pgx.CollectableRow) error {
				values := row.RawValues()

				id, err := strconv.Atoi(string(values[0]))
				if err != nil {
					return fmt.Errorf("invalid id %s: %w", string(values[0]), err)
				}

				idKey, err := dskey.FromParts(collection, id, "id")
				if err != nil {
					return fmt.Errorf("invalid id-key for id %d: %w", id, err)
				}
				keyValues[idKey] = []byte(strconv.Itoa(id))

				for i, value := range values {
					field := fieldDescription[i].Name
					key, err := dskey.FromParts(collection, id, field)
					if err != nil {
						return fmt.Errorf("invalid key on field %d: %w", i, err)
					}

					keyValues[key] = nil
					if value == nil {
						continue
					}

					keyValues[key], err = p.convertValue(value, fieldDescription[i].DataTypeOID)
					if err != nil {
						return fmt.Errorf("convert value for field %s/%s: %w", collection, field, err)
					}
				}

				return nil
			})
			if err != nil {
				return nil, fmt.Errorf("parse collection %s: %w", collection, err)
			}
		}
	}

	// The current implementation of cache expects to gets exactly the keys,
	// that where requested. If a key does not exist, nil has to be used.
	result := make(map[dskey.Key][]byte, len(keys))
	for _, key := range keys {
		result[key] = keyValues[key]
	}

	return result, nil
}

func (p *FlowPostgres) convertValue(value []byte, oid uint32) ([]byte, error) {
	switch oid {
	case pgtype.VarcharOID, pgtype.TextOID:
		return json.Marshal(string(value))

	case pgtype.Int4OID, pgtype.Float8OID, pgtype.JSONBOID:
		return bytes.Clone(value), nil

	case pgtype.Int4ArrayOID:
		result := make([]byte, len(value))
		copy(result, value)
		result[0] = '['
		result[len(result)-1] = ']'
		return result, nil

	case pgtype.BoolOID:
		if string(value) == "t" {
			return []byte("true"), nil
		}
		return []byte("false"), nil

	case pgtype.NumericOID:
		return fmt.Appendf(nil, `"%s"`, value), nil

	case pgtype.TimestamptzOID:
		timeValue, err := time.Parse("2006-01-02 15:04:05-07", string(value))
		if err != nil {
			return nil, fmt.Errorf("parsing time %s: %w", value, err)
		}
		return strconv.AppendInt(nil, timeValue.Unix(), 10), nil

	case pgtype.VarcharArrayOID, pgtype.TextArrayOID:
		return convertPGArray(string(value))

	default:
		if _, ok := p.enums[oid]; ok {
			return json.Marshal(string(value))
		}
		if _, ok := p.enumArray[oid]; ok {
			return convertPGArray(string(value))
		}

		return nil, fmt.Errorf("unsupported postgres type %d", oid)
	}
}

// convertPGArray transforms a postgres style array into a json array.
func convertPGArray(pgValue string) ([]byte, error) {
	strValue := strings.Trim(string(pgValue), "{}")
	if strValue == "" {
		return []byte("[]"), nil
	}
	strArray := strings.Split(strValue, ",")
	return json.Marshal(strArray)
}

// Update listens on pg notify to fetch updates.
func (p *FlowPostgres) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	var conn *pgx.Conn
	defer conn.Close(context.Background())

	lastXactID := 0
	for {
		if ctx.Err() != nil {
			return
		}

		if conn == nil || conn.IsClosed() {
			conn = getPostgresConnection(ctx, p.notifyConfig)
			if lastXactID > 0 {
				oslog.Info("Database reconnected")
			}

			_, err := conn.Exec(ctx, "LISTEN os_notify")
			if err != nil {
				err = fmt.Errorf("listen on channel os_notify: %w", err)
				updateFn(nil, err)
				panic(err)
			}
		}

		notification, err := conn.WaitForNotification(ctx)
		if err != nil {
			updateFn(nil, fmt.Errorf("wait for notification: %w", err))
			continue
		}

		var payload struct {
			XACTID int `json:"xactId"`
		}

		if err := json.Unmarshal([]byte(notification.Payload), &payload); err != nil {
			updateFn(nil, fmt.Errorf("unmarshal notify payload: %w", err))
			continue
		}

		var sql string
		args := []any{payload.XACTID}
		if lastXactID > 0 && lastXactID+1 < payload.XACTID {
			sql = `SELECT DISTINCT operation, fqid, updated_fields FROM os_notify_log_t WHERE xact_id <= $1::xid8 AND xact_id > $2::xid8;`
			args = append(args, lastXactID)
		} else {
			sql = `SELECT DISTINCT operation, fqid, updated_fields FROM os_notify_log_t WHERE xact_id = $1::xid8;`
		}
		rows, err := conn.Query(ctx, sql, args...)
		if err != nil {
			updateFn(nil, fmt.Errorf("query fqids for transactions %v: %w", args, err))
			continue
		}

		updateLogs, err := pgx.CollectRows(rows, pgx.RowToStructByName[struct {
			Operation     string
			Fqid          string
			UpdatedFields []string
		}])
		if err != nil {
			panic(fmt.Errorf("parse notify_log: %w", err))
		}

		nonFatalErrs := []error{}

		var deletedKeys []dskey.Key
		var updatedKeys []dskey.Key
		for _, updateLog := range updateLogs {
			collectionName, id, err := getCollectionNameAndID(updateLog.Fqid)
			if err != nil {
				nonFatalErrs = append(nonFatalErrs, fmt.Errorf("split fqid from %s: %w", updateLog.Fqid, err))
				continue
			}

			keys, err := createKeyList(collectionName, id, updateLog.UpdatedFields)
			if err != nil {
				nonFatalErrs = append(nonFatalErrs, fmt.Errorf("creating key list from notification: %w", err))
			}

			switch updateLog.Operation {
			case "delete":
				deletedKeys = append(deletedKeys, keys...)
			case "insert", "update":
				updatedKeys = append(updatedKeys, keys...)
			}
		}

		// TODO: don't use getWithConn for insert operation
		values, err := p.Get(ctx, updatedKeys...)
		if err != nil {
			updateFn(nil, fmt.Errorf("fetching keys %v: %w", updatedKeys, err))
			if conn.IsClosed() {
				continue
			} else {
				panic("error on healty connection - exiting")
			}
		}

		if values == nil && len(deletedKeys) != 0 {
			values = map[dskey.Key][]byte{}
		}

		for _, key := range deletedKeys {
			values[key] = nil
		}

		updateFn(values, errors.Join(nonFatalErrs...))
		lastXactID = payload.XACTID
	}
}

// WaitPostgresAvailable blocks until postgres db is availabe
func WaitPostgresAvailable(lookup environment.Environmenter) error {
	if _, forDocu := lookup.(*environment.ForDocu); forDocu {
		return nil
	}

	addr, err := postgresDSN(lookup)
	if err != nil {
		return fmt.Errorf("reading postgres config: %w", err)
	}

	config, err := pgx.ParseConfig(addr)
	if err != nil {
		return fmt.Errorf("parsing postgres config: %w", err)
	}

	ctx := context.Background()

	for {
		conn := getPostgresConnection(ctx, config)
		err := waitDatabaseInitialized(ctx, conn)
		_ = conn.Close(ctx)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		return nil
	}
}

// getPostgresConnection tries to connect to a database until it is successful
// TODO: Return error on credential related errors
func getPostgresConnection(ctx context.Context, connConfig *pgx.ConnConfig) *pgx.Conn {
	retryDelay := 1 * time.Second

	for {
		conn, err := pgx.ConnectConfig(ctx, connConfig)
		if err != nil {
			oslog.Error("Error connecting to db: %v", err)
			time.Sleep(retryDelay)
			continue
		}

		pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		err = conn.Ping(pingCtx)
		if err != nil {
			oslog.Info("Waiting for db to become ready")
			time.Sleep(retryDelay)
			continue
		}

		cancel()
		return conn
	}
}

// waitDatabaseInitialized checks if the version table is existing and the latest migration is finalized
func waitDatabaseInitialized(ctx context.Context, conn *pgx.Conn) error {
	for {
		var cnt int64
		err := conn.QueryRow(ctx, "SELECT COUNT(*) FROM version").Scan(&cnt)
		if err == nil && cnt >= 1 {
			return nil
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		if err != nil {
			oslog.Error("Could not request version table: %v", err)
		} else {
			oslog.Info("Waiting for db schema to become ready")
		}

		time.Sleep(1 * time.Second)
	}
}

func createKeyList(collection string, id int, fields []string) ([]dskey.Key, error) {
	if len(fields) == 0 {
		fields = collectionFields[collection]
	}

	keys := make([]dskey.Key, 0, len(fields))
	keyErrors := []error{}
	for _, field := range fields {
		key, err := dskey.FromParts(collection, id, field)
		if err != nil {
			keyErrors = append(keyErrors, err)
			continue
		}

		keys = append(keys, key)
	}

	var err error
	if len(keyErrors) > 0 {
		err = errors.Join(keyErrors...)
	}

	return keys, err
}

// forEachRow is like pgx.ForEachRow but uses CollectableRow instead of scan.
func forEachRow(rows pgx.Rows, fn func(row pgx.CollectableRow) error) error {
	defer rows.Close()

	for rows.Next() {
		if err := fn(rows); err != nil {
			return err
		}
	}
	return rows.Err()
}

// getCollectionNameAndID removes the suffix from the collection name.
func getCollectionNameAndID(keyStr string) (string, int, error) {
	idx1 := strings.IndexByte(keyStr, '/')

	if idx1 == -1 {
		return "", 0, fmt.Errorf("invalid key `%s`: missing slash", keyStr)
	}

	id, err := strconv.Atoi(keyStr[idx1+1:])
	if err != nil {
		return "", 0, fmt.Errorf("invalid key `%s`: id is not an integer", keyStr)
	}

	// Can be removed when this is merged:
	// https://github.com/OpenSlides/openslides-meta/pull/240
	return strings.TrimSuffix(keyStr[:idx1], "_t"), id, nil
}

// encodePostgresConfig encodes a string to be used in the postgres key value style.
//
// See: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
func encodePostgresConfig(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `'`, `\'`)
	return s
}

func postgresDSN(lookup environment.Environmenter) (string, error) {
	password, err := environment.ReadSecret(lookup, envPostgresPasswordFile)
	if err != nil {
		return "", fmt.Errorf("reading postgres password: %w", err)
	}

	return postgresConfigString(
		encodePostgresConfig(envPostgresHost.Value(lookup)),
		encodePostgresConfig(envPostgresPort.Value(lookup)),
		encodePostgresConfig(envPostgresDatabase.Value(lookup)),
		encodePostgresConfig(envPostgresUser.Value(lookup)),
		encodePostgresConfig(password),
	), nil
}

func postgresDSNNotify(lookup environment.Environmenter) (string, error) {
	password, err := environment.ReadSecretOr(lookup, envPostgresNotifyPasswordFile, envPostgresPasswordFile)
	if err != nil {
		return "", fmt.Errorf("reading postgres password: %w", err)
	}

	return postgresConfigString(
		encodePostgresConfig(envPostgresNotifyHost.ValueOr(lookup, envPostgresHost)),
		encodePostgresConfig(envPostgresNotifyPort.ValueOr(lookup, envPostgresPort)),
		encodePostgresConfig(envPostgresNotifyDatabase.ValueOr(lookup, envPostgresDatabase)),
		encodePostgresConfig(envPostgresNotifyUser.ValueOr(lookup, envPostgresUser)),
		encodePostgresConfig(password),
	), nil
}

func postgresConfigString(host, port, db, user, password string) string {
	return fmt.Sprintf(
		`user='%s' password='%s' host='%s' port='%s' dbname='%s'`,
		encodePostgresConfig(user),
		encodePostgresConfig(password),
		encodePostgresConfig(host),
		encodePostgresConfig(port),
		encodePostgresConfig(db),
	)
}
