# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Non3GppRegistration**](AMFRegistrationForNon3GPPAccessApi.md#Non3GppRegistration) | **Put** /{ueId}/registrations/amf-non-3gpp-access | register as AMF for non-3GPP access

# **Non3GppRegistration**
> AmfNon3GppAccessRegistration Non3GppRegistration(ctx, body, ueId)
register as AMF for non-3GPP access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AmfNon3GppAccessRegistration**](AmfNon3GppAccessRegistration.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 

### Return type

[**AmfNon3GppAccessRegistration**](AmfNon3GppAccessRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

