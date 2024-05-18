/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// It contains data which need to be changed.
type ChangeItem struct {
	Op *ChangeType `json:"op"`
	// contains a JSON pointer value (as defined in IETF RFC 6901) that references a target  location within the resource on which the change has been applied. 
	Path string `json:"path"`
	// indicates the path of the source JSON element (according to JSON Pointer syntax)  being moved or copied to the location indicated by the \"path\" attribute. It shall  be present if the \"op\" attribute is of value \"MOVE\". 
	From string `json:"from,omitempty"`
	OrigValue *Object `json:"origValue,omitempty"`
	NewValue *Object `json:"newValue,omitempty"`
}