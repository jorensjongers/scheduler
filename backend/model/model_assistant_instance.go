/*
 * Scheduler API
 *
 * Basic API for retrieving schedules, based on a ScheduleInput object.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// AssistantInstance - Holds all instance data for one assistant.
type AssistantInstance struct {

	// The identification number of this assistant.
	Id int32 `json:"id,omitempty"`

	Type AssistantType `json:"type,omitempty"`
}