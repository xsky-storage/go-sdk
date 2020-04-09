# \CloudPlatformsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudPlatform**](CloudPlatformsApi.md#CreateCloudPlatform) | **Post** /cloud-platforms/ | 
[**DeleteCloudPlatform**](CloudPlatformsApi.md#DeleteCloudPlatform) | **Delete** /cloud-platforms/{cloud_platform_id} | 
[**GetCloudPlatform**](CloudPlatformsApi.md#GetCloudPlatform) | **Get** /cloud-platforms/{cloud_platform_id} | 
[**ListCloudPlatforms**](CloudPlatformsApi.md#ListCloudPlatforms) | **Get** /cloud-platforms/ | 
[**SyncCloudPlatform**](CloudPlatformsApi.md#SyncCloudPlatform) | **Post** /cloud-platforms/{cloud_platform_id}:sync | 
[**UpdateCloudPlatform**](CloudPlatformsApi.md#UpdateCloudPlatform) | **Patch** /cloud-platforms/{cloud_platform_id} | 


# **CreateCloudPlatform**
> CloudPlatformResp CreateCloudPlatform(ctx, body)


Create cloud platform

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**CloudPlatformCreateReq**](CloudPlatformCreateReq.md)| cloud platform info | 

### Return type

[**CloudPlatformResp**](CloudPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudPlatform**
> DeleteCloudPlatform(ctx, cloudPlatformId)


Delete cloud platform

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudPlatformId** | **int64**| cloud platform id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudPlatform**
> CloudPlatformResp GetCloudPlatform(ctx, cloudPlatformId)


Get a cloud platform

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudPlatformId** | **int64**| cloud platform id | 

### Return type

[**CloudPlatformResp**](CloudPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCloudPlatforms**
> CloudPlatformsResp ListCloudPlatforms(ctx, optional)


List cloud platforms

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

[**CloudPlatformsResp**](CloudPlatformsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncCloudPlatform**
> CloudPlatformResp SyncCloudPlatform(ctx, cloudPlatformId)


Sync a cloud platform

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudPlatformId** | **int64**| cloud platform id | 

### Return type

[**CloudPlatformResp**](CloudPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCloudPlatform**
> CloudPlatformResp UpdateCloudPlatform(ctx, cloudPlatformId, body)


Update a cloud platform

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudPlatformId** | **int64**| cloud platform id | 
  **body** | [**CloudPlatformUpdateReq**](CloudPlatformUpdateReq.md)| cloud platform info | 

### Return type

[**CloudPlatformResp**](CloudPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

