# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSmfRegistration**](ParameterUpdateInTheSMFRegistrationApi.md#UpdateSmfRegistration) | **Patch** /{ueId}/registrations/smf-registrations/{pduSessionId} | update a parameter in the SMF registration

# **UpdateSmfRegistration**
> PatchResult UpdateSmfRegistration(ctx, body, ueId, pduSessionId, optional)
update a parameter in the SMF registration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmfRegistrationModification**](SmfRegistrationModification.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 
  **pduSessionId** | [**int32**](.md)| Identifier of the PDU session | 
 **optional** | ***ParameterUpdateInTheSMFRegistrationApiUpdateSmfRegistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ParameterUpdateInTheSMFRegistrationApiUpdateSmfRegistrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **supportedFeatures** | [**optional.Interface of string**](.md)| Features required to be supported by the target NF | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

