/*
 * Cloud Avenue API
 *
 * Cloud Avenue API allows you to create some ressources on vCloud Director.  ## Authentication In order to authenticate to this API, you must first authenticate through vCloudDirector API using **POST https://{{api_host}}/cloudapi/1.0.0/sessions**. Then use _X-VMWARE-VCLOUD-ACCESS-TOKEN_ from response header in your request.  If you need more information do not hesitate to read our dedicated page on [our Wiki](https://wiki.cloudavenue.orange-business.com/w/index.php/API_Cloud_Avenue) 
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cloudavenue

type OrgVdcsInner struct {
	// VMware unique ID. There is no ID in the body of the POST request, as VMware which generates it. This ID must be present for GET and PUT.
	Id string `json:"id,omitempty"`
	// Name of the new org vdc.
	Name string `json:"name"`
	// Description of the new org vdc.
	Description string `json:"description,omitempty"`
	IsEnabled bool `json:"isEnabled,omitempty"`
	// SILVER, GOLD and PLATINUM_3K storage types are available for all classes. PLATINUM_5K and PLATINUM_10K are available only for STD (Standard), HP (High Performance) and VOIP (VoIP).
	VdcServiceClass string `json:"vdcServiceClass"`
	// Allocation model used by this vdc. DUAL-ROOM isn't available for vdcServiceClass ECO.
	VdcDisponibilityClass string `json:"vdcDisponibilityClass"`
	// Choose Billing model of storage resources. PAYG and DRAAS only available for vdcServiceClass ECO and STD.
	VdcBillingModel string `json:"vdcBillingModel"`
	// Available for all vdcServiceClass.
	VdcStorageBillingModel string `json:"vdcStorageBillingModel"`
	// Maximum number of virtual NICs allowed in this vdc. Defaults to 0, which specifies an unlimited number.
	NicQuota float64 `json:"nicQuota,omitempty"`
	// Maximum number of network objects that can be deployed in this vdc. Defaults to 0, which means no networks can be deployed.
	NetworkQuota float64 `json:"networkQuota,omitempty"`
	// Maximum number of VMs that can be created in this vdc. Defaults to 0, which specifies an unlimited number.
	VmQuota float64 `json:"vmQuota,omitempty"`
	// Specifies the clock frequency, in Mhz, for any virtual CPU that is allocated to a VM.
	VcpuInMhz2 float64 `json:"vcpuInMhz2,omitempty"`
	// Capacity that is committed to be available or used as a limit in PAYG mode. Unit for compute capacity allocated to this vdc is MHz. Note. Reserved capacity is automatically set according to the service class.
	CpuAllocated float64 `json:"cpuAllocated"`
	// Memory capacity that is committed to be available or used as a limit in PAYG mode. Unit for memory capacity allocated to this vdc is Gb.
	MemoryAllocated float64 `json:"memoryAllocated"`
	VdcStorageProfiles *[]StorageProfile `json:"vdcStorageProfiles"`
	VdcEdges *[]VdcEdge `json:"vdcEdges,omitempty"`
}
