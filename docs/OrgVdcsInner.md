# OrgVdcsInner

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VMware unique ID. There is no ID in the body of the POST request, as VMware which generates it. This ID must be present for GET and PUT. | [optional] [default to null]
**Name** | **string** | Name of the new org vdc. | [default to null]
**Description** | **string** | Description of the new org vdc. | [optional] [default to null]
**IsEnabled** | **bool** |  | [optional] [default to true]
**VdcServiceClass** | **string** | SILVER, GOLD and PLATINUM_3K storage types are available for all classes. PLATINUM_5K and PLATINUM_10K are available only for STD (Standard), HP (High Performance) and VOIP (VoIP). | [default to null]
**VdcDisponibilityClass** | **string** | Allocation model used by this vdc. DUAL-ROOM isn&#x27;t available for vdcServiceClass ECO. | [default to null]
**VdcBillingModel** | **string** | Choose Billing model of storage resources. PAYG and DRAAS only available for vdcServiceClass ECO and STD. | [default to null]
**VdcStorageBillingModel** | **string** | Available for all vdcServiceClass. | [default to null]
**NicQuota** | **float64** | Maximum number of virtual NICs allowed in this vdc. Defaults to 0, which specifies an unlimited number. | [optional] [default to 0]
**NetworkQuota** | **float64** | Maximum number of network objects that can be deployed in this vdc. Defaults to 0, which means no networks can be deployed. | [optional] [default to 0]
**VmQuota** | **float64** | Maximum number of VMs that can be created in this vdc. Defaults to 0, which specifies an unlimited number. | [optional] [default to 0]
**VcpuInMhz2** | **float64** | Specifies the clock frequency, in Mhz, for any virtual CPU that is allocated to a VM. | [optional] [default to 2300]
**CpuAllocated** | **float64** | Capacity that is committed to be available or used as a limit in PAYG mode. Unit for compute capacity allocated to this vdc is MHz. Note. Reserved capacity is automatically set according to the service class. | [default to null]
**MemoryAllocated** | **float64** | Memory capacity that is committed to be available or used as a limit in PAYG mode. Unit for memory capacity allocated to this vdc is Gb. | [default to null]
**VdcStorageProfiles** | [***[]StorageProfile**](array.md) |  | [default to null]
**VdcEdges** | [***[]VdcEdge**](array.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

