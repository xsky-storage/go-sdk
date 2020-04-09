# \OsSamplesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSSamples**](OsSamplesApi.md#GetOSSamples) | **Get** /os-samples/ | 
[**GetOSSamplesByBucketName**](OsSamplesApi.md#GetOSSamplesByBucketName) | **Get** /os-samples/{user_name}/{bucket_name} | 
[**GetOSSamplesByUserName**](OsSamplesApi.md#GetOSSamplesByUserName) | **Get** /os-samples/{user_name} | 


# **GetOSSamples**
> OsSampleResp GetOSSamples(ctx, optional)


Get os samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time** | **string**| query time(url encode RFC3339) | 

### Return type

[**OsSampleResp**](OSSampleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOSSamplesByBucketName**
> OsSampleResp GetOSSamplesByBucketName(ctx, userName, bucketName, optional)


get os samples by os bucket name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **userName** | **string**| os user name | 
  **bucketName** | **string**| os bucket name | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string**| os user name | 
 **bucketName** | **string**| os bucket name | 
 **time** | **string**| query time(url encode RFC3339) | 

### Return type

[**OsSampleResp**](OSSampleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOSSamplesByUserName**
> OsSampleResp GetOSSamplesByUserName(ctx, userName, optional)


Get os samples by user name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **userName** | **string**| os user name | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string**| os user name | 
 **time** | **string**| query time(url encode RFC3339) | 

### Return type

[**OsSampleResp**](OSSampleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

