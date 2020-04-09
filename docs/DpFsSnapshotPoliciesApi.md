# \DpFsSnapshotPoliciesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpFSSnapshotPolicy**](DpFsSnapshotPoliciesApi.md#CreateDpFSSnapshotPolicy) | **Post** /dp-fs-snapshot-policies/ | 
[**DeleteDpFSSnapshotPolicy**](DpFsSnapshotPoliciesApi.md#DeleteDpFSSnapshotPolicy) | **Delete** /dp-fs-snapshot-policies/{policy_id} | 
[**GetDpFSSnapshotPolicy**](DpFsSnapshotPoliciesApi.md#GetDpFSSnapshotPolicy) | **Get** /dp-fs-snapshot-policies/{policy_id} | 
[**ListDpFSSnapshotPolicies**](DpFsSnapshotPoliciesApi.md#ListDpFSSnapshotPolicies) | **Get** /dp-fs-snapshot-policies/ | 
[**UpdateDpFSSnapshotPolicy**](DpFsSnapshotPoliciesApi.md#UpdateDpFSSnapshotPolicy) | **Patch** /dp-fs-snapshot-policies/{policy_id} | 


# **CreateDpFSSnapshotPolicy**
> DpFsSnapshotPolicyResp CreateDpFSSnapshotPolicy(ctx, body)


Create dp fs snapshot policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**DpFsSnapshotPolicyCreateReq**](DpFsSnapshotPolicyCreateReq.md)| policy info | 

### Return type

[**DpFsSnapshotPolicyResp**](DpFSSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDpFSSnapshotPolicy**
> DpFsSnapshotPolicyResp DeleteDpFSSnapshotPolicy(ctx, policyId, optional)


Delete dp fs snapshot policy

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

[**DpFsSnapshotPolicyResp**](DpFSSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDpFSSnapshotPolicy**
> DpFsSnapshotPolicyResp GetDpFSSnapshotPolicy(ctx, policyId)


Get dp fs snapshot policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyId** | **int64**| resource id | 

### Return type

[**DpFsSnapshotPolicyResp**](DpFSSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDpFSSnapshotPolicies**
> DpFsSnapshotPoliciesResp ListDpFSSnapshotPolicies(ctx, optional)


List dp fs snapshot policies

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

[**DpFsSnapshotPoliciesResp**](DpFSSnapshotPoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDpFSSnapshotPolicy**
> DpFsSnapshotPolicyResp UpdateDpFSSnapshotPolicy(ctx, policyId, body)


Update dp fs snapshot policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyId** | **int64**| resource id | 
  **body** | [**DpFsSnapshotPolicyUpdateReq**](DpFsSnapshotPolicyUpdateReq.md)| policy info | 

### Return type

[**DpFsSnapshotPolicyResp**](DpFSSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

