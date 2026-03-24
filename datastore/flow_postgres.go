package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/environment"
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
)

// FlowPostgres uses postgres to get the connections.
type FlowPostgres struct {
	Pool  *pgxpool.Pool
	enums map[uint32]struct{}
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

	return fmt.Sprintf(
		`user='%s' password='%s' host='%s' port='%s' dbname='%s'`,
		encodePostgresConfig(envPostgresUser.Value(lookup)),
		encodePostgresConfig(password),
		encodePostgresConfig(envPostgresHost.Value(lookup)),
		encodePostgresConfig(envPostgresPort.Value(lookup)),
		encodePostgresConfig(envPostgresDatabase.Value(lookup)),
	), nil
}

// NewFlowPostgres initializes a SourcePostgres.
func NewFlowPostgres(lookup environment.Environmenter) (*FlowPostgres, error) {
	addr, err := postgresDSN(lookup)
	if err != nil {
		return nil, fmt.Errorf("reading postgres password: %w", err)
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

	flow := FlowPostgres{Pool: pool}
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

	sql := `SELECT oid FROM pg_type WHERE typtype = 'e';`
	rows, err := c.Conn().Query(ctx, sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	p.enums = map[uint32]struct{}{}
	for rows.Next() {
		var oid uint32
		if err := rows.Scan(&oid); err != nil {
			return err
		}

		p.enums[oid] = struct{}{}
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
		collectionIDs[collection] = slices.Compact(collectionIDs[collection])
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
		strValue := strings.Trim(string(value), "{}")
		if strValue == "" {
			return []byte("[]"), nil
		}
		strArray := strings.Split(strValue, ",")
		return json.Marshal(strArray)

	default:
		if _, ok := p.enums[oid]; ok {
			return json.Marshal(string(value))
		}

		return nil, fmt.Errorf("unsupported postgres type %d", oid)
	}
}

// Update listens on pg notify to fetch updates.
func (p *FlowPostgres) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		updateFn(nil, fmt.Errorf("acquire connection: %w", err))
		return
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "LISTEN os_notify")
	if err != nil {
		updateFn(nil, fmt.Errorf("listen on channel os_notify: %w", err))
		return
	}

	for {
		notification, err := conn.Conn().WaitForNotification(ctx)
		if err != nil {
			updateFn(nil, fmt.Errorf("wait for notification: %w", err))
			return
		}

		var payload struct {
			XACTID int `json:"xactId"`
		}

		if err := json.Unmarshal([]byte(notification.Payload), &payload); err != nil {
			updateFn(nil, fmt.Errorf("unmarshal notify payload: %w", err))
			return
		}

		sql := `SELECT DISTINCT operation, fqid, updated_fields FROM os_notify_log_t WHERE xact_id = $1::xid8;`
		rows, err := conn.Conn().Query(ctx, sql, payload.XACTID)
		if err != nil {
			updateFn(nil, fmt.Errorf("query fqids for transaction %d: %w", payload.XACTID, err))
			return
		}

		updateLogs, err := pgx.CollectRows(rows, pgx.RowToStructByName[struct {
			Operation     string
			Fqid          string
			UpdatedFields []string
		}])
		if err != nil {
			updateFn(nil, fmt.Errorf("parse notify_log: %w", err))
			return
		}

		var deletedKeys []dskey.Key
		var updatedKeys []dskey.Key
		for _, updateLog := range updateLogs {
			collectionName, id, err := getCollectionNameAndID(updateLog.Fqid)
			if err != nil {
				updateFn(nil, fmt.Errorf("split fqid from %s: %w", updateLog.Fqid, err))
				return
			}

			keys, err := createKeyList(collectionName, id, updateLog.UpdatedFields)
			if err != nil {
				updateFn(nil, fmt.Errorf("creating key list from notification: %w", err))
				return
			}

			switch updateLog.Operation {
			case "delete":
				deletedKeys = append(deletedKeys, keys...)
			case "insert", "update":
				updatedKeys = append(updatedKeys, keys...)
			}
		}

		// TODO: don't use getWithConn for insert operation
		values, err := p.getWithConn(ctx, conn.Conn(), updatedKeys...)
		if err != nil {
			updateFn(nil, fmt.Errorf("fetching keys %v: %w", updatedKeys, err))
		}

		for _, key := range deletedKeys {
			values[key] = nil
		}

		updateFn(values, nil)
	}
}

// WaitPostgresAvailable blocks until postgres db is availabe
func WaitPostgresAvailable(lookup environment.Environmenter) error {
	if _, forDocu := lookup.(*environment.ForDocu); forDocu {
		return nil
	}

	addr, err := postgresDSN(lookup)
	if err != nil {
		return fmt.Errorf("reading postgres password: %w", err)
	}

	var conn *pgx.Conn
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		conn, err = pgx.Connect(ctx, addr)
		if err == nil {
			err = conn.Ping(ctx)
			_ = conn.Close(ctx)
		}

		cancel()
		if err == nil {
			return nil
		}

		time.Sleep(1 * time.Second)
	}
}

func createKeyList(collection string, id int, fields []string) ([]dskey.Key, error) {
	if len(fields) == 0 {
		fields = collectionFields[collection]
	}

	keys := make([]dskey.Key, len(fields))
	for i, field := range fields {
		key, err := dskey.FromParts(collection, id, field)
		if err != nil {
			continue
		}

		keys[i] = key
	}
	return keys, nil
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
