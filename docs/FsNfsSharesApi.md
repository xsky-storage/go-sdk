# \FsNfsSharesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFSNFSShareACLs**](FsNfsSharesApi.md#AddFSNFSShareACLs) | **Post** /fs-nfs-shares/{fs_nfs_share_id}:add-acls | 
[**CreateFSNFSShare**](FsNfsSharesApi.md#CreateFSNFSShare) | **Post** /fs-nfs-shares/ | 
[**DeleteFSNFSShare**](FsNfsSharesApi.md#DeleteFSNFSShare) | **Delete** /fs-nfs-shares/{fs_nfs_share_id} | 
[**GetFSNFSShare**](FsNfsSharesApi.md#GetFSNFSShare) | **Get** /fs-nfs-shares/{fs_nfs_share_id} | 
[**ListFSNFSShares**](FsNfsSharesApi.md#ListFSNFSShares) | **Get** /fs-nfs-shares/ | 
[**RemoveFSNFSShareACLs**](FsNfsSharesApi.md#RemoveFSNFSShareACLs) | **Post** /fs-nfs-shares/{fs_nfs_share_id}:remove-acls | 
[**UpdateFSNFSShareACLs**](FsNfsSharesApi.md#UpdateFSNFSShareACLs) | **Post** /fs-nfs-shares/{fs_nfs_share_id}:update-acls | 


# **AddFSNFSShareACLs**
> FsnfsShareResp AddFSNFSShareACLs(ctx, fsNfsShareId, body)


Add fs nfs shares acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsNfsShareId** | **int64**| fs nfs shares id | 
  **body** | [**FsnfsShareAddAcLsReq**](FsnfsShareAddAcLsReq.md)| share acls info | 

### Return type

[**FsnfsShareResp**](FSNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFSNFSShare**
> FsnfsShareResp CreateFSNFSShare(ctx, body)


Create fs nfs shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FsnfsShareCreateReq**](FsnfsShareCreateReq.md)| share info | 

### Return type

[**FsnfsShareResp**](FSNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFSNFSShare**
> FsnfsShareResp DeleteFSNFSShare(ctx, fsNfsShareId, optional)


delete fs nfs shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsNfsShareId** | **int64**| share id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fsNfsShareId** | **int64**| share id | 
 **force** | **bool**| force delete or not | 

### Return type

[**FsnfsShareResp**](FSNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFSNFSShare**
> FsnfsShareResp GetFSNFSShare(ctx, fsNfsShareId)


Get fs nfs shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsNfsShareId** | **int64**| share id | 

### Return type

[**FsnfsShareResp**](FSNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFSNFSShares**
> FsnfsSharesResp ListFSNFSShares(ctx, optional)


List fs nfs sharess

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
 **fsFolderId** | **int64**| fs nfs id | 
 **fsGatewayGroupId** | **int64**| file storage gateway group id | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**FsnfsSharesResp**](FSNFSSharesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveFSNFSShareACLs**
> FsnfsShareResp RemoveFSNFSShareACLs(ctx, fsNfsShareId, body)


remove fs nfs shares acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsNfsShareId** | **int64**| fs nfs shares id | 
  **body** | [**FsnfsShareRemoveAcLsReq**](FsnfsShareRemoveAcLsReq.md)| share acls info | 

### Return type

[**FsnfsShareResp**](FSNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFSNFSShareACLs**
> FsnfsShareResp UpdateFSNFSShareACLs(ctx, fsNfsShareId, body)


Update fs nfs share acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsNfsShareId** | **int64**| share id | 
  **body** | [**FsnfsShareUpdateAcLsReq**](FsnfsShareUpdateAcLsReq.md)| share info | 

### Return type

[**FsnfsShareResp**](FSNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

