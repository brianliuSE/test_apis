# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMultipleIdentifiers**](MultipleIdentifiersApi.md#GetMultipleIdentifiers) | **Get** /multiple-identifiers | Mapping of UE Identifiers

# **GetMultipleIdentifiers**
> map[string]SupiInfo GetMultipleIdentifiers(ctx, gpsiList, optional)
Mapping of UE Identifiers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gpsiList** | [**[]string**](string.md)| list of the GPSIs | 
 **optional** | ***MultipleIdentifiersApiGetMultipleIdentifiersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MultipleIdentifiersApiGetMultipleIdentifiersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 

### Return type

[**map[string]SupiInfo**](map.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

