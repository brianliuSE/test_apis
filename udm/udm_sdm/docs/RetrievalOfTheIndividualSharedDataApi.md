# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIndividualSharedData**](RetrievalOfTheIndividualSharedDataApi.md#GetIndividualSharedData) | **Get** /shared-data/{sharedDataId} | retrieve the individual shared data

# **GetIndividualSharedData**
> SharedData GetIndividualSharedData(ctx, sharedDataId, optional)
retrieve the individual shared data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sharedDataId** | [**[]string**](.md)| Id of the Shared data | 
 **optional** | ***RetrievalOfTheIndividualSharedDataApiGetIndividualSharedDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetrievalOfTheIndividualSharedDataApiGetIndividualSharedDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 
 **ifNoneMatch** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**SharedData**](SharedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

