# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIpSmGwRegistration**](IPSMGWRegistrationInfoRetrievalApi.md#GetIpSmGwRegistration) | **Get** /{ueId}/registrations/ip-sm-gw | Retrieve the IP-SM-GW registration information

# **GetIpSmGwRegistration**
> IpSmGwRegistration GetIpSmGwRegistration(ctx, ueId)
Retrieve the IP-SM-GW registration information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 

### Return type

[**IpSmGwRegistration**](IpSmGwRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

