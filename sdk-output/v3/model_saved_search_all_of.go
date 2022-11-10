/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
)

// SavedSearchAllOf struct for SavedSearchAllOf
type SavedSearchAllOf struct {
	// The saved search ID. 
	Id *string `json:"id,omitempty"`
	Owner *TypedReference `json:"owner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SavedSearchAllOf SavedSearchAllOf

// NewSavedSearchAllOf instantiates a new SavedSearchAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedSearchAllOf() *SavedSearchAllOf {
	this := SavedSearchAllOf{}
	return &this
}

// NewSavedSearchAllOfWithDefaults instantiates a new SavedSearchAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedSearchAllOfWithDefaults() *SavedSearchAllOf {
	this := SavedSearchAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SavedSearchAllOf) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearchAllOf) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SavedSearchAllOf) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SavedSearchAllOf) SetId(v string) {
	o.Id = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *SavedSearchAllOf) GetOwner() TypedReference {
	if o == nil || isNil(o.Owner) {
		var ret TypedReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearchAllOf) GetOwnerOk() (*TypedReference, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *SavedSearchAllOf) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given TypedReference and assigns it to the Owner field.
func (o *SavedSearchAllOf) SetOwner(v TypedReference) {
	o.Owner = &v
}

func (o SavedSearchAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SavedSearchAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSavedSearchAllOf := _SavedSearchAllOf{}

	if err = json.Unmarshal(bytes, &varSavedSearchAllOf); err == nil {
		*o = SavedSearchAllOf(varSavedSearchAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "owner")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSavedSearchAllOf struct {
	value *SavedSearchAllOf
	isSet bool
}

func (v NullableSavedSearchAllOf) Get() *SavedSearchAllOf {
	return v.value
}

func (v *NullableSavedSearchAllOf) Set(val *SavedSearchAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedSearchAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedSearchAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedSearchAllOf(val *SavedSearchAllOf) *NullableSavedSearchAllOf {
	return &NullableSavedSearchAllOf{value: val, isSet: true}
}

func (v NullableSavedSearchAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedSearchAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

