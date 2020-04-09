# \OsZoneTranslogsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSZoneTranslog**](OsZoneTranslogsApi.md#GetOSZoneTranslog) | **Get** /os-zone-translogs/{translog_uuid} | 
[**ListOSZoneTranslogs**](OsZoneTranslogsApi.md#ListOSZoneTranslogs) | **Get** /os-zone-translogs/ | 


# **GetOSZoneTranslog**
> OsZoneTranslogResp GetOSZoneTranslog(ctx, translogUuid)


Get a os zone translog

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **translogUuid** | **string**| translog uuid | 

### Return type

[**OsZoneTranslogResp**](OSZoneTranslogResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSZoneTranslogs**
> OsZoneTranslogsResp ListOSZoneTranslogs(ctx, optional)


List os zone translogs

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
 **epochUuid** | **string**| paging param | 

### Return type

[**OsZoneTranslogsResp**](OSZoneTranslogsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

