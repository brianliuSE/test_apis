# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**3GppSmsfRegistration**](SMSFRegistrationFor3GPPAccessApi.md#3GppSmsfRegistration) | **Put** /{ueId}/registrations/smsf-3gpp-access | register as SMSF for 3GPP access

# **3GppSmsfRegistration**
> SmsfRegistration 3GppSmsfRegistration(ctx, body, ueId)
register as SMSF for 3GPP access

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

