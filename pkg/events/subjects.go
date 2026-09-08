// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package events

// Event subjects: the logical event types. They are not the wire names; the bus
// owns the prefix it puts in front of them.
const (
	SubjectAuditRecorded = "audit.recorded"

	// Asset lifecycle (assets registry -> topology and other consumers).
	SubjectAssetAdded   = "asset.added"
	SubjectAssetChanged = "asset.changed"
	SubjectAssetRemoved = "asset.removed"

	// Vulnerability finding lifecycle (vulnerability service -> topology).
	SubjectVulnerabilityReported = "vulnerability.reported"
	SubjectVulnerabilityChanged  = "vulnerability.changed"
	SubjectVulnerabilityClosed   = "vulnerability.closed"

	// Misconfiguration finding lifecycle (posture service -> topology).
	SubjectMisconfigurationReported = "misconfiguration.reported"
	SubjectMisconfigurationChanged  = "misconfiguration.changed"
	SubjectMisconfigurationClosed   = "misconfiguration.closed"

	// Inventory lifecycle (inventory service -> topology).
	SubjectInventoryAdded   = "inventory.added"
	SubjectInventoryChanged = "inventory.changed"
	SubjectInventoryClosed  = "inventory.closed"

	SubjectSensorBound     = "sensor.bound"
	SubjectSensorHeartbeat = "sensor.heartbeat"
	SubjectSensorData      = "sensor.data"

	SubjectSensorDeployRequested = "sensor.deploy.requested"
	SubjectSensorDeployCompleted = "sensor.deploy.completed"

	SubjectInventorySnapshotCreated = "inventory.snapshot.created"

	SubjectFindingCreated  = "finding.created"
	SubjectFindingResolved = "finding.resolved"

	SubjectVulnerabilityCorrelated = "vulnerability.correlated"

	SubjectDiscoverySnapshotCompleted = "discovery.snapshot.completed"
	SubjectDiscoveryScopeRetired      = "discovery.scope.retired"
	SubjectAdapterChanged             = "adapter.changed"

	// Batch boundary (assets -> exposure): all per-asset events of one applied
	// discovery scope-run batch have been published.
	SubjectAssetBatchCompleted = "asset.batch.completed"

	SubjectFeedUpdated = "feed.updated"

	SubjectScanRequested = "scan.requested"

	SubjectExposureProjectionDirty    = "exposure.projection.dirty"
	SubjectExposureReprojectRequested = "exposure.reproject.requested"
	SubjectExposureAttackPathComputed = "exposure.attackpath.computed"
)
