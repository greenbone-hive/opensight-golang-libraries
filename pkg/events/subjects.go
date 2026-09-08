// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package events

// Event subjects: the logical event types. They are not the wire names; the bus
// owns the prefix it puts in front of them.
const (
	SubjectDiscoverySnapshotCompleted = "discovery.snapshot.completed"
	SubjectDiscoveryScopeRetired      = "discovery.scope.retired"
)
