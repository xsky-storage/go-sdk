# \OsReplicationPathsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSReplicationPath**](OsReplicationPathsApi.md#GetOSReplicationPath) | **Get** /os-replication-paths/{path_uuid} | 
[**ListOSReplicationPaths**](OsReplicationPathsApi.md#ListOSReplicationPaths) | **Get** /os-replication-paths/ | 


# **GetOSReplicationPath**
> OsReplicationPathResp GetOSReplicationPath(ctx, pathUuid)


Get a os replication path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pathUuid** | **string**| replication path uuid | 

### Return type

[**OsReplicationPathResp**](OSReplicationPathResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSReplicationPaths**
> OsReplicationPathsResp ListOSReplicationPaths(ctx, optional)


List os replication paths

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
 **marker** | **string**| paging param | 
 **replicationUuid** | **string**| os replication uuid | 
 **allowUnknownZone** | **bool**| allow has nil zone result(default true) | 

### Return type

[**OsReplicationPathsResp**](OSReplicationPathsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

