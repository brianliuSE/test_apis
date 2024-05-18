# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateNon3GppRegistration**](ParameterUpdateInTheAMFRegistrationForNon3GPPAccessApi.md#UpdateNon3GppRegistration) | **Patch** /{ueId}/registrations/amf-non-3gpp-access | update a parameter in the AMF registration for non-3GPP access

# **UpdateNon3GppRegistration**
> PatchResult UpdateNon3GppRegistration(ctx, body, ueId, optional)
update a parameter in the AMF registration for non-3GPP access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AmfNon3GppAccessRegistrationModification**](AmfNon3GppAccessRegistrationModification.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***ParameterUpdateInTheAMFRegistrationForNon3GPPAccessApiUpdateNon3GppRegistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ParameterUpdateInTheAMFRegistrationForNon3GPPAccessApiUpdateNon3GppRegistrationOpts struct
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

