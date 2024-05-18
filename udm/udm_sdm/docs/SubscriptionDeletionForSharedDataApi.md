# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UnsubscribeForSharedData**](SubscriptionDeletionForSharedDataApi.md#UnsubscribeForSharedData) | **Delete** /shared-data-subscriptions/{subscriptionId} | unsubscribe from notifications for shared data

# **UnsubscribeForSharedData**
> UnsubscribeForSharedData(ctx, subscriptionId)
unsubscribe from notifications for shared data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| Id of the Shared data Subscription | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

