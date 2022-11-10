/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"fmt"
)

// RoleCriteriaOperation An operation
type RoleCriteriaOperation string

// List of RoleCriteriaOperation
const (
	ROLECRITERIAOPERATION_EQUALS RoleCriteriaOperation = "EQUALS"
	ROLECRITERIAOPERATION_NOT_EQUALS RoleCriteriaOperation = "NOT_EQUALS"
	ROLECRITERIAOPERATION_CONTAINS RoleCriteriaOperation = "CONTAINS"
	ROLECRITERIAOPERATION_STARTS_WITH RoleCriteriaOperation = "STARTS_WITH"
	ROLECRITERIAOPERATION_ENDS_WITH RoleCriteriaOperation = "ENDS_WITH"
	ROLECRITERIAOPERATION_AND RoleCriteriaOperation = "AND"
	ROLECRITERIAOPERATION_OR RoleCriteriaOperation = "OR"
)

// All allowed values of RoleCriteriaOperation enum
var AllowedRoleCriteriaOperationEnumValues = []RoleCriteriaOperation{
	"EQUALS",
	"NOT_EQUALS",
	"CONTAINS",
	"STARTS_WITH",
	"ENDS_WITH",
	"AND",
	"OR",
}

func (v *RoleCriteriaOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoleCriteriaOperation(value)
	for _, existing := range AllowedRoleCriteriaOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoleCriteriaOperation", value)
}

// NewRoleCriteriaOperationFromValue returns a pointer to a valid RoleCriteriaOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoleCriteriaOperationFromValue(v string) (*RoleCriteriaOperation, error) {
	ev := RoleCriteriaOperation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoleCriteriaOperation: valid values are %v", v, AllowedRoleCriteriaOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoleCriteriaOperation) IsValid() bool {
	for _, existing := range AllowedRoleCriteriaOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoleCriteriaOperation value
func (v RoleCriteriaOperation) Ptr() *RoleCriteriaOperation {
	return &v
}

type NullableRoleCriteriaOperation struct {
	value *RoleCriteriaOperation
	isSet bool
}

func (v NullableRoleCriteriaOperation) Get() *RoleCriteriaOperation {
	return v.value
}

func (v *NullableRoleCriteriaOperation) Set(val *RoleCriteriaOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleCriteriaOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleCriteriaOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleCriteriaOperation(val *RoleCriteriaOperation) *NullableRoleCriteriaOperation {
	return &NullableRoleCriteriaOperation{value: val, isSet: true}
}

func (v NullableRoleCriteriaOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleCriteriaOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

