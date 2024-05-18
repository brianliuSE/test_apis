# SessionManagementSubscriptionData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SingleNssai** | [***Snssai**](Snssai.md) |  | [default to null]
**DnnConfigurations** | [**map[string]DnnConfiguration**](DnnConfiguration.md) | A map (list of key-value pairs where Dnn, or optionally the Wildcard DNN, serves as key) of DnnConfigurations | [optional] [default to null]
**InternalGroupIds** | **[]string** |  | [optional] [default to null]
**SharedVnGroupDataIds** | **map[string]string** | A map(list of key-value pairs) where GroupId serves as key of SharedDataId | [optional] [default to null]
**SharedDnnConfigurationsId** | **string** |  | [optional] [default to null]
**OdbPacketServices** | [***OdbPacketServices**](OdbPacketServices.md) |  | [optional] [default to null]
**TraceData** | [***TraceData**](TraceData.md) |  | [optional] [default to null]
**SharedTraceDataId** | **string** |  | [optional] [default to null]
**ExpectedUeBehavioursList** | [**map[string]ExpectedUeBehaviourData**](ExpectedUeBehaviourData.md) | A map(list of key-value pairs) where Dnn serves as key of ExpectedUeBehaviourData | [optional] [default to null]
**ExpectedUeBehaviourData** | [**map[string]map[string]ExpectedUeBehaviourData**](map.md) | A map(list of key-value pairs) where DNN serves as key | [optional] [default to null]
**AppSpecificExpectedUeBehaviourData** | [**map[string]map[string]AppSpecificExpectedUeBehaviourData**](map.md) | A map(list of key-value pairs) where DNN serves as key | [optional] [default to null]
**SuggestedPacketNumDlList** | [**map[string]SuggestedPacketNumDl**](SuggestedPacketNumDl.md) | A map(list of key-value pairs) where Dnn serves as key of SuggestedPacketNumDl | [optional] [default to null]
**Var3gppChargingCharacteristics** | **string** |  | [optional] [default to null]
**NsacMode** | [***NsacAdmissionMode**](NsacAdmissionMode.md) |  | [optional] [default to null]
**SessInactTimer** | **int32** |  | [optional] [default to null]
**OnDemand** | **bool** |  | [optional] [default to false]
**SupportedFeatures** | **string** |  | [optional] [default to null]
**AdditionalSharedDnnConfigurationsIds** | **[]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

