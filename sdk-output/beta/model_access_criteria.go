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

// AccessCriteria struct for AccessCriteria
type AccessCriteria struct {
	// Business name for the access construct list
	Name *string `json:"name,omitempty"`
	// List of criteria.  There is a min of 1 and max of 50 items in the list.
	CriteriaList []BaseReferenceDto `json:"criteriaList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessCriteria AccessCriteria

// NewAccessCriteria instantiates a new AccessCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessCriteria() *AccessCriteria {
	this := AccessCriteria{}
	return &this
}

// NewAccessCriteriaWithDefaults instantiates a new AccessCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessCriteriaWithDefaults() *AccessCriteria {
	this := AccessCriteria{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessCriteria) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCriteria) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessCriteria) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessCriteria) SetName(v string) {
	o.Name = &v
}

// GetCriteriaList returns the CriteriaList field value if set, zero value otherwise.
func (o *AccessCriteria) GetCriteriaList() []BaseReferenceDto {
	if o == nil || isNil(o.CriteriaList) {
		var ret []BaseReferenceDto
		return ret
	}
	return o.CriteriaList
}

// GetCriteriaListOk returns a tuple with the CriteriaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCriteria) GetCriteriaListOk() ([]BaseReferenceDto, bool) {
	if o == nil || isNil(o.CriteriaList) {
		return nil, false
	}
	return o.CriteriaList, true
}

// HasCriteriaList returns a boolean if a field has been set.
func (o *AccessCriteria) HasCriteriaList() bool {
	if o != nil && !isNil(o.CriteriaList) {
		return true
	}

	return false
}

// SetCriteriaList gets a reference to the given []BaseReferenceDto and assigns it to the CriteriaList field.
func (o *AccessCriteria) SetCriteriaList(v []BaseReferenceDto) {
	o.CriteriaList = v
}

func (o AccessCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.CriteriaList) {
		toSerialize["criteriaList"] = o.CriteriaList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessCriteria) UnmarshalJSON(bytes []byte) (err error) {
	varAccessCriteria := _AccessCriteria{}

	if err = json.Unmarshal(bytes, &varAccessCriteria); err == nil {
		*o = AccessCriteria(varAccessCriteria)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "criteriaList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessCriteria struct {
	value *AccessCriteria
	isSet bool
}

func (v NullableAccessCriteria) Get() *AccessCriteria {
	return v.value
}

func (v *NullableAccessCriteria) Set(val *AccessCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessCriteria(val *AccessCriteria) *NullableAccessCriteria {
	return &NullableAccessCriteria{value: val, isSet: true}
}

func (v NullableAccessCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


