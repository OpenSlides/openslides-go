package perm_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/datastore/dsmock"
	"github.com/OpenSlides/openslides-go/restrict/perm"
)

func TestHasSuperAdmin(t *testing.T) {
	ctx := context.Background()
	ds := dsmock.Stub(dsmock.YAMLData(`---
		user/1/organization_management_level: superadmin

		meeting/3/id: 3
	`))

	p, err := perm.New(ctx, dsfetch.New(ds), 1, 3)
	if err != nil {
		t.Fatalf("perm.New(): %v", err)
	}

	if !p.Has(perm.AgendaItemCanSee) {
		t.Errorf("p.Has(perm.AgendaItemCanSee) returned false, expected true for any perm")
	}
}
