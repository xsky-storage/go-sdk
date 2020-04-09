# \DomainUserValidatorsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomainUserValidator**](DomainUserValidatorsApi.md#CreateDomainUserValidator) | **Post** /domain-user-validators/ | 
[**GetDomainUserValidator**](DomainUserValidatorsApi.md#GetDomainUserValidator) | **Get** /domain-user-validators/{domain_user_validator_id} | 


# **CreateDomainUserValidator**
> DomainUserValidatorResp CreateDomainUserValidator(ctx, body)


Create domain user validator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**DomainUserValidatorCreateReq**](DomainUserValidatorCreateReq.md)| domain user validator info | 

### Return type

[**DomainUserValidatorResp**](DomainUserValidatorResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDomainUserValidator**
> DomainUserValidatorResp GetDomainUserValidator(ctx, domainUserValidatorId)


Get domain_user validator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **domainUserValidatorId** | **int64**| domain user validator id | 

### Return type

[**DomainUserValidatorResp**](DomainUserValidatorResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

