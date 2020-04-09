# \TrashesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CleanupTrash**](TrashesApi.md#CleanupTrash) | **Post** /trashes/{trash_id}:cleanup | 
[**GetTrash**](TrashesApi.md#GetTrash) | **Get** /trashes/{trash_id} | 
[**ListTrashes**](TrashesApi.md#ListTrashes) | **Get** /trashes/ | 
[**UpdateTrash**](TrashesApi.md#UpdateTrash) | **Patch** /trashes/{trash_id} | 


# **CleanupTrash**
> TrashResp CleanupTrash(ctx, trashId)


Cleanup trash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **trashId** | **int64**| trash id | 

### Return type

[**TrashResp**](TrashResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTrash**
> TrashResp GetTrash(ctx, trashId)


get a trash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **trashId** | **int64**| trash id | 

### Return type

[**TrashResp**](TrashResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTrashes**
> TrashesResp ListTrashes(ctx, optional)


List trashes

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
 **type_** | **string**| the type of trash | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**TrashesResp**](TrashesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTrash**
> TrashResp UpdateTrash(ctx, trashId, body)


Update trash info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **trashId** | **int64**| trash id | 
  **body** | [**TrashUpdateReq**](TrashUpdateReq.md)| trash info | 

### Return type

[**TrashResp**](TrashResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

