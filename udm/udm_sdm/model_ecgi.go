/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Contains the ECGI (E-UTRAN Cell Global Identity), as described in 3GPP 23.003
type Ecgi struct {
	PlmnId *PlmnId `json:"plmnId"`
	EutraCellId string `json:"eutraCellId"`
	Nid string `json:"nid,omitempty"`
}