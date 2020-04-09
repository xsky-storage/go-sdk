# \RoleMappingsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoleMapping**](RoleMappingsApi.md#CreateRoleMapping) | **Post** /role-mappings/ | 
[**DeleteRoleMapping**](RoleMappingsApi.md#DeleteRoleMapping) | **Delete** /role-mappings/{role_mapping_id} | 
[**GetRoleMapping**](RoleMappingsApi.md#GetRoleMapping) | **Get** /role-mappings/{role_mapping_id} | 
[**ListRoleMappings**](RoleMappingsApi.md#ListRoleMappings) | **Get** /role-mappings/ | 
[**UpdateRoleMapping**](RoleMappingsApi.md#UpdateRoleMapping) | **Patch** /role-mappings/{role_mapping_id} | 


# **CreateRoleMapping**
> RoleMappingResp CreateRoleMapping(ctx, body)


Create role mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RoleMappingCreateReq**](RoleMappingCreateReq.md)| role mapping info | 

### Return type

[**RoleMappingResp**](RoleMappingResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoleMapping**
> DeleteRoleMapping(ctx, roleMappingId)


Delete role mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **roleMappingId** | **int64**| role mapping id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleMapping**
> RoleMappingResp GetRoleMapping(ctx, roleMappingId)


Get a role mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **roleMappingId** | **int64**| role mapping id | 

### Return type

[**RoleMappingResp**](RoleMappingResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRoleMappings**
> RoleMappingsResp ListRoleMappings(ctx, optional)


List role mappings

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
 **identityPlatformId** | **int64**| identity platform id | 

### Return type

[**RoleMappingsResp**](RoleMappingsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoleMapping**
> RoleMappingResp UpdateRoleMapping(ctx, roleMappingId, body)


Update a role mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **roleMappingId** | **int64**| role mapping id | 
  **body** | [**RoleMappingUpdateReq**](RoleMappingUpdateReq.md)| role mapping info | 

### Return type

[**RoleMappingResp**](RoleMappingResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

