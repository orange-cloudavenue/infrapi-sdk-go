# {{classname}}

All URIs are relative to *https://console1.cloudavenue.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCustomersV20VcdaIpsGet**](VCDAApi.md#ApiCustomersV20VcdaIpsGet) | **Get** /api/customers/v2.0/vcda/ips | Get on premise IP addresses
[**ApiCustomersV20VcdaIpsIpAddressDelete**](VCDAApi.md#ApiCustomersV20VcdaIpsIpAddressDelete) | **Delete** /api/customers/v2.0/vcda/ips/{Ip-Address} | Remove on premise IP address
[**ApiCustomersV20VcdaIpsIpAddressPost**](VCDAApi.md#ApiCustomersV20VcdaIpsIpAddressPost) | **Post** /api/customers/v2.0/vcda/ips/{Ip-Address} | Add on premise IP address

# **ApiCustomersV20VcdaIpsGet**
> []string ApiCustomersV20VcdaIpsGet(ctx, )
Get on premise IP addresses

Get list of on premise IP addresses allowed for this organization's draas offer

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersV20VcdaIpsIpAddressDelete**
> string ApiCustomersV20VcdaIpsIpAddressDelete(ctx, ipAddress)
Remove on premise IP address

Remove an existing on premise IP address from this organization's draas offer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAddress** | [**string**](.md)|  | 

### Return type

**string**

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersV20VcdaIpsIpAddressPost**
> string ApiCustomersV20VcdaIpsIpAddressPost(ctx, ipAddress)
Add on premise IP address

Allow a new on premise IP address for this organization's draas offer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAddress** | [**string**](.md)|  | 

### Return type

**string**

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

