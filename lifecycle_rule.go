/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

// LifecycleRule is the details of object storage lifecycle rule
type LifecycleRule struct {

	Enabled bool `json:"enabled,omitempty"`

	Expiration *LifecycleExpiration `json:"expiration,omitempty"`

	Name string `json:"name,omitempty"`

	Prefix string `json:"prefix,omitempty"`

	Transitions []LifecycleTransition `json:"transitions,omitempty"`
}