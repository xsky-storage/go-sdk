# \OsZoneLocksApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOSZoneLock**](OsZoneLocksApi.md#CreateOSZoneLock) | **Post** /os-zone-locks/ | 
[**DeleteOSZoneLock**](OsZoneLocksApi.md#DeleteOSZoneLock) | **Delete** /os-zone-locks/{lock_uuid} | 
[**GetOSZoneLock**](OsZoneLocksApi.md#GetOSZoneLock) | **Get** /os-zone-locks/{lock_uuid} | 
[**ListOSZoneLocks**](OsZoneLocksApi.md#ListOSZoneLocks) | **Get** /os-zone-locks/ | 
[**RefreshOSZoneLock**](OsZoneLocksApi.md#RefreshOSZoneLock) | **Post** /os-zone-locks/{lock_uuid}:refresh | 


# **CreateOSZoneLock**
> OsZoneLockResp CreateOSZoneLock(ctx, body)


Create a os zone lock.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**OsZoneLockCreateReq**](OsZoneLockCreateReq.md)| os zone lock info | 

### Return type

[**OsZoneLockResp**](OSZoneLockResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOSZoneLock**
> DeleteOSZoneLock(ctx, lockUuid)


Delete a os zone lock.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lockUuid** | **string**| os zone lock uuid | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOSZoneLock**
> OsZoneLockResp GetOSZoneLock(ctx, lockUuid)


Get a os zone lock.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lockUuid** | **string**| os zone lock uuid | 

### Return type

[**OsZoneLockResp**](OSZoneLockResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSZoneLocks**
> OsZoneLocksResp ListOSZoneLocks(ctx, optional)


List os zone locks.

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
 **all** | **bool**| show all zone locks | 

### Return type

[**OsZoneLocksResp**](OSZoneLocksResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshOSZoneLock**
> OsZoneLockResp RefreshOSZoneLock(ctx, lockUuid)


Refresh a os zone lock.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lockUuid** | **string**| os zone lock uuid | 

### Return type

[**OsZoneLockResp**](OSZoneLockResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

