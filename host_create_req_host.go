/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type HostCreateReqHost struct {

	// admin ip
	AdminIp string `json:"admin_ip"`

	// host description
	Description string `json:"description,omitempty"`

	// gateway ips for s3
	GatewayIps []string `json:"gateway_ips,omitempty"`

	// meta device for docker
	MetaDevice string `json:"meta_device,omitempty"`

	// cluster private ip for internal data access
	PrivateIp string `json:"private_ip,omitempty"`

	// deprecated
	ProtectionDomainId int64 `json:"protection_domain_id,omitempty"`

	// public ip for outside data access
	PublicIp string `json:"public_ip,omitempty"`

	// host roles: admin,monitor,block_storage_gateway,file_storage_gateway,s3_gateway,nfs_gateway
	Roles []string `json:"roles,omitempty"`

	// storage server or storage client
	Type_ string `json:"type,omitempty"`
}
