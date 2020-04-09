# \FsFtpSessionsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFSFTPSessions**](FsFtpSessionsApi.md#ListFSFTPSessions) | **Get** /fs-ftp-sessions/ | 


# **ListFSFTPSessions**
> FsftpSessionsResp ListFSFTPSessions(ctx, optional)


List fs ftp sessions

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
 **fsFtpShareId** | **int64**| fs ftp share id | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**FsftpSessionsResp**](FSFTPSessionsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

