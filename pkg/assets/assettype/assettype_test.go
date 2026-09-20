// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package assettype

import (
	"testing"

	"github.com/greenbone-hive/opensight-golang-libraries/pkg/assets/assetcategory"
)

func TestAllHasNoDuplicates(t *testing.T) {
	seen := map[Type]struct{}{}
	for _, ty := range All {
		if ty == "" {
			t.Error("All contains an empty type")
		}
		if _, dup := seen[ty]; dup {
			t.Errorf("All lists %q more than once", ty)
		}
		seen[ty] = struct{}{}
	}
}

func TestEveryTypeHasACategory(t *testing.T) {
	for _, ty := range All {
		if !IsKnown(ty) {
			t.Errorf("type %q is missing from categoryByType", ty)
		}
		if CategoryOf(ty) == "" {
			t.Errorf("type %q has no category", ty)
		}
	}
	if len(categoryByType) != len(All) {
		t.Errorf("categoryByType has %d entries, All has %d (stale mapping?)", len(categoryByType), len(All))
	}
}

func TestATypeOutsideTheCatalogHasNoCategory(t *testing.T) {
	for _, ty := range []Type{"NotARealType", "Unknown", ""} {
		if IsKnown(ty) {
			t.Errorf("IsKnown(%q) should be false", ty)
		}
		if got := CategoryOf(ty); got != "" {
			t.Errorf("CategoryOf(%q) = %q, want no category", ty, got)
		}
	}
}

func TestEveryMachineIsAComputeTypeInTheCatalog(t *testing.T) {
	seen := map[Type]struct{}{}
	for _, m := range Machines {
		if _, dup := seen[m]; dup {
			t.Errorf("Machines lists %q more than once", m)
		}
		seen[m] = struct{}{}
		if CategoryOf(m) != assetcategory.Compute {
			t.Errorf("machine %q is in category %q, want Compute", m, CategoryOf(m))
		}
		if !IsMachine(m) {
			t.Errorf("IsMachine(%q) should be true", m)
		}
	}
	for _, ty := range []Type{AppService, BatchJob, LaunchTemplate, ObjectStorage, "NotARealType"} {
		if IsMachine(ty) {
			t.Errorf("IsMachine(%q) should be false", ty)
		}
	}
}
