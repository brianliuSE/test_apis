/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Ellipsoid point with uncertainty ellipse.
type PointUncertaintyEllipse struct {
	Point *GeographicalCoordinates `json:"point"`
	UncertaintyEllipse *UncertaintyEllipse `json:"uncertaintyEllipse"`
	Confidence int32 `json:"confidence"`
	Shape *SupportedGadShapes `json:"shape"`
}
