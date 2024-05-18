/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The response shall include a Location header field containing a different URI  (pointing to a different URI of an other service instance), or the same URI if a request  is redirected to the same target resource via a different SCP. 
type RedirectResponse struct {
	Cause string `json:"cause,omitempty"`
	TargetScp string `json:"targetScp,omitempty"`
	TargetSepp string `json:"targetSepp,omitempty"`
}
