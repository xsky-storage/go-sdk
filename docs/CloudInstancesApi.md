# \CloudInstancesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudInstance**](CloudInstancesApi.md#GetCloudInstance) | **Get** /cloud-instances/{cloud_instance_id} | 
[**GetCloudInstanceSamples**](CloudInstancesApi.md#GetCloudInstanceSamples) | **Get** /cloud-instances/{cloud_instance_id}/samples | 
[**ListCloudInstances**](CloudInstancesApi.md#ListCloudInstances) | **Get** /cloud-instances/ | 


# **GetCloudInstance**
> CloudInstanceResp GetCloudInstance(ctx, cloudInstanceId)


Get a cloud instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudInstanceId** | **int64**| cloud instance id | 

### Return type

[**CloudInstanceResp**](CloudInstanceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudInstanceSamples**
> CloudInstanceSamplesResp GetCloudInstanceSamples(ctx, cloudInstanceId, optional)


get a cloud instance's Samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudInstanceId** | **int64**| cloud instance id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudInstanceId** | **int64**| cloud instance id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**CloudInstanceSamplesResp**](CloudInstanceSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCloudInstances**
> CloudInstancesResp ListCloudInstances(ctx, optional)


List cloud instances

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
 **cloudPlatformId** | **int64**| cloud platform id | 
 **cloudPlatformType** | **string**| cloud platform type | 
 **cloudPlatformName** | **string**| cloud platform name | 

### Return type

[**CloudInstancesResp**](CloudInstancesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

