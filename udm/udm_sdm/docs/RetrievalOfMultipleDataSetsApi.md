# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataSets**](RetrievalOfMultipleDataSetsApi.md#GetDataSets) | **Get** /{supi} | retrieve multiple data sets

# **GetDataSets**
> SubscriptionDataSets GetDataSets(ctx, supi, datasetNames, optional)
retrieve multiple data sets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supi** | [**string**](.md)| Identifier of the UE | 
  **datasetNames** | [**[]DataSetName**](.md)| List of dataset names | 
 **optional** | ***RetrievalOfMultipleDataSetsApiGetDataSetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetrievalOfMultipleDataSetsApiGetDataSetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **plmnId** | [**optional.Interface of PlmnIdNid**](.md)| serving PLMN ID | 
 **adjacentPlmns** | [**optional.Interface of []PlmnId**](PlmnId.md)| List of PLMNs adjacent to the UE&#x27;s serving PLMN | 
 **singleNssai** | [**optional.Interface of Snssai**](.md)|  | 
 **dnn** | [**optional.Interface of string**](.md)|  | 
 **ucPurpose** | [**optional.Interface of UcPurpose**](.md)| User consent purpose | 
 **disasterRoamingInd** | **optional.Bool**| Indication whether Disaster Roaming service is applied or not | [default to false]
 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 
 **ifNoneMatch** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**SubscriptionDataSets**](SubscriptionDataSets.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

