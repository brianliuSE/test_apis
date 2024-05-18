# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSORInfo**](TriggerSORInfoUpdateApi.md#UpdateSORInfo) | **Post** /{supi}/am-data/update-sor | Nudm_Sdm custom operation to trigger SOR info update

# **UpdateSORInfo**
> SorInfo UpdateSORInfo(ctx, supi, optional)
Nudm_Sdm custom operation to trigger SOR info update

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supi** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***TriggerSORInfoUpdateApiUpdateSORInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggerSORInfoUpdateApiUpdateSORInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SorUpdateInfo**](SorUpdateInfo.md)|  | 

### Return type

[**SorInfo**](SorInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

