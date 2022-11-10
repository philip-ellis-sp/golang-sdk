/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointbetasdk

import (
	"encoding/json"
)

// RoleCriteriaLevel1 Defines STANDARD type Role membership
type RoleCriteriaLevel1 struct {
	Operation *RoleCriteriaOperation `json:"operation,omitempty"`
	Key *RoleCriteriaKey `json:"key,omitempty"`
	// String value to test the Identity attribute, Account attribute, or Entitlement specified in the key w/r/t the specified operation. If this criteria is a leaf node, that is, if the operation is one of EQUALS, NOT_EQUALS, CONTAINS, STARTS_WITH, or ENDS_WITH, this field is required. Otherwise, specifying it is an error.
	StringValue *string `json:"stringValue,omitempty"`
	// Array of child criteria. Required if the operation is AND or OR, otherwise it must be left null. A maximum of three levels of criteria are supported, including leaf nodes. Additionally, AND nodes can only be children or OR nodes and vice-versa.
	Children []RoleCriteriaLevel2 `json:"children,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleCriteriaLevel1 RoleCriteriaLevel1

// NewRoleCriteriaLevel1 instantiates a new RoleCriteriaLevel1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleCriteriaLevel1() *RoleCriteriaLevel1 {
	this := RoleCriteriaLevel1{}
	return &this
}

// NewRoleCriteriaLevel1WithDefaults instantiates a new RoleCriteriaLevel1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleCriteriaLevel1WithDefaults() *RoleCriteriaLevel1 {
	this := RoleCriteriaLevel1{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *RoleCriteriaLevel1) GetOperation() RoleCriteriaOperation {
	if o == nil || isNil(o.Operation) {
		var ret RoleCriteriaOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCriteriaLevel1) GetOperationOk() (*RoleCriteriaOperation, bool) {
	if o == nil || isNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *RoleCriteriaLevel1) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given RoleCriteriaOperation and assigns it to the Operation field.
func (o *RoleCriteriaLevel1) SetOperation(v RoleCriteriaOperation) {
	o.Operation = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RoleCriteriaLevel1) GetKey() RoleCriteriaKey {
	if o == nil || isNil(o.Key) {
		var ret RoleCriteriaKey
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCriteriaLevel1) GetKeyOk() (*RoleCriteriaKey, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RoleCriteriaLevel1) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given RoleCriteriaKey and assigns it to the Key field.
func (o *RoleCriteriaLevel1) SetKey(v RoleCriteriaKey) {
	o.Key = &v
}

// GetStringValue returns the StringValue field value if set, zero value otherwise.
func (o *RoleCriteriaLevel1) GetStringValue() string {
	if o == nil || isNil(o.StringValue) {
		var ret string
		return ret
	}
	return *o.StringValue
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCriteriaLevel1) GetStringValueOk() (*string, bool) {
	if o == nil || isNil(o.StringValue) {
		return nil, false
	}
	return o.StringValue, true
}

// HasStringValue returns a boolean if a field has been set.
func (o *RoleCriteriaLevel1) HasStringValue() bool {
	if o != nil && !isNil(o.StringValue) {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given string and assigns it to the StringValue field.
func (o *RoleCriteriaLevel1) SetStringValue(v string) {
	o.StringValue = &v
}

// GetChildren returns the Children field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleCriteriaLevel1) GetChildren() []RoleCriteriaLevel2 {
	if o == nil {
		var ret []RoleCriteriaLevel2
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleCriteriaLevel1) GetChildrenOk() ([]RoleCriteriaLevel2, bool) {
	if o == nil || isNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *RoleCriteriaLevel1) HasChildren() bool {
	if o != nil && isNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []RoleCriteriaLevel2 and assigns it to the Children field.
func (o *RoleCriteriaLevel1) SetChildren(v []RoleCriteriaLevel2) {
	o.Children = v
}

func (o RoleCriteriaLevel1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.StringValue) {
		toSerialize["stringValue"] = o.StringValue
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleCriteriaLevel1) UnmarshalJSON(bytes []byte) (err error) {
	varRoleCriteriaLevel1 := _RoleCriteriaLevel1{}

	if err = json.Unmarshal(bytes, &varRoleCriteriaLevel1); err == nil {
		*o = RoleCriteriaLevel1(varRoleCriteriaLevel1)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "operation")
		delete(additionalProperties, "key")
		delete(additionalProperties, "stringValue")
		delete(additionalProperties, "children")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleCriteriaLevel1 struct {
	value *RoleCriteriaLevel1
	isSet bool
}

func (v NullableRoleCriteriaLevel1) Get() *RoleCriteriaLevel1 {
	return v.value
}

func (v *NullableRoleCriteriaLevel1) Set(val *RoleCriteriaLevel1) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleCriteriaLevel1) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleCriteriaLevel1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleCriteriaLevel1(val *RoleCriteriaLevel1) *NullableRoleCriteriaLevel1 {
	return &NullableRoleCriteriaLevel1{value: val, isSet: true}
}

func (v NullableRoleCriteriaLevel1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleCriteriaLevel1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


