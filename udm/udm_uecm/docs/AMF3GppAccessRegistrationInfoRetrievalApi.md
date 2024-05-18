# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get3GppRegistration**](AMF3GppAccessRegistrationInfoRetrievalApi.md#Get3GppRegistration) | **Get** /{ueId}/registrations/amf-3gpp-access | retrieve the AMF registration for 3GPP access information

# **Get3GppRegistration**
> Amf3GppAccessRegistration Get3GppRegistration(ctx, ueId, optional)
retrieve the AMF registration for 3GPP access information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***AMF3GppAccessRegistrationInfoRetrievalApiGet3GppRegistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AMF3GppAccessRegistrationInfoRetrievalApiGet3GppRegistrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | [**optional.Interface of string**](.md)|  | 

### Return type

[**Amf3GppAccessRegistration**](Amf3GppAccessRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

