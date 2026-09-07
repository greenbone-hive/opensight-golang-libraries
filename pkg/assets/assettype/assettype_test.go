// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package assettype

import (
	"testing"

	"github.com/greenbone/opensight-golang-libraries/pkg/assets/assetcategory"
)

// TestAllHasNoDuplicates guards against a copy-paste slip in All.
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

// TestEveryTypeHasACategory proves the type->category mapping covers exactly the
// type set: every constant is IsKnown and lands on a real category. A missing
// entry would silently miscategorise an asset (CategoryOf falling through to
// Unknown); a stale entry means a deleted type lingering in the map.
func TestEveryTypeHasACategory(t *testing.T) {
	for _, ty := range All {
		if !IsKnown(ty) {
			t.Errorf("type %q is missing from categoryByType", ty)
		}
		if ty != Unknown && CategoryOf(ty) == assetcategory.Unknown {
			t.Errorf("type %q maps to the Unknown category", ty)
		}
	}
	if len(categoryByType) != len(All) {
		t.Errorf("categoryByType has %d entries, All has %d (stale mapping?)", len(categoryByType), len(All))
	}
}

// TestUnknownTypeFallsThrough proves a type outside the catalog is reported
// unknown rather than silently categorised.
func TestUnknownTypeFallsThrough(t *testing.T) {
	if IsKnown("NotARealType") {
		t.Error("IsKnown should be false for a type not in the catalog")
	}
	if CategoryOf("NotARealType") != assetcategory.Unknown {
		t.Error("CategoryOf should return Unknown for a type not in the catalog")
	}
}
