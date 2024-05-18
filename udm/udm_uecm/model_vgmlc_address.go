/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// This data type contains the address(es) of VGMLC.  Depending on the names of Vgmlcaddress, it could indicate either VGMLC IPv4 or IPv6 address. 
type VgmlcAddress struct {
	VgmlcAddressIpv4 string `json:"vgmlcAddressIpv4,omitempty"`
	VgmlcAddressIpv6 *Ipv6Addr `json:"vgmlcAddressIpv6,omitempty"`
	VgmlcFqdn string `json:"vgmlcFqdn,omitempty"`
}
