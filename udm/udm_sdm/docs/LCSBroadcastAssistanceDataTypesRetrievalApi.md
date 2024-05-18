# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLcsBcaData**](LCSBroadcastAssistanceDataTypesRetrievalApi.md#GetLcsBcaData) | **Get** /{supi}/lcs-bca-data | retrieve a UE&#x27;s LCS Broadcast Assistance Data Types Subscription Data

# **GetLcsBcaData**
> LcsBroadcastAssistanceTypesData GetLcsBcaData(ctx, supi, optional)
retrieve a UE's LCS Broadcast Assistance Data Types Subscription Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supi** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***LCSBroadcastAssistanceDataTypesRetrievalApiGetLcsBcaDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LCSBroadcastAssistanceDataTypesRetrievalApiGetLcsBcaDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 
 **plmnId** | [**optional.Interface of PlmnId**](.md)|  | 
 **ifNoneMatch** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**LcsBroadcastAssistanceTypesData**](LcsBroadcastAssistanceTypesData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

