# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscribeToSharedData**](SubscriptionCreationForSharedDataApi.md#SubscribeToSharedData) | **Post** /shared-data-subscriptions | subscribe to notifications for shared data

# **SubscribeToSharedData**
> SdmSubscription SubscribeToSharedData(ctx, body)
subscribe to notifications for shared data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdmSubscription**](SdmSubscription.md)|  | 

### Return type

[**SdmSubscription**](SdmSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

