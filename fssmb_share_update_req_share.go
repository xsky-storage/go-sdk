/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type FssmbShareUpdateReqShare struct {

	// default acl status
	AclInherited bool `json:"acl_inherited,omitempty"`

	// case sensitive
	CaseSensitive bool `json:"case_sensitive,omitempty"`

	// recycle status
	Recycled bool `json:"recycled,omitempty"`
}
