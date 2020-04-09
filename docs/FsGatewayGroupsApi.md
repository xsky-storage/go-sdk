# \FsGatewayGroupsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFSGateways**](FsGatewayGroupsApi.md#AddFSGateways) | **Post** /fs-gateway-groups/{fs_gateway_group_id}:add-gateways | 
[**CreateFSGatewayGroup**](FsGatewayGroupsApi.md#CreateFSGatewayGroup) | **Post** /fs-gateway-groups/ | 
[**DeleteFSGatewayGroup**](FsGatewayGroupsApi.md#DeleteFSGatewayGroup) | **Delete** /fs-gateway-groups/{fs_gateway_group_id} | 
[**GetFSGatewayGroup**](FsGatewayGroupsApi.md#GetFSGatewayGroup) | **Get** /fs-gateway-groups/{fs_gateway_group_id} | 
[**ListFSGatewayGroups**](FsGatewayGroupsApi.md#ListFSGatewayGroups) | **Get** /fs-gateway-groups/ | 
[**RemoveFSGateways**](FsGatewayGroupsApi.md#RemoveFSGateways) | **Post** /fs-gateway-groups/{fs_gateway_group_id}:remove-gateways | 
[**UpdateFSGatewayGroup**](FsGatewayGroupsApi.md#UpdateFSGatewayGroup) | **Patch** /fs-gateway-groups/{fs_gateway_group_id} | 


# **AddFSGateways**
> FsGatewayGroupResp AddFSGateways(ctx, fsGatewayGroupId, body)


add file storage gateways

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsGatewayGroupId** | **int64**| gateway group id | 
  **body** | [**FsGatewayGroupAddGatewaysReq**](FsGatewayGroupAddGatewaysReq.md)| gateways info | 

### Return type

[**FsGatewayGroupResp**](FSGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFSGatewayGroup**
> FsGatewayGroupResp CreateFSGatewayGroup(ctx, body)


Create file storage gateway group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FsGatewayGroupCreateReq**](FsGatewayGroupCreateReq.md)| gateway group info | 

### Return type

[**FsGatewayGroupResp**](FSGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFSGatewayGroup**
> FsGatewayGroupResp DeleteFSGatewayGroup(ctx, fsGatewayGroupId, optional)


delete file storage gateway group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsGatewayGroupId** | **int64**| gateway group id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fsGatewayGroupId** | **int64**| gateway group id | 
 **force** | **bool**| force delete or not | 

### Return type

[**FsGatewayGroupResp**](FSGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFSGatewayGroup**
> FsGatewayGroupResp GetFSGatewayGroup(ctx, fsGatewayGroupId)


Get file storage gateway group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsGatewayGroupId** | **int64**| gateway group id | 

### Return type

[**FsGatewayGroupResp**](FSGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFSGatewayGroups**
> FsGatewayGroupsResp ListFSGatewayGroups(ctx, optional)


List file storage gateway groups

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
 **type_** | **string**| type of file storage gateway group | 
 **security** | **string**| security of file storage gateway group | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**FsGatewayGroupsResp**](FSGatewayGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveFSGateways**
> FsGatewayGroupResp RemoveFSGateways(ctx, fsGatewayGroupId, body, optional)


remove gateways from gateway group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsGatewayGroupId** | **int64**| gateway group id | 
  **body** | [**FsGatewayGroupRemoveGatewaysReq**](FsGatewayGroupRemoveGatewaysReq.md)| gateways info | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fsGatewayGroupId** | **int64**| gateway group id | 
 **body** | [**FsGatewayGroupRemoveGatewaysReq**](FsGatewayGroupRemoveGatewaysReq.md)| gateways info | 
 **force** | **bool**| force delete or not | 

### Return type

[**FsGatewayGroupResp**](FSGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFSGatewayGroup**
> FsGatewayGroupResp UpdateFSGatewayGroup(ctx, fsGatewayGroupId, body)


Update file storage gateway group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsGatewayGroupId** | **int64**| gateway group id | 
  **body** | [**FsGatewayGroupUpdateReq**](FsGatewayGroupUpdateReq.md)| gateway group info | 

### Return type

[**FsGatewayGroupResp**](FSGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

