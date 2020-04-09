# S3LoadBalancerGroupCreateReqGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | group description | [optional] [default to null]
**HttpsPort** | **int64** | group access https port | [optional] [default to null]
**Name** | **string** | group name | [default to null]
**Port** | **int64** | group access http port | [optional] [default to null]
**Roles** | **[]string** | group roles | [optional] [default to null]
**S3LoadBalancers** | [**[]S3LoadBalancerGroupCreateReqGroupLoadBalancersElt**](S3LoadBalancerGroupCreateReq_Group_LoadBalancers_Elt.md) | s3 load balancers | [default to null]
**SearchHttpsPort** | **int64** | group search https port | [optional] [default to null]
**SearchPort** | **int64** | group search http port | [optional] [default to null]
**SyncPort** | **int64** | group sync http port | [optional] [default to null]
**WebConsoleEndpoint** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


