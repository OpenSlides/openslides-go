package dsmodels

import (
	"context"
	"reflect"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

type builderWrapperI interface {
	SetIds(ids []int)
	SetChild(builder builderWrapperI) builderWrapperI
	RelParent() builderWrapperI
	RelIdField() string
	RelField() string
	RelMany() bool
	loadChildren(ctx context.Context, parent any) error
	lazyAll(ctx context.Context) []any
}

type builderPtr[T any, M any] interface {
	lazy(ds *Fetch, id int) *M
	*T
}

type builder[C any, T builderPtr[C, M], M any] struct {
	ids      []int
	value    T
	parent   builderWrapperI
	children   map[string]builderWrapperI
	idField  string
	relField string
	many     bool
	loaded   bool
	fetch    *Fetch
}

func (b *builder[C, T, M]) SetIds(ids []int) {
	b.ids = ids
}

func (b *builder[C, T, M]) SetChild(builder builderWrapperI) builderWrapperI {
	if b.children == nil {
		b.children = map[string]builderWrapperI{}
	}

	if _, ok := b.children[builder.RelField()]; !ok {
		b.children[builder.RelField()] = builder
	}

	return b.children[builder.RelField()]
}

func (b *builder[C, T, M]) RelField() string {
	return b.relField
}

func (b *builder[C, T, M]) RelIdField() string {
	return b.idField
}

func (b *builder[C, T, M]) RelMany() bool {
	return b.many
}

func (b *builder[C, T, M]) RelParent() builderWrapperI {
	return b.parent
}

func (b *builder[C, T, M]) lazyAll(ctx context.Context) []any {
	items := []any{}
	for _, id := range b.ids {
		items = append(items, b.value.lazy(b.fetch, id))
	}
	return items
}

func (b *builder[C, T, M]) Preload(rel builderWrapperI) {
	children := []builderWrapperI{}
	for rel != b && rel != nil && rel.RelField() != "" {
		children = append([]builderWrapperI{rel}, children...)
		rel = rel.RelParent()
	}

	var cParent builderWrapperI
	cParent = b
	for _, cRel := range children {
		cParent = cParent.SetChild(cRel)
	}
}

func (b *builder[C, T, M]) loadChildren(ctx context.Context, parent any) error {
	if b.children == nil {
		return nil
	}

	rParent := reflect.ValueOf(parent).Elem()
	for _, child := range b.children {
		ids := []int{}
		idField := rParent.FieldByName(child.RelIdField())
		if child.RelMany() {
			ids = idField.Interface().([]int)
		} else if idField.Kind() == reflect.Int {
			ids = append(ids, int(idField.Int()))
		} else if idField.Type().Name() == "Maybe[int]" {
			id := idField.Interface().(dsfetch.Maybe[int])
			if val, set := id.Value(); set {
				ids = append(ids, val)
			}
		}
		child.SetIds(ids)

		items := child.lazyAll(ctx)
		if err := b.fetch.Execute(ctx); err != nil {
			return err
		}

		targetField := rParent.FieldByName(child.RelField())
		for _, item := range items {
			if err := child.loadChildren(ctx, item); err != nil {
				return err
			}

			if child.RelMany() {
				targetField.Set(reflect.Append(targetField, reflect.ValueOf(item).Elem()))
			} else {
				targetField.Set(reflect.ValueOf(item))
			}
		}
	}

	return nil
}

func (b *builder[C, T, M]) First(ctx context.Context) (M, error) {
	c := b.value.lazy(b.fetch, b.ids[0])

	if err := b.fetch.Execute(ctx); err != nil {
		var zero M
		return zero, err
	}

	if err := b.loadChildren(ctx, c); err != nil {
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

	for _, el := range items {
		if err := b.loadChildren(ctx, &el); err != nil {
			return []M{}, err
		}
	}

	return items, nil
}
