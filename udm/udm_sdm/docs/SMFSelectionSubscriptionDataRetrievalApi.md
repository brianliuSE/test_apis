# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSmfSelData**](SMFSelectionSubscriptionDataRetrievalApi.md#GetSmfSelData) | **Get** /{supi}/smf-select-data | retrieve a UE&#x27;s SMF Selection Subscription Data

# **GetSmfSelData**
> SmfSelectionSubscriptionData GetSmfSelData(ctx, supi, optional)
retrieve a UE's SMF Selection Subscription Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supi** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***SMFSelectionSubscriptionDataRetrievalApiGetSmfSelDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SMFSelectionSubscriptionDataRetrievalApiGetSmfSelDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 
 **plmnId** | [**optional.Interface of PlmnId**](.md)| serving PLMN ID | 
 **disasterRoamingInd** | **optional.Bool**| Indication whether Disaster Roaming service is applied or not | [default to false]
 **ifNoneMatch** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**SmfSelectionSubscriptionData**](SmfSelectionSubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

