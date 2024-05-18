/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SharedData struct {
	SharedDataId string `json:"sharedDataId"`
	SharedAmData *AccessAndMobilitySubscriptionData `json:"sharedAmData,omitempty"`
	SharedSmsSubsData *SmsSubscriptionData `json:"sharedSmsSubsData,omitempty"`
	SharedSmsMngSubsData *SmsManagementSubscriptionData `json:"sharedSmsMngSubsData,omitempty"`
	// A map(list of key-value pairs) where Dnn, or optionally the Wildcard DNN, serves as key of DnnConfiguration
	SharedDnnConfigurations map[string]DnnConfiguration `json:"sharedDnnConfigurations,omitempty"`
	SharedTraceData *TraceData `json:"sharedTraceData,omitempty"`
	// A map(list of key-value pairs) where singleNssai serves as key of SnssaiInfo
	SharedSnssaiInfos map[string]SnssaiInfo `json:"sharedSnssaiInfos,omitempty"`
	// A map(list of key-value pairs) where GroupId serves as key of VnGroupData
	SharedVnGroupDatas map[string]VnGroupData `json:"sharedVnGroupDatas,omitempty"`
	// A map(list of key-value pairs) where JSON pointer pointing to an attribute within the SharedData serves as key of SharedDataTreatmentInstruction
	TreatmentInstructions map[string]SharedDataTreatmentInstruction `json:"treatmentInstructions,omitempty"`
	SharedSmSubsData *SessionManagementSubscriptionData `json:"sharedSmSubsData,omitempty"`
	SharedEcsAddrConfigInfo *EcsAddrConfigInfo `json:"sharedEcsAddrConfigInfo,omitempty"`
}
