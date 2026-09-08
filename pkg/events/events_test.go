// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package events

import (
	"encoding/json"
	"testing"
	"time"
)

func TestMetaKey(t *testing.T) {
	m := Meta{EntityID: "asset-1", Version: 7}
	if got, want := m.Key(), "asset-1@7"; got != want {
		t.Errorf("Key() = %q, want %q", got, want)
	}
}

// TestEventJSONRoundTrip pins the wire shape of an event: it survives a
// marshal/unmarshal round trip, and the embedded Meta and Provenance fields are
// PROMOTED to the top level rather than nested under their type name. A
// consumer reads entity_id and target_scope_id as top-level keys; nesting them
// would break every consumer without breaking any producer.
func TestEventJSONRoundTrip(t *testing.T) {
	in := DiscoveryScopeRetired{
		Meta: Meta{
			ID:       "evt-1",
			Type:     SubjectDiscoveryScopeRetired,
			Source:   "discovery",
			Time:     time.Unix(1700000000, 0).UTC(),
			EntityID: "asset-1",
			Version:  3,
		},
		Provenance: Provenance{
			ConnectionID: "conn-1", ConnectionRevision: 2,
			TargetScopeID: "111111111111", RunID: "10", ScopeRunID: "77",
			Provider: "aws",
		},
		Reason: ReasonScopeExcluded,
	}

	data, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var out DiscoveryScopeRetired
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out.EntityID != "asset-1" || out.Version != 3 {
		t.Errorf("meta not round-tripped: %+v", out.Meta)
	}
	if out.Reason != ReasonScopeExcluded || out.Provenance != in.Provenance {
		t.Errorf("payload not round-tripped: %+v", out)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatal(err)
	}
	for _, nested := range []string{"Meta", "Provenance"} {
		if _, ok := raw[nested]; ok {
			t.Errorf("%s should be embedded/promoted, not nested under its own key", nested)
		}
	}
	if raw["entity_id"] != "asset-1" {
		t.Errorf("entity_id should be a top-level field, got %v", raw["entity_id"])
	}
}
