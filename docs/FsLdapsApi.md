# \FsLdapsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFSLdap**](FsLdapsApi.md#CreateFSLdap) | **Post** /fs-ldaps/ | 
[**DeleteFSLdap**](FsLdapsApi.md#DeleteFSLdap) | **Delete** /fs-ldaps/{fs_ldap_id} | 
[**GetFSLdap**](FsLdapsApi.md#GetFSLdap) | **Get** /fs-ldaps/{fs_ldap_id} | 
[**ListFSLdaps**](FsLdapsApi.md#ListFSLdaps) | **Get** /fs-ldaps/ | 
[**UpdateFSLdap**](FsLdapsApi.md#UpdateFSLdap) | **Patch** /fs-ldaps/{fs_ldap_id} | 
[**VerifyFSLdap**](FsLdapsApi.md#VerifyFSLdap) | **Post** /fs-ldaps/{fs_ldap_id}:verify | 


# **CreateFSLdap**
> FsLdapResp CreateFSLdap(ctx, body)


create file storage ldap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FsLdapCreateReq**](FsLdapCreateReq.md)| file storage ldap info | 

### Return type

[**FsLdapResp**](FSLdapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFSLdap**
> DeleteFSLdap(ctx, fsLdapId)


Delete file storage ldap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsLdapId** | **int64**| file storage ldap id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFSLdap**
> FsLdapResp GetFSLdap(ctx, fsLdapId)


Get a file storage ldap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsLdapId** | **int64**| file storage ldap id | 

### Return type

[**FsLdapResp**](FSLdapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFSLdaps**
> FsLdapsResp ListFSLdaps(ctx, optional)


List file storage ldaps

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

### Return type

[**FsLdapsResp**](FSLdapsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFSLdap**
> FsLdapResp UpdateFSLdap(ctx, fsLdapId, body)


Update a file storage ldap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsLdapId** | **int64**| file storage ldap id | 
  **body** | [**FsLdapUpdateReq**](FsLdapUpdateReq.md)| file storage ldap info | 

### Return type

[**FsLdapResp**](FSLdapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyFSLdap**
> FsLdapResp VerifyFSLdap(ctx, fsLdapId)


Verify a file storage ldap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsLdapId** | **int64**| file storage ldap id | 

### Return type

[**FsLdapResp**](FSLdapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

