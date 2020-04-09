# \HostsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHost**](HostsApi.md#CreateHost) | **Post** /hosts/ | 
[**DeleteHost**](HostsApi.md#DeleteHost) | **Delete** /hosts/{host_id} | 
[**GetHost**](HostsApi.md#GetHost) | **Get** /hosts/{host_id} | 
[**GetHostSamples**](HostsApi.md#GetHostSamples) | **Get** /hosts/{host_id}/samples | 
[**ListHosts**](HostsApi.md#ListHosts) | **Get** /hosts/ | 
[**UpdateHost**](HostsApi.md#UpdateHost) | **Patch** /hosts/{host_id} | 


# **CreateHost**
> HostResp CreateHost(ctx, body)


check env and install packages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**HostCreateReq**](HostCreateReq.md)| host info | 

### Return type

[**HostResp**](HostResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHost**
> HostResp DeleteHost(ctx, hostId, optional)


delete host

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hostId** | **int64**| host id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostId** | **int64**| host id | 
 **force** | **bool**| force delete or not | 

### Return type

[**HostResp**](HostResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHost**
> HostResp GetHost(ctx, hostId)


get a host info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hostId** | **int64**| the host id | 

### Return type

[**HostResp**](HostResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHostSamples**
> HostSamplesResp GetHostSamples(ctx, hostId, optional)


get a host's samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hostId** | **int64**| host id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostId** | **int64**| host id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**HostSamplesResp**](HostSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHosts**
> HostsResp ListHosts(ctx, optional)


List hosts by fileter

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
 **protectionDomainId** | **int64**| protection domain id | 
 **hostname** | **string**| host name | 
 **type_** | **string**| if it existed, value should be xdcactive | 
 **role** | **string**| filter by host role | 
 **fcAvailable** | **bool**| available host with fc port | 
 **replicationGatewayAvailable** | **bool**| available host for replication gateway | 
 **fsGatewayGroupId** | **int64**| file storage gateway group id | 
 **fsGatewayGroupUsed** | **bool**| used in file storage gateway group | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**HostsResp**](HostsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHost**
> HostResp UpdateHost(ctx, hostId, body)


update host info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hostId** | **int64**| host id | 
  **body** | [**HostUpdateReq**](HostUpdateReq.md)| host info | 

### Return type

[**HostResp**](HostResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

