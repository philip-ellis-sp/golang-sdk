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

// WorkItemType The type of the work item
type WorkItemType string

// List of WorkItemType
const (
	WORKITEMTYPE_UNKNOWN WorkItemType = "UNKNOWN"
	WORKITEMTYPE_GENERIC WorkItemType = "GENERIC"
	WORKITEMTYPE_CERTIFICATION WorkItemType = "CERTIFICATION"
	WORKITEMTYPE_REMEDIATION WorkItemType = "REMEDIATION"
	WORKITEMTYPE_DELEGATION WorkItemType = "DELEGATION"
	WORKITEMTYPE_APPROVAL WorkItemType = "APPROVAL"
	WORKITEMTYPE_VIOLATIONREVIEW WorkItemType = "VIOLATIONREVIEW"
	WORKITEMTYPE_FORM WorkItemType = "FORM"
	WORKITEMTYPE_POLICYVIOLATION WorkItemType = "POLICYVIOLATION"
	WORKITEMTYPE_CHALLENGE WorkItemType = "CHALLENGE"
	WORKITEMTYPE_IMPACTANALYSIS WorkItemType = "IMPACTANALYSIS"
	WORKITEMTYPE_SIGNOFF WorkItemType = "SIGNOFF"
	WORKITEMTYPE_EVENT WorkItemType = "EVENT"
	WORKITEMTYPE_MANUALACTION WorkItemType = "MANUALACTION"
	WORKITEMTYPE_TEST WorkItemType = "TEST"
)

// All allowed values of WorkItemType enum
var AllowedWorkItemTypeEnumValues = []WorkItemType{
	"UNKNOWN",
	"GENERIC",
	"CERTIFICATION",
	"REMEDIATION",
	"DELEGATION",
	"APPROVAL",
	"VIOLATIONREVIEW",
	"FORM",
	"POLICYVIOLATION",
	"CHALLENGE",
	"IMPACTANALYSIS",
	"SIGNOFF",
	"EVENT",
	"MANUALACTION",
	"TEST",
}

func (v *WorkItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WorkItemType(value)
	for _, existing := range AllowedWorkItemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WorkItemType", value)
}

// NewWorkItemTypeFromValue returns a pointer to a valid WorkItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWorkItemTypeFromValue(v string) (*WorkItemType, error) {
	ev := WorkItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WorkItemType: valid values are %v", v, AllowedWorkItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WorkItemType) IsValid() bool {
	for _, existing := range AllowedWorkItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkItemType value
func (v WorkItemType) Ptr() *WorkItemType {
	return &v
}

type NullableWorkItemType struct {
	value *WorkItemType
	isSet bool
}

func (v NullableWorkItemType) Get() *WorkItemType {
	return v.value
}

func (v *NullableWorkItemType) Set(val *WorkItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemType(val *WorkItemType) *NullableWorkItemType {
	return &NullableWorkItemType{value: val, isSet: true}
}

func (v NullableWorkItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

