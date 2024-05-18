/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TimeSyncData struct {
	Authorized bool `json:"authorized"`
	UuTimeSyncErrBdgt int32 `json:"uuTimeSyncErrBdgt,omitempty"`
	TempVals []TemporalValidity `json:"tempVals,omitempty"`
	CoverageArea []Tai `json:"coverageArea,omitempty"`
	ClockQualityDetailLevel *ClockQualityDetailLevel `json:"clockQualityDetailLevel,omitempty"`
	ClockQualityAcceptanceCriterion *ClockQualityAcceptanceCriterion `json:"clockQualityAcceptanceCriterion,omitempty"`
}
