# \DpBlockBackupPoliciesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpBlockBackupPolicy**](DpBlockBackupPoliciesApi.md#CreateDpBlockBackupPolicy) | **Post** /dp-block-backup-policies/ | 
[**DeleteDpBlockBackupPolicy**](DpBlockBackupPoliciesApi.md#DeleteDpBlockBackupPolicy) | **Delete** /dp-block-backup-policies/{policy_id} | 
[**GetDpBlockBackupPolicy**](DpBlockBackupPoliciesApi.md#GetDpBlockBackupPolicy) | **Get** /dp-block-backup-policies/{policy_id} | 
[**ListDpBlockBackupPolicies**](DpBlockBackupPoliciesApi.md#ListDpBlockBackupPolicies) | **Get** /dp-block-backup-policies/ | 
[**UpdateDpBlockBackupPolicy**](DpBlockBackupPoliciesApi.md#UpdateDpBlockBackupPolicy) | **Patch** /dp-block-backup-policies/{policy_id} | 


# **CreateDpBlockBackupPolicy**
> DpBlockBackupPolicyResp CreateDpBlockBackupPolicy(ctx, body)


Create dp block backup policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**DpBlockBackupPolicyCreateReq**](DpBlockBackupPolicyCreateReq.md)| policy info | 

### Return type

[**DpBlockBackupPolicyResp**](DpBlockBackupPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDpBlockBackupPolicy**
> DpBlockBackupPolicyResp DeleteDpBlockBackupPolicy(ctx, policyId, optional)


Delete dp block backup policy

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

[**DpBlockBackupPolicyResp**](DpBlockBackupPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDpBlockBackupPolicy**
> DpBlockBackupPolicyResp GetDpBlockBackupPolicy(ctx, policyId)


Get dp block backup policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyId** | **int64**| resource id | 

### Return type

[**DpBlockBackupPolicyResp**](DpBlockBackupPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDpBlockBackupPolicies**
> DpBlockBackupPoliciesResp ListDpBlockBackupPolicies(ctx, optional)


List dp block backup policies

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
 **blockVolumeId** | **int64**| show dp block backup policies of specific block volume | 

### Return type

[**DpBlockBackupPoliciesResp**](DpBlockBackupPoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDpBlockBackupPolicy**
> DpBlockBackupPolicyResp UpdateDpBlockBackupPolicy(ctx, policyId, body)


Update dp block backup policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyId** | **int64**| resource id | 
  **body** | [**DpBlockBackupPolicyUpdateReq**](DpBlockBackupPolicyUpdateReq.md)| policy info | 

### Return type

[**DpBlockBackupPolicyResp**](DpBlockBackupPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

