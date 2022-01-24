/*
 * Scheduler API
 *
 * API for getting generated schedules. Also used for getting and setting model parameters and instance data.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// IndividualSchedule - Represents the work schedule of an individual assistant.
type IndividualSchedule struct {

	// The identification number of the assistant for which this is an individual schedule.
	AssistantId int32 `json:"assistant_id"`

	// The workload of this inidividual schedule. Used when calculating the fairness score.
	Workload float32 `json:"workload"`

	// Contains all the individual assignments of this individual schedule.
	Assignments []Assignment `json:"assignments"`
}

// AssertIndividualScheduleRequired checks if the required fields are not zero-ed
func AssertIndividualScheduleRequired(obj IndividualSchedule) error {
	elements := map[string]interface{}{
		"assistant_id": obj.AssistantId,
		"workload": obj.Workload,
		"assignments": obj.Assignments,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Assignments {
		if err := AssertAssignmentRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseIndividualScheduleRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IndividualSchedule (e.g. [][]IndividualSchedule), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIndividualScheduleRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIndividualSchedule, ok := obj.(IndividualSchedule)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIndividualScheduleRequired(aIndividualSchedule)
	})
}