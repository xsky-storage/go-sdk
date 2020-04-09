# \DpBlockSnapshotPoliciesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpBlockSnapshotPolicy**](DpBlockSnapshotPoliciesApi.md#CreateDpBlockSnapshotPolicy) | **Post** /dp-block-snapshot-policies/ | 
[**DeleteDpBlockSnapshotPolicy**](DpBlockSnapshotPoliciesApi.md#DeleteDpBlockSnapshotPolicy) | **Delete** /dp-block-snapshot-policies/{policy_id} | 
[**GetDpBlockSnapshotPolicy**](DpBlockSnapshotPoliciesApi.md#GetDpBlockSnapshotPolicy) | **Get** /dp-block-snapshot-policies/{policy_id} | 
[**ListDpBlockSnapshotPolicies**](DpBlockSnapshotPoliciesApi.md#ListDpBlockSnapshotPolicies) | **Get** /dp-block-snapshot-policies/ | 
[**UpdateDpBlockSnapshotPolicy**](DpBlockSnapshotPoliciesApi.md#UpdateDpBlockSnapshotPolicy) | **Patch** /dp-block-snapshot-policies/{policy_id} | 


# **CreateDpBlockSnapshotPolicy**
> DpBlockSnapshotPolicyResp CreateDpBlockSnapshotPolicy(ctx, body)


Create dp block snapshot policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**DpBlockSnapshotPolicyCreateReq**](DpBlockSnapshotPolicyCreateReq.md)| policy info | 

### Return type

[**DpBlockSnapshotPolicyResp**](DpBlockSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDpBlockSnapshotPolicy**
> DpBlockSnapshotPolicyResp DeleteDpBlockSnapshotPolicy(ctx, policyId, optional)


Delete dp block snapshot policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyId** | **int64**| resource id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyId** | **int64**| resource id | 
 **force** | **bool**| force delete or not | 

### Return type

[**DpBlockSnapshotPolicyResp**](DpBlockSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDpBlockSnapshotPolicy**
> DpBlockSnapshotPolicyResp GetDpBlockSnapshotPolicy(ctx, policyId)


Get dp block snapshot policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyId** | **int64**| resource id | 

### Return type

[**DpBlockSnapshotPolicyResp**](DpBlockSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDpBlockSnapshotPolicies**
> DpBlockSnapshotPoliciesResp ListDpBlockSnapshotPolicies(ctx, optional)


List dp block snapshot policies

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

### Return type

[**DpBlockSnapshotPoliciesResp**](DpBlockSnapshotPoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDpBlockSnapshotPolicy**
> DpBlockSnapshotPolicyResp UpdateDpBlockSnapshotPolicy(ctx, policyId, body)


Update dp block snapshot policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyId** | **int64**| resource id | 
  **body** | [**DpBlockSnapshotPolicyUpdateReq**](DpBlockSnapshotPolicyUpdateReq.md)| policy info | 

### Return type

[**DpBlockSnapshotPolicyResp**](DpBlockSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

