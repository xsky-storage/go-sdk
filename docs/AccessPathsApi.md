# \AccessPathsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessPath**](AccessPathsApi.md#CreateAccessPath) | **Post** /access-paths/ | 
[**DeleteAccessPath**](AccessPathsApi.md#DeleteAccessPath) | **Delete** /access-paths/{access_path_id} | 
[**GetAccessPath**](AccessPathsApi.md#GetAccessPath) | **Get** /access-paths/{access_path_id} | 
[**ListAccessPaths**](AccessPathsApi.md#ListAccessPaths) | **Get** /access-paths/ | 
[**UpdateAccessPath**](AccessPathsApi.md#UpdateAccessPath) | **Patch** /access-paths/{access_path_id} | 


# **CreateAccessPath**
> AccessPathResp CreateAccessPath(ctx, accessPath)


Create an access path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accessPath** | [**AccessPathCreateReq**](AccessPathCreateReq.md)| access path info | 

### Return type

[**AccessPathResp**](AccessPathResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccessPath**
> AccessPathResp DeleteAccessPath(ctx, accessPathId)


Delete access path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accessPathId** | **int64**| access path id | 

### Return type

[**AccessPathResp**](AccessPathResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessPath**
> AccessPathResp GetAccessPath(ctx, accessPathId)


Get access path by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accessPathId** | **int64**| access path id | 

### Return type

[**AccessPathResp**](AccessPathResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccessPaths**
> AccessPathsResp ListAccessPaths(ctx, optional)


List access paths

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
 **clientGroupId** | **int64**| related client group id | 
 **all** | **bool**| show all access targets | 

### Return type

[**AccessPathsResp**](AccessPathsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccessPath**
> AccessPathResp UpdateAccessPath(ctx, accessPathId, accessPath)


update access path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accessPathId** | **int64**| access path id | 
  **accessPath** | [**AccessPathUpdateReq**](AccessPathUpdateReq.md)| access path info | 

### Return type

[**AccessPathResp**](AccessPathResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

