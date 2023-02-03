# AllOfcreateOrg2VdcsDualVdcsVdc1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the new org vdc. | [default to null]
**Description** | **string** | Description of the new org vdc. | [optional] [default to null]
**VdcServiceClass** | **string** |  | [default to null]
**VdcDisponibilityClass** | **string** |  | [default to null]
**VdcBillingModel** | **string** | Choose Billing model of compute resources. | [default to null]
**VcpuInMhz2** | **float64** | Specifies the clock frequency, in Mhz, for any virtual CPU that is allocated to a VM. | [default to 2200]
**CpuAllocated** | **float64** | Capacity that is committed to be available or used as a limit in PAYG mode. Unit for compute capacity allocated to this vdc is MHz. Note. Reserved capacity is automatically set according to the service class. | [default to null]
**MemoryAllocated** | **float64** | Memory capacity that is committed to be available or used as a limit in PAYG mode. Unit for memory capacity allocated to this vdc is Gb. | [default to null]
**VdcStorageBillingModel** | **string** | Available for all vdcServiceClass. | [default to null]
**VdcStorageProfiles** | [**[]VdcStorageProfilesV2**](vdcStorageProfilesV2.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

