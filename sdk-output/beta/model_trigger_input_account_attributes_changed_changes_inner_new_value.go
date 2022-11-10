/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointbetasdk

import (
	"encoding/json"
	"fmt"
)

// TriggerInputAccountAttributesChangedChangesInnerNewValue - The new value of the attribute.
type TriggerInputAccountAttributesChangedChangesInnerNewValue struct {
	ArrayOfstring *[]*string
	Bool *bool
	String *string
}

// []*stringAsTriggerInputAccountAttributesChangedChangesInnerNewValue is a convenience function that returns []*string wrapped in TriggerInputAccountAttributesChangedChangesInnerNewValue
func ArrayOfstringAsTriggerInputAccountAttributesChangedChangesInnerNewValue(v *[]*string) TriggerInputAccountAttributesChangedChangesInnerNewValue {
	return TriggerInputAccountAttributesChangedChangesInnerNewValue{
		ArrayOfstring: v,
	}
}

// boolAsTriggerInputAccountAttributesChangedChangesInnerNewValue is a convenience function that returns bool wrapped in TriggerInputAccountAttributesChangedChangesInnerNewValue
func BoolAsTriggerInputAccountAttributesChangedChangesInnerNewValue(v *bool) TriggerInputAccountAttributesChangedChangesInnerNewValue {
	return TriggerInputAccountAttributesChangedChangesInnerNewValue{
		Bool: v,
	}
}

// stringAsTriggerInputAccountAttributesChangedChangesInnerNewValue is a convenience function that returns string wrapped in TriggerInputAccountAttributesChangedChangesInnerNewValue
func StringAsTriggerInputAccountAttributesChangedChangesInnerNewValue(v *string) TriggerInputAccountAttributesChangedChangesInnerNewValue {
	return TriggerInputAccountAttributesChangedChangesInnerNewValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TriggerInputAccountAttributesChangedChangesInnerNewValue) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into ArrayOfstring
	err = newStrictDecoder(data).Decode(&dst.ArrayOfstring)
	if err == nil {
		jsonArrayOfstring, _ := json.Marshal(dst.ArrayOfstring)
		if string(jsonArrayOfstring) == "{}" { // empty struct
			dst.ArrayOfstring = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfstring = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfstring = nil
		dst.Bool = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TriggerInputAccountAttributesChangedChangesInnerNewValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TriggerInputAccountAttributesChangedChangesInnerNewValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TriggerInputAccountAttributesChangedChangesInnerNewValue) MarshalJSON() ([]byte, error) {
	if src.ArrayOfstring != nil {
		return json.Marshal(&src.ArrayOfstring)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TriggerInputAccountAttributesChangedChangesInnerNewValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfstring != nil {
		return obj.ArrayOfstring
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableTriggerInputAccountAttributesChangedChangesInnerNewValue struct {
	value *TriggerInputAccountAttributesChangedChangesInnerNewValue
	isSet bool
}

func (v NullableTriggerInputAccountAttributesChangedChangesInnerNewValue) Get() *TriggerInputAccountAttributesChangedChangesInnerNewValue {
	return v.value
}

func (v *NullableTriggerInputAccountAttributesChangedChangesInnerNewValue) Set(val *TriggerInputAccountAttributesChangedChangesInnerNewValue) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccountAttributesChangedChangesInnerNewValue) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccountAttributesChangedChangesInnerNewValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccountAttributesChangedChangesInnerNewValue(val *TriggerInputAccountAttributesChangedChangesInnerNewValue) *NullableTriggerInputAccountAttributesChangedChangesInnerNewValue {
	return &NullableTriggerInputAccountAttributesChangedChangesInnerNewValue{value: val, isSet: true}
}

func (v NullableTriggerInputAccountAttributesChangedChangesInnerNewValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccountAttributesChangedChangesInnerNewValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


