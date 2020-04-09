# \DpBlockReplicationPoliciesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpBlockReplicationPolicy**](DpBlockReplicationPoliciesApi.md#CreateDpBlockReplicationPolicy) | **Post** /dp-block-replication-policies/ | 
[**DeleteDpBlockReplicationPolicy**](DpBlockReplicationPoliciesApi.md#DeleteDpBlockReplicationPolicy) | **Delete** /dp-block-replication-policies/{policy_id} | 
[**GetDpBlockReplicationPolicy**](DpBlockReplicationPoliciesApi.md#GetDpBlockReplicationPolicy) | **Get** /dp-block-replication-policies/{policy_id} | 
[**ListDpBlockReplicationPolicies**](DpBlockReplicationPoliciesApi.md#ListDpBlockReplicationPolicies) | **Get** /dp-block-replication-policies/ | 
[**UpdateDpBlockReplicationPolicy**](DpBlockReplicationPoliciesApi.md#UpdateDpBlockReplicationPolicy) | **Patch** /dp-block-replication-policies/{policy_id} | 


# **CreateDpBlockReplicationPolicy**
> DpBlockReplicationPolicyResp CreateDpBlockReplicationPolicy(ctx, body)


Create dp block replication policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**DpBlockReplicationPolicyCreateReq**](DpBlockReplicationPolicyCreateReq.md)| policy info | 

### Return type

[**DpBlockReplicationPolicyResp**](DpBlockReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDpBlockReplicationPolicy**
> DeleteDpBlockReplicationPolicy(ctx, policyId)


Delete dp block replication policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyId** | **int64**| policy id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDpBlockReplicationPolicy**
> DpBlockReplicationPolicyResp GetDpBlockReplicationPolicy(ctx, policyId)


Get dp block replication policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyId** | **int64**| resource id | 

### Return type

[**DpBlockReplicationPolicyResp**](DpBlockReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDpBlockReplicationPolicies**
> DpBlockReplicationPoliciesResp ListDpBlockReplicationPolicies(ctx, optional)


List dp block replication policies

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

[**DpBlockReplicationPoliciesResp**](DpBlockReplicationPoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDpBlockReplicationPolicy**
> DpBlockReplicationPolicyResp UpdateDpBlockReplicationPolicy(ctx, policyId, body)


Update dp block replication policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyId** | **int64**| resource id | 
  **body** | [**DpBlockReplicationPolicyUpdateReq**](DpBlockReplicationPolicyUpdateReq.md)| policy info | 

### Return type

[**DpBlockReplicationPolicyResp**](DpBlockReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

