// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package computed

import "testing"

// TestIsCanonicalMatchesAll proves IsCanonical accepts exactly the members of All
// and rejects raw / unknown keys, so the discovery bucket split routes canonical
// fields to Computed and everything else to Properties.
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

// TestAllHasNoDuplicates guards against a copy-paste slip in All.
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

// TestPhase8ReachabilityKeysRegistered pins the phase-8 data-plane gate fields as
// canonical. If a producer emits one of these before it is in All, the discovery
// bucket split silently routes it to Properties and phase-10 reachability cannot
// see a public endpoint's allow-list (a false-negative exposure). This test fails
// the build if one is dropped from All.
func TestPhase8ReachabilityKeysRegistered(t *testing.T) {
	required := []string{
		PublicNetworkAccess,
		AuthorizedCidrs,
		APIServerPublic,
		APIServerAuthorizedCidrs,
		PrivateNodes,
	}
	for _, k := range required {
		if !IsCanonical(k) {
			t.Errorf("phase-8 reachability key %q is not canonical; register it in All before producers emit it", k)
		}
	}
}

// TestDerivedGraphFieldsAreNotCanonical proves the phase-10 engine's derived
// outputs are NOT in the producer contract, so a client that emits one by mistake
// is routed to Properties instead of masquerading as a canonical field.
func TestDerivedGraphFieldsAreNotCanonical(t *testing.T) {
	for _, k := range []string{"ingressAllowSet", "effectiveInbound"} {
		if IsCanonical(k) {
			t.Errorf("derived graph-only field %q must not be canonical (it is written by exposure, never emitted by a producer)", k)
		}
	}
}
