# \SslCertificatesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSLCertificate**](SslCertificatesApi.md#CreateSSLCertificate) | **Post** /ssl-certificates/ | 
[**DeleteSSLCertificate**](SslCertificatesApi.md#DeleteSSLCertificate) | **Delete** /ssl-certificates/{certificate_id} | 
[**GetSSLCertificate**](SslCertificatesApi.md#GetSSLCertificate) | **Get** /ssl-certificates/{certificate_id} | 
[**ListSSLCertificates**](SslCertificatesApi.md#ListSSLCertificates) | **Get** /ssl-certificates/ | 
[**UpdateSSLCertificate**](SslCertificatesApi.md#UpdateSSLCertificate) | **Patch** /ssl-certificates/{certificate_id} | 


# **CreateSSLCertificate**
> SslCertificateResp CreateSSLCertificate(ctx, body)


Create ssl certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**SslCertificateCreateReq**](SslCertificateCreateReq.md)| ssl certificate info | 

### Return type

[**SslCertificateResp**](SSLCertificateResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSSLCertificate**
> SslCertificateResp DeleteSSLCertificate(ctx, certificateId)


Delete ssl certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **certificateId** | **int64**| certificate id | 

### Return type

[**SslCertificateResp**](SSLCertificateResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSSLCertificate**
> SslCertificateResp GetSSLCertificate(ctx, certificateId, optional)


Get certificates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **certificateId** | **int64**| certificate id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateId** | **int64**| certificate id | 
 **limit** | **int64**| paging param | 
 **offset** | **int64**| paging param | 

### Return type

[**SslCertificateResp**](SSLCertificateResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSSLCertificates**
> SslCertificatesResp ListSSLCertificates(ctx, optional)


List certificates

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

[**SslCertificatesResp**](SSLCertificatesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSSLCertificate**
> SslCertificateResp UpdateSSLCertificate(ctx, certificateId, body)


Update ssl certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **certificateId** | **int64**| certificate id | 
  **body** | [**SslCertificateUpdateReq**](SslCertificateUpdateReq.md)| ssl certificate info | 

### Return type

[**SslCertificateResp**](SSLCertificateResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

