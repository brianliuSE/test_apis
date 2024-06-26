/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Contains the ProSe Subscription Data.
type ProseSubscriptionData struct {
	ProseServiceAuth *ProseServiceAuth `json:"proseServiceAuth,omitempty"`
	NrUePc5Ambr string `json:"nrUePc5Ambr,omitempty"`
	ProseAllowedPlmn []ProSeAllowedPlmn `json:"proseAllowedPlmn,omitempty"`
}
