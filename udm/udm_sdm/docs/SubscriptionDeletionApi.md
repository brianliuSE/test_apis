# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Unsubscribe**](SubscriptionDeletionApi.md#Unsubscribe) | **Delete** /{ueId}/sdm-subscriptions/{subscriptionId} | unsubscribe from notifications

# **Unsubscribe**
> Unsubscribe(ctx, ueId, subscriptionId)
unsubscribe from notifications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identity of the user | 
  **subscriptionId** | **string**| Id of the SDM Subscription | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

