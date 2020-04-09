# \OsLifecyclesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLifecycle**](OsLifecyclesApi.md#GetLifecycle) | **Get** /os-lifecycles/{lifecycle_id} | 
[**ListLifecycles**](OsLifecyclesApi.md#ListLifecycles) | **Get** /os-lifecycles/ | 


# **GetLifecycle**
> ObjectStorageLifecycleResp GetLifecycle(ctx, lifecycleId)


Get object storage lifecycle

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lifecycleId** | **int64**| lifecycle id | 

### Return type

[**ObjectStorageLifecycleResp**](ObjectStorageLifecycleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLifecycles**
> ObjectStorageLifecyclesResp ListLifecycles(ctx, optional)


List object storage lifecycles

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

### Return type

[**ObjectStorageLifecyclesResp**](ObjectStorageLifecyclesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

