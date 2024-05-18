/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Contains the Ranging/SL positioning Subscription Data.
type RangingSlPosSubscriptionData struct {
	RangingSlPosAuth *RangingSlPosAuth `json:"rangingSlPosAuth,omitempty"`
	RangingSlPosPlmn []RangingSlPosPlmn `json:"rangingSlPosPlmn,omitempty"`
	RangingSlPosQos *RangingSlPosQos `json:"rangingSlPosQos,omitempty"`
}
