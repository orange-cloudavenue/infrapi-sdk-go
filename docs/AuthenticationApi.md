# {{classname}}

All URIs are relative to *https://console1.cloudavenue.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetToken**](AuthenticationApi.md#GetToken) | **Post** /cloudapi/1.0.0/sessions | Login to your organization

# **GetToken**
> string GetToken(ctx, )
Login to your organization

This login method is available thanks to vCD API. Login must be like username@organization (example : r.dupont@cav01ev01ocb0000000). Once logged in, use use _X-VMWARE-VCLOUD-ACCESS-TOKEN_ from response header and add it in X-VMWARE-VCLOUD-ACCESS-TOKEN security parameter. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[vCloudDirector](../README.md#vCloudDirector)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

