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
	Pool *pgxpool.Pool
}

// encodePostgresConfig encodes a string to be used in the postgres key value style.
//
// See: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
func encodePostgresConfig(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `'`, `\'`)
	return s
}

// NewFlowPostgres initializes a SourcePostgres.
func NewFlowPostgres(lookup environment.Environmenter) (*FlowPostgres, error) {
	password, err := environment.ReadSecret(lookup, envPostgresPasswordFile)
	if err != nil {
		return nil, fmt.Errorf("reading postgres password: %w", err)
	}

	addr := fmt.Sprintf(
		`user='%s' password='%s' host='%s' port='%s' dbname='%s'`,
		encodePostgresConfig(envPostgresUser.Value(lookup)),
		encodePostgresConfig(password),
		encodePostgresConfig(envPostgresHost.Value(lookup)),
		encodePostgresConfig(envPostgresPort.Value(lookup)),
		encodePostgresConfig(envPostgresDatabase.Value(lookup)),
	)

	config, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %w", err)
	}

	flow := FlowPostgres{Pool: pool}

	return &flow, nil
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

	return getWithConn(ctx, conn.Conn(), keys...)
}

func getWithConn(ctx context.Context, conn *pgx.Conn, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	collectionIDs := make(map[string][]int)
	collectionFields := make(map[string][]string)
	for _, key := range keys {
		if !slices.Contains(collectionIDs[key.Collection()], key.ID()) {
			collectionIDs[key.Collection()] = append(collectionIDs[key.Collection()], key.ID())
		}

		if key.Field() != "id" && !slices.Contains(collectionFields[key.Collection()], key.Field()) {
			collectionFields[key.Collection()] = append(collectionFields[key.Collection()], key.Field())
		}
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

		rows, err := conn.Query(ctx, sql, ids)
		if err != nil {
			return nil, fmt.Errorf("sending query `%s`: %w", sql, err)
		}

		err = forEachRow(rows, func(row pgx.CollectableRow) error {
			values := row.RawValues()

			if row.FieldDescriptions()[0].Name != "id" {
				return fmt.Errorf("invalid row for collection %s, expect firts value to be the id. Got %s", collection, row.FieldDescriptions()[0].Name)
			}
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
				field := row.FieldDescriptions()[i].Name
				key, err := dskey.FromParts(collection, id, field)
				if err != nil {
					return fmt.Errorf("invalid key on field %d: %w", i, err)
				}

				keyValues[key] = nil
				if value == nil {
					continue
				}

				converted, err := convertValue(value, row.FieldDescriptions()[i].DataTypeOID)
				if err != nil {
					return fmt.Errorf("convert value for field %s/%s: %w", collection, field, err)
				}
				keyValues[key] = bytes.Clone(converted)
			}

			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("parse collection %s: %w", collection, err)
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

func convertValue(value []byte, oid uint32) ([]byte, error) {
	const (
		PSQLTypeVarChar     = 1043
		PSQLTypeInt         = 23
		PSQLTypeBool        = 16
		PSQLTypeText        = 25
		PSQLTypeIntList     = 1007
		PSQLTypeTimestamp   = 1184
		PSQLTypeDecimal     = 1700
		PSQLTypeJSON        = 3802
		PSQLTypeTextList    = 1009
		PSQLTypeVarCharList = 1015
		PSQLTypeFloat       = 701
	)

	switch oid {
	case PSQLTypeVarChar, PSQLTypeText:
		return json.Marshal(string(value))

	case PSQLTypeInt, PSQLTypeJSON, PSQLTypeFloat:
		return value, nil

	case PSQLTypeIntList:
		value[0] = '['
		value[len(value)-1] = ']'
		return value, nil

	case PSQLTypeBool:
		if string(value) == "t" {
			return []byte("true"), nil
		}
		return []byte("false"), nil

	case PSQLTypeDecimal:
		return fmt.Appendf([]byte{}, `"%s"`, value), nil

	case PSQLTypeTimestamp:
		timeValue, err := time.Parse("2006-01-02 15:04:05-07", string(value))
		if err != nil {
			return nil, fmt.Errorf("parsing time %s: %w", value, err)
		}

		return []byte(strconv.Itoa(int(timeValue.Unix()))), nil

	case PSQLTypeVarCharList, PSQLTypeTextList:
		strValue := strings.Trim(string(value), "{}")
		strArray := strings.Split(strValue, ",")
		return json.Marshal(strArray)

	default:
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

		sql := `SELECT DISTINCT fqid FROM os_notify_log_t WHERE xact_id = $1::xid8;`
		rows, err := conn.Conn().Query(ctx, sql, payload.XACTID)
		if err != nil {
			updateFn(nil, fmt.Errorf("query fqids for transaction %d: %w", payload.XACTID, err))
			return
		}

		fqids, err := pgx.CollectRows(rows, pgx.RowTo[string])
		if err != nil {
			updateFn(nil, fmt.Errorf("parse rows: %w", err))
			return
		}

		var allKeys []dskey.Key
		for _, fqid := range fqids {
			collectionName, id, err := getCollectionNameAndID(fqid)
			if err != nil {
				updateFn(nil, fmt.Errorf("split fqid from %s: %w", fqid, err))
				return
			}

			keys, err := createKeyList(collectionName, id)
			if err != nil {
				updateFn(nil, fmt.Errorf("creating key list from notification: %w", err))
				return
			}

			allKeys = append(allKeys, keys...)
		}

		values, err := getWithConn(ctx, conn.Conn(), allKeys...)
		if err != nil {
			updateFn(nil, fmt.Errorf("fetching keys %v: %w", allKeys, err))
		}

		updateFn(values, nil)
	}
}

func createKeyList(collection string, id int) ([]dskey.Key, error) {
	fields := collectionFields[collection]

	keys := make([]dskey.Key, len(fields))
	var err error
	for i, field := range fields {
		keys[i], err = dskey.FromParts(collection, id, field)
		if err != nil {
			return nil, fmt.Errorf("creating key from parts %q, %d, %q: %w", collection, id, field, err)
		}
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
