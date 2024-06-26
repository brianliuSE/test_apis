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

type SdmSubscription struct {
	NfInstanceId string `json:"nfInstanceId"`
	ImplicitUnsubscribe bool `json:"implicitUnsubscribe,omitempty"`
	Expires *time.Time `json:"expires,omitempty"`
	CallbackReference string `json:"callbackReference"`
	AmfServiceName *ServiceName `json:"amfServiceName,omitempty"`
	MonitoredResourceUris []string `json:"monitoredResourceUris"`
	SingleNssai *Snssai `json:"singleNssai,omitempty"`
	Dnn string `json:"dnn,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty"`
	PlmnId *PlmnId `json:"plmnId,omitempty"`
	ImmediateReport bool `json:"immediateReport,omitempty"`
	Report *ImmediateReport `json:"report,omitempty"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	ContextInfo *ContextInfo `json:"contextInfo,omitempty"`
	NfChangeFilter bool `json:"nfChangeFilter,omitempty"`
	UniqueSubscription bool `json:"uniqueSubscription,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
	UeConSmfDataSubFilter *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	AdjacentPlmns []PlmnId `json:"adjacentPlmns,omitempty"`
	DisasterRoamingInd bool `json:"disasterRoamingInd,omitempty"`
	DataRestorationCallbackUri string `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd bool `json:"udrRestartInd,omitempty"`
	// A map(list of key-value pairs) where a valid JSON pointer serves as key of expectedUeBehaviourThresholds
	ExpectedUeBehaviourThresholds map[string]ExpectedUeBehaviourThreshold `json:"expectedUeBehaviourThresholds,omitempty"`
}
