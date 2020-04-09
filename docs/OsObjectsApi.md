# \OsObjectsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOSObjectsBySQL**](OsObjectsApi.md#ListOSObjectsBySQL) | **Get** /os-objects/_sql | 
[**ListOSObjectsBySQL_0**](OsObjectsApi.md#ListOSObjectsBySQL_0) | **Post** /os-objects/_sql | 
[**ListOSObjectsBySearch**](OsObjectsApi.md#ListOSObjectsBySearch) | **Get** /os-objects/_search | 
[**ListOSObjectsBySearch_0**](OsObjectsApi.md#ListOSObjectsBySearch_0) | **Post** /os-objects/_search | 
[**ReportOSObjectsBySQL**](OsObjectsApi.md#ReportOSObjectsBySQL) | **Get** /os-objects/report/_sql | 
[**ReportOSObjectsBySQL_0**](OsObjectsApi.md#ReportOSObjectsBySQL_0) | **Post** /os-objects/report/_sql | 


# **ListOSObjectsBySQL**
> ListOSObjectsBySQL(ctx, )


List object storage objects by sql

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSObjectsBySQL_0**
> ListOSObjectsBySQL_0(ctx, )


List object storage objects by sql

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSObjectsBySearch**
> ListOSObjectsBySearch(ctx, )


List object storage objects by search

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSObjectsBySearch_0**
> ListOSObjectsBySearch_0(ctx, )


List object storage objects by search

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReportOSObjectsBySQL**
> string ReportOSObjectsBySQL(ctx, optional)


Download object storage objects report by sql

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sql** | **string**| select statement | 
 **osBuckets** | **string**| name of buckets joined by colon | 

### Return type

**string**

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReportOSObjectsBySQL_0**
> string ReportOSObjectsBySQL_0(ctx, optional)


Download object storage objects report by sql

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sql** | **string**| select statement | 
 **osBuckets** | **string**| name of buckets joined by colon | 

### Return type

**string**

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

