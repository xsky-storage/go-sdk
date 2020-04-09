# \S3LoadBalancerGroupsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddS3LoadBalancersToGroup**](S3LoadBalancerGroupsApi.md#AddS3LoadBalancersToGroup) | **Put** /s3-load-balancer-groups/{group_id}/s3-load-balancers | 
[**CreateS3LoadBalancerGroup**](S3LoadBalancerGroupsApi.md#CreateS3LoadBalancerGroup) | **Post** /s3-load-balancer-groups/ | 
[**DeleteS3LoadBalancerGroup**](S3LoadBalancerGroupsApi.md#DeleteS3LoadBalancerGroup) | **Delete** /s3-load-balancer-groups/{group_id} | 
[**GetS3LoadBalancerGroup**](S3LoadBalancerGroupsApi.md#GetS3LoadBalancerGroup) | **Get** /s3-load-balancer-groups/{group_id} | 
[**ListS3LoadBalancerGroups**](S3LoadBalancerGroupsApi.md#ListS3LoadBalancerGroups) | **Get** /s3-load-balancer-groups/ | 
[**RedeployS3LoadBalancerGroup**](S3LoadBalancerGroupsApi.md#RedeployS3LoadBalancerGroup) | **Post** /s3-load-balancer-groups/{group_id}:redeploy | 
[**RemoveS3LoadBalancersFromGroup**](S3LoadBalancerGroupsApi.md#RemoveS3LoadBalancersFromGroup) | **Delete** /s3-load-balancer-groups/{group_id}/s3-load-balancers | 
[**UpdateS3LoadBalancerGroup**](S3LoadBalancerGroupsApi.md#UpdateS3LoadBalancerGroup) | **Patch** /s3-load-balancer-groups/{group_id} | 


# **AddS3LoadBalancersToGroup**
> S3LoadBalancerGroupResp AddS3LoadBalancersToGroup(ctx, groupId, body)


add load balancers to group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| group id | 
  **body** | [**S3LoadBalancersAddReq**](S3LoadBalancersAddReq.md)| load balancer ids to add | 

### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateS3LoadBalancerGroup**
> S3LoadBalancerGroupResp CreateS3LoadBalancerGroup(ctx, body)


Create a s3 load balancer group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**S3LoadBalancerGroupCreateReq**](S3LoadBalancerGroupCreateReq.md)| s3 load balancer group info | 

### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteS3LoadBalancerGroup**
> S3LoadBalancerGroupResp DeleteS3LoadBalancerGroup(ctx, groupId, optional)


Delete s3 load balancer group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| s3 load balancer group id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int64**| s3 load balancer group id | 
 **force** | **bool**| forcedly delete or not | 

### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetS3LoadBalancerGroup**
> S3LoadBalancerGroupResp GetS3LoadBalancerGroup(ctx, groupId)


Get s3 load balancer group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| s3 load balancer group id | 

### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListS3LoadBalancerGroups**
> S3LoadBalancerGroupsResp ListS3LoadBalancerGroups(ctx, optional)


List s3 load balancer groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64**| paging param | 
 **offset** | **int64**| paging param | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**S3LoadBalancerGroupsResp**](S3LoadBalancerGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RedeployS3LoadBalancerGroup**
> S3LoadBalancerGroupResp RedeployS3LoadBalancerGroup(ctx, groupId)


Redeploy a s3 load balancer group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| s3 load balancer group id | 

### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveS3LoadBalancersFromGroup**
> S3LoadBalancerGroupResp RemoveS3LoadBalancersFromGroup(ctx, groupId, body, optional)


remove load balancers from group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| group id | 
  **body** | [**S3LoadBalancersRemoveReq**](S3LoadBalancersRemoveReq.md)| load balancer ids to remove | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int64**| group id | 
 **body** | [**S3LoadBalancersRemoveReq**](S3LoadBalancersRemoveReq.md)| load balancer ids to remove | 
 **force** | **bool**| forcedly remove or not | 

### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateS3LoadBalancerGroup**
> S3LoadBalancerGroupResp UpdateS3LoadBalancerGroup(ctx, groupId, body)


Update a s3 load balancer group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| s3 load balancer group id | 
  **body** | [**S3LoadBalancerGroupUpdateReq**](S3LoadBalancerGroupUpdateReq.md)| s3 load balancer group info | 

### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

