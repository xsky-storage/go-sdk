# \FsSmbShareAclsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFSSMBShareACLs**](FsSmbShareAclsApi.md#ListFSSMBShareACLs) | **Get** /fs-smb-share-acls/ | 


# **ListFSSMBShareACLs**
> FssmbShareAcLsResp ListFSSMBShareACLs(ctx, optional)


List fs smb share acls

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
 **fsSmbShareId** | **int64**| fs smb share id | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**FssmbShareAcLsResp**](FSSMBShareACLsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

