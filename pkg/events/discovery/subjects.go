// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package discovery

// Event subjects: the logical event types. They are not the wire names; the bus
// owns the prefix it puts in front of them.
const (
	SubjectSnapshotCompleted = "discovery.snapshot.completed"
	SubjectScopeRetired      = "discovery.scope.retired"
)
