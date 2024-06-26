/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SmsManagementSubscriptionData struct {
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	MtSmsSubscribed bool `json:"mtSmsSubscribed,omitempty"`
	MtSmsBarringAll bool `json:"mtSmsBarringAll,omitempty"`
	MtSmsBarringRoaming bool `json:"mtSmsBarringRoaming,omitempty"`
	MoSmsSubscribed bool `json:"moSmsSubscribed,omitempty"`
	MoSmsBarringAll bool `json:"moSmsBarringAll,omitempty"`
	MoSmsBarringRoaming bool `json:"moSmsBarringRoaming,omitempty"`
	SharedSmsMngDataIds []string `json:"sharedSmsMngDataIds,omitempty"`
	TraceData *TraceData `json:"traceData,omitempty"`
	SharedTraceDataId string `json:"sharedTraceDataId,omitempty"`
}
