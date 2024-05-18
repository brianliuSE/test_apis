/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The ACS information for the 5G-RG is defined in BBF TR-069 [42] or in BBF TR-369
type AcsInfo struct {
	AcsUrl string `json:"acsUrl,omitempty"`
	AcsIpv4Addr string `json:"acsIpv4Addr,omitempty"`
	AcsIpv6Addr *Ipv6Addr `json:"acsIpv6Addr,omitempty"`
}
