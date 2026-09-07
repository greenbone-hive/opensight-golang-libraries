// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package identifier is the canonical asset-identity claim taxonomy.
// Shared so discovery emits, and the identity / reconciliation engine and the
// sensor consume, ONE typed value set, instead of each service redefining the
// enum. The values are the exact snake_case strings stored and matched on end to
// end: no per-service translation. How strongly a claim weighs in a match is a
// consumer's policy, not part of the taxonomy, so it does not live here.
// Mirrors the provider package pattern.
package identifier

// Type is a canonical asset-identity claim type.
type Type string

// The canonical identity claim types, stored and matched as these exact strings.
const (
	Hostname           Type = "hostname"
	FQDN               Type = "fqdn"
	IPv4               Type = "ipv4"
	IPv6               Type = "ipv6"
	MACAddress         Type = "mac"
	BiosUUID           Type = "bios_uuid"
	SerialNumber       Type = "serial"
	ProviderResourceID Type = "provider_resource_id"
)

// MaxProviderResourceIDLen bounds a provider_resource_id value in bytes. It
// lives here because it is a CONTRACT BETWEEN TWO SERVICES: discovery's
// producer-side grammar gate accepts ids up to this length, and the
// reconciliation engine must accept exactly as many. A consumer bound that is
// lower than the producer's does not reject loudly, it drops the observation
// after the producer already blessed it. Keeping the number in one place is what
// makes that drift impossible; it had already happened once with 1024 on the
// producer and 500 on the consumer.
//
// THE BOUND IS PHYSICAL, NOT A POLICY. This column sits inside a composite
// btree UNIQUE index in both stores, and Postgres refuses to index an entry over
// 2,704 bytes. Measured on the dev instance (PG 17) against the real discovery
// constraint, with incompressible values:
//
//	2,048  ok
//	2,600  ok
//	2,700  ERROR: index row size 2736 exceeds btree version 4 maximum 2704
//	5,000  ERROR: index row size 5040 exceeds btree version 4 maximum 2704
//
// So the choice is never "bound or no bound", only "our bound or Postgres'".
// Ours rejects ONE resource, names it, and lets the rest of the scan land.
// Postgres' fires inside the apply transaction and rolls back the whole target,
// on every scan, with no self-heal, reporting an index-row-size error that no
// operator can act on.
//
// 2,048 sits above every id discovery emits (the longest is a GCP BigQuery
// dataset CAI name at 1,099 bytes: a 1,024-character dataset id under a 75-byte
// prefix) and far enough below the wall that a longer resource_type, an extra
// index column, or a page-size change cannot reach it. It is a ceiling to stay
// under, not a catalogue of forms to keep re-deriving.
const MaxProviderResourceIDLen = 2048

// All is every claim-type constant, for iteration and parity checks.
var All = []Type{
	Hostname,
	FQDN,
	IPv4,
	IPv6,
	MACAddress,
	BiosUUID,
	SerialNumber,
	ProviderResourceID,
}

// Valid reports whether t is a known claim type.
func Valid(t Type) bool {
	for _, v := range All {
		if t == v {
			return true
		}
	}

	return false
}
