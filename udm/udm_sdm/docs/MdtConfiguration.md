# MdtConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobType** | [***JobType**](JobType.md) |  | [default to null]
**ReportType** | [***ReportTypeMdt**](ReportTypeMdt.md) |  | [optional] [default to null]
**AreaScope** | [***AreaScope**](AreaScope.md) |  | [optional] [default to null]
**MeasurementLteList** | [**[]MeasurementLteForMdt**](MeasurementLteForMdt.md) |  | [optional] [default to null]
**MeasurementNrList** | [**[]MeasurementNrForMdt**](MeasurementNrForMdt.md) |  | [optional] [default to null]
**SensorMeasurementList** | [**[]SensorMeasurement**](SensorMeasurement.md) |  | [optional] [default to null]
**ReportingTriggerList** | [**[]ReportingTrigger**](ReportingTrigger.md) |  | [optional] [default to null]
**ReportInterval** | [***ReportIntervalMdt**](ReportIntervalMdt.md) |  | [optional] [default to null]
**ReportIntervalNr** | [***ReportIntervalNrMdt**](ReportIntervalNrMdt.md) |  | [optional] [default to null]
**ReportAmount** | [***ReportAmountMdt**](ReportAmountMdt.md) |  | [optional] [default to null]
**ReportAmountPerMeasurementLte** | [**map[string]ReportAmountMdt**](ReportAmountMdt.md) | A map (list of key-value pairs) where MeasurementLteForMdt serves as key;  | [optional] [default to null]
**ReportAmountPerMeasurementNr** | [**map[string]ReportAmountMdt**](ReportAmountMdt.md) | A map (list of key-value pairs) where MeasurementNrForMdt serves as key;  | [optional] [default to null]
**EventThresholdRsrp** | **int32** | This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in LTE. When present, this IE shall indicate the Event Threshold for RSRP, and the value shall be between 0-97.  | [optional] [default to null]
**EventThresholdRsrpNr** | **int32** | This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in NR. When present, this IE shall indicate the Event Threshold for RSRP, and the value shall be between 0-127.  | [optional] [default to null]
**EventThresholdRsrq** | **int32** | This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in LTE.When present, this IE shall indicate the Event Threshold for RSRQ, and the value shall be between 0-34.  | [optional] [default to null]
**EventThresholdRsrqNr** | **int32** | This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in NR.When present, this IE shall indicate the Event Threshold for RSRQ, and the value shall be between 0-127.  | [optional] [default to null]
**EventList** | [**[]EventForMdt**](EventForMdt.md) |  | [optional] [default to null]
**LoggingInterval** | [***LoggingIntervalMdt**](LoggingIntervalMdt.md) |  | [optional] [default to null]
**LoggingIntervalNr** | [***LoggingIntervalNrMdt**](LoggingIntervalNrMdt.md) |  | [optional] [default to null]
**LoggingDuration** | [***LoggingDurationMdt**](LoggingDurationMdt.md) |  | [optional] [default to null]
**LoggingDurationNr** | [***LoggingDurationNrMdt**](LoggingDurationNrMdt.md) |  | [optional] [default to null]
**PositioningMethod** | [***PositioningMethodMdt**](PositioningMethodMdt.md) |  | [optional] [default to null]
**AddPositioningMethodList** | [**[]PositioningMethodMdt**](PositioningMethodMdt.md) |  | [optional] [default to null]
**CollectionPeriodRmmLte** | [***CollectionPeriodRmmLteMdt**](CollectionPeriodRmmLteMdt.md) |  | [optional] [default to null]
**CollectionPeriodRmmNr** | [***CollectionPeriodRmmNrMdt**](CollectionPeriodRmmNrMdt.md) |  | [optional] [default to null]
**MeasurementPeriodLte** | [***MeasurementPeriodLteMdt**](MeasurementPeriodLteMdt.md) |  | [optional] [default to null]
**MdtAllowedPlmnIdList** | [**[]PlmnId**](PlmnId.md) |  | [optional] [default to null]
**MbsfnAreaList** | [**[]MbsfnArea**](MbsfnArea.md) |  | [optional] [default to null]
**InterFreqTargetList** | [**[]InterFreqTargetInfo**](InterFreqTargetInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

