/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type CloudPlatformCreateReqPlatform struct {

	// description of cloud platform
	Description string `json:"description,omitempty"`

	// extra properties of cloud platform
	ExtraProperties *CloudExtraProperties `json:"extra_properties,omitempty"`

	// name of cloud platform
	Name string `json:"name,omitempty"`

	// password of cloud platform user
	Password string `json:"password,omitempty"`

	// type of cloud platform
	Type_ string `json:"type,omitempty"`

	Url string `json:"url,omitempty"`

	Username string `json:"username,omitempty"`
}
