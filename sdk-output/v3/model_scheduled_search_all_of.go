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

// ScheduledSearchAllOf struct for ScheduledSearchAllOf
type ScheduledSearchAllOf struct {
	// The scheduled search ID. 
	Id *string `json:"id,omitempty"`
	Owner *TypedReference `json:"owner,omitempty"`
	// The ID of the scheduled search owner
	OwnerId *string `json:"ownerId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduledSearchAllOf ScheduledSearchAllOf

// NewScheduledSearchAllOf instantiates a new ScheduledSearchAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledSearchAllOf() *ScheduledSearchAllOf {
	this := ScheduledSearchAllOf{}
	return &this
}

// NewScheduledSearchAllOfWithDefaults instantiates a new ScheduledSearchAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledSearchAllOfWithDefaults() *ScheduledSearchAllOf {
	this := ScheduledSearchAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScheduledSearchAllOf) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearchAllOf) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScheduledSearchAllOf) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScheduledSearchAllOf) SetId(v string) {
	o.Id = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ScheduledSearchAllOf) GetOwner() TypedReference {
	if o == nil || isNil(o.Owner) {
		var ret TypedReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearchAllOf) GetOwnerOk() (*TypedReference, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ScheduledSearchAllOf) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given TypedReference and assigns it to the Owner field.
func (o *ScheduledSearchAllOf) SetOwner(v TypedReference) {
	o.Owner = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ScheduledSearchAllOf) GetOwnerId() string {
	if o == nil || isNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearchAllOf) GetOwnerIdOk() (*string, bool) {
	if o == nil || isNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ScheduledSearchAllOf) HasOwnerId() bool {
	if o != nil && !isNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ScheduledSearchAllOf) SetOwnerId(v string) {
	o.OwnerId = &v
}

func (o ScheduledSearchAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ScheduledSearchAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varScheduledSearchAllOf := _ScheduledSearchAllOf{}

	if err = json.Unmarshal(bytes, &varScheduledSearchAllOf); err == nil {
		*o = ScheduledSearchAllOf(varScheduledSearchAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "ownerId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScheduledSearchAllOf struct {
	value *ScheduledSearchAllOf
	isSet bool
}

func (v NullableScheduledSearchAllOf) Get() *ScheduledSearchAllOf {
	return v.value
}

func (v *NullableScheduledSearchAllOf) Set(val *ScheduledSearchAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledSearchAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledSearchAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledSearchAllOf(val *ScheduledSearchAllOf) *NullableScheduledSearchAllOf {
	return &NullableScheduledSearchAllOf{value: val, isSet: true}
}

func (v NullableScheduledSearchAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledSearchAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


