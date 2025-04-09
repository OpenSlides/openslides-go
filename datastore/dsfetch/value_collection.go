package dsfetch

import (
	"context"
	"fmt"
	"reflect"
	"strings"
)

type loader[T any] interface {
	lazy(ds *Fetch, id int)
	*T
}

// ValueCollection is a generic struct, where the loader interface is
// implemented by the pointer of C.
type ValueCollection[C any, T loader[C]] struct {
	id     int
	fetch  *Fetch
	loaded bool
	value  T
	load   map[string]map[string]struct{}
}

func (v *ValueCollection[T, P]) Value(ctx context.Context) (T, error) {
	var collection T
	v.Lazy(&collection)

	if err := v.fetch.Execute(ctx); err != nil {
		var zero T
		return zero, err
	}

	return collection, nil
}

func (v *ValueCollection[T, P]) Lazy(collection P) {
	collection.lazy(v.fetch, v.id)
}

func (v *ValueCollection[T, P]) Preload(field string) *ValueCollection[T, P] {
	parts := strings.SplitN(field, ".", 2)
	if v.load == nil {
		v.load = map[string]map[string]struct{}{}
	}

	if _, ok := v.load[parts[0]]; !ok {
		v.load[parts[0]] = map[string]struct{}{}
	}

	if len(parts) == 2 {
		v.load[parts[0]][parts[1]] = struct{}{}
		return v
	}

	return v
}

func (v *ValueCollection[T, P]) Load(ctx context.Context) (T, error) {
	collection, err := v.Value(ctx)
	if err != nil {
		return collection, err
	}

	v.value = &collection
	if v.load != nil {
		for field, subRelations := range v.load {
			ref := reflect.ValueOf(&collection).MethodByName(field).Call([]reflect.Value{})
			rel := ref[0]
			if rel.Kind() == reflect.Ptr {
				typeName := rel.Elem().Type().Name()
				if strings.HasPrefix(typeName, "RelationList") {
					// relation is represented by RelationList
					relListReturn := rel.MethodByName("Refs").Call([]reflect.Value{})
					relList := relListReturn[0]
					for i := range relList.Len() {
						valColl := relList.Index(i)
						for field := range subRelations {
							valColl.MethodByName("Preload").Call([]reflect.Value{reflect.ValueOf(field)})
						}
						valColl.MethodByName("Load").Call([]reflect.Value{reflect.ValueOf(ctx)})
					}
				} else if strings.HasPrefix(typeName, "ValueCollection") {
					// relation is represented by *ValueCollection
					for field := range subRelations {
						rel.MethodByName("Preload").Call([]reflect.Value{reflect.ValueOf(field)})
					}
					rel.MethodByName("Load").Call([]reflect.Value{reflect.ValueOf(ctx)})
				}
			} else if rel.Kind() == reflect.Struct {
				// relation is represented by Maybe[*ValueCollection]
				maybeVal := rel.MethodByName("Value").Call([]reflect.Value{})
				if maybeVal[1].Bool() {
					for field := range subRelations {
						maybeVal[0].MethodByName("Preload").Call([]reflect.Value{reflect.ValueOf(field)})
					}
					maybeVal[0].MethodByName("Load").Call([]reflect.Value{reflect.ValueOf(ctx)})
				}
			} else {
				var zero T
				return zero, fmt.Errorf("invalid field name added to preload")
			}
		}
	}

	v.loaded = true

	return collection, nil
}

func (v *ValueCollection[T, P]) Get() P {
	if !v.loaded {
		panic("accessing data that was not preloaded")
	}

	return v.value
}
