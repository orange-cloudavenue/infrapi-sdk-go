# {{classname}}

All URIs are relative to *https://console1.cloudavenue.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgVdc**](VDCApi.md#CreateOrgVdc) | **Post** /api/customers/v2.0/vdcs | Create a new Org VDC
[**DeleteOrgVdc**](VDCApi.md#DeleteOrgVdc) | **Delete** /api/customers/v2.0/vdcs/{vdc-name} | Delete a vDC
[**GetOrgVdcByName**](VDCApi.md#GetOrgVdcByName) | **Get** /api/customers/v2.0/vdcs/{vdc-name} | Get details about one vDC
[**GetOrgVdcs**](VDCApi.md#GetOrgVdcs) | **Get** /api/customers/v2.0/vdcs | List Org vDCs
[**UpdateOrgVdc**](VDCApi.md#UpdateOrgVdc) | **Put** /api/customers/v2.0/vdcs/{vdc-name} | Update a vDC

# **CreateOrgVdc**
> Jobcreated CreateOrgVdc(ctx, body)
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

# **DeleteOrgVdc**
> Jobcreated DeleteOrgVdc(ctx, vdcName)
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

# **GetOrgVdcByName**
> GetOrgVdcs2 GetOrgVdcByName(ctx, vdcName)
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

# **GetOrgVdcs**
> []VDcList GetOrgVdcs(ctx, )
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

# **UpdateOrgVdc**
> Jobcreated UpdateOrgVdc(ctx, body, vdcName)
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

