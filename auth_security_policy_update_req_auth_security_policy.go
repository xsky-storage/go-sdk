/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type AuthSecurityPolicyUpdateReqAuthSecurityPolicy struct {

	Enabled bool `json:"enabled,omitempty"`

	FreezeDuration int64 `json:"freeze_duration,omitempty"`

	MaxAttempts int64 `json:"max_attempts,omitempty"`

	PasswordExpirationWarning bool `json:"password_expiration_warning,omitempty"`

	PasswordLifetime int64 `json:"password_lifetime,omitempty"`
}
