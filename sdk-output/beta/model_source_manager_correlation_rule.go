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

// SourceManagerCorrelationRule Reference to the ManagerCorrelationRule, only used when a simple filter isn't sufficient.
type SourceManagerCorrelationRule struct {
	// The type of object being referenced
	Type *string `json:"type,omitempty"`
	// ID of the rule
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the rule
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourceManagerCorrelationRule SourceManagerCorrelationRule

// NewSourceManagerCorrelationRule instantiates a new SourceManagerCorrelationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceManagerCorrelationRule() *SourceManagerCorrelationRule {
	this := SourceManagerCorrelationRule{}
	return &this
}

// NewSourceManagerCorrelationRuleWithDefaults instantiates a new SourceManagerCorrelationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceManagerCorrelationRuleWithDefaults() *SourceManagerCorrelationRule {
	this := SourceManagerCorrelationRule{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SourceManagerCorrelationRule) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManagerCorrelationRule) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SourceManagerCorrelationRule) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SourceManagerCorrelationRule) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourceManagerCorrelationRule) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManagerCorrelationRule) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourceManagerCorrelationRule) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourceManagerCorrelationRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceManagerCorrelationRule) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceManagerCorrelationRule) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceManagerCorrelationRule) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceManagerCorrelationRule) SetName(v string) {
	o.Name = &v
}

func (o SourceManagerCorrelationRule) MarshalJSON() ([]byte, error) {
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

func (o *SourceManagerCorrelationRule) UnmarshalJSON(bytes []byte) (err error) {
	varSourceManagerCorrelationRule := _SourceManagerCorrelationRule{}

	if err = json.Unmarshal(bytes, &varSourceManagerCorrelationRule); err == nil {
		*o = SourceManagerCorrelationRule(varSourceManagerCorrelationRule)
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

type NullableSourceManagerCorrelationRule struct {
	value *SourceManagerCorrelationRule
	isSet bool
}

func (v NullableSourceManagerCorrelationRule) Get() *SourceManagerCorrelationRule {
	return v.value
}

func (v *NullableSourceManagerCorrelationRule) Set(val *SourceManagerCorrelationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceManagerCorrelationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceManagerCorrelationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceManagerCorrelationRule(val *SourceManagerCorrelationRule) *NullableSourceManagerCorrelationRule {
	return &NullableSourceManagerCorrelationRule{value: val, isSet: true}
}

func (v NullableSourceManagerCorrelationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceManagerCorrelationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


