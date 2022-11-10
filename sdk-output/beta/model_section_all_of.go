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

// SectionAllOf struct for SectionAllOf
type SectionAllOf struct {
	// Label of the section
	Label *string `json:"label,omitempty"`
	// List of FormItems. FormItems can be SectionDetails and/or FieldDetails
	FormItems []map[string]interface{} `json:"formItems,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SectionAllOf SectionAllOf

// NewSectionAllOf instantiates a new SectionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSectionAllOf() *SectionAllOf {
	this := SectionAllOf{}
	return &this
}

// NewSectionAllOfWithDefaults instantiates a new SectionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSectionAllOfWithDefaults() *SectionAllOf {
	this := SectionAllOf{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *SectionAllOf) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SectionAllOf) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *SectionAllOf) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *SectionAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetFormItems returns the FormItems field value if set, zero value otherwise.
func (o *SectionAllOf) GetFormItems() []map[string]interface{} {
	if o == nil || isNil(o.FormItems) {
		var ret []map[string]interface{}
		return ret
	}
	return o.FormItems
}

// GetFormItemsOk returns a tuple with the FormItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SectionAllOf) GetFormItemsOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.FormItems) {
		return nil, false
	}
	return o.FormItems, true
}

// HasFormItems returns a boolean if a field has been set.
func (o *SectionAllOf) HasFormItems() bool {
	if o != nil && !isNil(o.FormItems) {
		return true
	}

	return false
}

// SetFormItems gets a reference to the given []map[string]interface{} and assigns it to the FormItems field.
func (o *SectionAllOf) SetFormItems(v []map[string]interface{}) {
	o.FormItems = v
}

func (o SectionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.FormItems) {
		toSerialize["formItems"] = o.FormItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SectionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSectionAllOf := _SectionAllOf{}

	if err = json.Unmarshal(bytes, &varSectionAllOf); err == nil {
		*o = SectionAllOf(varSectionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "formItems")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSectionAllOf struct {
	value *SectionAllOf
	isSet bool
}

func (v NullableSectionAllOf) Get() *SectionAllOf {
	return v.value
}

func (v *NullableSectionAllOf) Set(val *SectionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSectionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSectionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSectionAllOf(val *SectionAllOf) *NullableSectionAllOf {
	return &NullableSectionAllOf{value: val, isSet: true}
}

func (v NullableSectionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSectionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


