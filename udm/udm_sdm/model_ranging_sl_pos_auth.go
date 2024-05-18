/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Indicates whether the UE is authorized to use related services. 
type RangingSlPosAuth struct {
	RgSlPosTargetAuth *UeAuth `json:"rgSlPosTargetAuth,omitempty"`
	RgSlPosSlRefAuth *UeAuth `json:"rgSlPosSlRefAuth,omitempty"`
	RgSlPosLocAuth *UeAuth `json:"rgSlPosLocAuth,omitempty"`
	RgSlPosClientAuth *UeAuth `json:"rgSlPosClientAuth,omitempty"`
	RgSlPosServerAuth *UeAuth `json:"rgSlPosServerAuth,omitempty"`
}