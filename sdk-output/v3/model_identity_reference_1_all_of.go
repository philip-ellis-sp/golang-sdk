/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
)

// IdentityReference1AllOf struct for IdentityReference1AllOf
type IdentityReference1AllOf struct {
	Type *DtoType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityReference1AllOf IdentityReference1AllOf

// NewIdentityReference1AllOf instantiates a new IdentityReference1AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityReference1AllOf() *IdentityReference1AllOf {
	this := IdentityReference1AllOf{}
	return &this
}

// NewIdentityReference1AllOfWithDefaults instantiates a new IdentityReference1AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityReference1AllOfWithDefaults() *IdentityReference1AllOf {
	this := IdentityReference1AllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityReference1AllOf) GetType() DtoType {
	if o == nil || isNil(o.Type) {
		var ret DtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityReference1AllOf) GetTypeOk() (*DtoType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityReference1AllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DtoType and assigns it to the Type field.
func (o *IdentityReference1AllOf) SetType(v DtoType) {
	o.Type = &v
}

func (o IdentityReference1AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityReference1AllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityReference1AllOf := _IdentityReference1AllOf{}

	if err = json.Unmarshal(bytes, &varIdentityReference1AllOf); err == nil {
		*o = IdentityReference1AllOf(varIdentityReference1AllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityReference1AllOf struct {
	value *IdentityReference1AllOf
	isSet bool
}

func (v NullableIdentityReference1AllOf) Get() *IdentityReference1AllOf {
	return v.value
}

func (v *NullableIdentityReference1AllOf) Set(val *IdentityReference1AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityReference1AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityReference1AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityReference1AllOf(val *IdentityReference1AllOf) *NullableIdentityReference1AllOf {
	return &NullableIdentityReference1AllOf{value: val, isSet: true}
}

func (v NullableIdentityReference1AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityReference1AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


