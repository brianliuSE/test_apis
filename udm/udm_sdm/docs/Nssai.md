# Nssai

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | **string** |  | [optional] [default to null]
**DefaultSingleNssais** | [**[]Snssai**](Snssai.md) |  | [default to null]
**SingleNssais** | [**[]Snssai**](Snssai.md) |  | [optional] [default to null]
**ProvisioningTime** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**AdditionalSnssaiData** | [**map[string]AdditionalSnssaiData**](AdditionalSnssaiData.md) | A map(list of key-value pairs) where singleNssai serves as key of AdditionalSnssaiData | [optional] [default to null]
**SuppressNssrgInd** | **bool** |  | [optional] [default to null]
**NssaiValidityTimeInfo** | [**map[string]time.Time**](time.Time.md) | A map(list of key-value pairs where single Nssai serves as key) of the current validity time  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

