# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAmData**](AccessAndMobilitySubscriptionDataRetrievalApi.md#GetAmData) | **Get** /{supi}/am-data | retrieve a UE&#x27;s Access and Mobility Subscription Data

# **GetAmData**
> AccessAndMobilitySubscriptionData GetAmData(ctx, supi, optional)
retrieve a UE's Access and Mobility Subscription Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supi** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***AccessAndMobilitySubscriptionDataRetrievalApiGetAmDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccessAndMobilitySubscriptionDataRetrievalApiGetAmDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 
 **plmnId** | [**optional.Interface of PlmnIdNid**](.md)| Serving PLMN ID or SNPN ID | 
 **adjacentPlmns** | [**optional.Interface of []PlmnId**](PlmnId.md)| List of PLMNs adjacent to the UE&#x27;s serving PLMN | 
 **disasterRoamingInd** | **optional.Bool**| Indication whether Disaster Roaming service is applied or not | [default to false]
 **sharedDataIds** | [**optional.Interface of []string**](string.md)| List of IDs identifying shared Access and Mobility Subscription Data already available at the NF service consumer  | 
 **ifNoneMatch** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**AccessAndMobilitySubscriptionData**](AccessAndMobilitySubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

