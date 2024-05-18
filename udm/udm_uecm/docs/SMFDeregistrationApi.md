# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmfDeregistration**](SMFDeregistrationApi.md#SmfDeregistration) | **Delete** /{ueId}/registrations/smf-registrations/{pduSessionId} | delete an SMF registration

# **SmfDeregistration**
> SmfDeregistration(ctx, ueId, pduSessionId, optional)
delete an SMF registration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 
  **pduSessionId** | [**int32**](.md)| Identifier of the PDU session | 
 **optional** | ***SMFDeregistrationApiSmfDeregistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SMFDeregistrationApiSmfDeregistrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **smfSetId** | [**optional.Interface of string**](.md)|  | 
 **smfInstanceId** | [**optional.Interface of string**](.md)|  | 
 **smfEventsImplicitlyUnsubscribed** | **optional.Bool**| Indication on SMF event subscriptions implicitly unsubscribed. | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

