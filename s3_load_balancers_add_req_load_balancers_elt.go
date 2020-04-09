/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type S3LoadBalancersAddReqLoadBalancersElt struct {

	// host of load balancer
	HostId int64 `json:"host_id"`

	// vip will be bound to interface, exclusive to ip
	InterfaceName string `json:"interface_name,omitempty"`

	// vip will be bound to interface of the gateway ip, exclusive to interface_name
	Ip string `json:"ip,omitempty"`

	// vip of load balancer
	Vip string `json:"vip"`
}
