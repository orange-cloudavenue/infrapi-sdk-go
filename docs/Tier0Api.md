# {{classname}}

All URIs are relative to *https://console1.cloudavenue.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTier0VrfByName**](Tier0Api.md#GetTier0VrfByName) | **Get** /api/customers/v2.0/tier-0-vrfs/{tier0_name} | Get Tier-0 Details
[**GetTier0Vrfs**](Tier0Api.md#GetTier0Vrfs) | **Get** /api/customers/v2.0/tier-0-vrfs | Get all Tier-0 Gateway

# **GetTier0VrfByName**
> Tier0Details GetTier0VrfByName(ctx, tier0Name)
Get Tier-0 Details

Get Tier-0 details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Name** | [**string**](.md)|  | 

### Return type

[**Tier0Details**](tier0Details.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier0Vrfs**
> []string GetTier0Vrfs(ctx, )
Get all Tier-0 Gateway

Get all Tier-0 gateway.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]string**](array.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

