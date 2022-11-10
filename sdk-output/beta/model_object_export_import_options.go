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

// ObjectExportImportOptions struct for ObjectExportImportOptions
type ObjectExportImportOptions struct {
	// Object ids to be included in an import or export.
	IncludedIds []string `json:"includedIds,omitempty"`
	// Object names to be included in an import or export.
	IncludedNames []string `json:"includedNames,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ObjectExportImportOptions ObjectExportImportOptions

// NewObjectExportImportOptions instantiates a new ObjectExportImportOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectExportImportOptions() *ObjectExportImportOptions {
	this := ObjectExportImportOptions{}
	return &this
}

// NewObjectExportImportOptionsWithDefaults instantiates a new ObjectExportImportOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectExportImportOptionsWithDefaults() *ObjectExportImportOptions {
	this := ObjectExportImportOptions{}
	return &this
}

// GetIncludedIds returns the IncludedIds field value if set, zero value otherwise.
func (o *ObjectExportImportOptions) GetIncludedIds() []string {
	if o == nil || isNil(o.IncludedIds) {
		var ret []string
		return ret
	}
	return o.IncludedIds
}

// GetIncludedIdsOk returns a tuple with the IncludedIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectExportImportOptions) GetIncludedIdsOk() ([]string, bool) {
	if o == nil || isNil(o.IncludedIds) {
		return nil, false
	}
	return o.IncludedIds, true
}

// HasIncludedIds returns a boolean if a field has been set.
func (o *ObjectExportImportOptions) HasIncludedIds() bool {
	if o != nil && !isNil(o.IncludedIds) {
		return true
	}

	return false
}

// SetIncludedIds gets a reference to the given []string and assigns it to the IncludedIds field.
func (o *ObjectExportImportOptions) SetIncludedIds(v []string) {
	o.IncludedIds = v
}

// GetIncludedNames returns the IncludedNames field value if set, zero value otherwise.
func (o *ObjectExportImportOptions) GetIncludedNames() []string {
	if o == nil || isNil(o.IncludedNames) {
		var ret []string
		return ret
	}
	return o.IncludedNames
}

// GetIncludedNamesOk returns a tuple with the IncludedNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectExportImportOptions) GetIncludedNamesOk() ([]string, bool) {
	if o == nil || isNil(o.IncludedNames) {
		return nil, false
	}
	return o.IncludedNames, true
}

// HasIncludedNames returns a boolean if a field has been set.
func (o *ObjectExportImportOptions) HasIncludedNames() bool {
	if o != nil && !isNil(o.IncludedNames) {
		return true
	}

	return false
}

// SetIncludedNames gets a reference to the given []string and assigns it to the IncludedNames field.
func (o *ObjectExportImportOptions) SetIncludedNames(v []string) {
	o.IncludedNames = v
}

func (o ObjectExportImportOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IncludedIds) {
		toSerialize["includedIds"] = o.IncludedIds
	}
	if !isNil(o.IncludedNames) {
		toSerialize["includedNames"] = o.IncludedNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ObjectExportImportOptions) UnmarshalJSON(bytes []byte) (err error) {
	varObjectExportImportOptions := _ObjectExportImportOptions{}

	if err = json.Unmarshal(bytes, &varObjectExportImportOptions); err == nil {
		*o = ObjectExportImportOptions(varObjectExportImportOptions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "includedIds")
		delete(additionalProperties, "includedNames")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableObjectExportImportOptions struct {
	value *ObjectExportImportOptions
	isSet bool
}

func (v NullableObjectExportImportOptions) Get() *ObjectExportImportOptions {
	return v.value
}

func (v *NullableObjectExportImportOptions) Set(val *ObjectExportImportOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectExportImportOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectExportImportOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectExportImportOptions(val *ObjectExportImportOptions) *NullableObjectExportImportOptions {
	return &NullableObjectExportImportOptions{value: val, isSet: true}
}

func (v NullableObjectExportImportOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectExportImportOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


