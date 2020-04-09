# \FsClientsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFSClient**](FsClientsApi.md#CreateFSClient) | **Post** /fs-clients/ | 
[**DeleteFSClient**](FsClientsApi.md#DeleteFSClient) | **Delete** /fs-clients/{fs_client_id} | 
[**GetFSClient**](FsClientsApi.md#GetFSClient) | **Get** /fs-clients/{fs_client_id} | 
[**ListFSClients**](FsClientsApi.md#ListFSClients) | **Get** /fs-clients/ | 
[**UpdateFSClient**](FsClientsApi.md#UpdateFSClient) | **Patch** /fs-clients/{fs_client_id} | 


# **CreateFSClient**
> FsClientResp CreateFSClient(ctx, body)


Create file storage client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FsClientCreateReq**](FsClientCreateReq.md)| client info | 

### Return type

[**FsClientResp**](FSClientResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFSClient**
> DeleteFSClient(ctx, fsClientId)


delete file storage client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsClientId** | **int64**| client id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFSClient**
> FsClientResp GetFSClient(ctx, fsClientId)


Get file storage client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsClientId** | **int64**| client id | 

### Return type

[**FsClientResp**](FSClientResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFSClients**
> FsClientsResp ListFSClients(ctx, optional)


List file storage clients

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
 **fsClientGroupId** | **int64**| file storage client group id | 
 **notFsClientGroupId** | **int64**| group id the client don&#39;t belong to | 

### Return type

[**FsClientsResp**](FSClientsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFSClient**
> FsClientResp UpdateFSClient(ctx, fsClientId, body)


Update file storage client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsClientId** | **int64**| client id | 
  **body** | [**FsClientUpdateReq**](FsClientUpdateReq.md)| client info | 

### Return type

[**FsClientResp**](FSClientResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

