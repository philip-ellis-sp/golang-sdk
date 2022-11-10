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

// TemplateBulkDeleteDto struct for TemplateBulkDeleteDto
type TemplateBulkDeleteDto struct {
	Key string `json:"key"`
	Medium *string `json:"medium,omitempty"`
	// The locale for the message text, a BCP 47 language tag.
	Locale *string `json:"locale,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateBulkDeleteDto TemplateBulkDeleteDto

// NewTemplateBulkDeleteDto instantiates a new TemplateBulkDeleteDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateBulkDeleteDto(key string) *TemplateBulkDeleteDto {
	this := TemplateBulkDeleteDto{}
	this.Key = key
	return &this
}

// NewTemplateBulkDeleteDtoWithDefaults instantiates a new TemplateBulkDeleteDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateBulkDeleteDtoWithDefaults() *TemplateBulkDeleteDto {
	this := TemplateBulkDeleteDto{}
	return &this
}

// GetKey returns the Key field value
func (o *TemplateBulkDeleteDto) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TemplateBulkDeleteDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TemplateBulkDeleteDto) SetKey(v string) {
	o.Key = v
}

// GetMedium returns the Medium field value if set, zero value otherwise.
func (o *TemplateBulkDeleteDto) GetMedium() string {
	if o == nil || isNil(o.Medium) {
		var ret string
		return ret
	}
	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateBulkDeleteDto) GetMediumOk() (*string, bool) {
	if o == nil || isNil(o.Medium) {
		return nil, false
	}
	return o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *TemplateBulkDeleteDto) HasMedium() bool {
	if o != nil && !isNil(o.Medium) {
		return true
	}

	return false
}

// SetMedium gets a reference to the given string and assigns it to the Medium field.
func (o *TemplateBulkDeleteDto) SetMedium(v string) {
	o.Medium = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *TemplateBulkDeleteDto) GetLocale() string {
	if o == nil || isNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateBulkDeleteDto) GetLocaleOk() (*string, bool) {
	if o == nil || isNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *TemplateBulkDeleteDto) HasLocale() bool {
	if o != nil && !isNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *TemplateBulkDeleteDto) SetLocale(v string) {
	o.Locale = &v
}

func (o TemplateBulkDeleteDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Medium) {
		toSerialize["medium"] = o.Medium
	}
	if !isNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TemplateBulkDeleteDto) UnmarshalJSON(bytes []byte) (err error) {
	varTemplateBulkDeleteDto := _TemplateBulkDeleteDto{}

	if err = json.Unmarshal(bytes, &varTemplateBulkDeleteDto); err == nil {
		*o = TemplateBulkDeleteDto(varTemplateBulkDeleteDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "medium")
		delete(additionalProperties, "locale")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateBulkDeleteDto struct {
	value *TemplateBulkDeleteDto
	isSet bool
}

func (v NullableTemplateBulkDeleteDto) Get() *TemplateBulkDeleteDto {
	return v.value
}

func (v *NullableTemplateBulkDeleteDto) Set(val *TemplateBulkDeleteDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateBulkDeleteDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateBulkDeleteDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateBulkDeleteDto(val *TemplateBulkDeleteDto) *NullableTemplateBulkDeleteDto {
	return &NullableTemplateBulkDeleteDto{value: val, isSet: true}
}

func (v NullableTemplateBulkDeleteDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateBulkDeleteDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


