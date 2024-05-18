# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get3GppSmsfRegistration**](SMSF3GPPAccessRegistrationInfoRetrievalApi.md#Get3GppSmsfRegistration) | **Get** /{ueId}/registrations/smsf-3gpp-access | retrieve the SMSF registration for 3GPP access information

# **Get3GppSmsfRegistration**
> SmsfRegistration Get3GppSmsfRegistration(ctx, ueId, optional)
retrieve the SMSF registration for 3GPP access information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***SMSF3GPPAccessRegistrationInfoRetrievalApiGet3GppSmsfRegistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SMSF3GPPAccessRegistrationInfoRetrievalApiGet3GppSmsfRegistrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | [**optional.Interface of string**](.md)|  | 

### Return type

[**SmsfRegistration**](SmsfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

