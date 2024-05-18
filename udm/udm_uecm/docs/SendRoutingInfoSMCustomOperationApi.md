# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendRoutingInfoSm**](SendRoutingInfoSMCustomOperationApi.md#SendRoutingInfoSm) | **Post** /{ueId}/registrations/send-routing-info-sm | Retreive addressing information for SMS delivery

# **SendRoutingInfoSm**
> RoutingInfoSmResponse SendRoutingInfoSm(ctx, body, ueId)
Retreive addressing information for SMS delivery

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutingInfoSmRequest**](RoutingInfoSmRequest.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 

### Return type

[**RoutingInfoSmResponse**](RoutingInfoSmResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

