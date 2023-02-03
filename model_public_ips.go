/*
 * Cloud Avenue API
 *
 * Cloud Avenue API allows you to create some ressources on vCloud Director. ## Authentication In order to authenticate to this API, you must first authenticate through vCloudDirector API using **POST https://{{api_host}}/cloudapi/1.0.0/sessions**. Then use _X-VMWARE-VCLOUD-ACCESS-TOKEN_ from response header in your request.  If you need more information do not hesitate to read our dedicated page on [our Wiki](https://wiki.cloudavenue.orange-business.com/w/index.php/API_Cloud_Avenue) 
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PublicIps struct {
	InternalIp string `json:"internalIp,omitempty"`
	NetworkConfig *PublicIpsNetworkConfig `json:"networkConfig,omitempty"`
}
