# Go API client for cloudavenue

Cloud Avenue API allows you to create some ressources on vCloud Director.  ## Authentication In order to authenticate to this API, you must first authenticate through vCloudDirector API using **POST https://{{api_host}}/cloudapi/1.0.0/sessions**. Then use _X-VMWARE-VCLOUD-ACCESS-TOKEN_ from response header in your request.  If you need more information do not hesitate to read our dedicated page on [our Wiki](https://wiki.cloudavenue.orange-business.com/w/index.php/API_Cloud_Avenue) 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./cloudavenue"
```

## Documentation for API Endpoints

All URIs are relative to *https://console1.cloudavenue.orange-business.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthenticationApi* | [**GetToken**](docs/AuthenticationApi.md#gettoken) | **Post** /cloudapi/1.0.0/sessions | Login to your organization
*EdgeGatewaysApi* | [**CreateVdcEdge**](docs/EdgeGatewaysApi.md#createvdcedge) | **Post** /api/customers/v2.0/vdcs/{vdc-name}/edges | Create one Edge gateway
*EdgeGatewaysApi* | [**CreateVdcGroupEdge**](docs/EdgeGatewaysApi.md#createvdcgroupedge) | **Post** /api/customers/v2.0/vdc-groups/{vdc-group-name}/edges | Create one Edge gateway in vDC group
*EdgeGatewaysApi* | [**DeleteEdge**](docs/EdgeGatewaysApi.md#deleteedge) | **Delete** /api/customers/v2.0/edges/{edge-id} | Remove Edge Gateway
*EdgeGatewaysApi* | [**GetEdgeById**](docs/EdgeGatewaysApi.md#getedgebyid) | **Get** /api/customers/v2.0/edges/{edge-id} | Get Edge gateway details
*EdgeGatewaysApi* | [**GetEdgeLoadBalancing**](docs/EdgeGatewaysApi.md#getedgeloadbalancing) | **Get** /api/customers/v2.0/edges/{edge-id}/load_balancing | Get Load Balancing Configuration on Edge Gateway
*EdgeGatewaysApi* | [**GetEdgeNetworks**](docs/EdgeGatewaysApi.md#getedgenetworks) | **Get** /api/customers/v2.0/edges/{edge-id}/networks | Get edge gateway network configuration
*EdgeGatewaysApi* | [**GetEdges**](docs/EdgeGatewaysApi.md#getedges) | **Get** /api/customers/v2.0/edges | List edge gateway of an organization
*EdgeGatewaysApi* | [**UpdateEdgeLoadBalancing**](docs/EdgeGatewaysApi.md#updateedgeloadbalancing) | **Put** /api/customers/v2.0/edges/{edge-id}/load_balancing | Modify Edge Gateway Load Balancing Configuration
*JobsApi* | [**GetJobById**](docs/JobsApi.md#getjobbyid) | **Get** /api/customers/v1.0/jobs/{Job-Id} | Returns a job status by ID.
*PublicIPApi* | [**CreatePublicIP**](docs/PublicIPApi.md#createpublicip) | **Post** /api/customers/v1.0/ip | Request and configure a new public IP address
*PublicIPApi* | [**DeletePublicIP**](docs/PublicIPApi.md#deletepublicip) | **Delete** /api/customers/v1.0/ip/{public-ip}/ | Remove all configuration related to a public IP address and free IP
*PublicIPApi* | [**GetPublicIPs**](docs/PublicIPApi.md#getpublicips) | **Get** /api/customers/v2.0/ip | Get Organization&#x27;s public IP addresses
*Tier0Api* | [**GetTier0VrfByName**](docs/Tier0Api.md#gettier0vrfbyname) | **Get** /api/customers/v2.0/tier-0-vrfs/{tier0_name} | Get Tier-0 Details
*Tier0Api* | [**GetTier0Vrfs**](docs/Tier0Api.md#gettier0vrfs) | **Get** /api/customers/v2.0/tier-0-vrfs | Get all Tier-0 Gateway
*UsersApi* | [**CreateFederation**](docs/UsersApi.md#createfederation) | **Post** /api/customers/v1.0/federate | Federate an existing organization
*UsersApi* | [**GetFederation**](docs/UsersApi.md#getfederation) | **Get** /api/customers/v1.0/federate | Return federation status of an organization
*UsersApi* | [**UpdateFederation**](docs/UsersApi.md#updatefederation) | **Put** /api/customers/v1.0/federate | Regenerate SAML certificate
*VCDAApi* | [**CreateVcdaIP**](docs/VCDAApi.md#createvcdaip) | **Post** /api/customers/v2.0/vcda/ips/{Ip-Address} | Add on premise IP address
*VCDAApi* | [**DeleteVcdaIP**](docs/VCDAApi.md#deletevcdaip) | **Delete** /api/customers/v2.0/vcda/ips/{Ip-Address} | Remove on premise IP address
*VCDAApi* | [**GetVcdaIPs**](docs/VCDAApi.md#getvcdaips) | **Get** /api/customers/v2.0/vcda/ips | Get on premise IP addresses
*VDCApi* | [**CreateOrgVdc**](docs/VDCApi.md#createorgvdc) | **Post** /api/customers/v2.0/vdcs | Create a new Org VDC
*VDCApi* | [**DeleteOrgVdc**](docs/VDCApi.md#deleteorgvdc) | **Delete** /api/customers/v2.0/vdcs/{vdc-name} | Delete a vDC
*VDCApi* | [**GetOrgVdcByName**](docs/VDCApi.md#getorgvdcbyname) | **Get** /api/customers/v2.0/vdcs/{vdc-name} | Get details about one vDC
*VDCApi* | [**GetOrgVdcs**](docs/VDCApi.md#getorgvdcs) | **Get** /api/customers/v2.0/vdcs | List Org vDCs
*VDCApi* | [**UpdateOrgVdc**](docs/VDCApi.md#updateorgvdc) | **Put** /api/customers/v2.0/vdcs/{vdc-name} | Update a vDC

## Documentation For Models

 - [Action](docs/Action.md)
 - [AllOfcreateOrg2VdcsDualVdcsVdc1](docs/AllOfcreateOrg2VdcsDualVdcsVdc1.md)
 - [AllOfcreateOrg2VdcsDualVdcsVdc2](docs/AllOfcreateOrg2VdcsDualVdcsVdc2.md)
 - [ApiError](docs/ApiError.md)
 - [BvpnType](docs/BvpnType.md)
 - [CreateOrg2Vdcs](docs/CreateOrg2Vdcs.md)
 - [CreateOrg2VdcsDualVdcs](docs/CreateOrg2VdcsDualVdcs.md)
 - [CreateOrgVdcV2](docs/CreateOrgVdcV2.md)
 - [CustomerInformation](docs/CustomerInformation.md)
 - [CustomerType](docs/CustomerType.md)
 - [EdgeGateway](docs/EdgeGateway.md)
 - [EdgeGatewayCreate](docs/EdgeGatewayCreate.md)
 - [EdgeGatewayLoadBalancing](docs/EdgeGatewayLoadBalancing.md)
 - [EdgeGatewayNetworkGet](docs/EdgeGatewayNetworkGet.md)
 - [FederationStatusType](docs/FederationStatusType.md)
 - [GetOrgVdcs2](docs/GetOrgVdcs2.md)
 - [IdentityProvider](docs/IdentityProvider.md)
 - [Jobcreated](docs/Jobcreated.md)
 - [Metadata](docs/Metadata.md)
 - [NetworkConfigInner](docs/NetworkConfigInner.md)
 - [OrgUser](docs/OrgUser.md)
 - [OrgUserRecord](docs/OrgUserRecord.md)
 - [OrgVdcV2](docs/OrgVdcV2.md)
 - [OrgVdcsInner](docs/OrgVdcsInner.md)
 - [ParentJobMetadata](docs/ParentJobMetadata.md)
 - [Pesg](docs/Pesg.md)
 - [PublicIps](docs/PublicIps.md)
 - [PublicIpsNetworkConfig](docs/PublicIpsNetworkConfig.md)
 - [Segment](docs/Segment.md)
 - [ServiceType](docs/ServiceType.md)
 - [SiteCode](docs/SiteCode.md)
 - [StorageProfile](docs/StorageProfile.md)
 - [StoragesInner](docs/StoragesInner.md)
 - [TagVrf](docs/TagVrf.md)
 - [Tier0Details](docs/Tier0Details.md)
 - [Tier0VrfClassOfService](docs/Tier0VrfClassOfService.md)
 - [UpdateOrgVdcV2](docs/UpdateOrgVdcV2.md)
 - [VCdPlugin](docs/VCdPlugin.md)
 - [VDcList](docs/VDcList.md)
 - [VdcEdge](docs/VdcEdge.md)
 - [VdcStorageProfilesV2](docs/VdcStorageProfilesV2.md)

## Documentation For Authorization

## X-VMWARE-VCLOUD-ACCESS-TOKEN
## vCloudDirector
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Author


