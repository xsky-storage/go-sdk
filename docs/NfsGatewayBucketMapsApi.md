# \NfsGatewayBucketMapsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNFSGatewayS3BucketMaps**](NfsGatewayBucketMapsApi.md#ListNFSGatewayS3BucketMaps) | **Get** /nfs-gateway-bucket-maps/ | 


# **ListNFSGatewayS3BucketMaps**
> NfsGatewayBucketMapsResp ListNFSGatewayS3BucketMaps(ctx, optional)


List nfs gateway s3 bucket maps

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
 **nfsGatewayId** | **int64**| nfs gateway id | 
 **osBucketId** | **int64**| s3 bucket id | 

### Return type

[**NfsGatewayBucketMapsResp**](NFSGatewayBucketMapsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

