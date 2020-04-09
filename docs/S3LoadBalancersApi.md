# \S3LoadBalancersApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchGetS3LoadBalancerSamples**](S3LoadBalancersApi.md#BatchGetS3LoadBalancerSamples) | **Get** /s3-load-balancers/samples | 
[**GetS3LoadBalancer**](S3LoadBalancersApi.md#GetS3LoadBalancer) | **Get** /s3-load-balancers/{load_balancer_id} | 
[**GetS3LoadBalancerSamples**](S3LoadBalancersApi.md#GetS3LoadBalancerSamples) | **Get** /s3-load-balancers/{s3_load_balancer_id}/samples | 
[**ListS3LoadBalancers**](S3LoadBalancersApi.md#ListS3LoadBalancers) | **Get** /s3-load-balancers/ | 


# **BatchGetS3LoadBalancerSamples**
> MultiS3LoadBalancersSamplesResp BatchGetS3LoadBalancerSamples(ctx, )


Get samples of multiple s3 load balancers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MultiS3LoadBalancersSamplesResp**](MultiS3LoadBalancersSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetS3LoadBalancer**
> S3LoadBalancerResp GetS3LoadBalancer(ctx, loadBalancerId)


Get s3 load balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **loadBalancerId** | **int64**| s3 load balancer id | 

### Return type

[**S3LoadBalancerResp**](S3LoadBalancerResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetS3LoadBalancerSamples**
> S3LoadBalancerSamplesResp GetS3LoadBalancerSamples(ctx, s3LoadBalancerId, optional)


get a s3 load balancer's samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **s3LoadBalancerId** | **int64**| s3 load balancer id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s3LoadBalancerId** | **int64**| s3 load balancer id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**S3LoadBalancerSamplesResp**](S3LoadBalancerSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListS3LoadBalancers**
> S3LoadBalancersResp ListS3LoadBalancers(ctx, optional)


List s3 load balancers

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
 **groupId** | **int64**| s3 load balancer group id | 

### Return type

[**S3LoadBalancersResp**](S3LoadBalancersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

