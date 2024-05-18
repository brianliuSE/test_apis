# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateNwdafRegistration**](ParameterUpdateInTheNWDAFRegistrationApi.md#UpdateNwdafRegistration) | **Patch** /{ueId}/registrations/nwdaf-registrations/{nwdafRegistrationId} | Update a parameter in the NWDAF registration

# **UpdateNwdafRegistration**
> InlineResponse200 UpdateNwdafRegistration(ctx, body, ueId, nwdafRegistrationId, optional)
Update a parameter in the NWDAF registration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NwdafRegistrationModification**](NwdafRegistrationModification.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 
  **nwdafRegistrationId** | **string**| NWDAF registration identifier | 
 **optional** | ***ParameterUpdateInTheNWDAFRegistrationApiUpdateNwdafRegistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ParameterUpdateInTheNWDAFRegistrationApiUpdateNwdafRegistrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

