/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Contains information related to the access token request
type AccessTokenReq struct {
	GrantType string `json:"grant_type"`
	NfInstanceId string `json:"nfInstanceId"`
	NfType *NfType `json:"nfType,omitempty"`
	TargetNfType *NfType `json:"targetNfType,omitempty"`
	Scope string `json:"scope"`
	TargetNfInstanceId string `json:"targetNfInstanceId,omitempty"`
	RequesterPlmn *PlmnId `json:"requesterPlmn,omitempty"`
	RequesterPlmnList []PlmnId `json:"requesterPlmnList,omitempty"`
	RequesterSnssaiList []Snssai `json:"requesterSnssaiList,omitempty"`
	RequesterFqdn string `json:"requesterFqdn,omitempty"`
	RequesterSnpnList []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetPlmn *PlmnId `json:"targetPlmn,omitempty"`
	TargetSnpn *PlmnIdNid `json:"targetSnpn,omitempty"`
	TargetSnssaiList []Snssai `json:"targetSnssaiList,omitempty"`
	TargetNsiList []string `json:"targetNsiList,omitempty"`
	TargetNfSetId string `json:"targetNfSetId,omitempty"`
	TargetNfServiceSetId string `json:"targetNfServiceSetId,omitempty"`
	HnrfAccessTokenUri string `json:"hnrfAccessTokenUri,omitempty"`
	SourceNfInstanceId string `json:"sourceNfInstanceId,omitempty"`
}
