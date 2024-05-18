# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNwdafRegistration**](NWDAFRegistrationInfoRetrievalApi.md#GetNwdafRegistration) | **Get** /{ueId}/registrations/nwdaf-registrations | retrieve the NWDAF registration

# **GetNwdafRegistration**
> []NwdafRegistration GetNwdafRegistration(ctx, ueId, optional)
retrieve the NWDAF registration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***NWDAFRegistrationInfoRetrievalApiGetNwdafRegistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NWDAFRegistrationInfoRetrievalApiGetNwdafRegistrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyticsIds** | [**optional.Interface of []EventId**](EventId.md)| List of analytics Id(s) provided by the consumers of NWDAF. | 
 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 

### Return type

[**[]NwdafRegistration**](NwdafRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

