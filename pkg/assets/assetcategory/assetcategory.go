// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package assetcategory is the single source of the canonical asset categories:
// the high-level group every asset Type belongs to. Callers write
// assetcategory.Tenancy. The Type values and the type->category mapping live in
// the sibling assettype package; nothing here references a parent catalog.
package assetcategory

// Category is a canonical, provider-neutral asset category.
type Category string

// The canonical asset categories. Unknown is the zero-value fallback.
const (
	Unknown         Category = "Unknown"
	Compute         Category = "Compute"
	Serverless      Category = "Serverless"
	Containers      Category = "Containers"
	Network         Category = "Network"
	IpAddress       Category = "IpAddress"
	Dns             Category = "Dns"
	LoadBalancing   Category = "LoadBalancing"
	ApiGateway      Category = "ApiGateway"
	Cdn             Category = "Cdn"
	EdgeSecurity    Category = "EdgeSecurity"
	Storage         Category = "Storage"
	Database        Category = "Database"
	Analytics       Category = "Analytics"
	Messaging       Category = "Messaging"
	Identity        Category = "Identity"
	Secrets         Category = "Secrets"
	SecurityTooling Category = "SecurityTooling"
	Observability   Category = "Observability"
	Tenancy         Category = "Tenancy"
	Geography       Category = "Geography"
	Governance      Category = "Governance"
	Migration       Category = "Migration"
	MachineLearning Category = "MachineLearning"
	DataProtection  Category = "DataProtection"
	Integration     Category = "Integration"
)

// All is every category constant, for iteration and parity checks.
var All = []Category{
	Unknown, Compute, Serverless, Containers, Network, IpAddress, Dns,
	LoadBalancing, ApiGateway, Cdn, EdgeSecurity, Storage, Database, Analytics,
	Messaging, Identity, Secrets, SecurityTooling, Observability, Tenancy,
	Geography, Governance, Migration, MachineLearning, DataProtection, Integration,
}
