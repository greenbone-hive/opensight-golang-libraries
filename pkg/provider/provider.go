// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package provider is the canonical cloud-provider taxonomy: the cloud an adapter
// connects to (and a resource belongs to). Shared so discovery, assets, and every
// other service key off one aws/azure/gcp value set and one validation, instead of
// each redefining the enum. Provider-neutral, mirroring the assets taxonomy
// packages (assettype / assetcategory / identifier / computed).
package provider

// Provider is a canonical cloud provider.
type Provider string

// The canonical providers, lowercase end to end; casing is a UI concern.
const (
	AWS   Provider = "aws"
	Azure Provider = "azure"
	GCP   Provider = "gcp"
)

// All is every provider constant, for iteration and parity checks.
var All = []Provider{AWS, Azure, GCP}

// Valid reports whether p is a known provider.
func Valid(p Provider) bool {
	for _, v := range All {
		if p == v {
			return true
		}
	}

	return false
}
