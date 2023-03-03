# {{classname}}

All URIs are relative to *https://console1.cloudavenue.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVdcEdge**](EdgeGatewaysApi.md#CreateVdcEdge) | **Post** /api/customers/v2.0/vdcs/{vdc-name}/edges | Create one Edge gateway
[**CreateVdcGroupEdge**](EdgeGatewaysApi.md#CreateVdcGroupEdge) | **Post** /api/customers/v2.0/vdc-groups/{vdc-group-name}/edges | Create one Edge gateway in vDC group
[**DeleteEdge**](EdgeGatewaysApi.md#DeleteEdge) | **Delete** /api/customers/v2.0/edges/{edge-id} | Remove Edge Gateway
[**GetEdgeById**](EdgeGatewaysApi.md#GetEdgeById) | **Get** /api/customers/v2.0/edges/{edge-id} | Get Edge gateway details
[**GetEdgeLoadBalancing**](EdgeGatewaysApi.md#GetEdgeLoadBalancing) | **Get** /api/customers/v2.0/edges/{edge-id}/load_balancing | Get Load Balancing Configuration on Edge Gateway
[**GetEdgeNetworks**](EdgeGatewaysApi.md#GetEdgeNetworks) | **Get** /api/customers/v2.0/edges/{edge-id}/networks | Get edge gateway network configuration
[**GetEdges**](EdgeGatewaysApi.md#GetEdges) | **Get** /api/customers/v2.0/edges | List edge gateway of an organization
[**UpdateEdgeLoadBalancing**](EdgeGatewaysApi.md#UpdateEdgeLoadBalancing) | **Put** /api/customers/v2.0/edges/{edge-id}/load_balancing | Modify Edge Gateway Load Balancing Configuration

# **CreateVdcEdge**
> Jobcreated CreateVdcEdge(ctx, body, vdcName)
Create one Edge gateway

When creating an Edge Gateway attached to a vDC, this vDC will be isolated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EdgeGatewayCreate**](EdgeGatewayCreate.md)| Edge Gateway creation body | 
  **vdcName** | **string**| VDC name. | 

### Return type

[**Jobcreated**](jobcreated.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: application/json;version=35.2
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVdcGroupEdge**
> Jobcreated CreateVdcGroupEdge(ctx, body, vdcGroupName)
Create one Edge gateway in vDC group

When creating an Edge Gateway attached to a vDC group, this Edge Gateway will be shared with other vDC inside the same vDC group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EdgeGatewayCreate**](EdgeGatewayCreate.md)| Edge Gateway creation body | 
  **vdcGroupName** | **string**| VDC Group name. | 

### Return type

[**Jobcreated**](jobcreated.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: application/json;version=35.2
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEdge**
> Jobcreated DeleteEdge(ctx, edgeId)
Remove Edge Gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeId** | **string**|  | 

### Return type

[**Jobcreated**](jobcreated.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEdgeById**
> EdgeGateway GetEdgeById(ctx, edgeId)
Get Edge gateway details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeId** | **string**|  | 

### Return type

[**EdgeGateway**](edgeGateway.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEdgeLoadBalancing**
> EdgeGatewayLoadBalancing GetEdgeLoadBalancing(ctx, edgeId)
Get Load Balancing Configuration on Edge Gateway

Retrieve the configuration of edge gateway network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeId** | **string**|  | 

### Return type

[**EdgeGatewayLoadBalancing**](edgeGatewayLoadBalancing.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEdgeNetworks**
> []EdgeGatewayNetworkGet GetEdgeNetworks(ctx, edgeId)
Get edge gateway network configuration

Retrieve the configuration of edge gateway network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeId** | **string**|  | 

### Return type

[**[]EdgeGatewayNetworkGet**](array.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEdges**
> []EdgeGateway GetEdges(ctx, )
List edge gateway of an organization

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EdgeGateway**](edgeGateway.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEdgeLoadBalancing**
> Jobcreated UpdateEdgeLoadBalancing(ctx, body, edgeId)
Modify Edge Gateway Load Balancing Configuration

Modification of the configuration of edge gateway load balancing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EdgeGatewayLoadBalancing**](EdgeGatewayLoadBalancing.md)| new edge gateway subnets configuration | 
  **edgeId** | **string**|  | 

### Return type

[**Jobcreated**](jobcreated.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: application/json;version=35.2
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

