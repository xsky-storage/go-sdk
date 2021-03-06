/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type S3LoadBalancerSamplesResp struct {

	// s3 load balancer samples
	S3LoadBalancerSamples []S3LoadBalancerStat `json:"s3_load_balancer_samples"`
}
