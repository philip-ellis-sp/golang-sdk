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

// ErrorMessageDto struct for ErrorMessageDto
type ErrorMessageDto struct {
	// The locale for the message text, a BCP 47 language tag.
	Locale *string `json:"locale,omitempty"`
	LocaleOrigin *LocaleOrigin `json:"localeOrigin,omitempty"`
	// Actual text of the error message in the indicated locale.
	Text *string `json:"text,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorMessageDto ErrorMessageDto

// NewErrorMessageDto instantiates a new ErrorMessageDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorMessageDto() *ErrorMessageDto {
	this := ErrorMessageDto{}
	return &this
}

// NewErrorMessageDtoWithDefaults instantiates a new ErrorMessageDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorMessageDtoWithDefaults() *ErrorMessageDto {
	this := ErrorMessageDto{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *ErrorMessageDto) GetLocale() string {
	if o == nil || isNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessageDto) GetLocaleOk() (*string, bool) {
	if o == nil || isNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *ErrorMessageDto) HasLocale() bool {
	if o != nil && !isNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *ErrorMessageDto) SetLocale(v string) {
	o.Locale = &v
}

// GetLocaleOrigin returns the LocaleOrigin field value if set, zero value otherwise.
func (o *ErrorMessageDto) GetLocaleOrigin() LocaleOrigin {
	if o == nil || isNil(o.LocaleOrigin) {
		var ret LocaleOrigin
		return ret
	}
	return *o.LocaleOrigin
}

// GetLocaleOriginOk returns a tuple with the LocaleOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessageDto) GetLocaleOriginOk() (*LocaleOrigin, bool) {
	if o == nil || isNil(o.LocaleOrigin) {
		return nil, false
	}
	return o.LocaleOrigin, true
}

// HasLocaleOrigin returns a boolean if a field has been set.
func (o *ErrorMessageDto) HasLocaleOrigin() bool {
	if o != nil && !isNil(o.LocaleOrigin) {
		return true
	}

	return false
}

// SetLocaleOrigin gets a reference to the given LocaleOrigin and assigns it to the LocaleOrigin field.
func (o *ErrorMessageDto) SetLocaleOrigin(v LocaleOrigin) {
	o.LocaleOrigin = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ErrorMessageDto) GetText() string {
	if o == nil || isNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessageDto) GetTextOk() (*string, bool) {
	if o == nil || isNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ErrorMessageDto) HasText() bool {
	if o != nil && !isNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *ErrorMessageDto) SetText(v string) {
	o.Text = &v
}

func (o ErrorMessageDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !isNil(o.LocaleOrigin) {
		toSerialize["localeOrigin"] = o.LocaleOrigin
	}
	if !isNil(o.Text) {
		toSerialize["text"] = o.Text
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ErrorMessageDto) UnmarshalJSON(bytes []byte) (err error) {
	varErrorMessageDto := _ErrorMessageDto{}

	if err = json.Unmarshal(bytes, &varErrorMessageDto); err == nil {
		*o = ErrorMessageDto(varErrorMessageDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "locale")
		delete(additionalProperties, "localeOrigin")
		delete(additionalProperties, "text")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorMessageDto struct {
	value *ErrorMessageDto
	isSet bool
}

func (v NullableErrorMessageDto) Get() *ErrorMessageDto {
	return v.value
}

func (v *NullableErrorMessageDto) Set(val *ErrorMessageDto) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorMessageDto) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorMessageDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorMessageDto(val *ErrorMessageDto) *NullableErrorMessageDto {
	return &NullableErrorMessageDto{value: val, isSet: true}
}

func (v NullableErrorMessageDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorMessageDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


