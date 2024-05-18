# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSupiOrGpsi**](GPSIToSUPITranslationOrSUPIToGPSITranslationApi.md#GetSupiOrGpsi) | **Get** /{ueId}/id-translation-result | retrieve a UE&#x27;s SUPI or GPSI

# **GetSupiOrGpsi**
> IdTranslationResult GetSupiOrGpsi(ctx, ueId, optional)
retrieve a UE's SUPI or GPSI

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***GPSIToSUPITranslationOrSUPIToGPSITranslationApiGetSupiOrGpsiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GPSIToSUPITranslationOrSUPIToGPSITranslationApiGetSupiOrGpsiOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 
 **afId** | **optional.String**| AF identifier | 
 **appPortId** | [**optional.Interface of AppPortId**](.md)| Application port identifier | 
 **afServiceId** | **optional.String**| AF Service Identifier | 
 **mtcProviderInfo** | [**optional.Interface of string**](.md)| MTC Provider Information | 
 **requestedGpsiType** | [**optional.Interface of GpsiType**](.md)| Requested GPSI Type | 
 **ifNoneMatch** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**IdTranslationResult**](IdTranslationResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

