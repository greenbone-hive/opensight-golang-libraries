// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package events

import "encoding/json"

// This file holds the typed event payloads. Each event embeds Meta. Payloads
// carry what consumers need to reconcile on their own; the discovery snapshot
// carries its full resource set inline. More events are added as services are
// built.

// Severity is a finding/vulnerability severity level.
type Severity string

// Severity levels, highest to lowest.
const (
	SeverityCritical Severity = "critical"
	SeverityHigh     Severity = "high"
	SeverityMedium   Severity = "medium"
	SeverityLow      Severity = "low"
	SeverityInfo     Severity = "info"
)

// AuditRecorded is published by every service for each user operation; consumed
// only by the audit service. Fire-and-forget — it never blocks the operation.
type AuditRecorded struct {
	Meta
	ActorID    string         `json:"actor_id"`
	Action     string         `json:"action"`
	EntityType string         `json:"entity_type"`
	Metadata   map[string]any `json:"metadata,omitempty"`
}

// AssetSnapshot is the asset content an asset lifecycle event carries inline:
// enough for the topology graph-build (node identity, canonical labels, and the
// computed FK fields) without a pull. Computed holds the canonical topology
// fields (assets/computed keys); it is empty for assets with no cloud envelope
// (e.g. manually registered ones).
type AssetSnapshot struct {
	AssetID       string         `json:"asset_id"`
	Name          string         `json:"name"`
	AssetType     string         `json:"asset_type"`
	AssetCategory string         `json:"asset_category"`
	Provider      string         `json:"provider,omitempty"`
	Status        string         `json:"status"`
	Computed      map[string]any `json:"computed,omitempty"`
}

// AssetEvent is the asset lifecycle payload, published by assets under
// asset.added / asset.changed / asset.removed. A removed event carries the last
// known snapshot so consumers can tear down derived state (e.g. the topology
// node) without a lookup. (entity_id = asset id; version orders per asset.)
//
// Graph identity downstream is (provider, resourceId). Scope is the discovery
// observation that caused this change (nil for manual/agent-sourced changes).
// Reason is set on asset.removed to say why (provider deletion vs claim
// retirement).
type AssetEvent struct {
	Meta
	Asset  AssetSnapshot   `json:"asset"`
	Reason LifecycleReason `json:"reason,omitempty"`
	Scope  *Provenance     `json:"scope,omitempty"`
}

// AssetBatchCompleted is published by assets after it finishes applying one
// discovery scope-run snapshot (all per-asset events of that batch precede it).
// Exposure uses it to run the structural rule pass once per completed batch at
// a real producer scan boundary instead of per asset event. It carries the
// originating scope-run provenance and coverage verdict verbatim. EntityID is
// Provenance.PartitionKey(); Version mirrors the consumed snapshot's version.
type AssetBatchCompleted struct {
	Meta
	Provenance
	Coverage   CoverageStatus      `json:"coverage"`
	Collectors []CollectorCoverage `json:"collectors,omitempty"`
	Counts     ChangeCounts        `json:"counts"`
}

// VulnerabilityEvent is the vulnerability finding lifecycle payload, published
// under vulnerability.reported / vulnerability.changed / vulnerability.closed.
// (entity_id = finding id.)
type VulnerabilityEvent struct {
	Meta
	FindingID string   `json:"finding_id"`
	AssetID   string   `json:"asset_id"`
	CVEID     string   `json:"cve_id,omitempty"`
	Title     string   `json:"title,omitempty"`
	Severity  Severity `json:"severity,omitempty"`
	Status    string   `json:"status,omitempty"`
}

// MisconfigurationEvent is the misconfiguration finding lifecycle payload,
// published under misconfiguration.reported / misconfiguration.changed /
// misconfiguration.closed. (entity_id = finding id.)
type MisconfigurationEvent struct {
	Meta
	FindingID string   `json:"finding_id"`
	AssetID   string   `json:"asset_id"`
	RuleID    string   `json:"rule_id,omitempty"`
	Title     string   `json:"title,omitempty"`
	Severity  Severity `json:"severity,omitempty"`
	Status    string   `json:"status,omitempty"`
}

// InventoryEvent is the inventory lifecycle payload, published under
// inventory.added / inventory.changed / inventory.closed.
// (entity_id = asset id.)
type InventoryEvent struct {
	Meta
	AssetID string          `json:"asset_id"`
	Dataset string          `json:"dataset,omitempty"`
	Counts  InventoryCounts `json:"counts,omitempty"`
}

// InventorySnapshotCreated is published by the inventory service when a sensor
// submission changes an asset's inventory (diff-based: an unchanged submission
// publishes nothing). It carries a small identity subset inline so the assets
// service can reconcile without a pull; the bulk record list stays in ClickHouse
// and is pulled by RunID via the claim-check endpoint (ADR-013).
// (entity_id = asset id; version orders snapshots per asset.)
type InventorySnapshotCreated struct {
	Meta
	AssetID         string            `json:"asset_id"`
	AgentID         string            `json:"agent_id"`
	RunID           string            `json:"run_id"`
	SyncType        string            `json:"sync_type"` // "full" | "delta"
	ChangedDatasets []string          `json:"changed_datasets"`
	Counts          InventoryCounts   `json:"counts"`
	Identity        InventoryIdentity `json:"identity"`
}

// InventoryCounts summarises the magnitude of an inventory snapshot's changes.
type InventoryCounts struct {
	Total   int `json:"total"`
	Added   int `json:"added"`
	Updated int `json:"updated"`
	Removed int `json:"removed"`
}

// InventoryIdentity is the small identity slice inlined into the snapshot event.
type InventoryIdentity struct {
	Hostname string   `json:"hostname,omitempty"`
	FQDN     string   `json:"fqdn,omitempty"`
	Serial   string   `json:"serial,omitempty"`
	MAC      string   `json:"mac,omitempty"`
	IPs      []string `json:"ips,omitempty"`
}

// FindingCreated is published by posture or vulnerability when a finding is
// raised. Carries the display fields needed to build a ticket (claim-check is
// unnecessary for these small fields).
type FindingCreated struct {
	Meta
	FindingID string   `json:"finding_id"`
	AssetID   string   `json:"asset_id"`
	AssetName string   `json:"asset_name"`
	RuleID    string   `json:"rule_id,omitempty"`
	CVEID     string   `json:"cve_id,omitempty"`
	Title     string   `json:"title"`
	Severity  Severity `json:"severity"`
}

// FindingResolved is published when a finding is resolved or risk-accepted.
type FindingResolved struct {
	Meta
	FindingID string `json:"finding_id"`
	AssetID   string `json:"asset_id"`
}

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

// ChangeCounts is the roll-up summary of an asset batch apply, computed by the
// assets service from its own reconciliation result.
type ChangeCounts struct {
	Added   int `json:"added"`
	Changed int `json:"changed"`
	Removed int `json:"removed"`
}
