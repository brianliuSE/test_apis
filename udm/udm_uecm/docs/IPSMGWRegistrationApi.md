# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IpSmGwRegistration**](IPSMGWRegistrationApi.md#IpSmGwRegistration) | **Put** /{ueId}/registrations/ip-sm-gw | Register an IP-SM-GW

# **IpSmGwRegistration**
> IpSmGwRegistration IpSmGwRegistration(ctx, body, ueId)
Register an IP-SM-GW

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSmGwRegistration**](IpSmGwRegistration.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 

### Return type

[**IpSmGwRegistration**](IpSmGwRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

