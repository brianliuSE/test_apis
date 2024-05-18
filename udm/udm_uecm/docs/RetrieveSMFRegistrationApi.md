# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveSmfRegistration**](RetrieveSMFRegistrationApi.md#RetrieveSmfRegistration) | **Get** /{ueId}/registrations/smf-registrations/{pduSessionId} | get an SMF registration

# **RetrieveSmfRegistration**
> SmfRegistration RetrieveSmfRegistration(ctx, ueId, pduSessionId)
get an SMF registration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 
  **pduSessionId** | [**int32**](.md)| Identifier of the PDU session | 

### Return type

[**SmfRegistration**](SmfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

