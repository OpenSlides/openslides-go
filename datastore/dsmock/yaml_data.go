package dsmock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/goccy/go-yaml"
)

// YAMLData creates key values from a yaml object.
//
// panics on error.
//
// Deprecated: Use [YAMLDataErr] instead. This function could be removed in the
// future.
func YAMLData(input string) map[dskey.Key][]byte {
	r, err := YAMLDataErr(input)
	if err != nil {
		panic(err)
	}
	return r
}

// YAMLDataErr creates key values from a yaml object.
func YAMLDataErr(input string) (map[dskey.Key][]byte, error) {
	input = strings.ReplaceAll(input, "\t", "  ")

	var db map[string]any
	if err := yaml.Unmarshal([]byte(input), &db); err != nil {
		return nil, fmt.Errorf("unmarshalling yaml: %w", err)
	}

	data := make(map[dskey.Key][]byte)
	for dbKey, dbValue := range db {
		parts := strings.Split(dbKey, "/")
		switch len(parts) {
		case 1:
			map1, ok := dbValue.(map[string]any)
			if !ok {
				return nil, fmt.Errorf("invalid type in db key %s: %T", dbKey, dbValue)
			}
			for rawID, rawObject := range map1 {
				id, err := strconv.Atoi(rawID)
				if err != nil {
					return nil, fmt.Errorf("invalid id type: got %T expected int", rawID)
				}
				field, ok := rawObject.(map[string]any)
				if !ok {
					return nil, fmt.Errorf("invalid object type: got %T, expected map[string]any", rawObject)
				}

				for fieldName, fieldValue := range field {
					key, err := dskey.FromParts(dbKey, id, fieldName)
					if err != nil {
						return nil, fmt.Errorf("creating key from %s, %d, %s: %w", dbKey, id, fieldName, err)
					}

					bs, err := json.Marshal(fieldValue)
					if err != nil {
						return nil, fmt.Errorf("creating value for key %s: %w", key, err)
					}

					data[key] = bs
				}

				idKey, err := dskey.FromParts(dbKey, id, "id")
				if err != nil {
					return nil, fmt.Errorf("creating id-key from %s, %d: %w", dbKey, id, err)
				}
				data[idKey] = []byte(strconv.Itoa(id))
			}

		case 2:
			field, ok := dbValue.(map[string]any)
			if !ok {
				return nil, fmt.Errorf("invalid object type: got %T, expected map[string]interface{}", dbValue)
			}

			id, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("converting to int %s: %w", parts[1], err)
			}

			for fieldName, fieldValue := range field {
				fqfield, err := dskey.FromParts(parts[0], id, fieldName)
				if err != nil {
					return nil, fmt.Errorf("creating key from %s, %d, %s: %w", parts[0], id, fieldName, err)
				}

				bs, err := json.Marshal(fieldValue)
				if err != nil {
					return nil, fmt.Errorf("creating value for key %s: %w", fqfield, err)
				}
				data[fqfield] = bs
			}

			idKey, err := dskey.FromParts(parts[0], id, "id")
			if err != nil {
				return nil, fmt.Errorf("creating id-key from %s, %d: %w", parts[0], id, err)
			}
			data[idKey] = []byte(parts[1])

		case 3:
			key := dskey.MustKey(dbKey)
			bs, err := json.Marshal(dbValue)
			if err != nil {
				return nil, fmt.Errorf("creating test db. Key %s: %w", dbKey, err)
			}

			data[key] = bs

			id, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("converting to int %s: %w", parts[1], err)
			}

			idKey, err := dskey.FromParts(parts[0], id, "id")
			if err != nil {
				return nil, fmt.Errorf("creating id-key from %s, %d: %w", parts[0], id, err)
			}
			data[idKey] = []byte(parts[1])
		default:
			return nil, fmt.Errorf("invalid db key %s", dbKey)
		}
	}

	for k, v := range data {
		if bytes.Equal(v, []byte("null")) {
			data[k] = nil
		}
	}

	return data, nil
}
