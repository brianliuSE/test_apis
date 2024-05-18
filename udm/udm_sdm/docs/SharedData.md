# SharedData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedDataId** | **string** |  | [default to null]
**SharedAmData** | [***AccessAndMobilitySubscriptionData**](AccessAndMobilitySubscriptionData.md) |  | [optional] [default to null]
**SharedSmsSubsData** | [***SmsSubscriptionData**](SmsSubscriptionData.md) |  | [optional] [default to null]
**SharedSmsMngSubsData** | [***SmsManagementSubscriptionData**](SmsManagementSubscriptionData.md) |  | [optional] [default to null]
**SharedDnnConfigurations** | [**map[string]DnnConfiguration**](DnnConfiguration.md) | A map(list of key-value pairs) where Dnn, or optionally the Wildcard DNN, serves as key of DnnConfiguration | [optional] [default to null]
**SharedTraceData** | [***TraceData**](TraceData.md) |  | [optional] [default to null]
**SharedSnssaiInfos** | [**map[string]SnssaiInfo**](SnssaiInfo.md) | A map(list of key-value pairs) where singleNssai serves as key of SnssaiInfo | [optional] [default to null]
**SharedVnGroupDatas** | [**map[string]VnGroupData**](VnGroupData.md) | A map(list of key-value pairs) where GroupId serves as key of VnGroupData | [optional] [default to null]
**TreatmentInstructions** | [**map[string]SharedDataTreatmentInstruction**](SharedDataTreatmentInstruction.md) | A map(list of key-value pairs) where JSON pointer pointing to an attribute within the SharedData serves as key of SharedDataTreatmentInstruction | [optional] [default to null]
**SharedSmSubsData** | [***SessionManagementSubscriptionData**](SessionManagementSubscriptionData.md) |  | [optional] [default to null]
**SharedEcsAddrConfigInfo** | [***EcsAddrConfigInfo**](EcsAddrConfigInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

