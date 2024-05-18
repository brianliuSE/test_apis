/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Contain the area based on Cells or Tracking Areas.
type AreaScope struct {
	EutraCellIdList []string `json:"eutraCellIdList,omitempty"`
	NrCellIdList []string `json:"nrCellIdList,omitempty"`
	TacList []string `json:"tacList,omitempty"`
	// A map (list of key-value pairs) where PlmnId converted to a string serves as key 
	TacInfoPerPlmn map[string]TacInfo `json:"tacInfoPerPlmn,omitempty"`
	// A map (list of key-value pairs) where PlmnId converted to a string serves as key 
	CagInfoPerPlmn map[string]CagInfo1 `json:"cagInfoPerPlmn,omitempty"`
	// A map (list of key-value pairs) where PlmnId converted to a string serves as key 
	NidInfoPerPlmn map[string]NidInfo `json:"nidInfoPerPlmn,omitempty"`
}
