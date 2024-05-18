/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UeContextInSmfData struct {
	// A map (list of key-value pairs where PduSessionId serves as key) of PduSessions
	PduSessions map[string]PduSession `json:"pduSessions,omitempty"`
	PgwInfo []PgwInfo `json:"pgwInfo,omitempty"`
	EmergencyInfo *EmergencyInfo `json:"emergencyInfo,omitempty"`
}