# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSmsf3GppRegistration**](ParameterUpdateInTheSMSFRegistrationFor3GPPAccessApi.md#UpdateSmsf3GppRegistration) | **Patch** /{ueId}/registrations/smsf-3gpp-access | update a parameter in the SMSF registration for 3GPP access

# **UpdateSmsf3GppRegistration**
> PatchResult UpdateSmsf3GppRegistration(ctx, body, ueId, optional)
update a parameter in the SMSF registration for 3GPP access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmsfRegistrationModification**](SmsfRegistrationModification.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***ParameterUpdateInTheSMSFRegistrationFor3GPPAccessApiUpdateSmsf3GppRegistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ParameterUpdateInTheSMSFRegistrationFor3GPPAccessApiUpdateSmsf3GppRegistrationOpts struct
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

