package maybe

import "encoding/json"

// Maybe holds a type or null.
type Maybe[T any] struct {
	value    T
	hasValue bool
}

// New initializes a Maybe with a value.
func New[T any](v T) Maybe[T] {
	return Maybe[T]{
		hasValue: true,
		value:    v,
	}
}

func (m Maybe[T]) Value() (T, bool) {
	if m.hasValue {
		return m.value, true
	}
	var zero T
	return zero, false
}

func (m Maybe[T]) Null() bool {
	return !m.hasValue
}

func (m *Maybe[T]) Set(v T) {
	m.value = v
	m.hasValue = true
}

func (m *Maybe[T]) SetNull() {
	m.hasValue = false
}

func (m *Maybe[T]) OrZero() T {
	if m.hasValue {
		return m.value
	}
	return *new(T)
}

func (m *Maybe[T]) OrElse(v T) T {
	if m.hasValue {
		return m.value
	}
	return v
}

func (m *Maybe[T]) UnmarshalJSON(bs []byte) error {
	if string(bs) == "null" {
		m.SetNull()
		return nil
	}

	var v T
	if err := json.Unmarshal(bs, &v); err != nil {
		return err
	}
	m.Set(v)
	return nil
}

func (m Maybe[T]) MarshalJSON() ([]byte, error) {
	if !m.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(m.value)
}
