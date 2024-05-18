# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSmfRegistration**](SMFSmfRegistrationApi.md#GetSmfRegistration) | **Get** /{ueId}/registrations/smf-registrations | retrieve the SMF registration information
[**Registration**](SMFSmfRegistrationApi.md#Registration) | **Put** /{ueId}/registrations/smf-registrations/{pduSessionId} | register as SMF

# **GetSmfRegistration**
> SmfRegistrationInfo GetSmfRegistration(ctx, ueId, optional)
retrieve the SMF registration information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***SMFSmfRegistrationApiGetSmfRegistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SMFSmfRegistrationApiGetSmfRegistrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **singleNssai** | [**optional.Interface of Snssai**](.md)|  | 
 **dnn** | [**optional.Interface of string**](.md)|  | 
 **supportedFeatures** | [**optional.Interface of string**](.md)|  | 

### Return type

[**SmfRegistrationInfo**](SmfRegistrationInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Registration**
> SmfRegistration Registration(ctx, body, ueId, pduSessionId)
register as SMF

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmfRegistration**](SmfRegistration.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 
  **pduSessionId** | [**int32**](.md)| Identifier of the PDU session | 

### Return type

[**SmfRegistration**](SmfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

