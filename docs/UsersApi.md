# {{classname}}

All URIs are relative to *https://console1.cloudavenue.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCustomersV10FederateGet**](UsersApi.md#ApiCustomersV10FederateGet) | **Get** /api/customers/v1.0/federate | Return federation status of an organization
[**ApiCustomersV10FederatePost**](UsersApi.md#ApiCustomersV10FederatePost) | **Post** /api/customers/v1.0/federate | Federate an existing organization
[**ApiCustomersV10FederatePut**](UsersApi.md#ApiCustomersV10FederatePut) | **Put** /api/customers/v1.0/federate | Regenerate SAML certificate

# **ApiCustomersV10FederateGet**
> FederationStatusType ApiCustomersV10FederateGet(ctx, )
Return federation status of an organization

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FederationStatusType**](federation-status-type.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersV10FederatePost**
> Jobcreated ApiCustomersV10FederatePost(ctx, )
Federate an existing organization

Federate an existing organization

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Jobcreated**](jobcreated.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersV10FederatePut**
> Jobcreated ApiCustomersV10FederatePut(ctx, )
Regenerate SAML certificate

Regenerate SAML certificate of an organization and updates the LemonLDAP entry with the new certificate.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Jobcreated**](jobcreated.md)

### Authorization

[X-VMWARE-VCLOUD-ACCESS-TOKEN](../README.md#X-VMWARE-VCLOUD-ACCESS-TOKEN)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

