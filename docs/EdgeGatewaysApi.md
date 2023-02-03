# {{classname}}

All URIs are relative to *https://console1.cloudavenue.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCustomersV20EdgesEdgeIdDelete**](EdgeGatewaysApi.md#ApiCustomersV20EdgesEdgeIdDelete) | **Delete** /api/customers/v2.0/edges/{edge-id} | Remove Edge Gateway
[**ApiCustomersV20EdgesEdgeIdGet**](EdgeGatewaysApi.md#ApiCustomersV20EdgesEdgeIdGet) | **Get** /api/customers/v2.0/edges/{edge-id} | Get Edge gateway details
[**ApiCustomersV20EdgesEdgeIdNetworksGet**](EdgeGatewaysApi.md#ApiCustomersV20EdgesEdgeIdNetworksGet) | **Get** /api/customers/v2.0/edges/{edge-id}/networks | Get edge gateway network configuration
[**ApiCustomersV20EdgesGet**](EdgeGatewaysApi.md#ApiCustomersV20EdgesGet) | **Get** /api/customers/v2.0/edges | List edge gateway of an organization
[**ApiCustomersV20VdcGroupsVdcGroupNameEdgesPost**](EdgeGatewaysApi.md#ApiCustomersV20VdcGroupsVdcGroupNameEdgesPost) | **Post** /api/customers/v2.0/vdc-groups/{vdc-group-name}/edges | Create one Edge gateway in vDC group
[**ApiCustomersV20VdcsVdcNameEdgesPost**](EdgeGatewaysApi.md#ApiCustomersV20VdcsVdcNameEdgesPost) | **Post** /api/customers/v2.0/vdcs/{vdc-name}/edges | Create one Edge gateway

# **ApiCustomersV20EdgesEdgeIdDelete**
> string ApiCustomersV20EdgesEdgeIdDelete(ctx, edgeId)
Remove Edge Gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeId** | **string**|  | 

### Return type

**string**

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersV20EdgesEdgeIdGet**
> EdgeGateway ApiCustomersV20EdgesEdgeIdGet(ctx, edgeId)
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

# **ApiCustomersV20EdgesEdgeIdNetworksGet**
> []EdgeGatewayNetworkGet ApiCustomersV20EdgesEdgeIdNetworksGet(ctx, edgeId)
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

# **ApiCustomersV20EdgesGet**
> []EdgeGateway ApiCustomersV20EdgesGet(ctx, )
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

# **ApiCustomersV20VdcGroupsVdcGroupNameEdgesPost**
> Jobcreated ApiCustomersV20VdcGroupsVdcGroupNameEdgesPost(ctx, body, vdcGroupName)
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

# **ApiCustomersV20VdcsVdcNameEdgesPost**
> Jobcreated ApiCustomersV20VdcsVdcNameEdgesPost(ctx, body, vdcName)
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

