/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type IdTranslationResult struct {
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	Supi string `json:"supi"`
	Gpsi string `json:"gpsi,omitempty"`
	AdditionalSupis []string `json:"additionalSupis,omitempty"`
	AdditionalGpsis []string `json:"additionalGpsis,omitempty"`
}