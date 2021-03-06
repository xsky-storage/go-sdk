/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type EmailSendReqEmail struct {

	Body string `json:"body"`

	// content type, 'text/plain' or 'text/html'
	ContentType string `json:"content_type,omitempty"`

	To []string `json:"to"`
}
