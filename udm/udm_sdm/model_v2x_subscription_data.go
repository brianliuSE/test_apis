/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type V2xSubscriptionData struct {
	NrV2xServicesAuth *NrV2xAuth `json:"nrV2xServicesAuth,omitempty"`
	LteV2xServicesAuth *LteV2xAuth `json:"lteV2xServicesAuth,omitempty"`
	NrUePc5Ambr string `json:"nrUePc5Ambr,omitempty"`
	LtePc5Ambr string `json:"ltePc5Ambr,omitempty"`
}
