# {{classname}}

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateRoamingInformation**](RoamingInformationUpdateApi.md#UpdateRoamingInformation) | **Post** /{ueId}/registrations/amf-3gpp-access/roaming-info-update | Update the Roaming Information

# **UpdateRoamingInformation**
> RoamingInfoUpdate UpdateRoamingInformation(ctx, body, ueId)
Update the Roaming Information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoamingInfoUpdate**](RoamingInfoUpdate.md)|  | 
  **ueId** | [**string**](.md)| Identifier of the UE | 

### Return type

[**RoamingInfoUpdate**](RoamingInfoUpdate.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

