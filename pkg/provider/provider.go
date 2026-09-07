// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package provider is the canonical taxonomy of the systems assets are
// discovered from: what an adapter connects to, and where a resource lives.
// Shared so discovery, assets, and every other service key off one value set and
// one validation, instead of each redefining the enum.
//
// The set today is the three public clouds. It is not limited to them by design:
// on-premises sources are part of the same discovery surface, and each becomes a
// constant here alongside the clouds rather than a separate parallel enum.
// Mirrors the assets taxonomy packages (assettype / assetcategory / identifier /
// computed).
package provider

// Provider is a canonical source of discovered assets.
type Provider string

// The canonical providers, lowercase end to end; casing is a UI concern.
// Public clouds today; on-premises sources join this block as they land.
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
