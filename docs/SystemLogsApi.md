# \SystemLogsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadSystemLogs**](SystemLogsApi.md#DownloadSystemLogs) | **Get** /system-logs/archive | 
[**ListSystemLogs**](SystemLogsApi.md#ListSystemLogs) | **Get** /system-logs/ | 


# **DownloadSystemLogs**
> string DownloadSystemLogs(ctx, )


download system logs

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSystemLogs**
> SystemLogsResp ListSystemLogs(ctx, hostId, catalog, optional)


List system logs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hostId** | **int64**| The id of host system logs belong to | 
  **catalog** | **string**| The name of catalog system logs belong to | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostId** | **int64**| The id of host system logs belong to | 
 **catalog** | **string**| The name of catalog system logs belong to | 
 **limit** | **int64**| paging param | 
 **offset** | **int64**| paging param | 

### Return type

[**SystemLogsResp**](SystemLogsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

