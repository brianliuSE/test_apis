# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUeCtxInSmfData**](UEContextInSMFDataRetrievalApi.md#GetUeCtxInSmfData) | **Get** /{supi}/ue-context-in-smf-data | retrieve a UE&#x27;s UE Context In SMF Data

# **GetUeCtxInSmfData**
> UeContextInSmfData GetUeCtxInSmfData(ctx, supi, optional)
retrieve a UE's UE Context In SMF Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supi** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***UEContextInSMFDataRetrievalApiGetUeCtxInSmfDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UEContextInSMFDataRetrievalApiGetUeCtxInSmfDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 

### Return type

[**UeContextInSmfData**](UeContextInSmfData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

