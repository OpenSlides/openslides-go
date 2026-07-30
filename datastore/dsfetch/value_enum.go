package dsfetch

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
)

// ValueEnum[T] is a value from the datastore.
type ValueEnum[T any] struct {
	err error

	key      dskey.Key
	required bool

	lazies []*T

	fetch *Fetch
}

// Value returns the value.
func (v *ValueEnum[T]) Value(ctx context.Context) (T, error) {
	var zero T
	if err := v.err; err != nil {
		return zero, v.err
	}

	rawValue, err := v.fetch.getOneKey(ctx, v.key)
	if err != nil {
		return zero, err
	}

	value, err := v.convert(rawValue)
	if err != nil {
		return zero, fmt.Errorf("converting raw value: %w", err)
	}

	return value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueEnum[T]) Lazy(value *T) {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
	v.lazies = append(v.lazies, value)
}

// convert converts the json value to the type.
func (v *ValueEnum[T]) convert(p []byte) (T, error) {
	var zero T
	if p == nil {
		if v.required {
			return zero, fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
		return zero, nil
	}
	var value T
	if err := json.Unmarshal(p, &value); err != nil {
		return zero, fmt.Errorf("decoding value %q: %w", p, err)
	}
	return value, nil
}

// setLazy sets the lazy values defiend with Lazy.
func (v *ValueEnum[T]) setLazy(p []byte) error {
	value, err := v.convert(p)
	if err != nil {
		return fmt.Errorf("converting value: %w", err)
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = value
	}

	return nil
}
