# \NetworkAddressesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkAddress**](NetworkAddressesApi.md#GetNetworkAddress) | **Get** /network-addresses/{network_address_id} | 
[**ListNetworkAddresses**](NetworkAddressesApi.md#ListNetworkAddresses) | **Get** /network-addresses/ | 


# **GetNetworkAddress**
> NetworkAddressResp GetNetworkAddress(ctx, networkAddressId)


Get a network address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **networkAddressId** | **int64**| network address id | 

### Return type

[**NetworkAddressResp**](NetworkAddressResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNetworkAddresses**
> NetworkAddressesResp ListNetworkAddresses(ctx, optional)


List network addresses

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
 **networkInterfaceId** | **int64**| network interface id | 
 **hostId** | **int64**| host id | 
 **role** | **string**| network address role | 
 **vipGroupId** | **int64**| vip group id | 

### Return type

[**NetworkAddressesResp**](NetworkAddressesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

