# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthTrigger**](TriggerThePrimaryReAuthenticationApi.md#AuthTrigger) | **Post** /{ueId}/registrations/auth-trigger | trigger the primary (re-)authentication

# **AuthTrigger**
> AuthTrigger(ctx, body, ueId)
trigger the primary (re-)authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthTriggerInfo**](AuthTriggerInfo.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

