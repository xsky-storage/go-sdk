# \PlacementNodesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlacementNode**](PlacementNodesApi.md#CreatePlacementNode) | **Post** /placement-nodes/ | 
[**DeletePlacementNode**](PlacementNodesApi.md#DeletePlacementNode) | **Delete** /placement-nodes/{placement_node_id} | 
[**GetPlacementNode**](PlacementNodesApi.md#GetPlacementNode) | **Get** /placement-nodes/{placement_node_id} | 
[**GetPlacementNodeTopology**](PlacementNodesApi.md#GetPlacementNodeTopology) | **Get** /placement-nodes/{placement_node_id}/topology | 
[**UpdatePlacementNode**](PlacementNodesApi.md#UpdatePlacementNode) | **Patch** /placement-nodes/{placement_node_id} | 


# **CreatePlacementNode**
> PlacementNodeResp CreatePlacementNode(ctx, body)


Create placement node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**PlacementNodeCreateReq**](PlacementNodeCreateReq.md)| placement node info | 

### Return type

[**PlacementNodeResp**](PlacementNodeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePlacementNode**
> DeletePlacementNode(ctx, placementNodeId)


delete placement node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **placementNodeId** | **int64**| placement node id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlacementNode**
> PlacementNodeResp GetPlacementNode(ctx, placementNodeId)


Get placement node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **placementNodeId** | **int64**| placement node id | 

### Return type

[**PlacementNodeResp**](PlacementNodeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlacementNodeTopology**
> PlacementNodeTopologyResp GetPlacementNodeTopology(ctx, placementNodeId, optional)


Get subtree topology of placement node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **placementNodeId** | **int64**| placement node id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **placementNodeId** | **int64**| placement node id | 
 **osdType** | **string**| osd type | 
 **osdRole** | **string**| osd role | 
 **withCompound** | **bool**| with compound osd | 
 **withHybrid** | **bool**| with hybrid osd | 

### Return type

[**PlacementNodeTopologyResp**](PlacementNodeTopologyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePlacementNode**
> PlacementNodeResp UpdatePlacementNode(ctx, placementNodeId, body)


update placement node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **placementNodeId** | **int64**| the placement node id | 
  **body** | [**PlacementNodeUpdateReq**](PlacementNodeUpdateReq.md)| the placement node info | 

### Return type

[**PlacementNodeResp**](PlacementNodeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

