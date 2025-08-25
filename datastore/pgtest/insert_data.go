package pgtest

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmock"
	"github.com/jackc/pgx/v5"
)

func insertTestData(ctx context.Context, conn *pgx.Conn, testData string) error {
	yml := dsmock.YAMLData(testData)
	dataByTable := make(map[string]map[int]map[string]any)

	for key, jsonValue := range yml {
		collection := key.Collection()
		id := key.ID()
		field := key.Field()

		if dataByTable[collection] == nil {
			dataByTable[collection] = make(map[int]map[string]any)
		}
		if dataByTable[collection][id] == nil {
			dataByTable[collection][id] = make(map[string]any)
		}

		var value any
		if err := json.Unmarshal(jsonValue, &value); err != nil {
			return fmt.Errorf("failed to unmarshal JSON for %s.%d.%s: %w", collection, id, field, err)
		}

		dataByTable[collection][id][field] = value
	}

	for tableName, rows := range dataByTable {
		if err := insertRowsForTable(ctx, conn, tableName, rows); err != nil {
			return fmt.Errorf("failed to insert data for table %s: %w", tableName, err)
		}
	}

	return nil
}

func insertRowsForTable(ctx context.Context, conn *pgx.Conn, tableName string, rows map[int]map[string]any) error {
	if len(rows) == 0 {
		return nil
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

	batch := &pgx.Batch{}
	placeholders := make([]string, len(columns))
	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) ON CONFLICT (id) DO UPDATE SET %s",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
		buildUpdateClause(columns),
	)

	// Füge alle Zeilen zum Batch hinzu
	for id, rowData := range rows {
		values := make([]interface{}, len(columns))

		for i, column := range columns {
			if column == "id" {
				values[i] = id
			} else if value, exists := rowData[column]; exists {
				values[i] = value
			} else {
				values[i] = nil
			}
		}

		batch.Queue(query, values...)
	}

	results := conn.SendBatch(ctx, batch)
	defer results.Close()

	for i := 0; i < len(rows); i++ {
		_, err := results.Exec()
		if err != nil {
			return fmt.Errorf("failed to execute batch insert row %d: %w", i, err)
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
