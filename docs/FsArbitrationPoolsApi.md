# \FsArbitrationPoolsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFSArbitrationPool**](FsArbitrationPoolsApi.md#CreateFSArbitrationPool) | **Post** /fs-arbitration-pools/ | 
[**DeleteFSArbitrationPool**](FsArbitrationPoolsApi.md#DeleteFSArbitrationPool) | **Delete** /fs-arbitration-pools/{fs_arbitration_pool_id} | 
[**GetFSArbitrationPool**](FsArbitrationPoolsApi.md#GetFSArbitrationPool) | **Get** /fs-arbitration-pools/{fs_arbitration_pool_id} | 
[**ListFSArbitrationPools**](FsArbitrationPoolsApi.md#ListFSArbitrationPools) | **Get** /fs-arbitration-pools/ | 
[**UpdateFSArbitrationPool**](FsArbitrationPoolsApi.md#UpdateFSArbitrationPool) | **Patch** /fs-arbitration-pools/{fs_arbitration_pool_id} | 


# **CreateFSArbitrationPool**
> FsArbitrationPoolResp CreateFSArbitrationPool(ctx, body)


create file storage arbitration pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FsArbitrationPoolCreateReq**](FsArbitrationPoolCreateReq.md)| file system arbitration pool info | 

### Return type

[**FsArbitrationPoolResp**](FSArbitrationPoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFSArbitrationPool**
> DeleteFSArbitrationPool(ctx, fsArbitrationPoolId)


Delete file storage arbitration pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsArbitrationPoolId** | **int64**| file storage arbitration pool id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFSArbitrationPool**
> FsArbitrationPoolResp GetFSArbitrationPool(ctx, fsArbitrationPoolId)


Get file storage arbitration pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsArbitrationPoolId** | **int64**| file storage arbitration pool id | 

### Return type

[**FsArbitrationPoolResp**](FSArbitrationPoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFSArbitrationPools**
> FsArbitrationPoolsResp ListFSArbitrationPools(ctx, optional)


List file storage arbitration pools

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

[**FsArbitrationPoolsResp**](FSArbitrationPoolsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFSArbitrationPool**
> FsArbitrationPoolResp UpdateFSArbitrationPool(ctx, fsArbitrationPoolId, body)


Update file system arbitration pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsArbitrationPoolId** | **int64**| file system arbitration pool id | 
  **body** | [**FsArbitrationPoolUpdateReq**](FsArbitrationPoolUpdateReq.md)| file system arbitration pool info | 

### Return type

[**FsArbitrationPoolResp**](FSArbitrationPoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

