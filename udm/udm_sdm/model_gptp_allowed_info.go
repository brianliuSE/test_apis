/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// AF authorization information for gPTP
type GptpAllowedInfo struct {
	Dnn string `json:"dnn,omitempty"`
	SNssai *Snssai `json:"sNssai,omitempty"`
	GptpAllowed bool `json:"gptpAllowed"`
	CoverageArea []Tai `json:"coverageArea,omitempty"`
	UuTimeSyncErrBdgt int32 `json:"uuTimeSyncErrBdgt,omitempty"`
	TempVals []TemporalValidity `json:"tempVals,omitempty"`
}
