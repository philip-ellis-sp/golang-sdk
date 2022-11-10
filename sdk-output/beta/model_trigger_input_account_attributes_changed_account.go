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

// TriggerInputAccountAttributesChangedAccount Details of the account where the attributes changed.
type TriggerInputAccountAttributesChangedAccount struct {
	// SailPoint generated unique identifier.
	Id string `json:"id"`
	// The source's unique identifier for the account. UUID is generated by the source system.
	Uuid NullableString `json:"uuid"`
	// Name of the account.
	Name string `json:"name"`
	// Unique ID of the account on the source.
	NativeIdentity string `json:"nativeIdentity"`
	// The type of the account
	Type map[string]interface{} `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputAccountAttributesChangedAccount TriggerInputAccountAttributesChangedAccount

// NewTriggerInputAccountAttributesChangedAccount instantiates a new TriggerInputAccountAttributesChangedAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputAccountAttributesChangedAccount(id string, uuid NullableString, name string, nativeIdentity string, type_ map[string]interface{}) *TriggerInputAccountAttributesChangedAccount {
	this := TriggerInputAccountAttributesChangedAccount{}
	this.Id = id
	this.Uuid = uuid
	this.Name = name
	this.NativeIdentity = nativeIdentity
	this.Type = type_
	return &this
}

// NewTriggerInputAccountAttributesChangedAccountWithDefaults instantiates a new TriggerInputAccountAttributesChangedAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputAccountAttributesChangedAccountWithDefaults() *TriggerInputAccountAttributesChangedAccount {
	this := TriggerInputAccountAttributesChangedAccount{}
	return &this
}

// GetId returns the Id field value
func (o *TriggerInputAccountAttributesChangedAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountAttributesChangedAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TriggerInputAccountAttributesChangedAccount) SetId(v string) {
	o.Id = v
}

// GetUuid returns the Uuid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TriggerInputAccountAttributesChangedAccount) GetUuid() string {
	if o == nil || o.Uuid.Get() == nil {
		var ret string
		return ret
	}

	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputAccountAttributesChangedAccount) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// SetUuid sets field value
func (o *TriggerInputAccountAttributesChangedAccount) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// GetName returns the Name field value
func (o *TriggerInputAccountAttributesChangedAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountAttributesChangedAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerInputAccountAttributesChangedAccount) SetName(v string) {
	o.Name = v
}

// GetNativeIdentity returns the NativeIdentity field value
func (o *TriggerInputAccountAttributesChangedAccount) GetNativeIdentity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NativeIdentity
}

// GetNativeIdentityOk returns a tuple with the NativeIdentity field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountAttributesChangedAccount) GetNativeIdentityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NativeIdentity, true
}

// SetNativeIdentity sets field value
func (o *TriggerInputAccountAttributesChangedAccount) SetNativeIdentity(v string) {
	o.NativeIdentity = v
}

// GetType returns the Type field value
func (o *TriggerInputAccountAttributesChangedAccount) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountAttributesChangedAccount) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *TriggerInputAccountAttributesChangedAccount) SetType(v map[string]interface{}) {
	o.Type = v
}

func (o TriggerInputAccountAttributesChangedAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["nativeIdentity"] = o.NativeIdentity
	}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputAccountAttributesChangedAccount) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputAccountAttributesChangedAccount := _TriggerInputAccountAttributesChangedAccount{}

	if err = json.Unmarshal(bytes, &varTriggerInputAccountAttributesChangedAccount); err == nil {
		*o = TriggerInputAccountAttributesChangedAccount(varTriggerInputAccountAttributesChangedAccount)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nativeIdentity")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputAccountAttributesChangedAccount struct {
	value *TriggerInputAccountAttributesChangedAccount
	isSet bool
}

func (v NullableTriggerInputAccountAttributesChangedAccount) Get() *TriggerInputAccountAttributesChangedAccount {
	return v.value
}

func (v *NullableTriggerInputAccountAttributesChangedAccount) Set(val *TriggerInputAccountAttributesChangedAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccountAttributesChangedAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccountAttributesChangedAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccountAttributesChangedAccount(val *TriggerInputAccountAttributesChangedAccount) *NullableTriggerInputAccountAttributesChangedAccount {
	return &NullableTriggerInputAccountAttributesChangedAccount{value: val, isSet: true}
}

func (v NullableTriggerInputAccountAttributesChangedAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccountAttributesChangedAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


