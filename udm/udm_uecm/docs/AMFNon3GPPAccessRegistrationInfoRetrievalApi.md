# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNon3GppRegistration**](AMFNon3GPPAccessRegistrationInfoRetrievalApi.md#GetNon3GppRegistration) | **Get** /{ueId}/registrations/amf-non-3gpp-access | retrieve the AMF registration for non-3GPP access information

# **GetNon3GppRegistration**
> AmfNon3GppAccessRegistration GetNon3GppRegistration(ctx, ueId, optional)
retrieve the AMF registration for non-3GPP access information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***AMFNon3GPPAccessRegistrationInfoRetrievalApiGetNon3GppRegistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AMFNon3GPPAccessRegistrationInfoRetrievalApiGetNon3GppRegistrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | [**optional.Interface of string**](.md)|  | 

### Return type

[**AmfNon3GppAccessRegistration**](AmfNon3GppAccessRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

