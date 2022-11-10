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

// ActivateCampaignOptions struct for ActivateCampaignOptions
type ActivateCampaignOptions struct {
	// The timezone must be in a valid ISO 8601 format. Timezones in ISO 8601 are represented as UTC (represented as 'Z') or as an offset from UTC. The offset format can be +/-hh:mm, +/-hhmm, or +/-hh.
	TimeZone *string `json:"timeZone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ActivateCampaignOptions ActivateCampaignOptions

// NewActivateCampaignOptions instantiates a new ActivateCampaignOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivateCampaignOptions() *ActivateCampaignOptions {
	this := ActivateCampaignOptions{}
	var timeZone string = "Z"
	this.TimeZone = &timeZone
	return &this
}

// NewActivateCampaignOptionsWithDefaults instantiates a new ActivateCampaignOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivateCampaignOptionsWithDefaults() *ActivateCampaignOptions {
	this := ActivateCampaignOptions{}
	var timeZone string = "Z"
	this.TimeZone = &timeZone
	return &this
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *ActivateCampaignOptions) GetTimeZone() string {
	if o == nil || isNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateCampaignOptions) GetTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *ActivateCampaignOptions) HasTimeZone() bool {
	if o != nil && !isNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *ActivateCampaignOptions) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o ActivateCampaignOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ActivateCampaignOptions) UnmarshalJSON(bytes []byte) (err error) {
	varActivateCampaignOptions := _ActivateCampaignOptions{}

	if err = json.Unmarshal(bytes, &varActivateCampaignOptions); err == nil {
		*o = ActivateCampaignOptions(varActivateCampaignOptions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "timeZone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableActivateCampaignOptions struct {
	value *ActivateCampaignOptions
	isSet bool
}

func (v NullableActivateCampaignOptions) Get() *ActivateCampaignOptions {
	return v.value
}

func (v *NullableActivateCampaignOptions) Set(val *ActivateCampaignOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableActivateCampaignOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableActivateCampaignOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivateCampaignOptions(val *ActivateCampaignOptions) *NullableActivateCampaignOptions {
	return &NullableActivateCampaignOptions{value: val, isSet: true}
}

func (v NullableActivateCampaignOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivateCampaignOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


