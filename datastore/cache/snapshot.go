package cache

import (
	"context"
	"errors"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/benbjohnson/immutable"
)

var ErrIncompleteSnapshot = errors.New("snapshot is incomplete")

// Snapshot implements the flow.Getter interface over an immutable map.
type Snapshot struct {
	data *immutable.Map[dskey.Key, []byte]
}

// Get returns keys from the snapshot.
func (s Snapshot) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	out := make(map[dskey.Key][]byte, len(keys))
	for _, key := range keys {
		value, ok := s.data.Get(key)
		if !ok {
			return nil, ErrIncompleteSnapshot
		}
		out[key] = value
	}

	return out, nil
}
