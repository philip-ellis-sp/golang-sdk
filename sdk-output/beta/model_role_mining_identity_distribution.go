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

// RoleMiningIdentityDistribution struct for RoleMiningIdentityDistribution
type RoleMiningIdentityDistribution struct {
	// Id of the potential role
	AttributeName *string `json:"attributeName,omitempty"`
	Distribution []map[string]string `json:"distribution,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningIdentityDistribution RoleMiningIdentityDistribution

// NewRoleMiningIdentityDistribution instantiates a new RoleMiningIdentityDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningIdentityDistribution() *RoleMiningIdentityDistribution {
	this := RoleMiningIdentityDistribution{}
	return &this
}

// NewRoleMiningIdentityDistributionWithDefaults instantiates a new RoleMiningIdentityDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningIdentityDistributionWithDefaults() *RoleMiningIdentityDistribution {
	this := RoleMiningIdentityDistribution{}
	return &this
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *RoleMiningIdentityDistribution) GetAttributeName() string {
	if o == nil || isNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningIdentityDistribution) GetAttributeNameOk() (*string, bool) {
	if o == nil || isNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *RoleMiningIdentityDistribution) HasAttributeName() bool {
	if o != nil && !isNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *RoleMiningIdentityDistribution) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetDistribution returns the Distribution field value if set, zero value otherwise.
func (o *RoleMiningIdentityDistribution) GetDistribution() []map[string]string {
	if o == nil || isNil(o.Distribution) {
		var ret []map[string]string
		return ret
	}
	return o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningIdentityDistribution) GetDistributionOk() ([]map[string]string, bool) {
	if o == nil || isNil(o.Distribution) {
		return nil, false
	}
	return o.Distribution, true
}

// HasDistribution returns a boolean if a field has been set.
func (o *RoleMiningIdentityDistribution) HasDistribution() bool {
	if o != nil && !isNil(o.Distribution) {
		return true
	}

	return false
}

// SetDistribution gets a reference to the given []map[string]string and assigns it to the Distribution field.
func (o *RoleMiningIdentityDistribution) SetDistribution(v []map[string]string) {
	o.Distribution = v
}

func (o RoleMiningIdentityDistribution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !isNil(o.Distribution) {
		toSerialize["distribution"] = o.Distribution
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleMiningIdentityDistribution) UnmarshalJSON(bytes []byte) (err error) {
	varRoleMiningIdentityDistribution := _RoleMiningIdentityDistribution{}

	if err = json.Unmarshal(bytes, &varRoleMiningIdentityDistribution); err == nil {
		*o = RoleMiningIdentityDistribution(varRoleMiningIdentityDistribution)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "attributeName")
		delete(additionalProperties, "distribution")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningIdentityDistribution struct {
	value *RoleMiningIdentityDistribution
	isSet bool
}

func (v NullableRoleMiningIdentityDistribution) Get() *RoleMiningIdentityDistribution {
	return v.value
}

func (v *NullableRoleMiningIdentityDistribution) Set(val *RoleMiningIdentityDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningIdentityDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningIdentityDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningIdentityDistribution(val *RoleMiningIdentityDistribution) *NullableRoleMiningIdentityDistribution {
	return &NullableRoleMiningIdentityDistribution{value: val, isSet: true}
}

func (v NullableRoleMiningIdentityDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningIdentityDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


