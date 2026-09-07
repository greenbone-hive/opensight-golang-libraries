// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package assettype is the single source of the canonical asset types: the
// fixed, provider-neutral type catalog plus each type's category. Callers write
// assettype.Organization. Categories live in the sibling assetcategory package;
// nothing here references a parent catalog.
package assettype

import "github.com/greenbone/opensight-golang-libraries/pkg/assets/assetcategory"

// Type is a canonical, provider-neutral asset type.
type Type string

// The canonical asset types. Unknown is the zero-value fallback for a resource
// whose type has not been mapped.
const (
	Unknown                   Type = "Unknown"
	Server                    Type = "Server"
	VirtualMachine            Type = "VirtualMachine"
	Endpoint                  Type = "Endpoint"
	ContainerHost             Type = "ContainerHost"
	AppService                Type = "AppService"
	BatchJob                  Type = "BatchJob"
	VirtualDesktop            Type = "VirtualDesktop"
	AutoScalingGroup          Type = "AutoScalingGroup"
	LaunchTemplate            Type = "LaunchTemplate"
	ScheduledTask             Type = "ScheduledTask"
	ServerlessFunction        Type = "ServerlessFunction"
	KubernetesCluster         Type = "KubernetesCluster"
	KubernetesNodePool        Type = "KubernetesNodePool"
	ContainerService          Type = "ContainerService"
	ContainerTaskDefinition   Type = "ContainerTaskDefinition"
	ContainerRegistry         Type = "ContainerRegistry"
	ContainerImage            Type = "ContainerImage"
	FargateProfile            Type = "FargateProfile"
	VirtualNetwork            Type = "VirtualNetwork"
	Subnet                    Type = "Subnet"
	RouteTable                Type = "RouteTable"
	Route                     Type = "Route"
	NetworkAcl                Type = "NetworkAcl"
	NetworkSecurityGroup      Type = "NetworkSecurityGroup"
	ApplicationSecurityGroup  Type = "ApplicationSecurityGroup"
	NetworkInterface          Type = "NetworkInterface"
	InternetGateway           Type = "InternetGateway"
	EgressOnlyInternetGateway Type = "EgressOnlyInternetGateway"
	NatGateway                Type = "NatGateway"
	CloudRouter               Type = "CloudRouter"
	BastionHost               Type = "BastionHost"
	NetworkWatcher            Type = "NetworkWatcher"
	DhcpOptionsSet            Type = "DhcpOptionsSet"
	PrefixList                Type = "PrefixList"
	Internet                  Type = "Internet"
	VpcPeering                Type = "VpcPeering"
	TransitHub                Type = "TransitHub"
	VpcEndpoint               Type = "VpcEndpoint"
	VpcEndpointService        Type = "VpcEndpointService"
	PrivateLinkService        Type = "PrivateLinkService"
	ServiceAttachment         Type = "ServiceAttachment"
	PrivateEndpoint           Type = "PrivateEndpoint"
	VpnGateway                Type = "VpnGateway"
	VpnConnection             Type = "VpnConnection"
	ClientVpnEndpoint         Type = "ClientVpnEndpoint"
	DedicatedCircuit          Type = "DedicatedCircuit"
	OnPremSite                Type = "OnPremSite"
	PublicIpAddress           Type = "PublicIpAddress"
	DnsZone                   Type = "DnsZone"
	DnsRecord                 Type = "DnsRecord"
	DnsResolver               Type = "DnsResolver"
	LoadBalancer              Type = "LoadBalancer"
	ApplicationGateway        Type = "ApplicationGateway"
	LbListener                Type = "LbListener"
	LbBackendPool             Type = "LbBackendPool"
	GlobalAcceleratorEndpoint Type = "GlobalAcceleratorEndpoint"
	TrafficManager            Type = "TrafficManager"
	ApiGateway                Type = "ApiGateway"
	Cdn                       Type = "Cdn"
	WafPolicy                 Type = "WafPolicy"
	ManagedFirewall           Type = "ManagedFirewall"
	FirewallPolicy            Type = "FirewallPolicy"
	DdosProtectionPlan        Type = "DdosProtectionPlan"
	BlockStorage              Type = "BlockStorage"
	BlockStorageSnapshot      Type = "BlockStorageSnapshot"
	FileStorage               Type = "FileStorage"
	ArchiveStorage            Type = "ArchiveStorage"
	MachineImage              Type = "MachineImage"
	StorageAccessPoint        Type = "StorageAccessPoint"
	ObjectStorage             Type = "ObjectStorage"
	DatabaseServer            Type = "DatabaseServer"
	DatabaseCluster           Type = "DatabaseCluster"
	DatabaseNoSql             Type = "DatabaseNoSql"
	DatabaseSnapshot          Type = "DatabaseSnapshot"
	DataWarehouse             Type = "DataWarehouse"
	Cache                     Type = "Cache"
	SearchService             Type = "SearchService"
	DataLake                  Type = "DataLake"
	DataCatalog               Type = "DataCatalog"
	DataPipeline              Type = "DataPipeline"
	MessageQueue              Type = "MessageQueue"
	Topic                     Type = "Topic"
	StreamService             Type = "StreamService"
	EventBus                  Type = "EventBus"
	IotHub                    Type = "IotHub"
	IdentityProvider          Type = "IdentityProvider"
	UserIdentity              Type = "UserIdentity"
	IamGroup                  Type = "IamGroup"
	IamPolicy                 Type = "IamPolicy"
	RoleDefinition            Type = "RoleDefinition"
	RoleAssignment            Type = "RoleAssignment"
	ServicePrincipal          Type = "ServicePrincipal"
	WorkloadIdentityPool      Type = "WorkloadIdentityPool"
	SamlProvider              Type = "SamlProvider"
	OidcProvider              Type = "OidcProvider"
	AccessKey                 Type = "AccessKey"
	ServiceAccountKey         Type = "ServiceAccountKey"
	Application               Type = "Application"
	IdentityDevice            Type = "IdentityDevice"
	ExternalPrincipal         Type = "ExternalPrincipal"
	PublicPrincipal           Type = "PublicPrincipal"
	UncollectedPrincipal      Type = "UncollectedPrincipal"
	InstanceProfile           Type = "InstanceProfile"
	PermissionSet             Type = "PermissionSet"
	PolicyStatement           Type = "PolicyStatement"
	ResourcePolicy            Type = "ResourcePolicy"
	GovernanceBinding         Type = "GovernanceBinding"
	AppRoleAssignment         Type = "AppRoleAssignment"
	RoleEligibilitySchedule   Type = "RoleEligibilitySchedule"
	DenyAssignment            Type = "DenyAssignment"
	DenyPolicy                Type = "DenyPolicy"
	DenyRule                  Type = "DenyRule"
	ResourceShare             Type = "ResourceShare"
	IdpConfig                 Type = "IdpConfig"
	MfaDevice                 Type = "MfaDevice"
	PasswordPolicy            Type = "PasswordPolicy"
	ConditionalAccessPolicy   Type = "ConditionalAccessPolicy"
	DirectoryService          Type = "DirectoryService"
	KeyVault                  Type = "KeyVault"
	EncryptionKey             Type = "EncryptionKey"
	Secret                    Type = "Secret"
	Certificate               Type = "Certificate"
	HsmCluster                Type = "HsmCluster"
	SecurityCenter            Type = "SecurityCenter"
	SecurityFinding           Type = "SecurityFinding"
	ThreatDetector            Type = "ThreatDetector"
	AccessAnalyzer            Type = "AccessAnalyzer"
	ConfigAudit               Type = "ConfigAudit"
	FirewallManagerPolicy     Type = "FirewallManagerPolicy"
	LogWorkspace              Type = "LogWorkspace"
	LogGroup                  Type = "LogGroup"
	LogMetricFilter           Type = "LogMetricFilter"
	AuditTrail                Type = "AuditTrail"
	MonitorAlert              Type = "MonitorAlert"
	Dashboard                 Type = "Dashboard"
	Account                   Type = "Account"
	OrganizationalUnit        Type = "OrganizationalUnit"
	ResourceGroup             Type = "ResourceGroup"
	Organization              Type = "Organization"
	Region                    Type = "Region"
	ServiceControlPolicy      Type = "ServiceControlPolicy"
	ResourceControlPolicy     Type = "ResourceControlPolicy"
	PolicyAssignment          Type = "PolicyAssignment"
	OrganizationPolicy        Type = "OrganizationPolicy"
	MigrationProject          Type = "MigrationProject"
	MigrationTask             Type = "MigrationTask"
	ReplicationConfiguration  Type = "ReplicationConfiguration"

	// Crown-jewel plane additions (secrets/keys, backup, ML, integration).
	BackupVault          Type = "BackupVault"
	MlWorkspace          Type = "MlWorkspace"
	MlModel              Type = "MlModel"
	MlDataset            Type = "MlDataset"
	CertificateAuthority Type = "CertificateAuthority"
	ApiKey               Type = "ApiKey"
	PolicyStore          Type = "PolicyStore"
	AutomationService    Type = "AutomationService"
	AppConfiguration     Type = "AppConfiguration"
)

// All is every asset-type constant, for iteration and parity checks.
var All = []Type{
	Unknown,
	Server,
	VirtualMachine,
	Endpoint,
	ContainerHost,
	AppService,
	BatchJob,
	VirtualDesktop,
	AutoScalingGroup,
	LaunchTemplate,
	ScheduledTask,
	ServerlessFunction,
	KubernetesCluster,
	KubernetesNodePool,
	ContainerService,
	ContainerTaskDefinition,
	ContainerRegistry,
	ContainerImage,
	FargateProfile,
	VirtualNetwork,
	Subnet,
	RouteTable,
	Route,
	NetworkAcl,
	NetworkSecurityGroup,
	ApplicationSecurityGroup,
	NetworkInterface,
	InternetGateway,
	EgressOnlyInternetGateway,
	NatGateway,
	CloudRouter,
	BastionHost,
	NetworkWatcher,
	DhcpOptionsSet,
	PrefixList,
	Internet,
	VpcPeering,
	TransitHub,
	VpcEndpoint,
	VpcEndpointService,
	PrivateLinkService,
	ServiceAttachment,
	PrivateEndpoint,
	VpnGateway,
	VpnConnection,
	ClientVpnEndpoint,
	DedicatedCircuit,
	OnPremSite,
	PublicIpAddress,
	DnsZone,
	DnsRecord,
	DnsResolver,
	LoadBalancer,
	ApplicationGateway,
	LbListener,
	LbBackendPool,
	GlobalAcceleratorEndpoint,
	TrafficManager,
	ApiGateway,
	Cdn,
	WafPolicy,
	ManagedFirewall,
	FirewallPolicy,
	DdosProtectionPlan,
	BlockStorage,
	BlockStorageSnapshot,
	FileStorage,
	ArchiveStorage,
	MachineImage,
	StorageAccessPoint,
	ObjectStorage,
	DatabaseServer,
	DatabaseCluster,
	DatabaseNoSql,
	DatabaseSnapshot,
	DataWarehouse,
	Cache,
	SearchService,
	DataLake,
	DataCatalog,
	DataPipeline,
	MessageQueue,
	Topic,
	StreamService,
	EventBus,
	IotHub,
	IdentityProvider,
	UserIdentity,
	IamGroup,
	IamPolicy,
	RoleDefinition,
	RoleAssignment,
	ServicePrincipal,
	WorkloadIdentityPool,
	SamlProvider,
	OidcProvider,
	AccessKey,
	ServiceAccountKey,
	Application,
	IdentityDevice,
	ExternalPrincipal,
	PublicPrincipal,
	UncollectedPrincipal,
	InstanceProfile,
	PermissionSet,
	PolicyStatement,
	ResourcePolicy,
	GovernanceBinding,
	AppRoleAssignment,
	RoleEligibilitySchedule,
	DenyAssignment,
	DenyPolicy,
	DenyRule,
	ResourceShare,
	IdpConfig,
	MfaDevice,
	PasswordPolicy,
	ConditionalAccessPolicy,
	DirectoryService,
	KeyVault,
	EncryptionKey,
	Secret,
	Certificate,
	HsmCluster,
	SecurityCenter,
	SecurityFinding,
	ThreatDetector,
	AccessAnalyzer,
	ConfigAudit,
	FirewallManagerPolicy,
	LogWorkspace,
	LogGroup,
	LogMetricFilter,
	AuditTrail,
	MonitorAlert,
	Dashboard,
	Account,
	OrganizationalUnit,
	ResourceGroup,
	Organization,
	Region,
	ServiceControlPolicy,
	ResourceControlPolicy,
	PolicyAssignment,
	OrganizationPolicy,
	MigrationProject,
	MigrationTask,
	ReplicationConfiguration,
	BackupVault,
	MlWorkspace,
	MlModel,
	MlDataset,
	CertificateAuthority,
	ApiKey,
	PolicyStore,
	AutomationService,
	AppConfiguration,
}

// categoryByType maps each known type to its category (single source; keys are the
// assettype constants, values the assetcategory constants).
var categoryByType = map[Type]assetcategory.Category{
	Unknown:                   assetcategory.Unknown,
	Server:                    assetcategory.Compute,
	VirtualMachine:            assetcategory.Compute,
	Endpoint:                  assetcategory.Compute,
	ContainerHost:             assetcategory.Compute,
	AppService:                assetcategory.Compute,
	BatchJob:                  assetcategory.Compute,
	VirtualDesktop:            assetcategory.Compute,
	AutoScalingGroup:          assetcategory.Compute,
	LaunchTemplate:            assetcategory.Compute,
	ScheduledTask:             assetcategory.Compute,
	ServerlessFunction:        assetcategory.Serverless,
	KubernetesCluster:         assetcategory.Containers,
	KubernetesNodePool:        assetcategory.Containers,
	ContainerService:          assetcategory.Containers,
	ContainerTaskDefinition:   assetcategory.Containers,
	ContainerRegistry:         assetcategory.Containers,
	ContainerImage:            assetcategory.Containers,
	FargateProfile:            assetcategory.Containers,
	VirtualNetwork:            assetcategory.Network,
	Subnet:                    assetcategory.Network,
	RouteTable:                assetcategory.Network,
	Route:                     assetcategory.Network,
	NetworkAcl:                assetcategory.Network,
	NetworkSecurityGroup:      assetcategory.Network,
	ApplicationSecurityGroup:  assetcategory.Network,
	NetworkInterface:          assetcategory.Network,
	InternetGateway:           assetcategory.Network,
	EgressOnlyInternetGateway: assetcategory.Network,
	NatGateway:                assetcategory.Network,
	CloudRouter:               assetcategory.Network,
	BastionHost:               assetcategory.Network,
	NetworkWatcher:            assetcategory.Network,
	DhcpOptionsSet:            assetcategory.Network,
	PrefixList:                assetcategory.Network,
	Internet:                  assetcategory.Network,
	VpcPeering:                assetcategory.Network,
	TransitHub:                assetcategory.Network,
	VpcEndpoint:               assetcategory.Network,
	VpcEndpointService:        assetcategory.Network,
	PrivateLinkService:        assetcategory.Network,
	ServiceAttachment:         assetcategory.Network,
	PrivateEndpoint:           assetcategory.Network,
	VpnGateway:                assetcategory.Network,
	VpnConnection:             assetcategory.Network,
	ClientVpnEndpoint:         assetcategory.Network,
	DedicatedCircuit:          assetcategory.Network,
	OnPremSite:                assetcategory.Network,
	PublicIpAddress:           assetcategory.IpAddress,
	DnsZone:                   assetcategory.Dns,
	DnsRecord:                 assetcategory.Dns,
	DnsResolver:               assetcategory.Dns,
	LoadBalancer:              assetcategory.LoadBalancing,
	ApplicationGateway:        assetcategory.LoadBalancing,
	LbListener:                assetcategory.LoadBalancing,
	LbBackendPool:             assetcategory.LoadBalancing,
	GlobalAcceleratorEndpoint: assetcategory.LoadBalancing,
	TrafficManager:            assetcategory.LoadBalancing,
	ApiGateway:                assetcategory.ApiGateway,
	Cdn:                       assetcategory.Cdn,
	WafPolicy:                 assetcategory.EdgeSecurity,
	ManagedFirewall:           assetcategory.EdgeSecurity,
	FirewallPolicy:            assetcategory.EdgeSecurity,
	DdosProtectionPlan:        assetcategory.EdgeSecurity,
	BlockStorage:              assetcategory.Storage,
	BlockStorageSnapshot:      assetcategory.Storage,
	FileStorage:               assetcategory.Storage,
	ArchiveStorage:            assetcategory.Storage,
	MachineImage:              assetcategory.Storage,
	StorageAccessPoint:        assetcategory.Storage,
	ObjectStorage:             assetcategory.Storage,
	DatabaseServer:            assetcategory.Database,
	DatabaseCluster:           assetcategory.Database,
	DatabaseNoSql:             assetcategory.Database,
	DatabaseSnapshot:          assetcategory.Database,
	DataWarehouse:             assetcategory.Database,
	Cache:                     assetcategory.Database,
	SearchService:             assetcategory.Database,
	DataLake:                  assetcategory.Analytics,
	DataCatalog:               assetcategory.Analytics,
	DataPipeline:              assetcategory.Integration,
	MessageQueue:              assetcategory.Messaging,
	Topic:                     assetcategory.Messaging,
	StreamService:             assetcategory.Messaging,
	EventBus:                  assetcategory.Messaging,
	IotHub:                    assetcategory.Messaging,
	IdentityProvider:          assetcategory.Identity,
	UserIdentity:              assetcategory.Identity,
	IamGroup:                  assetcategory.Identity,
	IamPolicy:                 assetcategory.Identity,
	RoleDefinition:            assetcategory.Identity,
	RoleAssignment:            assetcategory.Identity,
	ServicePrincipal:          assetcategory.Identity,
	WorkloadIdentityPool:      assetcategory.Identity,
	SamlProvider:              assetcategory.Identity,
	OidcProvider:              assetcategory.Identity,
	AccessKey:                 assetcategory.Identity,
	ServiceAccountKey:         assetcategory.Identity,
	Application:               assetcategory.Identity,
	IdentityDevice:            assetcategory.Identity,
	ExternalPrincipal:         assetcategory.Identity,
	PublicPrincipal:           assetcategory.Identity,
	UncollectedPrincipal:      assetcategory.Identity,
	InstanceProfile:           assetcategory.Identity,
	PermissionSet:             assetcategory.Identity,
	PolicyStatement:           assetcategory.Identity,
	ResourcePolicy:            assetcategory.Identity,
	GovernanceBinding:         assetcategory.Identity,
	AppRoleAssignment:         assetcategory.Identity,
	RoleEligibilitySchedule:   assetcategory.Identity,
	DenyAssignment:            assetcategory.Identity,
	DenyPolicy:                assetcategory.Identity,
	DenyRule:                  assetcategory.Identity,
	ResourceShare:             assetcategory.Identity,
	IdpConfig:                 assetcategory.Identity,
	MfaDevice:                 assetcategory.Identity,
	PasswordPolicy:            assetcategory.Identity,
	ConditionalAccessPolicy:   assetcategory.Identity,
	DirectoryService:          assetcategory.Identity,
	KeyVault:                  assetcategory.Secrets,
	EncryptionKey:             assetcategory.Secrets,
	Secret:                    assetcategory.Secrets,
	Certificate:               assetcategory.Secrets,
	HsmCluster:                assetcategory.Secrets,
	SecurityCenter:            assetcategory.SecurityTooling,
	SecurityFinding:           assetcategory.SecurityTooling,
	ThreatDetector:            assetcategory.SecurityTooling,
	AccessAnalyzer:            assetcategory.SecurityTooling,
	ConfigAudit:               assetcategory.SecurityTooling,
	FirewallManagerPolicy:     assetcategory.SecurityTooling,
	LogWorkspace:              assetcategory.Observability,
	LogGroup:                  assetcategory.Observability,
	LogMetricFilter:           assetcategory.Observability,
	AuditTrail:                assetcategory.Observability,
	MonitorAlert:              assetcategory.Observability,
	Dashboard:                 assetcategory.Observability,
	Account:                   assetcategory.Tenancy,
	OrganizationalUnit:        assetcategory.Tenancy,
	ResourceGroup:             assetcategory.Tenancy,
	Organization:              assetcategory.Tenancy,
	Region:                    assetcategory.Geography,
	ServiceControlPolicy:      assetcategory.Governance,
	ResourceControlPolicy:     assetcategory.Governance,
	PolicyAssignment:          assetcategory.Governance,
	OrganizationPolicy:        assetcategory.Governance,
	MigrationProject:          assetcategory.Migration,
	MigrationTask:             assetcategory.Migration,
	ReplicationConfiguration:  assetcategory.Migration,
	BackupVault:               assetcategory.DataProtection,
	MlWorkspace:               assetcategory.MachineLearning,
	MlModel:                   assetcategory.MachineLearning,
	MlDataset:                 assetcategory.MachineLearning,
	CertificateAuthority:      assetcategory.Secrets,
	ApiKey:                    assetcategory.Identity,
	PolicyStore:               assetcategory.Governance,
	AutomationService:         assetcategory.Integration,
	AppConfiguration:          assetcategory.Integration,
}

// CategoryOf returns the category for a type, or assetcategory.Unknown if the
// type is not in the catalog.
func CategoryOf(t Type) assetcategory.Category {
	if c, ok := categoryByType[t]; ok {
		return c
	}

	return assetcategory.Unknown
}

// IsKnown reports whether t is in the catalog.
func IsKnown(t Type) bool { _, ok := categoryByType[t]; return ok }
