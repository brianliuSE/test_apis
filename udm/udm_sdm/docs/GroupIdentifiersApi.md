# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupIdentifiers**](GroupIdentifiersApi.md#GetGroupIdentifiers) | **Get** /group-data/group-identifiers | Mapping of Group Identifiers

# **GetGroupIdentifiers**
> GroupIdentifiers GetGroupIdentifiers(ctx, optional)
Mapping of Group Identifiers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupIdentifiersApiGetGroupIdentifiersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupIdentifiersApiGetGroupIdentifiersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extGroupId** | [**optional.Interface of string**](.md)| External Group Identifier | 
 **intGroupId** | [**optional.Interface of string**](.md)| Internal Group Identifier | 
 **ueIdInd** | **optional.Bool**| Indication whether UE identifiers are required or not | [default to false]
 **supportedFeatures** | [**optional.Interface of string**](.md)| Supported Features | 
 **afId** | **optional.String**| AF identifier | 
 **ifNoneMatch** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **optional.String**| Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**GroupIdentifiers**](GroupIdentifiers.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

