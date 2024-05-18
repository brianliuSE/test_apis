/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Local 3D point with uncertainty ellipsoid
type Local3dPointUncertaintyEllipsoid struct {
	LocalOrigin *LocalOrigin `json:"localOrigin"`
	Point *RelativeCartesianLocation `json:"point"`
	UncertaintyEllipsoid *UncertaintyEllipsoid `json:"uncertaintyEllipsoid"`
	Confidence int32 `json:"confidence"`
	Shape *SupportedGadShapes `json:"shape"`
}
