# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeregAMF**](TriggerAMFFor3GPPAccessDeregistrationApi.md#DeregAMF) | **Post** /{ueId}/registrations/amf-3gpp-access/dereg-amf | trigger AMF for 3GPP access deregistration

# **DeregAMF**
> DeregAMF(ctx, body, ueId)
trigger AMF for 3GPP access deregistration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AmfDeregInfo**](AmfDeregInfo.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

