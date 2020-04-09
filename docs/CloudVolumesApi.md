# \CloudVolumesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudVolume**](CloudVolumesApi.md#GetCloudVolume) | **Get** /cloud-volumes/{cloud_volume_id} | 
[**ListCloudVolumes**](CloudVolumesApi.md#ListCloudVolumes) | **Get** /cloud-volumes/ | 


# **GetCloudVolume**
> CloudVolumeResp GetCloudVolume(ctx, cloudVolumeId)


Get a cloud volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudVolumeId** | **int64**| cloud volume id | 

### Return type

[**CloudVolumeResp**](CloudVolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCloudVolumes**
> CloudVolumesResp ListCloudVolumes(ctx, optional)


List cloud volumes

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
 **cloudInstanceId** | **int64**| cloud instance id | 

### Return type

[**CloudVolumesResp**](CloudVolumesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

