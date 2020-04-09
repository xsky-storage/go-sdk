/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type FsGatewayGroupCreateReqGatewayGroup struct {

	// description of gateway group
	Description string `json:"description,omitempty"`

	// ftp encoding format, default is utf8
	Encoding string `json:"encoding,omitempty"`

	// file storage gateways list
	FsGateways []FsGatewayReq `json:"fs_gateways"`

	// name of gateway group
	Name string `json:"name"`

	// nfs versions of nfs supported
	NfsVersions []string `json:"nfs_versions,omitempty"`

	// smb security type
	Security string `json:"security,omitempty"`

	// smb version 1.0 enabled
	Smb1Enabled bool `json:"smb1_enabled,omitempty"`

	// smb ports
	SmbPorts []int64 `json:"smb_ports,omitempty"`

	// types of supported (smb, nfs, ftp)
	Types []string `json:"types"`

	// virtual ip of gateway group
	Vip string `json:"vip"`
}