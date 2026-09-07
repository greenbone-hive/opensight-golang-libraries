// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package provider

type Provider string

const (
	AWS          Provider = "aws"
	Azure        Provider = "azure"
	GCP          Provider = "gcp"
	GB_APPLIANCE Provider = "greenbone_appliance"
)

var All = []Provider{AWS, Azure, GCP}

func Valid(p Provider) bool {
	for _, v := range All {
		if p == v {
			return true
		}
	}

	return false
}
