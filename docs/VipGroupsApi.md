# \VipGroupsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVIPGroup**](VipGroupsApi.md#CreateVIPGroup) | **Post** /vip-groups/ | 
[**DeleteVIPGroup**](VipGroupsApi.md#DeleteVIPGroup) | **Delete** /vip-groups/{vip_group_id} | 
[**GetVIPGroup**](VipGroupsApi.md#GetVIPGroup) | **Get** /vip-groups/{vip_group_id} | 
[**ListVIPGroups**](VipGroupsApi.md#ListVIPGroups) | **Get** /vip-groups/ | 
[**RedeployVIPGroup**](VipGroupsApi.md#RedeployVIPGroup) | **Post** /vip-groups/{vip_group_id}:redeploy | 
[**UpdateVIPGroup**](VipGroupsApi.md#UpdateVIPGroup) | **Patch** /vip-groups/{vip_group_id} | 


# **CreateVIPGroup**
> VipGroupResp CreateVIPGroup(ctx, vipGroup)


Create a vip group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **vipGroup** | [**VipGroupCreateReq**](VipGroupCreateReq.md)| vip group info | 

### Return type

[**VipGroupResp**](VIPGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVIPGroup**
> VipGroupResp DeleteVIPGroup(ctx, vipGroupId)


Delete a vip group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **vipGroupId** | **int64**| vip group id | 

### Return type

[**VipGroupResp**](VIPGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVIPGroup**
> VipGroupResp GetVIPGroup(ctx, vipGroupId)


Get a vip group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **vipGroupId** | **int64**| vip group id | 

### Return type

[**VipGroupResp**](VIPGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVIPGroups**
> VipGroupResps ListVIPGroups(ctx, optional)


List vip groups

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
 **resourceType** | **string**| resource type | 
 **resourceId** | **int64**| resource id | 

### Return type

[**VipGroupResps**](VIPGroupResps.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RedeployVIPGroup**
> VipGroupResp RedeployVIPGroup(ctx, vipGroupId)


Redeploy a vip group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **vipGroupId** | **int64**| vip group id | 

### Return type

[**VipGroupResp**](VIPGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVIPGroup**
> VipGroupResp UpdateVIPGroup(ctx, vipGroupId, vipGroup)


Update a vip group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **vipGroupId** | **int64**| vip group id | 
  **vipGroup** | [**VipGroupUpdateReq**](VipGroupUpdateReq.md)| vip group info | 

### Return type

[**VipGroupResp**](VIPGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

