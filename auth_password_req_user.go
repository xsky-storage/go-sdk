/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type AuthPasswordReqUser struct {

	// user email for auth
	Email string `json:"email,omitempty"`

	// user id for auth
	Id int64 `json:"id,omitempty"`

	// user name or email for auth
	Name string `json:"name,omitempty"`

	// password for auth
	Password string `json:"password"`
}
