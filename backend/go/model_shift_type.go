/*
 * Scheduler API
 *
 * API for getting generated schedules. Also used for getting and setting model parameters and instance data.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ShiftType : Identifies a specific shift type. Includes the FREE shift type.
type ShiftType string

// List of ShiftType
const (
	JAEV ShiftType = "JAEV"
	JAWE ShiftType = "JAWE"
	JAHO ShiftType = "JAHO"
	JANW ShiftType = "JANW"
	SAEW ShiftType = "SAEW"
	SAWE ShiftType = "SAWE"
	SAHO ShiftType = "SAHO"
	TPWE ShiftType = "TPWE"
	TPHO ShiftType = "TPHO"
	CALL ShiftType = "CALL"
	FREE ShiftType = "FREE"
)

// AssertShiftTypeRequired checks if the required fields are not zero-ed
func AssertShiftTypeRequired(obj ShiftType) error {
	return nil
}

// AssertRecurseShiftTypeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ShiftType (e.g. [][]ShiftType), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseShiftTypeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aShiftType, ok := obj.(ShiftType)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertShiftTypeRequired(aShiftType)
	})
}
