/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// UE Time Synchronization Subscription Data
type TimeSyncSubscriptionData struct {
	AfReqAuthorizations *AfRequestAuthorization `json:"afReqAuthorizations"`
	ServiceIds []TimeSyncServiceId `json:"serviceIds"`
}
