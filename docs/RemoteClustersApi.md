# \RemoteClustersApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRemoteCluster**](RemoteClustersApi.md#CreateRemoteCluster) | **Post** /remote-clusters/ | 
[**DeleteRemoteCluster**](RemoteClustersApi.md#DeleteRemoteCluster) | **Delete** /remote-clusters/{cluster_id} | 
[**GetRemoteCluster**](RemoteClustersApi.md#GetRemoteCluster) | **Get** /remote-clusters/{cluster_id} | 
[**ListRemoteClusters**](RemoteClustersApi.md#ListRemoteClusters) | **Get** /remote-clusters/ | 
[**UpdateRemoteCluster**](RemoteClustersApi.md#UpdateRemoteCluster) | **Patch** /remote-clusters/{cluster_id} | 


# **CreateRemoteCluster**
> RemoteClusterResp CreateRemoteCluster(ctx, body)


Create a remote cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RemoteClusterCreateReq**](RemoteClusterCreateReq.md)| remote cluster info | 

### Return type

[**RemoteClusterResp**](RemoteClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRemoteCluster**
> RemoteClusterResp DeleteRemoteCluster(ctx, clusterId)


Delete a remote cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterId** | **int64**| remote cluster id | 

### Return type

[**RemoteClusterResp**](RemoteClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRemoteCluster**
> RemoteClusterResp GetRemoteCluster(ctx, clusterId)


Get remote cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterId** | **int64**| remote cluster id | 

### Return type

[**RemoteClusterResp**](RemoteClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRemoteClusters**
> RemoteClustersResp ListRemoteClusters(ctx, optional)


List remote clusters

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
 **fsId** | **string**| fsid of cluster | 

### Return type

[**RemoteClustersResp**](RemoteClustersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRemoteCluster**
> RemoteClusterResp UpdateRemoteCluster(ctx, clusterId, body)


Update a remote cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterId** | **int64**| remote cluster id | 
  **body** | [**RemoteClusterUpdateReq**](RemoteClusterUpdateReq.md)| remote cluster info | 

### Return type

[**RemoteClusterResp**](RemoteClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

