// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package events

import (
	"encoding/json"
	"testing"
	"time"
)

// Meta.Key forms the (entity_id, version) idempotency key.
func TestMetaKey(t *testing.T) {
	m := Meta{EntityID: "asset-1", Version: 7}
	if got, want := m.Key(), "asset-1@7"; got != want {
		t.Errorf("Key() = %q, want %q", got, want)
	}
}

// An event round-trips through JSON with Meta promoted to top-level fields.
func TestEventJSONRoundTrip(t *testing.T) {
	in := AssetEvent{
		Meta: Meta{
			ID:       "evt-1",
			Type:     SubjectAssetAdded,
			Source:   "assets",
			Time:     time.Unix(1700000000, 0).UTC(),
			EntityID: "asset-1",
			Version:  3,
		},
		Asset: AssetSnapshot{
			AssetID:   "asset-1",
			AssetType: "VirtualMachine",
			Name:      "web-01",
			Status:    "active",
		},
	}

	data, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var out AssetEvent
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out.EntityID != "asset-1" || out.Version != 3 {
		t.Errorf("meta not round-tripped: %+v", out.Meta)
	}
	if out.Asset.AssetType != "VirtualMachine" || out.Asset.Name != "web-01" {
		t.Errorf("payload not round-tripped: %+v", out)
	}

	// Meta fields are promoted to the top level (not nested).
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatal(err)
	}
	if _, nested := raw["Meta"]; nested {
		t.Error("Meta should be embedded/promoted, not nested under a Meta key")
	}
	if raw["entity_id"] != "asset-1" {
		t.Errorf("entity_id should be a top-level field, got %v", raw["entity_id"])
	}
}
