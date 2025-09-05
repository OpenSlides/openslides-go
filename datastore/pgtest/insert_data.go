package pgtest

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/dsmock"
	"github.com/OpenSlides/openslides-go/fastjson"
	"github.com/OpenSlides/openslides-go/metagen"
	"github.com/jackc/pgx/v5"
)

func insertTestData(ctx context.Context, conn *pgx.Conn, testData string) error {
	yml := dsmock.YAMLData(testData)
	dataByTable := make(map[string]map[int]map[string]any)
	var nmData []nmInfo

	for key, jsonValue := range yml {
		collection := key.Collection()
		id := key.ID()
		field := key.Field()

		if strings.HasSuffix(metagen.RelationListFields[key.CollectionField()], "_ids") {
			// nm-tale
			nm, err := genNMInfo(key, jsonValue)
			if err != nil {
				return fmt.Errorf("generate NM info for %s/%d/%s: %w", collection, id, field, err)
			}

			nmData = append(nmData, nm)
			continue
		}

		if dataByTable[collection] == nil {
			dataByTable[collection] = make(map[int]map[string]any)
		}
		if dataByTable[collection][id] == nil {
			dataByTable[collection][id] = make(map[string]any)
		}

		var value any
		if err := json.Unmarshal(jsonValue, &value); err != nil {
			return fmt.Errorf("unmarshal JSON for %s/%d/%s: %w", collection, id, field, err)
		}

		dataByTable[collection][id][field] = value
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	if _, err = tx.Exec(ctx, "SET CONSTRAINTS ALL DEFERRED"); err != nil {
		return fmt.Errorf("set constraints deferred: %w", err)
	}

	if err := insertData(ctx, tx, dataByTable); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}

	for _, nm := range nmData {
		if err := inserNMData(ctx, tx, nm); err != nil {
			return fmt.Errorf("insert nm-data for %s: %w", nm.table, err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func insertData(ctx context.Context, tx pgx.Tx, dataByTable map[string]map[int]map[string]any) error {
	for tableName, rows := range dataByTable {
		if len(rows) == 0 {
			continue
		}

		columnSet := make(map[string]bool)
		columnSet["id"] = true
		for _, row := range rows {
			for field := range row {
				columnSet[field] = true
			}
		}
		columns := make([]string, 0, len(columnSet))
		for column := range columnSet {
			columns = append(columns, column)
		}
		sort.Strings(columns)

		if len(columns) == 1 {
			// Only id field. Skip this
			continue
		}

		placeholders := make([]string, len(columns))
		for i := range placeholders {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
		}

		query := fmt.Sprintf(
			"INSERT INTO \"%s\" (%s) VALUES (%s) ON CONFLICT (id) DO UPDATE SET %s",
			tableName,
			strings.Join(columns, ", "),
			strings.Join(placeholders, ", "),
			buildUpdateClause(columns),
		)

		for id, rowData := range rows {
			values := make([]any, len(columns))
			for i, column := range columns {
				if column == "id" {
					values[i] = id
				} else if value, exists := rowData[column]; exists {
					values[i] = value
				} else {
					values[i] = nil
				}
			}

			if _, err := tx.Exec(ctx, query, values...); err != nil {
				return fmt.Errorf("execute insert for id %d: %w", id, err)
			}
		}
	}

	return nil
}

func buildUpdateClause(columns []string) string {
	var updates []string
	for _, column := range columns {
		if column != "id" {
			updates = append(updates, fmt.Sprintf("%s = EXCLUDED.%s", column, column))
		}
	}
	return strings.Join(updates, ", ")
}

type nmInfo struct {
	table string
	id1   []int
	id2   []int
}

func genNMInfo(key dskey.Key, value []byte) (nmInfo, error) {
	cfs := make([]string, 2)
	cfs[0] = key.CollectionField()
	cfs[1] = metagen.RelationListFields[key.CollectionField()]
	reverse := false
	if cfs[0] > cfs[1] {
		cfs[0], cfs[1] = cfs[1], cfs[0]
		reverse = true
	}

	c1, f1 := splitCollectionName(cfs[0])
	c2, _ := splitCollectionName(cfs[1])
	nmTableName := fmt.Sprintf("nm_%s_%s_%s_t", c1, f1, c2)

	otherIDs, err := fastjson.DecodeIntList(value)
	if err != nil {
		return nmInfo{}, fmt.Errorf("decode value: %w", err)
	}

	if reverse {
		return nmInfo{
			table: nmTableName,
			id1:   otherIDs,
			id2:   []int{key.ID()},
		}, nil
	}
	return nmInfo{
		table: nmTableName,
		id1:   []int{key.ID()},
		id2:   otherIDs,
	}, nil
}

func splitCollectionName(cf string) (string, string) {
	parts := strings.Split(cf, "/")
	if len(parts) != 2 {
		panic(fmt.Sprintf("invalid collection name %s", cf))
	}
	return parts[0], parts[1]
}

func inserNMData(ctx context.Context, tx pgx.Tx, nm nmInfo) error {
	batch := &pgx.Batch{}

	for _, id1 := range nm.id1 {
		for _, id2 := range nm.id2 {
			query := fmt.Sprintf("INSERT INTO \"%s\" VALUES ($1, $2)", nm.table)
			batch.Queue(query, id1, id2)
		}
	}

	results := tx.SendBatch(ctx, batch)
	defer results.Close()

	for i := range len(nm.id1) * len(nm.id2) {
		_, err := results.Exec()
		if err != nil {
			return fmt.Errorf("execute batch insert row %d: %w", i, err)
		}
	}

	return nil
}
