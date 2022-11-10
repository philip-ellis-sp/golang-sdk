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

// TriggerOutputAccessRequestDynamicApprover struct for TriggerOutputAccessRequestDynamicApprover
type TriggerOutputAccessRequestDynamicApprover struct {
	// The unique ID of the identity to add to the approver list for the access request.
	Id string `json:"id"`
	// The name of the identity to add to the approver list for the access request.
	Name string `json:"name"`
	// The type of object being referenced.
	Type map[string]interface{} `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _TriggerOutputAccessRequestDynamicApprover TriggerOutputAccessRequestDynamicApprover

// NewTriggerOutputAccessRequestDynamicApprover instantiates a new TriggerOutputAccessRequestDynamicApprover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerOutputAccessRequestDynamicApprover(id string, name string, type_ map[string]interface{}) *TriggerOutputAccessRequestDynamicApprover {
	this := TriggerOutputAccessRequestDynamicApprover{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewTriggerOutputAccessRequestDynamicApproverWithDefaults instantiates a new TriggerOutputAccessRequestDynamicApprover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerOutputAccessRequestDynamicApproverWithDefaults() *TriggerOutputAccessRequestDynamicApprover {
	this := TriggerOutputAccessRequestDynamicApprover{}
	return &this
}

// GetId returns the Id field value
func (o *TriggerOutputAccessRequestDynamicApprover) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerOutputAccessRequestDynamicApprover) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerOutputAccessRequestDynamicApprover) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TriggerOutputAccessRequestDynamicApprover) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerOutputAccessRequestDynamicApprover) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerOutputAccessRequestDynamicApprover) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *TriggerOutputAccessRequestDynamicApprover) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerOutputAccessRequestDynamicApprover) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *TriggerOutputAccessRequestDynamicApprover) SetType(v map[string]interface{}) {
	o.Type = v
}

func (o TriggerOutputAccessRequestDynamicApprover) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerOutputAccessRequestDynamicApprover) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerOutputAccessRequestDynamicApprover := _TriggerOutputAccessRequestDynamicApprover{}

	if err = json.Unmarshal(bytes, &varTriggerOutputAccessRequestDynamicApprover); err == nil {
		*o = TriggerOutputAccessRequestDynamicApprover(varTriggerOutputAccessRequestDynamicApprover)
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

type NullableTriggerOutputAccessRequestDynamicApprover struct {
	value *TriggerOutputAccessRequestDynamicApprover
	isSet bool
}

func (v NullableTriggerOutputAccessRequestDynamicApprover) Get() *TriggerOutputAccessRequestDynamicApprover {
	return v.value
}

func (v *NullableTriggerOutputAccessRequestDynamicApprover) Set(val *TriggerOutputAccessRequestDynamicApprover) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerOutputAccessRequestDynamicApprover) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerOutputAccessRequestDynamicApprover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerOutputAccessRequestDynamicApprover(val *TriggerOutputAccessRequestDynamicApprover) *NullableTriggerOutputAccessRequestDynamicApprover {
	return &NullableTriggerOutputAccessRequestDynamicApprover{value: val, isSet: true}
}

func (v NullableTriggerOutputAccessRequestDynamicApprover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerOutputAccessRequestDynamicApprover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


