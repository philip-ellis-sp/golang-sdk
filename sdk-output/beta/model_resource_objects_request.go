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

// ResourceObjectsRequest Request model for peek resource objects from source connectors.
type ResourceObjectsRequest struct {
	// The type of resource objects to iterate over.
	ObjectType *string `json:"objectType,omitempty"`
	// The maximum number of resource objects to iterate over and return.
	MaxCount *int32 `json:"maxCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceObjectsRequest ResourceObjectsRequest

// NewResourceObjectsRequest instantiates a new ResourceObjectsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceObjectsRequest() *ResourceObjectsRequest {
	this := ResourceObjectsRequest{}
	var objectType string = "account"
	this.ObjectType = &objectType
	var maxCount int32 = 25
	this.MaxCount = &maxCount
	return &this
}

// NewResourceObjectsRequestWithDefaults instantiates a new ResourceObjectsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceObjectsRequestWithDefaults() *ResourceObjectsRequest {
	this := ResourceObjectsRequest{}
	var objectType string = "account"
	this.ObjectType = &objectType
	var maxCount int32 = 25
	this.MaxCount = &maxCount
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *ResourceObjectsRequest) GetObjectType() string {
	if o == nil || isNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObjectsRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil || isNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *ResourceObjectsRequest) HasObjectType() bool {
	if o != nil && !isNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *ResourceObjectsRequest) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetMaxCount returns the MaxCount field value if set, zero value otherwise.
func (o *ResourceObjectsRequest) GetMaxCount() int32 {
	if o == nil || isNil(o.MaxCount) {
		var ret int32
		return ret
	}
	return *o.MaxCount
}

// GetMaxCountOk returns a tuple with the MaxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObjectsRequest) GetMaxCountOk() (*int32, bool) {
	if o == nil || isNil(o.MaxCount) {
		return nil, false
	}
	return o.MaxCount, true
}

// HasMaxCount returns a boolean if a field has been set.
func (o *ResourceObjectsRequest) HasMaxCount() bool {
	if o != nil && !isNil(o.MaxCount) {
		return true
	}

	return false
}

// SetMaxCount gets a reference to the given int32 and assigns it to the MaxCount field.
func (o *ResourceObjectsRequest) SetMaxCount(v int32) {
	o.MaxCount = &v
}

func (o ResourceObjectsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !isNil(o.MaxCount) {
		toSerialize["maxCount"] = o.MaxCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceObjectsRequest) UnmarshalJSON(bytes []byte) (err error) {
	varResourceObjectsRequest := _ResourceObjectsRequest{}

	if err = json.Unmarshal(bytes, &varResourceObjectsRequest); err == nil {
		*o = ResourceObjectsRequest(varResourceObjectsRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "objectType")
		delete(additionalProperties, "maxCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceObjectsRequest struct {
	value *ResourceObjectsRequest
	isSet bool
}

func (v NullableResourceObjectsRequest) Get() *ResourceObjectsRequest {
	return v.value
}

func (v *NullableResourceObjectsRequest) Set(val *ResourceObjectsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceObjectsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceObjectsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceObjectsRequest(val *ResourceObjectsRequest) *NullableResourceObjectsRequest {
	return &NullableResourceObjectsRequest{value: val, isSet: true}
}

func (v NullableResourceObjectsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceObjectsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


