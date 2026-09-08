// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package computed

import "testing"

func TestIsCanonicalMatchesAll(t *testing.T) {
	for _, k := range All {
		if !IsCanonical(k) {
			t.Errorf("IsCanonical(%q) = false, want true (it is in All)", k)
		}
	}
	for _, k := range []string{"", "displayName", "parent", "projectNumber", "vpcId", "ParentID"} {
		if IsCanonical(k) {
			t.Errorf("IsCanonical(%q) = true, want false (not a canonical key)", k)
		}
	}
}

func TestAllHasNoDuplicates(t *testing.T) {
	seen := map[string]struct{}{}
	for _, k := range All {
		if k == "" {
			t.Error("All contains an empty key")
		}
		if _, dup := seen[k]; dup {
			t.Errorf("All lists %q more than once", k)
		}
		seen[k] = struct{}{}
	}
}

func TestDataPlaneGateKeysRegistered(t *testing.T) {
	required := []string{
		PublicNetworkAccess,
		AuthorizedCidrs,
		APIServerPublic,
		APIServerAuthorizedCidrs,
		PrivateNodes,
	}
	for _, k := range required {
		if !IsCanonical(k) {
			t.Errorf("data-plane gate key %q is not canonical; register it in All before producers emit it", k)
		}
	}
}

func TestDerivedGraphFieldsAreNotCanonical(t *testing.T) {
	for _, k := range []string{"ingressAllowSet", "effectiveInbound"} {
		if IsCanonical(k) {
			t.Errorf("derived graph-only field %q must not be canonical (it is written by exposure, never emitted by a producer)", k)
		}
	}
}
