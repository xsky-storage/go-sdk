# \OsRemotePoliciesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSRemotePolicy**](OsRemotePoliciesApi.md#GetOSRemotePolicy) | **Get** /os-remote-policies/{policy_uuid} | 
[**ListOSRemotePolicies**](OsRemotePoliciesApi.md#ListOSRemotePolicies) | **Get** /os-remote-policies/ | 


# **GetOSRemotePolicy**
> OsRemotePolicyResp GetOSRemotePolicy(ctx, policyUuid)


Get a os remote policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyUuid** | **string**| policy uuid | 

### Return type

[**OsRemotePolicyResp**](OSRemotePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSRemotePolicies**
> OsRemotePoliciesResp ListOSRemotePolicies(ctx, optional)


List os remote policies

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
 **marker** | **string**| paging param | 
 **zoneUuid** | **string**| zone uuid | 

### Return type

[**OsRemotePoliciesResp**](OSRemotePoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

