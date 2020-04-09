# \HostInitializationsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HostInitialization**](HostInitializationsApi.md#HostInitialization) | **Post** /host-initializations/ | 


# **HostInitialization**
> HostInitializationResp HostInitialization(ctx, body)


Create host initialization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**HostInitializationReq**](HostInitializationReq.md)| host initialization info | 

### Return type

[**HostInitializationResp**](HostInitializationResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

