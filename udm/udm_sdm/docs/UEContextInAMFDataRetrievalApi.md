# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUeCtxInAmfData**](UEContextInAMFDataRetrievalApi.md#GetUeCtxInAmfData) | **Get** /{supi}/ue-context-in-amf-data | retrieve a UE&#x27;s UE Context In AMF Data

# **GetUeCtxInAmfData**
> UeContextInAmfData GetUeCtxInAmfData(ctx, supi, optional)
retrieve a UE's UE Context In AMF Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supi** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***UEContextInAMFDataRetrievalApiGetUeCtxInAmfDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UEContextInAMFDataRetrievalApiGetUeCtxInAmfDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 

### Return type

[**UeContextInAmfData**](UeContextInAmfData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

