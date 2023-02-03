# Pesg

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | internal object identifier | [optional] [default to null]
**Name** | **string** | This action implies the creation of a Logical Switch and an External Network. | [optional] [default to null]
**Type_** | **string** | For external customers &#x60;ADMIN&#x60;, &#x60;INET&#x60;. &lt;br/&gt; For internal customers &#x60;PFS&#x60;, &#x60;RBCI&#x60;, &#x60;EIN&#x60;, &#x60;EXPL&#x60;, &#x60;ST&#x60;, &#x60;GINIS&#x60;, &#x60;GINEFA&#x60;, &#x60;GINCMZ&#x60;, &#x60;NFS&#x60;, &#x60;BACKUP&#x60;. | [default to null]
**IntercoIp** | **string** | Contains public-ip in case of INET PESG and the associated network mask | [optional] [default to null]
**VlanId** | **int32** | Only and required for NFS PESG | [optional] [default to null]
**InternalIp** | **string** | Internal Ip connected with associated network mmask (will be the one belonging to the external network of the organization) | [optional] [default to null]
**NetworkConfig** | [***[]NetworkConfigInner**](array.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

