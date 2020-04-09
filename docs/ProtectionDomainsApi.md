# \ProtectionDomainsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProtectionDomain**](ProtectionDomainsApi.md#GetProtectionDomain) | **Get** /protection-domains/{protection_domain_id} | 
[**ListProtectionDomains**](ProtectionDomainsApi.md#ListProtectionDomains) | **Get** /protection-domains/ | 


# **GetProtectionDomain**
> ProtectionDomainRecordResp GetProtectionDomain(ctx, protectionDomainId)


get protection domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **protectionDomainId** | **int64**| protection domain id | 

### Return type

[**ProtectionDomainRecordResp**](ProtectionDomainRecordResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProtectionDomains**
> ProtectionDomainRecordsResp ListProtectionDomains(ctx, optional)


List protection domains

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

[**ProtectionDomainRecordsResp**](ProtectionDomainRecordsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

