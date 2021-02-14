/*
 * Scheduler API
 *
 * Basic API for retrieving schedules, based on a ScheduleInput object.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// AssistantType : The different assistant types.
type AssistantType string

// List of AssistantType
const (
	JA       AssistantType = "JA"
	JA_F     AssistantType = "JA_F"
	SA       AssistantType = "SA"
	SA_F     AssistantType = "SA_F"
	SA_NEO   AssistantType = "SA_NEO"
	SA_F_NEO AssistantType = "SA_F_NEO"
)