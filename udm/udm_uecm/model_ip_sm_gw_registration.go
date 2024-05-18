/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// This data type contains the IP-SW-GW routing information.
type IpSmGwRegistration struct {
	IpSmGwMapAddress string `json:"ipSmGwMapAddress,omitempty"`
	IpSmGwDiameterAddress *NetworkNodeDiameterAddress `json:"ipSmGwDiameterAddress,omitempty"`
	IpsmgwIpv4 string `json:"ipsmgwIpv4,omitempty"`
	IpsmgwIpv6 *Ipv6Addr `json:"ipsmgwIpv6,omitempty"`
	IpsmgwFqdn string `json:"ipsmgwFqdn,omitempty"`
	NfInstanceId string `json:"nfInstanceId,omitempty"`
	UnriIndicator bool `json:"unriIndicator,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
	IpSmGwSbiSupInd bool `json:"ipSmGwSbiSupInd,omitempty"`
}