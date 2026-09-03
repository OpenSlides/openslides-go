// This tool generates the enum types for the datastore.
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
	"strings"
	"text/template"

	"github.com/OpenSlides/openslides-go/collection"
	"github.com/OpenSlides/openslides-go/datastore/dsgen"
)

//go:embed header.go.tmpl
var tmplHeader string

//go:embed enum.go.tmpl
var tmplEnum string

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

	if err := genEnums(buf, fromYml); err != nil {
		return fmt.Errorf("generate enums types: %w", err)
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
				name := dsgen.EnumName(colName, fieldName, field)
				enumMap[name] = []enumField{}
				for _, enumValue := range field.Enum.Values {
					goEnumName := dsgen.GoName(strings.ReplaceAll(enumValue, "-", "_"))
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

	tmpl, err := template.New("enum.go.tmpl").Parse(tmplEnum)
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
