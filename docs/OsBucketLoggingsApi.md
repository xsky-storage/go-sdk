# \OsBucketLoggingsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOSBucketLoggings**](OsBucketLoggingsApi.md#ListOSBucketLoggings) | **Get** /os-bucket-loggings/ | 


# **ListOSBucketLoggings**
> OsBucketLoggingsResp ListOSBucketLoggings(ctx, optional)


List os bucket loggings

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
 **sourceBucketName** | **string**| source bucket name | 
 **targetBucketName** | **string**| target bucket name | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**OsBucketLoggingsResp**](OSBucketLoggingsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

