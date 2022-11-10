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

// IdentityAttributePreview struct for IdentityAttributePreview
type IdentityAttributePreview struct {
	// Name of the attribute that is being previewed.
	Name *string `json:"name,omitempty"`
	// Value that was derived during the preview.
	Value map[string]interface{} `json:"value,omitempty"`
	// The value of the attribute before the preview.
	PreviousValue map[string]interface{} `json:"previousValue,omitempty"`
	ErrorMessages []ErrorMessageDto `json:"errorMessages,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityAttributePreview IdentityAttributePreview

// NewIdentityAttributePreview instantiates a new IdentityAttributePreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAttributePreview() *IdentityAttributePreview {
	this := IdentityAttributePreview{}
	return &this
}

// NewIdentityAttributePreviewWithDefaults instantiates a new IdentityAttributePreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAttributePreviewWithDefaults() *IdentityAttributePreview {
	this := IdentityAttributePreview{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityAttributePreview) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttributePreview) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityAttributePreview) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityAttributePreview) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IdentityAttributePreview) GetValue() map[string]interface{} {
	if o == nil || isNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttributePreview) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IdentityAttributePreview) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *IdentityAttributePreview) SetValue(v map[string]interface{}) {
	o.Value = v
}

// GetPreviousValue returns the PreviousValue field value if set, zero value otherwise.
func (o *IdentityAttributePreview) GetPreviousValue() map[string]interface{} {
	if o == nil || isNil(o.PreviousValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.PreviousValue
}

// GetPreviousValueOk returns a tuple with the PreviousValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttributePreview) GetPreviousValueOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.PreviousValue) {
		return map[string]interface{}{}, false
	}
	return o.PreviousValue, true
}

// HasPreviousValue returns a boolean if a field has been set.
func (o *IdentityAttributePreview) HasPreviousValue() bool {
	if o != nil && !isNil(o.PreviousValue) {
		return true
	}

	return false
}

// SetPreviousValue gets a reference to the given map[string]interface{} and assigns it to the PreviousValue field.
func (o *IdentityAttributePreview) SetPreviousValue(v map[string]interface{}) {
	o.PreviousValue = v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise.
func (o *IdentityAttributePreview) GetErrorMessages() []ErrorMessageDto {
	if o == nil || isNil(o.ErrorMessages) {
		var ret []ErrorMessageDto
		return ret
	}
	return o.ErrorMessages
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttributePreview) GetErrorMessagesOk() ([]ErrorMessageDto, bool) {
	if o == nil || isNil(o.ErrorMessages) {
		return nil, false
	}
	return o.ErrorMessages, true
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *IdentityAttributePreview) HasErrorMessages() bool {
	if o != nil && !isNil(o.ErrorMessages) {
		return true
	}

	return false
}

// SetErrorMessages gets a reference to the given []ErrorMessageDto and assigns it to the ErrorMessages field.
func (o *IdentityAttributePreview) SetErrorMessages(v []ErrorMessageDto) {
	o.ErrorMessages = v
}

func (o IdentityAttributePreview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.PreviousValue) {
		toSerialize["previousValue"] = o.PreviousValue
	}
	if !isNil(o.ErrorMessages) {
		toSerialize["errorMessages"] = o.ErrorMessages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityAttributePreview) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityAttributePreview := _IdentityAttributePreview{}

	if err = json.Unmarshal(bytes, &varIdentityAttributePreview); err == nil {
		*o = IdentityAttributePreview(varIdentityAttributePreview)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		delete(additionalProperties, "previousValue")
		delete(additionalProperties, "errorMessages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityAttributePreview struct {
	value *IdentityAttributePreview
	isSet bool
}

func (v NullableIdentityAttributePreview) Get() *IdentityAttributePreview {
	return v.value
}

func (v *NullableIdentityAttributePreview) Set(val *IdentityAttributePreview) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAttributePreview) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAttributePreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAttributePreview(val *IdentityAttributePreview) *NullableIdentityAttributePreview {
	return &NullableIdentityAttributePreview{value: val, isSet: true}
}

func (v NullableIdentityAttributePreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAttributePreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


