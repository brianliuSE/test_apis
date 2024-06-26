/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Ellipsoid point with altitude and uncertainty ellipsoid.
type PointAltitudeUncertainty struct {
	Point *GeographicalCoordinates `json:"point"`
	Altitude float64 `json:"altitude"`
	UncertaintyEllipse *UncertaintyEllipse `json:"uncertaintyEllipse"`
	UncertaintyAltitude float32 `json:"uncertaintyAltitude"`
	Confidence int32 `json:"confidence"`
	VConfidence int32 `json:"vConfidence,omitempty"`
	Shape *SupportedGadShapes `json:"shape"`
}
