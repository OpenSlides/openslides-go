package dsmodels_test

import (
	"testing"

	"github.com/OpenSlides/openslides-go/datastore/dsmock"
	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

func TestRequestSingleModel(t *testing.T) {
	ds := dsmodels.New(dsmock.Stub(dsmock.YAMLData(`---
	topic/1:
		sequential_number: 1
		title: foo
		meeting_id: 1
		agenda_item_id: 1
		list_of_speakers_id: 1
	`)))

	_, err := ds.Topic(1).First(t.Context())
	if err != nil {
		t.Errorf("Topic 1 returned unexpected error: %v", err)
	}
}

func TestRequestSingleModelWithRequiredRelation(t *testing.T) {
	ds := dsmodels.New(dsmock.Stub(dsmock.YAMLData(`---
	topic/1:
		sequential_number: 1
		title: foo
		meeting_id: 1
		agenda_item_id: 1
		list_of_speakers_id: 1
	agenda_item/1/content_object_id: topic/1
	agenda_item/1/meeting_id: 1
	`)))

	tQ := ds.Topic(1)
	_, err := tQ.Preload(tQ.AgendaItem()).First(t.Context())
	if err != nil {
		t.Errorf("Topic 1 with agenda item returned unexpected error: %v", err)
	}
}

func TestRequestSingleModelWithExistingMaybeRelation(t *testing.T) {
	ds := dsmodels.New(dsmock.Stub(dsmock.YAMLData(`---
	agenda_item/1:
		content_object_id: topic/1
		meeting_id: 1
		parent_id: 2
	agenda_item/2/content_object_id: topic/2
	agenda_item/2/meeting_id: 1
	`)))

	q := ds.AgendaItem(1)
	res, err := q.Preload(q.Parent()).First(t.Context())
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

func TestRequestModelsWithExistingMaybeRelation(t *testing.T) {
	ds := dsmodels.New(dsmock.Stub(dsmock.YAMLData(`---
	agenda_item/1:
		content_object_id: topic/1
		meeting_id: 1
		parent_id: 2
	agenda_item/2/content_object_id: topic/2
	agenda_item/2/meeting_id: 1
	`)))

	q := ds.AgendaItem(1)
	resAll, err := q.Preload(q.Parent()).Get(t.Context())
	if err != nil {
		t.Errorf("Agenda item returned unexpected error: %v", err)
	}

	if len(resAll) != 1 {
		t.Errorf("len(resAll) = %d, expected 1", len(resAll))
		return
	}

	res := resAll[0]
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
	res, err := q.Preload(q.Parent()).First(t.Context())
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

func TestRequestSingleModelWithEmptyListRelation(t *testing.T) {
	ds := dsmodels.New(dsmock.Stub(dsmock.YAMLData(`---
	agenda_item/1/content_object_id: topic/1
	agenda_item/1/meeting_id: 1
	`)))

	q := ds.AgendaItem(1)
	res, err := q.Preload(q.ChildList()).First(t.Context())
	if err != nil {
		t.Errorf("Agenda item returned unexpected error: %v", err)
	}

	if res.ContentObjectID != "topic/1" {
		t.Errorf("res.ContentObjectID = %s, expected topic/1", res.ContentObjectID)
	}
}

func TestRequestSingleModelWithListRelation(t *testing.T) {
	ds := dsmodels.New(dsmock.Stub(dsmock.YAMLData(`---
	agenda_item/1:
		content_object_id: topic/1
		meeting_id: 1
		child_ids: [2, 3]
	agenda_item/2:
		content_object_id: topic/2
		meeting_id: 1
		parent_id: 1
	agenda_item/3:
		content_object_id: topic/3
		meeting_id: 1
		parent_id: 1
	`)))

	q := ds.AgendaItem(1)
	res, err := q.Preload(q.ChildList()).First(t.Context())
	if err != nil {
		t.Errorf("Agenda item returned unexpected error: %v", err)
	}

	if len(res.ChildList) != 2 {
		t.Errorf("len(res.ChildList) = %d, expected 2", len(res.ChildList))
	}

	if res.ChildList[0].ContentObjectID != "topic/2" {
		t.Errorf("res.ChildList[0].ContentObjectID = %s, expected topic/2", res.ChildList[0].ContentObjectID)
	}

	if res.ChildList[1].ContentObjectID != "topic/3" {
		t.Errorf("res.ChildList[1].ContentObjectID = %s, expected topic/3", res.ChildList[1].ContentObjectID)
	}
}

func TestRequestSingleModelWithNestedRelation(t *testing.T) {
	ds := dsmodels.New(dsmock.Stub(dsmock.YAMLData(`---
	topic/1:
		sequential_number: 1
		title: foo
		meeting_id: 1
		agenda_item_id: 1
		list_of_speakers_id: 1
	agenda_item/1:
		content_object_id: topic/1
		meeting_id: 1
		parent_id: 2
	agenda_item/2:
		content_object_id: topic/2
		meeting_id: 1
		parent_id: 3
	agenda_item/3/content_object_id: topic/3
	agenda_item/3/meeting_id: 1
	`)))

	tQ := ds.Topic(1)
	res, err := tQ.Preload(tQ.AgendaItem().Parent().Parent()).First(t.Context())
	if err != nil {
		t.Errorf("Topic 1 with agenda item returned unexpected error: %v", err)
	}

	parent, isSet := res.AgendaItem.Parent.Value()
	if !isSet {
		t.Errorf("res.AgendaItem.Parent is empty")
	}

	parent, isSet = parent.Parent.Value()
	if !isSet {
		t.Errorf("res.AgendaItem.Parent.Parent is empty")
	}

	if parent.ContentObjectID != "topic/3" {
		t.Errorf("parent.ContentObjectID = %s, expected topic/3", parent.ContentObjectID)
	}
}

func TestRequestSinglePreloadMulti(t *testing.T) {
	ds := dsmodels.New(dsmock.Stub(dsmock.YAMLData(`---
	topic/1:
		sequential_number: 1
		title: foo
		meeting_id: 1
		agenda_item_id: 1
		list_of_speakers_id: 1
	agenda_item/1/content_object_id: topic/1
	agenda_item/1/meeting_id: 1
	list_of_speakers/1:
		sequential_number: 1
		content_object_id: topic/1
		meeting_id: 1
	`)))

	tQ := ds.Topic(1)
	res, err := tQ.Preload(tQ.AgendaItem()).Preload(tQ.ListOfSpeakers()).First(t.Context())
	if err != nil {
		t.Errorf("Topic 1 with agenda item returned unexpected error: %v", err)
	}

	if res.AgendaItem.ContentObjectID != "topic/1" {
		t.Errorf("res.AgendaItem.ContentObjectID = %s, expected topic/1", res.AgendaItem.ContentObjectID)
	}

	if res.ListOfSpeakers.ContentObjectID != "topic/1" {
		t.Errorf("res.ListOfSpeakers.ContentObjectID = %s, expected topic/1", res.ListOfSpeakers.ContentObjectID)
	}
}
