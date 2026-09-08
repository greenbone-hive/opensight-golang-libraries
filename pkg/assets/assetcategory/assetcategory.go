// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package assetcategory

type Category string

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

var All = []Category{
	Unknown, Compute, Serverless, Containers, Network, IpAddress, Dns,
	LoadBalancing, ApiGateway, Cdn, EdgeSecurity, Storage, Database, Analytics,
	Messaging, Identity, Secrets, SecurityTooling, Observability, Tenancy,
	Geography, Governance, Migration, MachineLearning, DataProtection, Integration,
}
