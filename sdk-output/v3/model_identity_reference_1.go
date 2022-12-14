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

// IdentityReference1 struct for IdentityReference1
type IdentityReference1 struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *DtoType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityReference1 IdentityReference1

// NewIdentityReference1 instantiates a new IdentityReference1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityReference1() *IdentityReference1 {
	this := IdentityReference1{}
	return &this
}

// NewIdentityReference1WithDefaults instantiates a new IdentityReference1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityReference1WithDefaults() *IdentityReference1 {
	this := IdentityReference1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityReference1) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityReference1) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityReference1) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityReference1) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityReference1) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityReference1) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityReference1) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityReference1) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityReference1) GetType() DtoType {
	if o == nil || isNil(o.Type) {
		var ret DtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityReference1) GetTypeOk() (*DtoType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityReference1) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DtoType and assigns it to the Type field.
func (o *IdentityReference1) SetType(v DtoType) {
	o.Type = &v
}

func (o IdentityReference1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityReference1) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityReference1 := _IdentityReference1{}

	if err = json.Unmarshal(bytes, &varIdentityReference1); err == nil {
		*o = IdentityReference1(varIdentityReference1)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityReference1 struct {
	value *IdentityReference1
	isSet bool
}

func (v NullableIdentityReference1) Get() *IdentityReference1 {
	return v.value
}

func (v *NullableIdentityReference1) Set(val *IdentityReference1) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityReference1) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityReference1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityReference1(val *IdentityReference1) *NullableIdentityReference1 {
	return &NullableIdentityReference1{value: val, isSet: true}
}

func (v NullableIdentityReference1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityReference1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


