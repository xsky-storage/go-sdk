/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type ChangePasswordReqReq struct {

	// encrypted original password for auth
	EncryptedOriginalPassword string `json:"encrypted_original_password,omitempty"`

	// encrypted password for auth
	EncryptedPassword string `json:"encrypted_password,omitempty"`

	// original password of user
	OriginalPassword string `json:"original_password"`

	// new password of user
	Password string `json:"password"`

	// rsa key id
	RsaKeyId string `json:"rsa_key_id,omitempty"`
}
