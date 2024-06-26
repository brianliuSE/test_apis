/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// This datatype contains SMSF registration for 3GPP access that is created or updated with the received information. 
type SmsfRegistration struct {
	SmsfInstanceId string `json:"smsfInstanceId"`
	SmsfSetId string `json:"smsfSetId,omitempty"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	PlmnId *PlmnId `json:"plmnId"`
	SmsfMAPAddress string `json:"smsfMAPAddress,omitempty"`
	SmsfDiameterAddress *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	RegistrationTime *time.Time `json:"registrationTime,omitempty"`
	ContextInfo *ContextInfo `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri string `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
	SmsfSbiSupInd bool `json:"smsfSbiSupInd,omitempty"`
	UdrRestartInd bool `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime *time.Time `json:"lastSynchronizationTime,omitempty"`
	UeMemoryAvailableInd bool `json:"ueMemoryAvailableInd,omitempty"`
}
