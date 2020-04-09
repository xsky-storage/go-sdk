# \TrashResourcesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTrashResource**](TrashResourcesApi.md#DeleteTrashResource) | **Delete** /trash-resources/{trash_resource_id} | 
[**GetTrashResource**](TrashResourcesApi.md#GetTrashResource) | **Get** /trash-resources/{trash_resource_id} | 
[**ListTrashResources**](TrashResourcesApi.md#ListTrashResources) | **Get** /trash-resources/ | 
[**RestoreTrashResource**](TrashResourcesApi.md#RestoreTrashResource) | **Post** /trash-resources/{trash_resource_id}:restore | 


# **DeleteTrashResource**
> TrashResourceResp DeleteTrashResource(ctx, trashResourceId)


Delete trash resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **trashResourceId** | **int64**| trash resource id | 

### Return type

[**TrashResourceResp**](TrashResourceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTrashResource**
> TrashResourceResp GetTrashResource(ctx, trashResourceId)


get a trash resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **trashResourceId** | **int64**| trash resource id | 

### Return type

[**TrashResourceResp**](TrashResourceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTrashResources**
> TrashResourcesResp ListTrashResources(ctx, optional)


List trash resources

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
 **type_** | **string**| the type of trash resources | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**TrashResourcesResp**](TrashResourcesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreTrashResource**
> TrashResourceResp RestoreTrashResource(ctx, trashResourceId)


Restore trash resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **trashResourceId** | **int64**| trash resource id | 

### Return type

[**TrashResourceResp**](TrashResourceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

