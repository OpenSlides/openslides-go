package dsmodels_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-go/datastore/dsmock"
	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

func TestRequestSingleModel(t *testing.T) {
	ds := dsmodels.New(dsmock.Stub(dsmock.YAMLData(`---
	topic/1/sequential_number: 1
	topic/1/title: foo
	topic/1/meeting_id: 1
	topic/1/agenda_item_id: 1
	topic/1/list_of_speakers_id: 1
	`)))

	_, err := ds.Topic(1).First(context.Background())
	if err != nil {
		t.Errorf("Topic 1 returned unexpected error: %v", err)
	}
}

func TestRequestSingleModelWithRequiredRelation(t *testing.T) {
	ds := dsmodels.New(dsmock.Stub(dsmock.YAMLData(`---
	topic/1/sequential_number: 1
	topic/1/title: foo
	topic/1/meeting_id: 1
	topic/1/agenda_item_id: 1
	topic/1/list_of_speakers_id: 1
	agenda_item/1/content_object_id: topic/1
	agenda_item/1/meeting_id: 1
	`)))

	tQ := ds.Topic(1)
	_, err := tQ.Preload(tQ.AgendaItem()).First(context.Background())
	if err != nil {
		t.Errorf("Topic 1 with agenda item returned unexpected error: %v", err)
	}
}

func TestRequestSingleModelWithExistingMaybeRelation(t *testing.T) {
	ds := dsmodels.New(dsmock.Stub(dsmock.YAMLData(`---
	agenda_item/1/content_object_id: topic/1
	agenda_item/1/meeting_id: 1
	agenda_item/1/parent_id: 2
	agenda_item/2/content_object_id: topic/2
	agenda_item/2/meeting_id: 1
	`)))

	q := ds.AgendaItem(1)
	res, err := q.Preload(q.Parent()).First(context.Background())
	if err != nil {
		t.Errorf("Agenda item returned unexpected error: %v", err)
	}

	if res.ContentObjectID != "topic/1" {
		t.Errorf("res.ContentObjectID = %s, expected topic/1", res.ContentObjectID)
	}

	if val, isSet := res.Parent.Value(); !isSet {
		t.Errorf("parent is not set")
	} else if val.ContentObjectID != "topic/2" {
		t.Errorf("res.ContentObjectID = %s, expected topic/2", val.ContentObjectID)
	}
}

func TestRequestSingleModelWithNonExistingMaybeRelation(t *testing.T) {
	ds := dsmodels.New(dsmock.Stub(dsmock.YAMLData(`---
	agenda_item/1/content_object_id: topic/1
	agenda_item/1/meeting_id: 1
	`)))

	q := ds.AgendaItem(1)
	res, err := q.Preload(q.Parent()).First(context.Background())
	if err != nil {
		t.Errorf("Agenda item returned unexpected error: %v", err)
	}

	if res.ContentObjectID != "topic/1" {
		t.Errorf("res.ContentObjectID = %s, expected topic/1", res.ContentObjectID)
	}

	if _, isSet := res.Parent.Value(); isSet {
		t.Errorf("parent should be empty")
	}
}
