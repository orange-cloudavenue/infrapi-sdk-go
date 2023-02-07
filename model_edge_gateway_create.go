/*
 * Cloud Avenue API
 *
 * Cloud Avenue API allows you to create some ressources on vCloud Director. ## Authentication In order to authenticate to this API, you must first authenticate through vCloudDirector API using **POST https://{{api_host}}/cloudapi/1.0.0/sessions**. Then use _X-VMWARE-VCLOUD-ACCESS-TOKEN_ from response header in your request.  If you need more information do not hesitate to read our dedicated page on [our Wiki](https://wiki.cloudavenue.orange-business.com/w/index.php/API_Cloud_Avenue) 
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cloudavenue

// Edge Gateway (NSX-T Tier-1 Gateway)
type EdgeGatewayCreate struct {
	// Identifier of Tier-0 VRF to connect the Edge Gateway
	Tier0VrfId string `json:"tier0VrfId"`
	// Enable the Load Balancing NSX AVI Feature
	EnableLoadBalancing bool `json:"enableLoadBalancing,omitempty"`
}
