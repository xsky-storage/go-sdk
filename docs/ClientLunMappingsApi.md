# \ClientLunMappingsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClientLunMapping**](ClientLunMappingsApi.md#GetClientLunMapping) | **Get** /client-lun-mappings/{client_lun_mapping_id} | 
[**ListClientLunMappings**](ClientLunMappingsApi.md#ListClientLunMappings) | **Get** /client-lun-mappings/ | 


# **GetClientLunMapping**
> ClientLunMappingResp GetClientLunMapping(ctx, clientLunMappingId)


get a client lun mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clientLunMappingId** | **int64**| client lun mapping id | 

### Return type

[**ClientLunMappingResp**](ClientLunMappingResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClientLunMappings**
> ClientLunMappingsResp ListClientLunMappings(ctx, mappingGroupId, optional)


List client lun mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingGroupId** | **int64**| mapping group id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappingGroupId** | **int64**| mapping group id | 
 **limit** | **int64**| paging param | 
 **offset** | **int64**| paging param | 

### Return type

[**ClientLunMappingsResp**](ClientLunMappingsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

