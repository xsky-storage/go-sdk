# \HostEncSpecsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostEncSpec**](HostEncSpecsApi.md#CreateHostEncSpec) | **Post** /host-enc-specs/ | 
[**DeleteHostEncSpec**](HostEncSpecsApi.md#DeleteHostEncSpec) | **Delete** /host-enc-specs/{spec_id} | 
[**GetHostEncSpec**](HostEncSpecsApi.md#GetHostEncSpec) | **Get** /host-enc-specs/{spec_id} | 
[**ListHostEncSpecs**](HostEncSpecsApi.md#ListHostEncSpecs) | **Get** /host-enc-specs/ | 


# **CreateHostEncSpec**
> HostEncSpecResp CreateHostEncSpec(ctx, body)


Create a host enclosure spec

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**HostEncSpecCreateReq**](HostEncSpecCreateReq.md)| host enclosure spec info | 

### Return type

[**HostEncSpecResp**](HostEncSpecResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHostEncSpec**
> DeleteHostEncSpec(ctx, specId)


Delete a host enclosure spec

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **specId** | **int64**| host enclosure spec id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHostEncSpec**
> HostEncSpecResp GetHostEncSpec(ctx, specId)


Get host enclosure spec

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **specId** | **int64**| host enclosure spec id | 

### Return type

[**HostEncSpecResp**](HostEncSpecResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHostEncSpecs**
> HostEncSpecsResp ListHostEncSpecs(ctx, optional)


List host enclosure specs

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
 **model** | **string**| host model | 
 **vendor** | **string**| host vendor | 

### Return type

[**HostEncSpecsResp**](HostEncSpecsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

