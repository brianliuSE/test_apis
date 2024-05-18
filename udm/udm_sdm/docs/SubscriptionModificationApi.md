# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Modify**](SubscriptionModificationApi.md#Modify) | **Patch** /{ueId}/sdm-subscriptions/{subscriptionId} | modify the subscription
[**ModifySharedDataSubs**](SubscriptionModificationApi.md#ModifySharedDataSubs) | **Patch** /shared-data-subscriptions/{subscriptionId} | modify the subscription

# **Modify**
> InlineResponse200 Modify(ctx, body, ueId, subscriptionId, optional)
modify the subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdmSubsModification**](SdmSubsModification.md)|  | 
  **ueId** | [**string**](.md)| Identity of the user | 
  **subscriptionId** | **string**| Id of the SDM Subscription | 
 **optional** | ***SubscriptionModificationApiModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionModificationApiModifyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **supportedFeatures** | [**optional.Interface of string**](.md)| Features required to be supported by the target NF | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifySharedDataSubs**
> InlineResponse200 ModifySharedDataSubs(ctx, body, subscriptionId, optional)
modify the subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdmSubsModification**](SdmSubsModification.md)|  | 
  **subscriptionId** | **string**| Id of the SDM Subscription | 
 **optional** | ***SubscriptionModificationApiModifySharedDataSubsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionModificationApiModifySharedDataSubsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **supportedFeatures** | [**optional.Interface of string**](.md)| Features required to be supported by the target NF | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

