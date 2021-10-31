package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// ListOfSpeakers handels the restriction for the list_of_speakers collection.
type ListOfSpeakers struct{}

// Modes returns the restrictions modes for the meeting collection.
func (los ListOfSpeakers) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return los.see
	}
	return nil
}

func (los ListOfSpeakers) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, losID int) (bool, error) {
	mid, err := los.meetingID(ctx, ds, losID)
	if err != nil {
		return false, fmt.Errorf("fetching meeting id for los %d: %w", losID, err)
	}

	perms, err := mperms.Meeting(ctx, mid)
	if err != nil {
		return false, fmt.Errorf("getting perms for meetind %d: %w", mid, err)
	}

	return perms.Has(perm.ListOfSpeakersCanSee), nil
}

func (los ListOfSpeakers) meetingID(ctx context.Context, ds *datastore.Request, id int) (int, error) {
	mid, err := ds.ListOfSpeakers_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("fetching meeting_id for the list of speakers %d: %w", id, err)
	}
	return mid, nil
}
