/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type IpAddress struct {
	Ipv4Addr string `json:"ipv4Addr,omitempty"`
	Ipv6Addr *Ipv6Addr `json:"ipv6Addr,omitempty"`
	Ipv6Prefix *Ipv6Prefix `json:"ipv6Prefix,omitempty"`
}
