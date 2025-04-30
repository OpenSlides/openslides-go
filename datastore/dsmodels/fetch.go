package dsmodels

import (
	"context"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/flow"
)

//go:generate sh -c "go run gen_models/main.go > models_generated.go"

// Fetch provides functions to access the fields of the datastore.
//
// Fetch is not save for concurent use. One Fetch object AND its value can only be
// used in one goroutine.
type Fetch struct {
	getter flow.Getter
	dsfetch.Fetch
}

// New initializes a Request object.
func New(getter flow.Getter) *Fetch {
	r := Fetch{
		getter: getter,
		Fetch:  *dsfetch.New(getter),
	}

	return &r
}

// Get calls the getter the flow was created with.
func (f *Fetch) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	return f.getter.Get(ctx, keys...)
}
