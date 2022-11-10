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

// TriggerInputAccessRequestDynamicApproverRequestedFor The identity for whom the access is requested for.
type TriggerInputAccessRequestDynamicApproverRequestedFor struct {
	// The type of object that is referenced
	Type map[string]interface{} `json:"type,omitempty"`
	// ID of the object to which this reference applies
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the object to which this reference applies
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputAccessRequestDynamicApproverRequestedFor TriggerInputAccessRequestDynamicApproverRequestedFor

// NewTriggerInputAccessRequestDynamicApproverRequestedFor instantiates a new TriggerInputAccessRequestDynamicApproverRequestedFor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputAccessRequestDynamicApproverRequestedFor() *TriggerInputAccessRequestDynamicApproverRequestedFor {
	this := TriggerInputAccessRequestDynamicApproverRequestedFor{}
	return &this
}

// NewTriggerInputAccessRequestDynamicApproverRequestedForWithDefaults instantiates a new TriggerInputAccessRequestDynamicApproverRequestedFor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputAccessRequestDynamicApproverRequestedForWithDefaults() *TriggerInputAccessRequestDynamicApproverRequestedFor {
	this := TriggerInputAccessRequestDynamicApproverRequestedFor{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TriggerInputAccessRequestDynamicApproverRequestedFor) GetType() map[string]interface{} {
	if o == nil || isNil(o.Type) {
		var ret map[string]interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestDynamicApproverRequestedFor) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Type) {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TriggerInputAccessRequestDynamicApproverRequestedFor) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given map[string]interface{} and assigns it to the Type field.
func (o *TriggerInputAccessRequestDynamicApproverRequestedFor) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TriggerInputAccessRequestDynamicApproverRequestedFor) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestDynamicApproverRequestedFor) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TriggerInputAccessRequestDynamicApproverRequestedFor) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TriggerInputAccessRequestDynamicApproverRequestedFor) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TriggerInputAccessRequestDynamicApproverRequestedFor) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerInputAccessRequestDynamicApproverRequestedFor) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TriggerInputAccessRequestDynamicApproverRequestedFor) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TriggerInputAccessRequestDynamicApproverRequestedFor) SetName(v string) {
	o.Name = &v
}

func (o TriggerInputAccessRequestDynamicApproverRequestedFor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
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

func (o *TriggerInputAccessRequestDynamicApproverRequestedFor) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputAccessRequestDynamicApproverRequestedFor := _TriggerInputAccessRequestDynamicApproverRequestedFor{}

	if err = json.Unmarshal(bytes, &varTriggerInputAccessRequestDynamicApproverRequestedFor); err == nil {
		*o = TriggerInputAccessRequestDynamicApproverRequestedFor(varTriggerInputAccessRequestDynamicApproverRequestedFor)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputAccessRequestDynamicApproverRequestedFor struct {
	value *TriggerInputAccessRequestDynamicApproverRequestedFor
	isSet bool
}

func (v NullableTriggerInputAccessRequestDynamicApproverRequestedFor) Get() *TriggerInputAccessRequestDynamicApproverRequestedFor {
	return v.value
}

func (v *NullableTriggerInputAccessRequestDynamicApproverRequestedFor) Set(val *TriggerInputAccessRequestDynamicApproverRequestedFor) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccessRequestDynamicApproverRequestedFor) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccessRequestDynamicApproverRequestedFor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccessRequestDynamicApproverRequestedFor(val *TriggerInputAccessRequestDynamicApproverRequestedFor) *NullableTriggerInputAccessRequestDynamicApproverRequestedFor {
	return &NullableTriggerInputAccessRequestDynamicApproverRequestedFor{value: val, isSet: true}
}

func (v NullableTriggerInputAccessRequestDynamicApproverRequestedFor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccessRequestDynamicApproverRequestedFor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

