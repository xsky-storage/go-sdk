/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

// CloudExtraProperties contains extra properties
type CloudExtraProperties struct {

	// domain id for openstack
	DomainId string `json:"domain_id,omitempty"`

	// domain name for openstack
	DomainName string `json:"domain_name,omitempty"`

	// regions for openstack
	Regions []string `json:"regions,omitempty"`

	// tenant id for openstack
	TenantId string `json:"tenant_id,omitempty"`

	// tenant name for openstack
	TenantName string `json:"tenant_name,omitempty"`
}
