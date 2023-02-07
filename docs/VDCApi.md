# {{classname}}

All URIs are relative to *https://console1.cloudavenue.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCustomersV20VdcsGet**](VDCApi.md#ApiCustomersV20VdcsGet) | **Get** /api/customers/v2.0/vdcs | List Org vDCs
[**ApiCustomersV20VdcsPost**](VDCApi.md#ApiCustomersV20VdcsPost) | **Post** /api/customers/v2.0/vdcs | Create a new Org VDC
[**ApiCustomersV20VdcsVdcNameDelete**](VDCApi.md#ApiCustomersV20VdcsVdcNameDelete) | **Delete** /api/customers/v2.0/vdcs/{vdc-name} | Delete a vDC
[**ApiCustomersV20VdcsVdcNameGet**](VDCApi.md#ApiCustomersV20VdcsVdcNameGet) | **Get** /api/customers/v2.0/vdcs/{vdc-name} | Get details about one vDC
[**ApiCustomersV20VdcsVdcNamePut**](VDCApi.md#ApiCustomersV20VdcsVdcNamePut) | **Put** /api/customers/v2.0/vdcs/{vdc-name} | Update a vDC

# **ApiCustomersV20VdcsGet**
> []VDcList ApiCustomersV20VdcsGet(ctx, )
List Org vDCs

List all VDC inside an Organization

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]VDcList**](vDCList.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersV20VdcsPost**
> Jobcreated ApiCustomersV20VdcsPost(ctx, body)
Create a new Org VDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOrgVdcV2**](CreateOrgVdcV2.md)| Tenant technical details | 

### Return type

[**Jobcreated**](jobcreated.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: application/json;version=35.2
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersV20VdcsVdcNameDelete**
> Jobcreated ApiCustomersV20VdcsVdcNameDelete(ctx, vdcName)
Delete a vDC

Delete an empty vDC, will return an error if the vDC has Edges or vApps (must remove before)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcName** | **string**| VDC name. | 

### Return type

[**Jobcreated**](jobcreated.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersV20VdcsVdcNameGet**
> GetOrgVdcs2 ApiCustomersV20VdcsVdcNameGet(ctx, vdcName)
Get details about one vDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcName** | [**string**](.md)|  | 

### Return type

[**GetOrgVdcs2**](GetOrgVdcs2.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersV20VdcsVdcNamePut**
> Jobcreated ApiCustomersV20VdcsVdcNamePut(ctx, body, vdcName)
Update a vDC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateOrgVdcV2**](UpdateOrgVdcV2.md)| shaping details | 
  **vdcName** | [**string**](.md)|  | 

### Return type

[**Jobcreated**](jobcreated.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: application/json;version=35.2
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

