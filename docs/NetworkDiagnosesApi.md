# \NetworkDiagnosesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkDiagnosis**](NetworkDiagnosesApi.md#CreateNetworkDiagnosis) | **Post** /network-diagnoses/ | 
[**DeleteNetworkDiagnosis**](NetworkDiagnosesApi.md#DeleteNetworkDiagnosis) | **Delete** /network-diagnoses/{network_diagnosis_id} | 
[**GetNetworkDiagnosis**](NetworkDiagnosesApi.md#GetNetworkDiagnosis) | **Get** /network-diagnoses/{network_diagnosis_id} | 
[**ListNetworkDiagnoses**](NetworkDiagnosesApi.md#ListNetworkDiagnoses) | **Get** /network-diagnoses/ | 
[**StopNetworkDiagnosis**](NetworkDiagnosesApi.md#StopNetworkDiagnosis) | **Post** /network-diagnoses/{network_diagnosis_id}:stop | 


# **CreateNetworkDiagnosis**
> NetworkDiagnosisResp CreateNetworkDiagnosis(ctx, body)


Create network diagnosis

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**NetworkDiagnosisCreateReq**](NetworkDiagnosisCreateReq.md)| network diagnosis info | 

### Return type

[**NetworkDiagnosisResp**](NetworkDiagnosisResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetworkDiagnosis**
> DeleteNetworkDiagnosis(ctx, networkDiagnosisId)


delete network diagnosis

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **networkDiagnosisId** | **int64**| network diagnosis id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkDiagnosis**
> NetworkDiagnosisResp GetNetworkDiagnosis(ctx, networkDiagnosisId)


Get network diagnosis

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **networkDiagnosisId** | **int64**| network diagnosis id | 

### Return type

[**NetworkDiagnosisResp**](NetworkDiagnosisResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNetworkDiagnoses**
> NetworkDiagnosesResp ListNetworkDiagnoses(ctx, optional)


List network diagnoses

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

[**NetworkDiagnosesResp**](NetworkDiagnosesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopNetworkDiagnosis**
> NetworkDiagnosisResp StopNetworkDiagnosis(ctx, networkDiagnosisId)


stop network diagnosis

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **networkDiagnosisId** | **int64**| network diagnosis id | 

### Return type

[**NetworkDiagnosisResp**](NetworkDiagnosisResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

