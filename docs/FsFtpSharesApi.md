# \FsFtpSharesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFSFTPShareACLs**](FsFtpSharesApi.md#AddFSFTPShareACLs) | **Post** /fs-ftp-shares/{fs_ftp_share_id}:add-acls | 
[**CreateFSFTPShare**](FsFtpSharesApi.md#CreateFSFTPShare) | **Post** /fs-ftp-shares/ | 
[**DeleteFSFTPShare**](FsFtpSharesApi.md#DeleteFSFTPShare) | **Delete** /fs-ftp-shares/{fs_ftp_share_id} | 
[**GetFSFTPShare**](FsFtpSharesApi.md#GetFSFTPShare) | **Get** /fs-ftp-shares/{fs_ftp_share_id} | 
[**ListFSFTPShares**](FsFtpSharesApi.md#ListFSFTPShares) | **Get** /fs-ftp-shares/ | 
[**RemoveFSFTPShareACLs**](FsFtpSharesApi.md#RemoveFSFTPShareACLs) | **Post** /fs-ftp-shares/{fs_ftp_share_id}:remove-acls | 
[**UpdateFSFTPShareACLs**](FsFtpSharesApi.md#UpdateFSFTPShareACLs) | **Post** /fs-ftp-shares/{fs_ftp_share_id}:update-acls | 


# **AddFSFTPShareACLs**
> FsftpShareResp AddFSFTPShareACLs(ctx, fsFtpShareId, body)


Add fs ftp share acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFtpShareId** | **int64**| fs ftp share id | 
  **body** | [**FsftpShareAddAcLsReq**](FsftpShareAddAcLsReq.md)| ftp share acls info | 

### Return type

[**FsftpShareResp**](FSFTPShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFSFTPShare**
> FsftpShareResp CreateFSFTPShare(ctx, body)


Create fs ftp share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FsftpShareCreateReq**](FsftpShareCreateReq.md)| share info | 

### Return type

[**FsftpShareResp**](FSFTPShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFSFTPShare**
> FsftpShareResp DeleteFSFTPShare(ctx, fsFtpShareId, optional)


delete fs ftp share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFtpShareId** | **int64**| share id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fsFtpShareId** | **int64**| share id | 
 **force** | **bool**| force delete or not | 

### Return type

[**FsftpShareResp**](FSFTPShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFSFTPShare**
> FsftpShareResp GetFSFTPShare(ctx, fsFtpShareId)


Get fs ftp share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFtpShareId** | **int64**| share id | 

### Return type

[**FsftpShareResp**](FSFTPShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFSFTPShares**
> FsftpSharesResp ListFSFTPShares(ctx, optional)


List fs ftp shares

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
 **fsFolderId** | **int64**| fs smb id | 
 **fsGatewayGroupId** | **int64**| fs gateway group id | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**FsftpSharesResp**](FSFTPSharesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveFSFTPShareACLs**
> FsftpShareResp RemoveFSFTPShareACLs(ctx, fsFtpShareId, body)


remove fs ftp share acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFtpShareId** | **int64**| fs ftp share id | 
  **body** | [**FsftpShareRemoveAcLsReq**](FsftpShareRemoveAcLsReq.md)| share acls info | 

### Return type

[**FsftpShareResp**](FSFTPShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFSFTPShareACLs**
> FsftpShareResp UpdateFSFTPShareACLs(ctx, fsFtpShareId, body)


Update fs ftp share acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFtpShareId** | **int64**| ftp share id | 
  **body** | [**FsftpShareUpdateAcLsReq**](FsftpShareUpdateAcLsReq.md)| ftp share acls info | 

### Return type

[**FsftpShareResp**](FSFTPShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

