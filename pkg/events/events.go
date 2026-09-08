// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package events

import "encoding/json"

// This file holds the typed event payloads of the discovery -> asset service
// contract. Each event embeds Meta and Provenance; the snapshot carries its
// full resource set inline so the consumer can reconcile on its own.

// DiscoverySnapshotCompleted is emitted by discovery after every scope run and
// pushed to the assets service with the partition's COMPLETE live resource set
// inline. Discovery computes no diff and keeps no resource state: the consumer
// owns reconciliation, diffing the snapshot against its own previous state.
//
// Absence semantics: a resource missing from Resources is a deletion assertion
// ONLY when Coverage is complete. A partial/failed snapshot says nothing about
// absence (a failed collector's resources are simply missing), so consumers
// MUST apply it upsert-only and never reap on it. Scope exclusion, connection
// deletion and target moves/closures are NOT absences; they travel as
// DiscoveryScopeRetired with their own reason.
//
// Ordering: EntityID is Provenance.PartitionKey() and Version is monotonic per
// partition (the scope-run sequence, never a timestamp), so snapshots of
// different targets are ordered independently and may arrive in any order
// without gating each other.
type DiscoverySnapshotCompleted struct {
	Meta
	Provenance
	Account       string              `json:"account"`
	TriggerSource string              `json:"trigger_source"`
	Coverage      CoverageStatus      `json:"coverage"`
	Collectors    []CollectorCoverage `json:"collectors,omitempty"`
	Resources     []SnapshotResource  `json:"resources"`
}

// SnapshotResource is one live resource in a discovery snapshot: its identity
// within the partition plus the marshaled resource body (name, resourceName,
// assetType, computed, tags, identifiers).
type SnapshotResource struct {
	Type               string          `json:"type"`
	ProviderResourceID string          `json:"provider_resource_id"`
	Body               json.RawMessage `json:"body"`
}

// DiscoveryScopeRetired is emitted by discovery when a target partition leaves
// a connection's coverage for a control-plane reason: the operator excluded it,
// the connection was deleted, or hierarchy reconciliation found the target
// moved or closed. Consumers retire the partition's source claims with the
// carried reason; it never asserts provider deletion (see LifecycleReason).
// EntityID is Provenance.PartitionKey(); Version continues the partition's
// monotonic sequence.
type DiscoveryScopeRetired struct {
	Meta
	Provenance
	Reason LifecycleReason `json:"reason"`
}
