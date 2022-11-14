/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointccsdk

import (
	"encoding/json"
)

// CreateClusterRequest struct for CreateClusterRequest
type CreateClusterRequest struct {
	Name *string `json:"name,omitempty"`
	GmtOffset *float32 `json:"gmtOffset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateClusterRequest CreateClusterRequest

// NewCreateClusterRequest instantiates a new CreateClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterRequest() *CreateClusterRequest {
	this := CreateClusterRequest{}
	return &this
}

// NewCreateClusterRequestWithDefaults instantiates a new CreateClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterRequestWithDefaults() *CreateClusterRequest {
	this := CreateClusterRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateClusterRequest) SetName(v string) {
	o.Name = &v
}

// GetGmtOffset returns the GmtOffset field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetGmtOffset() float32 {
	if o == nil || isNil(o.GmtOffset) {
		var ret float32
		return ret
	}
	return *o.GmtOffset
}

// GetGmtOffsetOk returns a tuple with the GmtOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetGmtOffsetOk() (*float32, bool) {
	if o == nil || isNil(o.GmtOffset) {
		return nil, false
	}
	return o.GmtOffset, true
}

// HasGmtOffset returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasGmtOffset() bool {
	if o != nil && !isNil(o.GmtOffset) {
		return true
	}

	return false
}

// SetGmtOffset gets a reference to the given float32 and assigns it to the GmtOffset field.
func (o *CreateClusterRequest) SetGmtOffset(v float32) {
	o.GmtOffset = &v
}

func (o CreateClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.GmtOffset) {
		toSerialize["gmtOffset"] = o.GmtOffset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateClusterRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateClusterRequest := _CreateClusterRequest{}

	if err = json.Unmarshal(bytes, &varCreateClusterRequest); err == nil {
		*o = CreateClusterRequest(varCreateClusterRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "gmtOffset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateClusterRequest struct {
	value *CreateClusterRequest
	isSet bool
}

func (v NullableCreateClusterRequest) Get() *CreateClusterRequest {
	return v.value
}

func (v *NullableCreateClusterRequest) Set(val *CreateClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterRequest(val *CreateClusterRequest) *NullableCreateClusterRequest {
	return &NullableCreateClusterRequest{value: val, isSet: true}
}

func (v NullableCreateClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

