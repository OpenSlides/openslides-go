package dsmodels

import (
	"context"
)

type builderWrapperI interface {
	SetIds(ids []int)
}

type builderPtr[T any, M any] interface {
	lazy(ds *Fetch, id int) *M
	*T
}

type builder[C any, T builderPtr[C, M], M any] struct {
	ids      []int
	value    T
	idField  string
	relField string
	many     bool
	loaded   bool
	fetch    *Fetch
}

func (b *builder[C, T, M]) SetIds(ids []int) {
	b.ids = ids
}

func (b *builder[C, T, M]) Preload(rel builderWrapperI) {

}

func (b *builder[C, T, M]) First(ctx context.Context) (M, error) {
	c := b.value.lazy(b.fetch, b.ids[0])

	if err := b.fetch.Execute(ctx); err != nil {
		var zero M
		return zero, err
	}

	return *c, nil
}

func (b *builder[C, T, M]) Get(ctx context.Context) ([]M, error) {
	items := []M{}
	for _, id := range b.ids {
		items = append(items, *b.value.lazy(b.fetch, id))
	}

	if err := b.fetch.Execute(ctx); err != nil {
		return []M{}, err
	}

	return items, nil
}

/*
{{ range .Relations}}
func (c *{{$.GoName}}) {{.MethodName}}() {{.ResultType}}{
    if c.{{.StructFieldName}} == nil {
        {{- if .IsList}}
        refs := make([]*ValueCollection[{{.Type}}, *{{.Type}}], len(c.{{.FieldName}}))
        for i, id := range c.{{.FieldName}}{
            refs[i] = &ValueCollection[{{.Type}}, *{{.Type}}]{
                id: id,
                fetch: c.fetch,
            }
        }
        c.{{.StructFieldName}} = &RelationList[{{.Type}}, *{{.Type}}]{refs}
        {{- else if .Required}}
        c.{{.StructFieldName}} = &ValueCollection[{{.Type}}, *{{.Type}}]{
		    id: c.{{.FieldName}},
		    fetch: c.fetch,
	    }
	    {{- else}}
        var ref dsfetch.Maybe[*ValueCollection[{{.Type}}, *{{.Type}}]]
	    id, hasValue := c.{{.FieldName}}.Value()
	    if hasValue {
	        value := &ValueCollection[{{.Type}}, *{{.Type}}]{
		        id:    id,
		        fetch: c.fetch,
	        }
	        ref.Set(value)
	    }
        c.{{.StructFieldName}} = &MaybeRelation[{{.Type}}, *{{.Type}}]{ref}
	    {{- end}}
	}
    return c.{{.StructFieldName}}
}
{{end}}
*/
