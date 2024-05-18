/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ExpectedUeBehaviourThreshold struct {
	ExpecedUeBehaviourDatasets []ExpecedUeBehaviourDataset `json:"expecedUeBehaviourDatasets,omitempty"`
	SingleNssais []Snssai `json:"singleNssais,omitempty"`
	Dnns []string `json:"dnns,omitempty"`
	ConfidenceLevel string `json:"confidenceLevel,omitempty"`
	AccuracyLevel string `json:"accuracyLevel,omitempty"`
}
