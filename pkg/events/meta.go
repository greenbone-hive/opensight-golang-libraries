// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package events holds the envelope every event carries, whatever contract it
// belongs to. The payloads live in the subpackages, one per producer-consumer
// contract. The structs are format-neutral: the wire codec is the transport's
// concern, so the serialization format can change without touching a schema.
package events

import (
	"strconv"
	"time"
)

// Meta is the common metadata embedded in every event. EntityID and Version
// form the idempotency/ordering key consumers use to drop duplicate or
// out-of-order events.
//
// Ordering is strictly per EntityID stream: streams never gate each other.
// Partitioned events (a discovery scan or retirement, say) use
// Provenance.PartitionKey() as EntityID with a per-partition monotonic
// sequence as Version, never a timestamp, which cannot order parallel
// producers of one partition.
type Meta struct {
	ID       string    `json:"id"`        // unique event id
	Type     string    `json:"type"`      // subject, e.g. "discovery.scan.completed"
	Source   string    `json:"source"`    // publishing service
	Time     time.Time `json:"time"`      // occurred-at (UTC)
	EntityID string    `json:"entity_id"` // idempotency/ordering key
	Version  int64     `json:"version"`   // idempotency/ordering key
}

// Key returns the idempotency/ordering key for the event. Consumers store the
// highest processed Version per EntityID and ignore anything not newer.
func (m *Meta) Key() string {
	return m.EntityID + "@" + strconv.FormatInt(m.Version, 10)
}
