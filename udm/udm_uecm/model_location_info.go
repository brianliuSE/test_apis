/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// This data type contains the UE's location information in 5GC.
type LocationInfo struct {
	Supi string `json:"supi,omitempty"`
	Gpsi string `json:"gpsi,omitempty"`
	RegistrationLocationInfoList []RegistrationLocationInfo `json:"registrationLocationInfoList"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
