// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package assetcategory

import "testing"

func TestAllHasNoDuplicates(t *testing.T) {
	seen := map[Category]struct{}{}
	for _, c := range All {
		if c == "" {
			t.Error("All contains an empty category")
		}
		if _, dup := seen[c]; dup {
			t.Errorf("All lists %q more than once", c)
		}
		seen[c] = struct{}{}
	}
}
