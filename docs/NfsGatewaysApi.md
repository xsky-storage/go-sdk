# \NfsGatewaysApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNFSGateway**](NfsGatewaysApi.md#CreateNFSGateway) | **Post** /nfs-gateways/ | 
[**CreateNFSGatewayBucketMap**](NfsGatewaysApi.md#CreateNFSGatewayBucketMap) | **Put** /nfs-gateways/{gateway_id}/buckets/{bucket_id} | 
[**DeleteNFSGateway**](NfsGatewaysApi.md#DeleteNFSGateway) | **Delete** /nfs-gateways/{gateway_id} | 
[**DeleteNFSGatewayBucketMap**](NfsGatewaysApi.md#DeleteNFSGatewayBucketMap) | **Delete** /nfs-gateways/{gateway_id}/buckets/{bucket_id} | 
[**DoNFSGateway**](NfsGatewaysApi.md#DoNFSGateway) | **Put** /nfs-gateways/{gateway_id} | 
[**GetNFSGateway**](NfsGatewaysApi.md#GetNFSGateway) | **Get** /nfs-gateways/{gateway_id} | 
[**GetNFSGatewayBucketMap**](NfsGatewaysApi.md#GetNFSGatewayBucketMap) | **Get** /nfs-gateways/{gateway_id}/buckets/{bucket_id} | 
[**GetNFSGatewaySamples**](NfsGatewaysApi.md#GetNFSGatewaySamples) | **Get** /nfs-gateways/{gateway_id}/samples | 
[**ListNFSGatewayBucketMaps**](NfsGatewaysApi.md#ListNFSGatewayBucketMaps) | **Get** /nfs-gateways/{gateway_id}/buckets | 
[**ListNFSGateways**](NfsGatewaysApi.md#ListNFSGateways) | **Get** /nfs-gateways/ | 
[**UpdateNFSGateway**](NfsGatewaysApi.md#UpdateNFSGateway) | **Patch** /nfs-gateways/{gateway_id} | 
[**UpdateNFSGatewayBucketMap**](NfsGatewaysApi.md#UpdateNFSGatewayBucketMap) | **Patch** /nfs-gateways/{gateway_id}/buckets/{bucket_id} | 


# **CreateNFSGateway**
> NfsGatewayResp CreateNFSGateway(ctx, body)


create nfs gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**NfsGatewayCreateReq**](NfsGatewayCreateReq.md)| nfs gateway info | 

### Return type

[**NfsGatewayResp**](NFSGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNFSGatewayBucketMap**
> NfsGatewayBucketMapResp CreateNFSGatewayBucketMap(ctx, gatewayId, bucketId)


add bucket to nfs gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| nfs gateway id | 
  **bucketId** | **int64**| bucket id | 

### Return type

[**NfsGatewayBucketMapResp**](NFSGatewayBucketMapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNFSGateway**
> NfsGatewayResp DeleteNFSGateway(ctx, gatewayId, optional)


delete nfs gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| nfs gateway id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayId** | **int64**| nfs gateway id | 
 **force** | **bool**| force delete or not | 

### Return type

[**NfsGatewayResp**](NFSGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNFSGatewayBucketMap**
> NfsGatewayBucketMapResp DeleteNFSGatewayBucketMap(ctx, gatewayId, bucketId, optional)


remove bucket from nfs gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| nfs gateway id | 
  **bucketId** | **int64**| bucket id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayId** | **int64**| nfs gateway id | 
 **bucketId** | **int64**| bucket id | 
 **force** | **bool**| force delete or no | 

### Return type

[**NfsGatewayBucketMapResp**](NFSGatewayBucketMapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoNFSGateway**
> NfsGatewayResp DoNFSGateway(ctx, gatewayId, body, optional)


start/stop nfs gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| nfs gateway id | 
  **body** | [**NfsGatewayActionReq**](NfsGatewayActionReq.md)| nfs gateway action info | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayId** | **int64**| nfs gateway id | 
 **body** | [**NfsGatewayActionReq**](NfsGatewayActionReq.md)| nfs gateway action info | 
 **force** | **bool**| force stop or no | 

### Return type

[**NfsGatewayResp**](NFSGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNFSGateway**
> NfsGatewayResp GetNFSGateway(ctx, gatewayId)


show nfs gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| nfs gateway id | 

### Return type

[**NfsGatewayResp**](NFSGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNFSGatewayBucketMap**
> NfsGatewayBucketMapResp GetNFSGatewayBucketMap(ctx, gatewayId, bucketId)


get nfs gateway bucket map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| nfs gateway id | 
  **bucketId** | **int64**| bucket id | 

### Return type

[**NfsGatewayBucketMapResp**](NFSGatewayBucketMapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNFSGatewaySamples**
> NfsGatewaySamplesResp GetNFSGatewaySamples(ctx, gatewayId, optional)


Get nfs gateway's samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| gateway id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayId** | **int64**| gateway id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**NfsGatewaySamplesResp**](NFSGatewaySamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNFSGatewayBucketMaps**
> NfsGatewayBucketMapsResp ListNFSGatewayBucketMaps(ctx, gatewayId, optional)


List nfs gateway bucket maps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| nfs gateway id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayId** | **int64**| nfs gateway id | 
 **limit** | **int64**| paging param | 
 **offset** | **int64**| paging param | 

### Return type

[**NfsGatewayBucketMapsResp**](NFSGatewayBucketMapsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNFSGateways**
> NfsGatewaysResp ListNFSGateways(ctx, optional)


List all nfs gateways

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
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**NfsGatewaysResp**](NFSGatewaysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNFSGateway**
> NfsGatewayResp UpdateNFSGateway(ctx, gatewayId, body)


update nfs gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| nfs gateway id | 
  **body** | [**NfsGatewayUpdateReq**](NfsGatewayUpdateReq.md)| nfs gateway info | 

### Return type

[**NfsGatewayResp**](NFSGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNFSGatewayBucketMap**
> NfsGatewayBucketMapResp UpdateNFSGatewayBucketMap(ctx, gatewayId, bucketId, body, optional)


update nfs gateway bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| nfs gateway id | 
  **bucketId** | **int64**| bucket id | 
  **body** | [**NfsGatewayBucketMapUpdateReq**](NfsGatewayBucketMapUpdateReq.md)| nfs gateway bucket update info | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayId** | **int64**| nfs gateway id | 
 **bucketId** | **int64**| bucket id | 
 **body** | [**NfsGatewayBucketMapUpdateReq**](NfsGatewayBucketMapUpdateReq.md)| nfs gateway bucket update info | 
 **force** | **bool**| force update bucket map | 

### Return type

[**NfsGatewayBucketMapResp**](NFSGatewayBucketMapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

