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
