# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRegistrations**](UECMRegistrationInfoRetrievalApi.md#GetRegistrations) | **Get** /{ueId}/registrations | retrieve UE registration data sets

# **GetRegistrations**
> RegistrationDataSets GetRegistrations(ctx, ueId, registrationDatasetNames, optional)
retrieve UE registration data sets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 
  **registrationDatasetNames** | [**[]RegistrationDataSetName**](.md)| List of UECM registration dataset names | 
 **optional** | ***UECMRegistrationInfoRetrievalApiGetRegistrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UECMRegistrationInfoRetrievalApiGetRegistrationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **supportedFeatures** | [**optional.Interface of string**](.md)|  | 
 **singleNssai** | [**optional.Interface of Snssai**](.md)|  | 
 **dnn** | [**optional.Interface of string**](.md)|  | 

### Return type

[**RegistrationDataSets**](RegistrationDataSets.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

