package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-go/restrict/collection"
	"github.com/OpenSlides/openslides-go/restrict/perm"
)

func TestStructureLevelModeA(t *testing.T) {
	f := collection.StructureLevel{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		structure_level/1/meeting_id: 30
		`,
	)

	testCase(
		"list_of_speakers.can_see",
		t,
		f,
		true,
		`---
		structure_level/1/meeting_id: 30
		`,
		withPerms(30, perm.ListOfSpeakersCanSee),
	)

	testCase(
		"user.can_see",
		t,
		f,
		true,
		`---
		structure_level/1/meeting_id: 30
		`,
		withPerms(30, perm.UserCanSee),
	)
}
