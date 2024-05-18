/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Contains attributes of NwdafRegistration that can be modified using PATCH
type NwdafRegistrationModification struct {
	NwdafInstanceId string `json:"nwdafInstanceId"`
	NwdafSetId string `json:"nwdafSetId,omitempty"`
	AnalyticsIds []EventId `json:"analyticsIds,omitempty"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
