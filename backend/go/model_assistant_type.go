/*
 * Scheduler API
 *
 * API for getting generated schedules. Also used for getting and setting model parameters and instance data.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// AssistantType : The skill category of an assistant.
type AssistantType string

// List of AssistantType
const (
	JA AssistantType = "JA"
	JA_F AssistantType = "JA_F"
	SA AssistantType = "SA"
	SA_F AssistantType = "SA_F"
	SA_NEO AssistantType = "SA_NEO"
	SA_F_NEO AssistantType = "SA_F_NEO"
)

// AssertAssistantTypeRequired checks if the required fields are not zero-ed
func AssertAssistantTypeRequired(obj AssistantType) error {
	return nil
}

// AssertRecurseAssistantTypeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AssistantType (e.g. [][]AssistantType), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAssistantTypeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAssistantType, ok := obj.(AssistantType)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAssistantTypeRequired(aAssistantType)
	})
}