# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**3GppRegistration**](AMFRegistrationFor3GPPAccessApi.md#3GppRegistration) | **Put** /{ueId}/registrations/amf-3gpp-access | register as AMF for 3GPP access

# **3GppRegistration**
> Amf3GppAccessRegistration 3GppRegistration(ctx, body, ueId)
register as AMF for 3GPP access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Amf3GppAccessRegistration**](Amf3GppAccessRegistration.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 

### Return type

[**Amf3GppAccessRegistration**](Amf3GppAccessRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

