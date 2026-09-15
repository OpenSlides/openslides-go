package dsgen

import (
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-go/collection"
)

// GoName formats a snake_case or identifier string to Go CamelCase (e.g. "meeting_id" -> "MeetingID").
func GoName(name string) string {
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

// FirstLower returns the string with the first letter in lowercase.
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

// WithoutID removes "ID" or "IDs" suffix and appends "List" if "IDs" was removed.
func WithoutID(in string) string {
	result := strings.TrimSuffix(strings.TrimSuffix(in, "ID"), "IDs")
	if strings.HasSuffix(in, "IDs") {
		return result + "List"
	}
	return result
}

// EnumName returns the Go type name for a field enum or global enum.
func EnumName(collection, fieldName string, field *collection.Field) string {
	enumName := GoName(collection) + "_" + GoName(fieldName)
	if field.Enum.GlobalName != "" {
		enumName = GoName(field.Enum.GlobalName)
	}

	return enumName
}

// ZeroValue returns the Go zero value string representation for a given Go type.
func ZeroValue(t string) string {
	switch t {
	case "int", "float64", "float32":
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

// ValueType returns the fetch Value type name for a field and the related go
// type. Handling relations, base types, and enums.
func ValueType(collection, fieldName string, field *collection.Field) (string, string) {
	//required := field.Required
	collectionType := field.Type

	if field.Enum.GlobalName != "" || len(field.Enum.Values) != 0 {
		goType := fmt.Sprintf("dstypes.%s", EnumName(collection, fieldName, field))
		return fmt.Sprintf("ValueEnum[%s]", goType), goType
	}

	switch collectionType {
	case "number", "relation", "timestamp":
		return "ValueInt", "int"
	case "string", "text", "HTMLStrict", "color", "HTMLPermissive", "generic-relation", "template", "timezone":
		return "ValueString", "string"
	case "decimal(6)":
		return "ValueDecimal", "decimal.Decimal"
	case "boolean":
		return "ValueBool", "bool"
	case "float":
		return "ValueFloat", "float64"
	case "relation-list", "number[]":
		return "ValueIntSlice", "[]int"
	case "JSON":
		return "ValueJSON", "json.RawMessage"
	case "string[]", "text[]", "generic-relation-list":
		return "ValueStringSlice", "[]string"
	default:
		panic(fmt.Sprintf("Unknown type %q", collectionType))
	}
}
