# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IpSmGwDeregistration**](IPSMGWDeregistrationApi.md#IpSmGwDeregistration) | **Delete** /{ueId}/registrations/ip-sm-gw | Delete the IP-SM-GW registration

# **IpSmGwDeregistration**
> IpSmGwDeregistration(ctx, ueId)
Delete the IP-SM-GW registration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ueId** | [**string**](.md)| Identifier of the UE | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

