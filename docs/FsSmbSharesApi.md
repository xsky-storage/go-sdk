# \FsSmbSharesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFSSMBShareACLs**](FsSmbSharesApi.md#AddFSSMBShareACLs) | **Post** /fs-smb-shares/{fs_smb_share_id}:add-acls | 
[**CreateFSSMBShare**](FsSmbSharesApi.md#CreateFSSMBShare) | **Post** /fs-smb-shares/ | 
[**DeleteFSSMBShare**](FsSmbSharesApi.md#DeleteFSSMBShare) | **Delete** /fs-smb-shares/{fs_smb_share_id} | 
[**GetFSSMBShare**](FsSmbSharesApi.md#GetFSSMBShare) | **Get** /fs-smb-shares/{fs_smb_share_id} | 
[**ListFSSMBShares**](FsSmbSharesApi.md#ListFSSMBShares) | **Get** /fs-smb-shares/ | 
[**RemoveFSSMBShareACLs**](FsSmbSharesApi.md#RemoveFSSMBShareACLs) | **Post** /fs-smb-shares/{fs_smb_share_id}:remove-acls | 
[**UpdateFSSMBShare**](FsSmbSharesApi.md#UpdateFSSMBShare) | **Patch** /fs-smb-shares/{fs_smb_share_id} | 
[**UpdateFSSMBShareACLs**](FsSmbSharesApi.md#UpdateFSSMBShareACLs) | **Post** /fs-smb-shares/{fs_smb_share_id}:update-acls | 


# **AddFSSMBShareACLs**
> FssmbShareResp AddFSSMBShareACLs(ctx, fsSmbShareId, body)


Add fs smb share acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsSmbShareId** | **int64**| fs smb share id | 
  **body** | [**FssmbShareAddAcLsReq**](FssmbShareAddAcLsReq.md)| share acls info | 

### Return type

[**FssmbShareResp**](FSSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFSSMBShare**
> FssmbShareResp CreateFSSMBShare(ctx, body)


Create fs smb share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FssmbShareCreateReq**](FssmbShareCreateReq.md)| share info | 

### Return type

[**FssmbShareResp**](FSSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFSSMBShare**
> FssmbShareResp DeleteFSSMBShare(ctx, fsSmbShareId, optional)


delete fs smb share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsSmbShareId** | **int64**| share id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fsSmbShareId** | **int64**| share id | 
 **force** | **bool**| force delete or not | 

### Return type

[**FssmbShareResp**](FSSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFSSMBShare**
> FssmbShareResp GetFSSMBShare(ctx, fsSmbShareId)


Get fs smb share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsSmbShareId** | **int64**| share id | 

### Return type

[**FssmbShareResp**](FSSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFSSMBShares**
> FssmbSharesResp ListFSSMBShares(ctx, optional)


List fs smb shares

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
 **fsFolderId** | **int64**| fs folder id | 
 **fsGatewayGroupId** | **int64**| fs gateway group id | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**FssmbSharesResp**](FSSMBSharesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveFSSMBShareACLs**
> FssmbShareResp RemoveFSSMBShareACLs(ctx, fsSmbShareId, body)


remove fs smb share acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsSmbShareId** | **int64**| fs smb share id | 
  **body** | [**FssmbShareRemoveAcLsReq**](FssmbShareRemoveAcLsReq.md)| share acls info | 

### Return type

[**FssmbShareResp**](FSSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFSSMBShare**
> FssmbShareResp UpdateFSSMBShare(ctx, fsSmbShareId, body)


Update fs smb share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsSmbShareId** | **int64**| share id | 
  **body** | [**FssmbShareUpdateReq**](FssmbShareUpdateReq.md)| share info | 

### Return type

[**FssmbShareResp**](FSSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFSSMBShareACLs**
> FssmbShareResp UpdateFSSMBShareACLs(ctx, fsSmbShareId, body)


Update fs smb share ACL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsSmbShareId** | **int64**| share id | 
  **body** | [**FssmbShareUpdateAcLsReq**](FssmbShareUpdateAcLsReq.md)| share acls info | 

### Return type

[**FssmbShareResp**](FSSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

