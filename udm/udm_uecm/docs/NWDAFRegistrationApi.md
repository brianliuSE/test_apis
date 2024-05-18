# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NwdafRegistration**](NWDAFRegistrationApi.md#NwdafRegistration) | **Put** /{ueId}/registrations/nwdaf-registrations/{nwdafRegistrationId} | register as NWDAF

# **NwdafRegistration**
> NwdafRegistration NwdafRegistration(ctx, body, ueId, nwdafRegistrationId)
register as NWDAF

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NwdafRegistration**](NwdafRegistration.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 
  **nwdafRegistrationId** | **string**| NWDAF registration identifier | 

### Return type

[**NwdafRegistration**](NwdafRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

