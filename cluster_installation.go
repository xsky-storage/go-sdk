/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

// ClusterInstallation returns installation status of the cluster
type ClusterInstallation struct {

	Finished bool `json:"finished,omitempty"`

	Steps *ClusterInstallationSteps `json:"steps,omitempty"`
}
