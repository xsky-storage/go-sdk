# \OsStorageClassesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSStorageClass**](OsStorageClassesApi.md#GetOSStorageClass) | **Get** /os-storage-classes/{storage_class_id} | 
[**ListOSStorageClasses**](OsStorageClassesApi.md#ListOSStorageClasses) | **Get** /os-storage-classes/ | 
[**UpdateOSStorageClass**](OsStorageClassesApi.md#UpdateOSStorageClass) | **Patch** /os-storage-classes/{storage_class_id} | 


# **GetOSStorageClass**
> OsStorageClassResp GetOSStorageClass(ctx, storageClassId)


Get a os storage class

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storageClassId** | **string**| storage class id | 

### Return type

[**OsStorageClassResp**](OSStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSStorageClasses**
> OsStorageClassesResp ListOSStorageClasses(ctx, optional)


List os storage classes

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
 **osPolicyId** | **int64**| os policy id | 

### Return type

[**OsStorageClassesResp**](OSStorageClassesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOSStorageClass**
> OsStorageClassResp UpdateOSStorageClass(ctx, storageClassId, body)


update storage class

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storageClassId** | **int64**| storage class id | 
  **body** | [**OsStorageClassUpdateReq**](OsStorageClassUpdateReq.md)| storage class info | 

### Return type

[**OsStorageClassResp**](OSStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

