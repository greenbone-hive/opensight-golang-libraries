// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package events

import (
	"errors"
	"fmt"
)

// CoverageStatus is the completeness verdict of one scan partition (a target
// scope, or a collector×region within it). Only CoverageComplete authorizes
// authoritative removals; partial/failed coverage can only upsert.
type CoverageStatus string

// The coverage verdicts.
const (
	// CoverageComplete: every selected collector finished all pages and nested
	// reads for the partition. The partition's absence set is authoritative.
	CoverageComplete CoverageStatus = "complete"
	// CoveragePartial: at least one collector failed or was interrupted; the
	// produced set is a lower bound and asserts nothing about absence.
	CoveragePartial CoverageStatus = "partial"
	// CoverageFailed: the partition produced no usable state (auth failure,
	// throttling, cancellation, missing role, disabled API).
	CoverageFailed CoverageStatus = "failed"
)

// LifecycleReason says WHY a resource, source claim or target scope left a
// consumer's view. Only ReasonResourceDeleted asserts the resource is gone at
// the cloud provider; every other reason is a control-plane or authorization
// change and must never be treated as provider deletion.
type LifecycleReason string

// The lifecycle reasons.
const (
	// ReasonResourceDeleted: a COMPLETE scan of the owning partition no longer
	// observed the resource. The only reason that asserts provider deletion.
	ReasonResourceDeleted LifecycleReason = "resource_deleted"
	// ReasonScopeExcluded: the operator removed the target/region/collector from
	// the connection's selection. Claims retire; the cloud resource may live on.
	ReasonScopeExcluded LifecycleReason = "scope_excluded"
	// ReasonConnectionDeleted: the connection was deleted; all its claims retire.
	ReasonConnectionDeleted LifecycleReason = "connection_deleted"
	// ReasonAuthorizationLost: the producer can no longer read the scope. Claims
	// are kept; coverage turns unknown/stale. Never retires a claim by itself.
	ReasonAuthorizationLost LifecycleReason = "authorization_lost"
	// ReasonTargetMoved: hierarchy reconciliation moved the target out of the
	// selected subtree; treated like an exclusion, not a deletion.
	ReasonTargetMoved LifecycleReason = "target_moved"
	// ReasonTargetClosed: the provider reports the account/subscription/project
	// as closed or suspended.
	ReasonTargetClosed LifecycleReason = "target_closed"
)

// RetiresClaim reports whether the reason retires a source claim. Authorization
// loss keeps the claim (coverage goes stale instead), and resource deletion is
// carried per-resource by snapshot diffs, not by a scope retirement.
func (r LifecycleReason) RetiresClaim() bool {
	switch r {
	case ReasonScopeExcluded, ReasonConnectionDeleted, ReasonTargetMoved, ReasonTargetClosed:
		return true
	case ReasonResourceDeleted, ReasonAuthorizationLost:
		return false
	}

	return false
}

// IsProviderDeletion reports whether the reason asserts the resource no longer
// exists at the cloud provider.
func (r LifecycleReason) IsProviderDeletion() bool { return r == ReasonResourceDeleted }

// Provenance identifies exactly which observation of the cloud produced an
// event: which connection (and its config revision), target scope
// (account/subscription/project), logical run and per-target scope run, and
// provider. It is embedded in every discovery-sourced event and travels intact
// through Discovery -> Asset Management -> Exposure.
//
// ConnectionID is provenance, NOT canonical identity: canonical resource/asset/
// graph identity is (provider, canonicalResourceId). Two connections observing
// one resource yield two source claims and one asset.
type Provenance struct {
	ConnectionID       string `json:"connection_id"`
	ConnectionRevision int64  `json:"connection_revision"`
	TargetScopeID      string `json:"target_scope_id"`
	RunID              string `json:"run_id"`
	ScopeRunID         string `json:"scope_run_id"`
	Provider           string `json:"provider"`
}

// PartitionKey is the event-stream identity of one target partition. Events for
// one partition are ordered by a per-partition monotonic Meta.Version; events
// for different partitions are independent and must never gate each other.
// Producers of partitioned events MUST set Meta.EntityID to this key.
func (p *Provenance) PartitionKey() string {
	return p.ConnectionID + "|" + p.TargetScopeID
}

// validate checks the mandatory provenance fields.
func (p *Provenance) validate() error {
	switch {
	case p.ConnectionID == "":
		return errors.New("events: provenance missing connection_id")
	case p.TargetScopeID == "":
		return errors.New("events: provenance missing target_scope_id")
	case p.Provider == "":
		return errors.New("events: provenance missing provider")
	}

	return nil
}

// CollectorCoverage is one collector×region completeness verdict inside a scope
// run. RegionOrGlobal is the region name or "global" for global collectors.
type CollectorCoverage struct {
	Collector      string         `json:"collector"`
	RegionOrGlobal string         `json:"region_or_global"`
	Status         CoverageStatus `json:"status"`
	Error          string         `json:"error,omitempty"`
}

// Validate enforces the snapshot contract: complete provenance, the partition
// ordering key, and per-resource identity. The absence safety rule (only a
// complete snapshot asserts deletions) cannot be validated here because absence
// is implicit; consumers MUST gate their reap on Coverage == complete.
func (e *DiscoverySnapshotCompleted) Validate() error {
	if err := e.validate(); err != nil {
		return err
	}
	if e.RunID == "" || e.ScopeRunID == "" {
		return errors.New("events: discovery snapshot missing run/scope-run id")
	}
	if e.EntityID != e.PartitionKey() {
		return fmt.Errorf("events: discovery snapshot entity_id %q must be the partition key %q", e.EntityID, e.PartitionKey())
	}
	if e.Version <= 0 {
		return errors.New("events: discovery snapshot needs a positive per-partition version")
	}
	for i := range e.Resources {
		if e.Resources[i].Type == "" || e.Resources[i].ProviderResourceID == "" {
			return fmt.Errorf("events: discovery snapshot resource %d missing type or provider_resource_id", i)
		}
	}

	return nil
}

// Validate enforces the retirement contract: complete provenance, partition
// ordering key, and a reason that actually retires claims. resource_deleted
// travels per-resource in snapshot diffs and authorization_lost never retires.
func (e *DiscoveryScopeRetired) Validate() error {
	if err := e.validate(); err != nil {
		return err
	}
	if !e.Reason.RetiresClaim() {
		return fmt.Errorf("events: scope retirement reason %q does not retire claims", e.Reason)
	}
	if e.EntityID != e.PartitionKey() {
		return fmt.Errorf("events: scope retirement entity_id %q must be the partition key %q", e.EntityID, e.PartitionKey())
	}

	return nil
}
