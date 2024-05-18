# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**3GppSmsfDeregistration**](SMSFDeregistrationFor3GPPAccessApi.md#3GppSmsfDeregistration) | **Delete** /{ueId}/registrations/smsf-3gpp-access | delete the SMSF registration for 3GPP access

# **3GppSmsfDeregistration**
> 3GppSmsfDeregistration(ctx, ueId, optional)
delete the SMSF registration for 3GPP access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***SMSFDeregistrationFor3GPPAccessApi3GppSmsfDeregistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SMSFDeregistrationFor3GPPAccessApi3GppSmsfDeregistrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smsfSetId** | [**optional.Interface of string**](.md)|  | 
 **ifMatch** | **optional.String**| Validator for conditional requests, as described in IETF RFC 9110, 3.1 | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

