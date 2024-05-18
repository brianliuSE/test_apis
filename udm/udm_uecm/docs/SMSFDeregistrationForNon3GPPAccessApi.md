# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Non3GppSmsfDeregistration**](SMSFDeregistrationForNon3GPPAccessApi.md#Non3GppSmsfDeregistration) | **Delete** /{ueId}/registrations/smsf-non-3gpp-access | delete SMSF registration for non 3GPP access

# **Non3GppSmsfDeregistration**
> Non3GppSmsfDeregistration(ctx, ueId, optional)
delete SMSF registration for non 3GPP access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***SMSFDeregistrationForNon3GPPAccessApiNon3GppSmsfDeregistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SMSFDeregistrationForNon3GPPAccessApiNon3GppSmsfDeregistrationOpts struct
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

