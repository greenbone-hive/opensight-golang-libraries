// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package identifier

type Type string

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

func Valid(t Type) bool {
	for _, v := range All {
		if t == v {
			return true
		}
	}

	return false
}
