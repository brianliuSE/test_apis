# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Subscribe**](SubscriptionCreationApi.md#Subscribe) | **Post** /{ueId}/sdm-subscriptions | subscribe to notifications

# **Subscribe**
> SdmSubscription Subscribe(ctx, body, ueId, optional)
subscribe to notifications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdmSubscription**](SdmSubscription.md)|  | 
  **ueId** | [**string**](.md)| Identity of the user | 
 **optional** | ***SubscriptionCreationApiSubscribeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionCreationApiSubscribeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sharedDataIds** | [**optional.Interface of []string**](string.md)| List of IDs identifying shared Data already available at and subscribed by the NF service consumer  | 

### Return type

[**SdmSubscription**](SdmSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

