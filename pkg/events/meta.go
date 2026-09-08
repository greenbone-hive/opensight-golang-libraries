// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package events holds the envelope every event carries. The payloads live in
// the subpackages, one per producer-consumer contract.
package events

import (
	"strconv"
	"time"
)

// Meta is embedded in every event. EntityID and Version are the
// idempotency/ordering key: consumers keep the highest Version per EntityID and
// drop anything not newer. Ordering is per EntityID stream, and Version is a
// monotonic sequence, never a timestamp, which cannot order parallel producers
// of one stream.
type Meta struct {
	ID       string    `json:"id"`        // unique event id
	Type     string    `json:"type"`      // subject, e.g. "discovery.scan.completed"
	Source   string    `json:"source"`    // publishing service
	Time     time.Time `json:"time"`      // occurred-at (UTC)
	EntityID string    `json:"entity_id"` // idempotency/ordering key
	Version  int64     `json:"version"`   // idempotency/ordering key
}

func (m *Meta) Key() string {
	return m.EntityID + "@" + strconv.FormatInt(m.Version, 10)
}
