# StorageProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VMware unique ID. There is no ID in the body of the POST request, as VMware which generates it. This ID must be present for GET and PUT. | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Limit** | **int32** | Max number of units allocated for this storage profile. In Gb | [default to 2000]
**Used** | **int32** | Current usage of this StorageProfile in Gb. Deletion is allowed only if there is no usage of the StorageProfile | [optional] [default to null]
**Default_** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

