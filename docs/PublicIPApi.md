# {{classname}}

All URIs are relative to *https://console1.cloudavenue.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCustomersV10IpPost**](PublicIPApi.md#ApiCustomersV10IpPost) | **Post** /api/customers/v1.0/ip | Request and configure a new public IP address
[**ApiCustomersV10IpPublicIpDelete**](PublicIPApi.md#ApiCustomersV10IpPublicIpDelete) | **Delete** /api/customers/v1.0/ip/{public-ip} | Remove all configuration related to a public IP address and free IP
[**ApiCustomersV20IpGet**](PublicIPApi.md#ApiCustomersV20IpGet) | **Get** /api/customers/v2.0/ip | Get Organization&#x27;s public IP addresses

# **ApiCustomersV10IpPost**
> Jobcreated ApiCustomersV10IpPost(ctx, optional)
Request and configure a new public IP address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicIPApiApiCustomersV10IpPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicIPApiApiCustomersV10IpPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNattedIP** | **optional.String**| Target IP to configure NAT. If not provided, it configure [Double NAT](https://wiki.flexible-computing-advanced.orange-business.com/wiki/Le_r%C3%A9seau#Double_NAT) with a new IP on INET VDC Edge. If the IP or IP range provided is in 100.64.102.1-100.64.102.253 [Double NAT](https://wiki.flexible-computing-advanced.orange-business.com/wiki/Le_r%C3%A9seau#Double_NAT) is configured. If the IP is a private IP [Direct NAT](https://wiki.flexible-computing-advanced.orange-business.com/wiki/Le_r%C3%A9seau#NAT_Direct) is configured. | 
 **xVDCEdgeName** | **optional.String**| Public IP is natted toward the specified VDC Edge. This parameter overrides X-VDC-Name. | 
 **xVDCName** | **optional.String**| Public IP is natted toward the INET VDC Edge in the specified VDC Name. This parameter helps to find target VDC Edge in case of multiples INET VDC Edges with same names. | 

### Return type

[**Jobcreated**](jobcreated.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersV10IpPublicIpDelete**
> Jobcreated ApiCustomersV10IpPublicIpDelete(ctx, publicIp)
Remove all configuration related to a public IP address and free IP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publicIp** | **string**| Public IP address to free | 

### Return type

[**Jobcreated**](jobcreated.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersV20IpGet**
> PublicIps ApiCustomersV20IpGet(ctx, )
Get Organization's public IP addresses

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PublicIps**](publicIps.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

