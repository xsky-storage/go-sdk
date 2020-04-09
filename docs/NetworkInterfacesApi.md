# \NetworkInterfacesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkInterface**](NetworkInterfacesApi.md#GetNetworkInterface) | **Get** /network-interfaces/{network_interface_id} | 
[**GetNetworkInterfaceSamples**](NetworkInterfacesApi.md#GetNetworkInterfaceSamples) | **Get** /network-interfaces/{network_interface_id}/samples | 
[**ListNetworkInterfaces**](NetworkInterfacesApi.md#ListNetworkInterfaces) | **Get** /network-interfaces/ | 


# **GetNetworkInterface**
> NetworkInterfaceResp GetNetworkInterface(ctx, networkInterfaceId)


Get a network interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **networkInterfaceId** | **int64**| network interface id | 

### Return type

[**NetworkInterfaceResp**](NetworkInterfaceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkInterfaceSamples**
> NetworkInterfaceSamplesResp GetNetworkInterfaceSamples(ctx, networkInterfaceId, optional)


get a network interface's Samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **networkInterfaceId** | **int64**| network interface id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkInterfaceId** | **int64**| network interface id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**NetworkInterfaceSamplesResp**](NetworkInterfaceSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNetworkInterfaces**
> NetworkInterfacesResp ListNetworkInterfaces(ctx, optional)


List network interfaces

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
 **masterNetworkInterfaceId** | **int64**| master network interface id | 
 **hostId** | **int64**| host id | 
 **role** | **string**| network interface role | 

### Return type

[**NetworkInterfacesResp**](NetworkInterfacesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

