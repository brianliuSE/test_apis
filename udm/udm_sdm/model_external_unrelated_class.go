/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ExternalUnrelatedClass struct {
	LcsClientExternals []LcsClientExternal `json:"lcsClientExternals,omitempty"`
	AfExternals []AfExternal `json:"afExternals,omitempty"`
	LcsClientGroupExternals []LcsClientGroupExternal `json:"lcsClientGroupExternals,omitempty"`
}
