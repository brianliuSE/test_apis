# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUeCtxInSmsfData**](UEContextInSMSFDataRetrievalApi.md#GetUeCtxInSmsfData) | **Get** /{supi}/ue-context-in-smsf-data | retrieve a UE&#x27;s UE Context In SMSF Data

# **GetUeCtxInSmsfData**
> UeContextInSmsfData GetUeCtxInSmsfData(ctx, supi, optional)
retrieve a UE's UE Context In SMSF Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supi** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***UEContextInSMSFDataRetrievalApiGetUeCtxInSmsfDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UEContextInSMSFDataRetrievalApiGetUeCtxInSmsfDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 

### Return type

[**UeContextInSmsfData**](UeContextInSmsfData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

