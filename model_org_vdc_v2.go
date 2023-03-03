/*
 * Cloud Avenue API
 *
 * Cloud Avenue API allows you to create some ressources on vCloud Director.  ## Authentication In order to authenticate to this API, you must first authenticate through vCloudDirector API using **POST https://{{api_host}}/cloudapi/1.0.0/sessions**. Then use _X-VMWARE-VCLOUD-ACCESS-TOKEN_ from response header in your request.  If you need more information do not hesitate to read our dedicated page on [our Wiki](https://wiki.cloudavenue.orange-business.com/w/index.php/API_Cloud_Avenue) 
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cloudavenue

type OrgVdcV2 struct {
	// Name of the new org vdc.
	Name string `json:"name"`
	// Description of the new org vdc.
	Description string `json:"description,omitempty"`
	VdcServiceClass string `json:"vdcServiceClass"`
	VdcDisponibilityClass string `json:"vdcDisponibilityClass"`
	// Choose Billing model of compute resources.
	VdcBillingModel string `json:"vdcBillingModel"`
	// Specifies the clock frequency, in Mhz, for any virtual CPU that is allocated to a VM.
	VcpuInMhz2 float64 `json:"vcpuInMhz2"`
	// Capacity that is committed to be available or used as a limit in PAYG mode. Unit for compute capacity allocated to this vdc is MHz. Note. Reserved capacity is automatically set according to the service class.
	CpuAllocated float64 `json:"cpuAllocated"`
	// Memory capacity that is committed to be available or used as a limit in PAYG mode. Unit for memory capacity allocated to this vdc is Gb.
	MemoryAllocated float64 `json:"memoryAllocated"`
	// Available for all vdcServiceClass.
	VdcStorageBillingModel string `json:"vdcStorageBillingModel"`
	VdcStorageProfiles []VdcStorageProfilesV2 `json:"vdcStorageProfiles"`
}
