package maybe_test

import (
	"encoding/json"
	"testing"

	"github.com/OpenSlides/openslides-go/maybe"
)

func TestMaybe(t *testing.T) {
	var maybeInt maybe.Maybe[int]

	if got := maybeInt.Null(); !got {
		t.Errorf("empty Maybe[int].Null() == %v, expected true", got)
	}

	if _, got := maybeInt.Value(); got {
		t.Errorf("empty Maybe[int].Value() == _, %v, expected false", got)
	}

	maybeInt.Set(0)

	if got := maybeInt.Null(); got {
		t.Errorf("setted Maybe[int].Null() == %v, expected false", got)
	}

	if _, got := maybeInt.Value(); !got {
		t.Errorf("setted Maybe[int].Value() == _, %v, expected true", got)
	}

	if got, _ := maybeInt.Value(); got != 0 {
		t.Errorf("setted Maybe[int].Value() == %v, expected 0", got)
	}

	other := maybe.New(0)

	if other != maybeInt {
		t.Errorf("setted Maybe[int] != MaybeValue: %v, %v", other, maybeInt)
	}
}

func TestMaybeUnmarshall(t *testing.T) {
	jsonzero := []byte(`null`)

	var value maybe.Maybe[int]
	if err := json.Unmarshal(jsonzero, &value); err != nil {
		t.Fatalf("Error: %v", err)
	}
	if !value.Null() {
		v, isNull := value.Value()
		t.Errorf("value is not null, got: %v, %v", v, isNull)
	}
}
