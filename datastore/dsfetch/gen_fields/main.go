// This tool generates the code needed for the request object.
//
// To call it, just call "go generate ./..." in the root folder of the repository
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
	"text/template"

	"github.com/OpenSlides/openslides-go/collection"
)

//go:embed value.go.tmpl
var tmplValue string

//go:embed header.go.tmpl
var tmplHeader string

//go:embed enum.go.tmpl
var tmplEnum string

//go:embed field.go.tmpl
var tmplField string

func main() {
	if err := run(os.Stdout); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run(w io.Writer) error {
	fromYml, err := collection.Collections("../../meta/")
	if err != nil {
		return fmt.Errorf("parse collections: %w", err)
	}

	buf := new(bytes.Buffer)

	if err := genHeader(buf); err != nil {
		return fmt.Errorf("generate file header: %w", err)
	}

	if err := genValueTypes(buf); err != nil {
		return fmt.Errorf("generate value types: %w", err)
	}

	if err := genEnums(buf, fromYml); err != nil {
		return fmt.Errorf("generate enums types: %w", err)
	}

	if err := genFieldMethods(buf, fromYml); err != nil {
		return fmt.Errorf("generate field methods: %w", err)
	}

	formated, err := format.Source(buf.Bytes())
	if err != nil {
		if _, err := w.Write(buf.Bytes()); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
		return fmt.Errorf("formating code: %w", err)
	}

	if _, err := w.Write(formated); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}
	return nil
}

func genHeader(buf *bytes.Buffer) error {
	tmpl, err := template.New("header.go").Parse(tmplHeader)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	if err := tmpl.Execute(buf, nil); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	return nil
}

var typesToGo = map[string]string{
	"ValueInt":         "int",
	"ValueMaybeInt":    "Maybe[int]",
	"ValueString":      "string",
	"ValueMaybeString": "Maybe[string]",
	"ValueDecimal":     "decimal.Decimal",
	"ValueBool":        "bool",
	"ValueFloat":       "float64",
	"ValueJSON":        "json.RawMessage",
	"ValueIntSlice":    "[]int",
	"ValueStringSlice": "[]string",
}

func genValueTypes(buf *bytes.Buffer) error {
	// Make sure the types are in the same order every time go generate runs.
	var types []string
	for k := range typesToGo {
		types = append(types, k)
	}
	sort.Strings(types)

	tmpl, err := template.New("value.go").Parse(tmplValue)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	for _, name := range types {
		data := struct {
			TypeName  string
			GoType    string
			Zero      string
			MaybeType bool
		}{
			name,
			typesToGo[name],
			zeroValue(typesToGo[name]),
			strings.HasPrefix(name, "ValueMaybe"),
		}
		if err := tmpl.Execute(buf, data); err != nil {
			return fmt.Errorf("executing template: %w", err)
		}
	}
	return nil
}

// zeroValue returns the zero value for a go type
func zeroValue(t string) string {
	switch t {
	case "int", "float64":
		return "0"
	case "string":
		return `""`
	case "bool":
		return "false"
	case "json.RawMessage", "[]int", "[]string":
		return "nil"
	}
	return "unknown type " + t
}

func genEnums(buf *bytes.Buffer, fromYML map[string]collection.Collection) error {
	// Make sure the types are in the same order every time go generate runs.
	type enumField struct {
		GoName   string
		RawValue string
	}

	enumMap := make(map[string][]enumField)
	for colName, col := range fromYML {
		for fieldName, field := range col.Fields {
			if len(field.Enum.Values) > 0 || field.Enum.GlobalName != "" {
				name := enumName(colName, fieldName, field)
				enumMap[name] = []enumField{}
				for _, enumValue := range field.Enum.Values {
					goEnumName := goName(strings.ReplaceAll(enumValue, "-", "_"))
					if goEnumName == "" {
						goEnumName = "empty"
					}

					enumMap[name] = append(enumMap[name], enumField{
						GoName:   goEnumName,
						RawValue: enumValue,
					})
				}
			}
		}
	}

	type enumData struct {
		GoName string
		Fields []enumField
	}

	enums := []enumData{}
	for name, values := range enumMap {
		enums = append(enums, enumData{
			GoName: name,
			Fields: values,
		})
	}

	slices.SortFunc(enums, func(a, b enumData) int {
		return strings.Compare(a.GoName, b.GoName)
	})

	tmpl, err := template.New("value_enum.go").Parse(tmplEnum)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	for _, e := range enums {
		if err := tmpl.Execute(buf, map[string]any{
			"Name":   e.GoName,
			"Values": e.Fields,
		}); err != nil {
			return fmt.Errorf("executing template: %w", err)
		}
	}

	return nil
}

func genFieldMethods(buf *bytes.Buffer, fromYML map[string]collection.Collection) error {
	fields, err := toFields(fromYML)
	if err != nil {
		return fmt.Errorf("generate field definitions: %w", err)
	}

	tmpl, err := template.New("field.go").Parse(tmplField)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	for _, field := range fields {
		if err := tmpl.Execute(buf, field); err != nil {
			return fmt.Errorf("executing template: %w", err)
		}
	}

	return nil
}

// toFields returns all fields all collections with there go-type as string.
func toFields(raw map[string]collection.Collection) ([]field, error) {
	var fields []field
	for collectionName, collection := range raw {
		for fieldName, collectionField := range collection.Fields {
			f := field{}
			f.GoName = goName(collectionName) + "_" + goName(fieldName)
			f.ValueType = valueType(goName(collectionName), fieldName, collectionField)
			f.Collection = firstLower(goName(collectionName))
			f.CollectionName = collectionName
			f.FieldName = fieldName
			f.Required = collectionField.Required

			if collectionField.Type == "relation" || collectionField.Type == "generic-relation" {
				f.SingleRelation = true
			}

			fields = append(fields, f)
		}
	}

	sort.Slice(fields, func(i, j int) bool {
		return fields[i].GoName < fields[j].GoName
	})

	return fields, nil
}

type field struct {
	GoName         string
	ValueType      string
	Collection     string
	CollectionName string
	FieldName      string
	Required       bool
	SingleRelation bool
}

func goName(name string) string {
	if name == "id" {
		return "ID"
	}

	name = strings.ReplaceAll(name, "_$", "")

	parts := strings.Split(name, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	name = strings.Join(parts, "")

	name = strings.ReplaceAll(name, "Id", "ID")
	return name
}

func firstLower(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}

func enumName(collection string, fieldName string, field *collection.Field) string {
	enumName := goName(collection) + "_" + goName(fieldName)
	if field.Enum.GlobalName != "" {
		enumName = goName(field.Enum.GlobalName)
	}

	return enumName
}

func valueType(collection string, fieldName string, field *collection.Field) string {
	required := field.Required
	collectionType := field.Type

	if !required && collectionType == "relation" {
		return "ValueMaybeInt"
	}

	if !required && collectionType == "generic-relation" {
		return "ValueMaybeString"
	}

	if field.Enum.GlobalName != "" || len(field.Enum.Values) != 0 {
		return fmt.Sprintf("ValueEnum[%s]", enumName(collection, fieldName, field))
	}

	switch collectionType {
	case "number", "relation", "timestamp":
		return "ValueInt"

	case "string", "text", "HTMLStrict", "color", "HTMLPermissive", "generic-relation", "template", "timezone":
		return "ValueString"

	case "decimal(6)":
		return "ValueDecimal"

	case "boolean":
		return "ValueBool"

	case "float":
		return "ValueFloat"

	case "relation-list", "number[]":
		return "ValueIntSlice"

	case "JSON":
		return "ValueJSON"

	case "string[]", "text[]", "generic-relation-list":
		return "ValueStringSlice"

	default:
		panic(fmt.Sprintf("Unknown type %q", collectionType))
	}
}
