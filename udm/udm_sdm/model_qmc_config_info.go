/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// It contains the configuration information for signaling-based activation of the  Quality of Experience (QoE) Measurements Collection (QMC) functionality.  
type QmcConfigInfo struct {
	QoeReference string `json:"qoeReference"`
	ServiceType *QoeServiceType `json:"serviceType,omitempty"`
	SliceScope []Snssai `json:"sliceScope,omitempty"`
	AreaScope *QmcAreaScope `json:"areaScope,omitempty"`
	QoeCollectionEntityAddress *IpAddr `json:"qoeCollectionEntityAddress,omitempty"`
	QoeTarget *QoeTarget `json:"qoeTarget,omitempty"`
	MdtAlignmentInfo *Object `json:"mdtAlignmentInfo,omitempty"`
	AvailableRanVisibleQoeMetrics []AvailableRanVisibleQoeMetric `json:"availableRanVisibleQoeMetrics,omitempty"`
	ContainerForAppLayerMeasConfig string `json:"containerForAppLayerMeasConfig,omitempty"`
	MbsCommunicationServiceType *MbsServiceType `json:"mbsCommunicationServiceType,omitempty"`
}
