/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RecurTime struct {
	RecurTimeWindow *ValidTimePeriod `json:"recurTimeWindow,omitempty"`
	RecurType *RecurType `json:"recurType,omitempty"`
	RecurMonth []int32 `json:"recurMonth,omitempty"`
	RecurWeek []int32 `json:"recurWeek,omitempty"`
	RecurDay []int32 `json:"recurDay,omitempty"`
	RecurDate []time.Time `json:"recurDate,omitempty"`
}
