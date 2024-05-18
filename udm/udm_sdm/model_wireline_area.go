/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// One and only one of the \"globLineIds\", \"hfcNIds\", \"areaCodeB\",d \"areaCodeC\" and  combGciAndHfcNIds attributes        shall be included in a WirelineArea data structure 
type WirelineArea struct {
	GlobalLineIds []Bytes `json:"globalLineIds,omitempty"`
	HfcNIds []string `json:"hfcNIds,omitempty"`
	AreaCodeB string `json:"areaCodeB,omitempty"`
	AreaCodeC string `json:"areaCodeC,omitempty"`
	CombGciAndHfcNIds []CombGciAndHfcNIds `json:"combGciAndHfcNIds,omitempty"`
}