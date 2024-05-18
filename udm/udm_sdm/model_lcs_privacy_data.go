/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type LcsPrivacyData struct {
	Lpi *Lpi `json:"lpi,omitempty"`
	UnrelatedClass *UnrelatedClass `json:"unrelatedClass,omitempty"`
	PlmnOperatorClasses []PlmnOperatorClass `json:"plmnOperatorClasses,omitempty"`
	EvtRptExpectedArea *GeographicArea `json:"evtRptExpectedArea,omitempty"`
	AreaUsageInd *AllOfLcsPrivacyDataAreaUsageInd `json:"areaUsageInd,omitempty"`
	UpLocRepIndAf *AllOfLcsPrivacyDataUpLocRepIndAf `json:"upLocRepIndAf,omitempty"`
}
