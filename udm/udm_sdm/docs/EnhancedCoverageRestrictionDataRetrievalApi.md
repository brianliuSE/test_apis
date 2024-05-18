# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEcrData**](EnhancedCoverageRestrictionDataRetrievalApi.md#GetEcrData) | **Get** /{supi}/am-data/ecr-data | retrieve a UE&#x27;s subscribed Enhanced Coverage Restriction Data

# **GetEcrData**
> EnhancedCoverageRestrictionData GetEcrData(ctx, supi, optional)
retrieve a UE's subscribed Enhanced Coverage Restriction Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supi** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***EnhancedCoverageRestrictionDataRetrievalApiGetEcrDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnhancedCoverageRestrictionDataRetrievalApiGetEcrDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 
 **ifNoneMatch** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**EnhancedCoverageRestrictionData**](EnhancedCoverageRestrictionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

