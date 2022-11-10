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

// IdentityDeleted struct for IdentityDeleted
type IdentityDeleted struct {
	Identity TriggerInputIdentityDeletedIdentity `json:"identity"`
	// The attributes assigned to the identity.  Attributes are determined by the identity profile.
	Attributes map[string]interface{} `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _IdentityDeleted IdentityDeleted

// NewIdentityDeleted instantiates a new IdentityDeleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityDeleted(identity TriggerInputIdentityDeletedIdentity, attributes map[string]interface{}) *IdentityDeleted {
	this := IdentityDeleted{}
	this.Identity = identity
	this.Attributes = attributes
	return &this
}

// NewIdentityDeletedWithDefaults instantiates a new IdentityDeleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityDeletedWithDefaults() *IdentityDeleted {
	this := IdentityDeleted{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *IdentityDeleted) GetIdentity() TriggerInputIdentityDeletedIdentity {
	if o == nil {
		var ret TriggerInputIdentityDeletedIdentity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *IdentityDeleted) GetIdentityOk() (*TriggerInputIdentityDeletedIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *IdentityDeleted) SetIdentity(v TriggerInputIdentityDeletedIdentity) {
	o.Identity = v
}

// GetAttributes returns the Attributes field value
func (o *IdentityDeleted) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IdentityDeleted) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *IdentityDeleted) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o IdentityDeleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identity"] = o.Identity
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityDeleted) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityDeleted := _IdentityDeleted{}

	if err = json.Unmarshal(bytes, &varIdentityDeleted); err == nil {
		*o = IdentityDeleted(varIdentityDeleted)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityDeleted struct {
	value *IdentityDeleted
	isSet bool
}

func (v NullableIdentityDeleted) Get() *IdentityDeleted {
	return v.value
}

func (v *NullableIdentityDeleted) Set(val *IdentityDeleted) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityDeleted) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityDeleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityDeleted(val *IdentityDeleted) *NullableIdentityDeleted {
	return &NullableIdentityDeleted{value: val, isSet: true}
}

func (v NullableIdentityDeleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityDeleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


