// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package assettype

import (
	"testing"

	"github.com/greenbone/opensight-golang-libraries/pkg/assets/assetcategory"
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
		if ty != Unknown && CategoryOf(ty) == assetcategory.Unknown {
			t.Errorf("type %q maps to the Unknown category", ty)
		}
	}
	if len(categoryByType) != len(All) {
		t.Errorf("categoryByType has %d entries, All has %d (stale mapping?)", len(categoryByType), len(All))
	}
}

func TestUnknownTypeFallsThrough(t *testing.T) {
	if IsKnown("NotARealType") {
		t.Error("IsKnown should be false for a type not in the catalog")
	}
	if CategoryOf("NotARealType") != assetcategory.Unknown {
		t.Error("CategoryOf should return Unknown for a type not in the catalog")
	}
}
