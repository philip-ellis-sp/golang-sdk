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

// TranslationMessage struct for TranslationMessage
type TranslationMessage struct {
	// The key of the translation message
	Key *string `json:"key,omitempty"`
	// The values corresponding to the translation messages
	Values []string `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TranslationMessage TranslationMessage

// NewTranslationMessage instantiates a new TranslationMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranslationMessage() *TranslationMessage {
	this := TranslationMessage{}
	return &this
}

// NewTranslationMessageWithDefaults instantiates a new TranslationMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranslationMessageWithDefaults() *TranslationMessage {
	this := TranslationMessage{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TranslationMessage) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationMessage) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TranslationMessage) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *TranslationMessage) SetKey(v string) {
	o.Key = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *TranslationMessage) GetValues() []string {
	if o == nil || isNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationMessage) GetValuesOk() ([]string, bool) {
	if o == nil || isNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *TranslationMessage) HasValues() bool {
	if o != nil && !isNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *TranslationMessage) SetValues(v []string) {
	o.Values = v
}

func (o TranslationMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Values) {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TranslationMessage) UnmarshalJSON(bytes []byte) (err error) {
	varTranslationMessage := _TranslationMessage{}

	if err = json.Unmarshal(bytes, &varTranslationMessage); err == nil {
		*o = TranslationMessage(varTranslationMessage)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTranslationMessage struct {
	value *TranslationMessage
	isSet bool
}

func (v NullableTranslationMessage) Get() *TranslationMessage {
	return v.value
}

func (v *NullableTranslationMessage) Set(val *TranslationMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableTranslationMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableTranslationMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranslationMessage(val *TranslationMessage) *NullableTranslationMessage {
	return &NullableTranslationMessage{value: val, isSet: true}
}

func (v NullableTranslationMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranslationMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


