# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Non3GppSmsfRegistration**](SMSFRegistrationForNon3GPPAccessApi.md#Non3GppSmsfRegistration) | **Put** /{ueId}/registrations/smsf-non-3gpp-access | register as SMSF for non-3GPP access

# **Non3GppSmsfRegistration**
> SmsfRegistration Non3GppSmsfRegistration(ctx, body, ueId)
register as SMSF for non-3GPP access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmsfRegistration**](SmsfRegistration.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 

### Return type

[**SmsfRegistration**](SmsfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

