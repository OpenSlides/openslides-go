package cache

import (
	"context"
	"fmt"
	"maps"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/benbjohnson/immutable"
)

// Snapshot implements the flow.Getter interface over an immutable map.
type Snapshot struct {
	data            *immutable.Map[dskey.Key, []byte]
	notFoundHandler flow.Getter
}

// Get returns keys from the snapshot.
func (s Snapshot) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	out := make(map[dskey.Key][]byte, len(keys))
	var notFound []dskey.Key
	for _, key := range keys {
		value, ok := s.data.Get(key)
		if !ok {
			notFound = append(notFound, key)
			continue
		}
		out[key] = value
	}

	if len(notFound) > 0 {
		if s.notFoundHandler == nil {
			return nil, IncompleteSnapshotError{notFound}
		}
		found, err := s.notFoundHandler.Get(ctx, notFound...)
		if err != nil {
			return nil, fmt.Errorf("not found handler: %w", err)
		}

		maps.Copy(out, found)
	}

	return out, nil
}

// IncompleteSnapshotError is returned when a snapshot is ask for data, that does
// not exist in the snapshot.
type IncompleteSnapshotError struct {
	Keys []dskey.Key
}

// Error returns a string representation of the error.
func (err IncompleteSnapshotError) Error() string {
	return fmt.Sprintf("did not found %d keys in snapshot", len(err.Keys))
}
