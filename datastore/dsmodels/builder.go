package dsmodels

import (
	"context"
	"reflect"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

type builderWrapperI interface {
	SetIds(ids []int)
	setChild(builder builderWrapperI) builderWrapperI
	getParent() builderWrapperI
	getIDField() string
	getRelField() string
	getMany() bool
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
	children map[string]builderWrapperI
	idField  string
	relField string
	many     bool
	fetch    *Fetch
}

func (b *builder[C, T, M]) SetIds(ids []int) {
	b.ids = ids
}

func (b *builder[C, T, M]) setChild(builder builderWrapperI) builderWrapperI {
	if b.children == nil {
		b.children = map[string]builderWrapperI{}
	}

	if _, ok := b.children[builder.getRelField()]; !ok {
		b.children[builder.getRelField()] = builder
	}

	return b.children[builder.getRelField()]
}

func (b *builder[C, T, M]) getRelField() string {
	return b.relField
}

func (b *builder[C, T, M]) getIDField() string {
	return b.idField
}

func (b *builder[C, T, M]) getMany() bool {
	return b.many
}

func (b *builder[C, T, M]) getParent() builderWrapperI {
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
	for rel != b && rel != nil && rel.getRelField() != "" {
		children = append([]builderWrapperI{rel}, children...)
		rel = rel.getParent()
	}

	var cParent builderWrapperI
	cParent = b
	for _, cRel := range children {
		cParent = cParent.setChild(cRel)
	}
}

func (b *builder[C, T, M]) loadChildren(ctx context.Context, parent any) error {
	if b.children == nil {
		return nil
	}

	rParent := reflect.ValueOf(parent).Elem()
	for _, child := range b.children {
		ids := []int{}
		idField := rParent.FieldByName(child.getIDField())
		targetField := rParent.FieldByName(child.getRelField())
		if child.getMany() {
			ids = idField.Interface().([]int)
		} else if idField.Kind() == reflect.Int {
			ids = append(ids, int(idField.Int()))
		} else if idField.Type().Name() == "Maybe[int]" {
			relMaybeType := targetField.Type().Elem()
			relValue := reflect.New(relMaybeType)
			relValue.MethodByName("SetNull").Call([]reflect.Value{})
			targetField.Set(relValue)

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

		for _, item := range items {
			if err := child.loadChildren(ctx, item); err != nil {
				return err
			}

			if child.getMany() {
				targetField.Set(reflect.Append(targetField, reflect.ValueOf(item).Elem()))
			} else if idField.Type().Name() == "Maybe[int]" {
				targetField.MethodByName("Set").Call([]reflect.Value{reflect.ValueOf(item).Elem()})
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
