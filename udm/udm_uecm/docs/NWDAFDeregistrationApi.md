# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NwdafDeregistration**](NWDAFDeregistrationApi.md#NwdafDeregistration) | **Delete** /{ueId}/registrations/nwdaf-registrations/{nwdafRegistrationId} | delete an NWDAF registration

# **NwdafDeregistration**
> NwdafDeregistration(ctx, ueId, nwdafRegistrationId)
delete an NWDAF registration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 
  **nwdafRegistrationId** | **string**| NWDAF registration identifier | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

