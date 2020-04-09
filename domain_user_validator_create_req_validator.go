/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type DomainUserValidatorCreateReqValidator struct {

	// Domain groups
	Groups []string `json:"groups,omitempty"`

	// Domain type
	Type_ string `json:"type"`

	// Domain users
	Users []string `json:"users,omitempty"`
}
