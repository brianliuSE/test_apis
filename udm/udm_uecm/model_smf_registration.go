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

// This datatype contains a complete set of mandatory information relevant to an SMF  serving the UE. 
type SmfRegistration struct {
	SmfInstanceId string `json:"smfInstanceId"`
	SmfSetId string `json:"smfSetId,omitempty"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	PduSessionId int32 `json:"pduSessionId"`
	SingleNssai *Snssai `json:"singleNssai"`
	Dnn string `json:"dnn,omitempty"`
	EmergencyServices bool `json:"emergencyServices,omitempty"`
	PcscfRestorationCallbackUri string `json:"pcscfRestorationCallbackUri,omitempty"`
	PlmnId *PlmnId `json:"plmnId"`
	PgwFqdn string `json:"pgwFqdn,omitempty"`
	PgwIpAddr *IpAddress `json:"pgwIpAddr,omitempty"`
	EpdgInd bool `json:"epdgInd,omitempty"`
	DeregCallbackUri string `json:"deregCallbackUri,omitempty"`
	RegistrationReason *RegistrationReason `json:"registrationReason,omitempty"`
	RegistrationTime *time.Time `json:"registrationTime,omitempty"`
	ContextInfo *ContextInfo `json:"contextInfo,omitempty"`
	PcfId string `json:"pcfId,omitempty"`
	DataRestorationCallbackUri string `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
	UdrRestartInd bool `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime *time.Time `json:"lastSynchronizationTime,omitempty"`
	PduSessionReActivationRequired bool `json:"pduSessionReActivationRequired,omitempty"`
	StaleCheckCallbackUri string `json:"staleCheckCallbackUri,omitempty"`
	UdmStaleCheckCallbackUri string `json:"udmStaleCheckCallbackUri,omitempty"`
	WildcardInd bool `json:"wildcardInd,omitempty"`
}
