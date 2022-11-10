/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
)

// BaseReferenceDto1 struct for BaseReferenceDto1
type BaseReferenceDto1 struct {
	// the application ID
	Id *string `json:"id,omitempty"`
	// the application name
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseReferenceDto1 BaseReferenceDto1

// NewBaseReferenceDto1 instantiates a new BaseReferenceDto1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseReferenceDto1() *BaseReferenceDto1 {
	this := BaseReferenceDto1{}
	return &this
}

// NewBaseReferenceDto1WithDefaults instantiates a new BaseReferenceDto1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseReferenceDto1WithDefaults() *BaseReferenceDto1 {
	this := BaseReferenceDto1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseReferenceDto1) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReferenceDto1) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseReferenceDto1) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseReferenceDto1) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseReferenceDto1) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReferenceDto1) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseReferenceDto1) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseReferenceDto1) SetName(v string) {
	o.Name = &v
}

func (o BaseReferenceDto1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseReferenceDto1) UnmarshalJSON(bytes []byte) (err error) {
	varBaseReferenceDto1 := _BaseReferenceDto1{}

	if err = json.Unmarshal(bytes, &varBaseReferenceDto1); err == nil {
		*o = BaseReferenceDto1(varBaseReferenceDto1)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseReferenceDto1 struct {
	value *BaseReferenceDto1
	isSet bool
}

func (v NullableBaseReferenceDto1) Get() *BaseReferenceDto1 {
	return v.value
}

func (v *NullableBaseReferenceDto1) Set(val *BaseReferenceDto1) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseReferenceDto1) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseReferenceDto1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseReferenceDto1(val *BaseReferenceDto1) *NullableBaseReferenceDto1 {
	return &NullableBaseReferenceDto1{value: val, isSet: true}
}

func (v NullableBaseReferenceDto1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseReferenceDto1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


