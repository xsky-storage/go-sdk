# \OsPoliciesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicy**](OsPoliciesApi.md#CreatePolicy) | **Post** /os-policies/ | 
[**DeletePolicy**](OsPoliciesApi.md#DeletePolicy) | **Delete** /os-policies/{policy_id} | 
[**GetPolicy**](OsPoliciesApi.md#GetPolicy) | **Get** /os-policies/{policy_id} | 
[**ListPolicies**](OsPoliciesApi.md#ListPolicies) | **Get** /os-policies/ | 
[**UpdatePolicy**](OsPoliciesApi.md#UpdatePolicy) | **Patch** /os-policies/{policy_id} | 


# **CreatePolicy**
> ObjectStoragePolicyResp CreatePolicy(ctx, body)


Create object storage policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ObjectStoragePolicyCreateReq**](ObjectStoragePolicyCreateReq.md)| policy info | 

### Return type

[**ObjectStoragePolicyResp**](ObjectStoragePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicy**
> ObjectStoragePolicyResp DeletePolicy(ctx, policyId)


Delete object storage policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyId** | **int64**| policy id | 

### Return type

[**ObjectStoragePolicyResp**](ObjectStoragePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicy**
> ObjectStoragePolicyResp GetPolicy(ctx, policyId)


Get object storage policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyId** | **int64**| policy id | 

### Return type

[**ObjectStoragePolicyResp**](ObjectStoragePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicies**
> ObjectStoragePoliciesResp ListPolicies(ctx, optional)


List object storage policies

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

[**ObjectStoragePoliciesResp**](ObjectStoragePoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicy**
> ObjectStoragePolicyResp UpdatePolicy(ctx, policyId, body)


Update object storage policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policyId** | **int64**| policy id | 
  **body** | [**ObjectStoragePolicyUpdateReq**](ObjectStoragePolicyUpdateReq.md)| policy info | 

### Return type

[**ObjectStoragePolicyResp**](ObjectStoragePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

