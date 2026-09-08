// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package discovery

// Subjects are the logical event names, not the wire names: the bus owns the
// prefix it puts in front of them.
const (
	SubjectScanCompleted = "discovery.scan.completed"
	SubjectScopeRetired  = "discovery.scope.retired"
)
