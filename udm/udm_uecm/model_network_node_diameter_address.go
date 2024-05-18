/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// This data type is a part of smsfDiameterAddress and it should be present whenever smsf supports Diameter protocol. 
type NetworkNodeDiameterAddress struct {
	Name *Fqdn `json:"name"`
	Realm *Fqdn `json:"realm"`
}