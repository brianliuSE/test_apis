/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// This information element contains the associations between APN/DNN and PGW-C+SMF selected by the AMF for EPS interworking. 
type EpsInterworkingInfo struct {
	// A map (list of key-value pairs where Dnn serves as key) of EpsIwkPgws
	EpsIwkPgws map[string]EpsIwkPgw `json:"epsIwkPgws,omitempty"`
	RegistrationTime *time.Time `json:"registrationTime,omitempty"`
}
