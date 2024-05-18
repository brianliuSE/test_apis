/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type LcsSubscriptionData struct {
	ConfiguredLmfId string `json:"configuredLmfId,omitempty"`
	PruInd *PruInd `json:"pruInd,omitempty"`
	LpHapType *LpHapType `json:"lpHapType,omitempty"`
	UserPlanePosIndLmf bool `json:"userPlanePosIndLmf,omitempty"`
}
