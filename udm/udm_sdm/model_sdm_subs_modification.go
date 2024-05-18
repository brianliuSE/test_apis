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

type SdmSubsModification struct {
	Expires *time.Time `json:"expires,omitempty"`
	MonitoredResourceUris []string `json:"monitoredResourceUris,omitempty"`
	// A map(list of key-value pairs) where a valid JSON pointer serves as key of expectedUeBehaviourThresholds
	ExpectedUeBehaviourThresholds map[string]ExpectedUeBehaviourThreshold `json:"expectedUeBehaviourThresholds,omitempty"`
}
