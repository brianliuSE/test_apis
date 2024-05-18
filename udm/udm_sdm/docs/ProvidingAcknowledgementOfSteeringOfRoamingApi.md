# {{classname}}

All URIs are relative to *{apiRoot}/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SorAckInfo**](ProvidingAcknowledgementOfSteeringOfRoamingApi.md#SorAckInfo) | **Put** /{supi}/am-data/sor-ack | Nudm_Sdm Info service operation

# **SorAckInfo**
> SorAckInfo(ctx, supi, optional)
Nudm_Sdm Info service operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supi** | [**string**](.md)| Identifier of the UE | 
 **optional** | ***ProvidingAcknowledgementOfSteeringOfRoamingApiSorAckInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvidingAcknowledgementOfSteeringOfRoamingApiSorAckInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AcknowledgeInfo**](AcknowledgeInfo.md)|  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

