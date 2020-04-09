# \FsQuotaTreesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQuotaTree**](FsQuotaTreesApi.md#GetQuotaTree) | **Get** /fs-quota-trees/{fs_quota_tree_id} | 
[**ListQuotaTrees**](FsQuotaTreesApi.md#ListQuotaTrees) | **Get** /fs-quota-trees/ | 
[**UpdateQuotaTree**](FsQuotaTreesApi.md#UpdateQuotaTree) | **Patch** /fs-quota-trees/{fs_quota_tree_id} | 


# **GetQuotaTree**
> FsQuotaTreeResp GetQuotaTree(ctx, fsQuotaTreeId)


Get file storage quota tree

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsQuotaTreeId** | **int64**| quota tree id | 

### Return type

[**FsQuotaTreeResp**](FSQuotaTreeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListQuotaTrees**
> FsQuotaTreesResp ListQuotaTrees(ctx, optional)


List file storage quota trees

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
 **fsFolderId** | **int64**| file storage folder id | 

### Return type

[**FsQuotaTreesResp**](FSQuotaTreesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuotaTree**
> FsQuotaTreeResp UpdateQuotaTree(ctx, fsQuotaTreeId, body)


Update file storage quota tree

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsQuotaTreeId** | **int64**| quota tree id | 
  **body** | [**FsQuotaTreeUpdateReq**](FsQuotaTreeUpdateReq.md)| quota tree info | 

### Return type

[**FsQuotaTreeResp**](FSQuotaTreeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

