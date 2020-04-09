# \OsExternalStorageClassesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalStorageClass**](OsExternalStorageClassesApi.md#CreateExternalStorageClass) | **Post** /os-external-storage-classes/ | 
[**DeleteOSExternalStorageClass**](OsExternalStorageClassesApi.md#DeleteOSExternalStorageClass) | **Delete** /os-external-storage-classes/{external_storage_class_id} | 
[**GetOSExternalStorageClass**](OsExternalStorageClassesApi.md#GetOSExternalStorageClass) | **Get** /os-external-storage-classes/{external_storage_class_id} | 
[**ListOSExternalStorageClasses**](OsExternalStorageClassesApi.md#ListOSExternalStorageClasses) | **Get** /os-external-storage-classes/ | 
[**UpdateOSExternalStorageClass**](OsExternalStorageClassesApi.md#UpdateOSExternalStorageClass) | **Patch** /os-external-storage-classes/{external_storage_class_id} | 


# **CreateExternalStorageClass**
> OsExternalStorageClassResp CreateExternalStorageClass(ctx, body)


Create os external storage class

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**OsExternalStorageClassCreateReq**](OsExternalStorageClassCreateReq.md)| external storage class info | 

### Return type

[**OsExternalStorageClassResp**](OSExternalStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOSExternalStorageClass**
> OsExternalStorageClassResp DeleteOSExternalStorageClass(ctx, externalStorageClassId)


Delete an os external storage class

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **externalStorageClassId** | **int64**| external storage class id | 

### Return type

[**OsExternalStorageClassResp**](OSExternalStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOSExternalStorageClass**
> OsExternalStorageClassResp GetOSExternalStorageClass(ctx, externalStorageClassId)


Get a os external storage class

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **externalStorageClassId** | **int64**| external storage class id | 

### Return type

[**OsExternalStorageClassResp**](OSExternalStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSExternalStorageClasses**
> OsExternalStorageClassesResp ListOSExternalStorageClasses(ctx, optional)


List os external storage classes

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
 **generalStatus** | **bool**| query records with active or error status | 

### Return type

[**OsExternalStorageClassesResp**](OSExternalStorageClassesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOSExternalStorageClass**
> OsExternalStorageClassResp UpdateOSExternalStorageClass(ctx, externalStorageClassId, body)


update os external storage class

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **externalStorageClassId** | **int64**| external storage class id | 
  **body** | [**OsExternalStorageClassUpdateReq**](OsExternalStorageClassUpdateReq.md)| external storage class info | 

### Return type

[**OsExternalStorageClassResp**](OSExternalStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

