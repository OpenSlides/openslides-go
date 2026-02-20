package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"log"
	"os"
	"slices"

	"github.com/OpenSlides/openslides-go/collection"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	td, err := parse("../meta")
	if err != nil {
		return fmt.Errorf("parse collections: %w", err)
	}

	if err := write(os.Stdout, td); err != nil {
		return fmt.Errorf("write to stdout: %w", err)
	}

	return nil
}

type templateData struct {
	Collection map[string][]string
}

func parse(path string) (templateData, error) {
	inData, err := collection.Collections(path)
	if err != nil {
		return templateData{}, fmt.Errorf("parse collections: %w", err)
	}

	td := templateData{
		Collection: make(map[string][]string),
	}
	for collectionName, collection := range inData {
		for fieldName := range collection.Fields {
			td.Collection[collectionName] = append(td.Collection[collectionName], fieldName)
		}

		slices.Sort(td.Collection[collectionName])
	}

	return td, nil
}

const tpl = `// Code generated from meta collections. DO NOT EDIT.
package datastore

var collectionFields = map[string][]string{
	{{- range $key, $value := .Collection}}
		"{{$key}}": { {{range $field := $value}} "{{$field}}", {{end}} },
	{{- end}}
}
`

func write(w io.Writer, td templateData) error {
	t := template.New("t")
	t, err := t.Parse(tpl)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	buf := new(bytes.Buffer)

	if err := t.Execute(buf, td); err != nil {
		return fmt.Errorf("writing template: %w", err)
	}

	formated, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("formating code: %w", err)
	}

	if _, err := w.Write(formated); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}

	return nil
}
