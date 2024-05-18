# SdmSubscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** |  | [default to null]
**ImplicitUnsubscribe** | **bool** |  | [optional] [default to null]
**Expires** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**CallbackReference** | **string** |  | [default to null]
**AmfServiceName** | [***ServiceName**](ServiceName.md) |  | [optional] [default to null]
**MonitoredResourceUris** | **[]string** |  | [default to null]
**SingleNssai** | [***Snssai**](Snssai.md) |  | [optional] [default to null]
**Dnn** | **string** |  | [optional] [default to null]
**SubscriptionId** | **string** |  | [optional] [default to null]
**PlmnId** | [***PlmnId**](PlmnId.md) |  | [optional] [default to null]
**ImmediateReport** | **bool** |  | [optional] [default to false]
**Report** | [***ImmediateReport**](ImmediateReport.md) |  | [optional] [default to null]
**SupportedFeatures** | **string** |  | [optional] [default to null]
**ContextInfo** | [***ContextInfo**](ContextInfo.md) |  | [optional] [default to null]
**NfChangeFilter** | **bool** |  | [optional] [default to false]
**UniqueSubscription** | **bool** |  | [optional] [default to null]
**ResetIds** | **[]string** |  | [optional] [default to null]
**UeConSmfDataSubFilter** | [***UeContextInSmfDataSubFilter**](UeContextInSmfDataSubFilter.md) |  | [optional] [default to null]
**AdjacentPlmns** | [**[]PlmnId**](PlmnId.md) |  | [optional] [default to null]
**DisasterRoamingInd** | **bool** |  | [optional] [default to false]
**DataRestorationCallbackUri** | **string** |  | [optional] [default to null]
**UdrRestartInd** | **bool** |  | [optional] [default to false]
**ExpectedUeBehaviourThresholds** | [**map[string]ExpectedUeBehaviourThreshold**](ExpectedUeBehaviourThreshold.md) | A map(list of key-value pairs) where a valid JSON pointer serves as key of expectedUeBehaviourThresholds | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

