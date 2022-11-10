/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointbetasdk

import (
	"encoding/json"
)

// AccessItemRemoved struct for AccessItemRemoved
type AccessItemRemoved struct {
	AccessItem *AccessItemAssociatedAccessItem `json:"accessItem,omitempty"`
	// the identity id
	IdentityId *string `json:"identityId,omitempty"`
	// the event type
	EventType *string `json:"eventType,omitempty"`
	// the date of event
	Dt *string `json:"dt,omitempty"`
	GovernanceEvent *CorrelatedGovernanceEvent `json:"governanceEvent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessItemRemoved AccessItemRemoved

// NewAccessItemRemoved instantiates a new AccessItemRemoved object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessItemRemoved() *AccessItemRemoved {
	this := AccessItemRemoved{}
	return &this
}

// NewAccessItemRemovedWithDefaults instantiates a new AccessItemRemoved object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessItemRemovedWithDefaults() *AccessItemRemoved {
	this := AccessItemRemoved{}
	return &this
}

// GetAccessItem returns the AccessItem field value if set, zero value otherwise.
func (o *AccessItemRemoved) GetAccessItem() AccessItemAssociatedAccessItem {
	if o == nil || isNil(o.AccessItem) {
		var ret AccessItemAssociatedAccessItem
		return ret
	}
	return *o.AccessItem
}

// GetAccessItemOk returns a tuple with the AccessItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRemoved) GetAccessItemOk() (*AccessItemAssociatedAccessItem, bool) {
	if o == nil || isNil(o.AccessItem) {
		return nil, false
	}
	return o.AccessItem, true
}

// HasAccessItem returns a boolean if a field has been set.
func (o *AccessItemRemoved) HasAccessItem() bool {
	if o != nil && !isNil(o.AccessItem) {
		return true
	}

	return false
}

// SetAccessItem gets a reference to the given AccessItemAssociatedAccessItem and assigns it to the AccessItem field.
func (o *AccessItemRemoved) SetAccessItem(v AccessItemAssociatedAccessItem) {
	o.AccessItem = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *AccessItemRemoved) GetIdentityId() string {
	if o == nil || isNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRemoved) GetIdentityIdOk() (*string, bool) {
	if o == nil || isNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *AccessItemRemoved) HasIdentityId() bool {
	if o != nil && !isNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *AccessItemRemoved) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *AccessItemRemoved) GetEventType() string {
	if o == nil || isNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRemoved) GetEventTypeOk() (*string, bool) {
	if o == nil || isNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *AccessItemRemoved) HasEventType() bool {
	if o != nil && !isNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *AccessItemRemoved) SetEventType(v string) {
	o.EventType = &v
}

// GetDt returns the Dt field value if set, zero value otherwise.
func (o *AccessItemRemoved) GetDt() string {
	if o == nil || isNil(o.Dt) {
		var ret string
		return ret
	}
	return *o.Dt
}

// GetDtOk returns a tuple with the Dt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRemoved) GetDtOk() (*string, bool) {
	if o == nil || isNil(o.Dt) {
		return nil, false
	}
	return o.Dt, true
}

// HasDt returns a boolean if a field has been set.
func (o *AccessItemRemoved) HasDt() bool {
	if o != nil && !isNil(o.Dt) {
		return true
	}

	return false
}

// SetDt gets a reference to the given string and assigns it to the Dt field.
func (o *AccessItemRemoved) SetDt(v string) {
	o.Dt = &v
}

// GetGovernanceEvent returns the GovernanceEvent field value if set, zero value otherwise.
func (o *AccessItemRemoved) GetGovernanceEvent() CorrelatedGovernanceEvent {
	if o == nil || isNil(o.GovernanceEvent) {
		var ret CorrelatedGovernanceEvent
		return ret
	}
	return *o.GovernanceEvent
}

// GetGovernanceEventOk returns a tuple with the GovernanceEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemRemoved) GetGovernanceEventOk() (*CorrelatedGovernanceEvent, bool) {
	if o == nil || isNil(o.GovernanceEvent) {
		return nil, false
	}
	return o.GovernanceEvent, true
}

// HasGovernanceEvent returns a boolean if a field has been set.
func (o *AccessItemRemoved) HasGovernanceEvent() bool {
	if o != nil && !isNil(o.GovernanceEvent) {
		return true
	}

	return false
}

// SetGovernanceEvent gets a reference to the given CorrelatedGovernanceEvent and assigns it to the GovernanceEvent field.
func (o *AccessItemRemoved) SetGovernanceEvent(v CorrelatedGovernanceEvent) {
	o.GovernanceEvent = &v
}

func (o AccessItemRemoved) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccessItem) {
		toSerialize["accessItem"] = o.AccessItem
	}
	if !isNil(o.IdentityId) {
		toSerialize["identityId"] = o.IdentityId
	}
	if !isNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !isNil(o.Dt) {
		toSerialize["dt"] = o.Dt
	}
	if !isNil(o.GovernanceEvent) {
		toSerialize["governanceEvent"] = o.GovernanceEvent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessItemRemoved) UnmarshalJSON(bytes []byte) (err error) {
	varAccessItemRemoved := _AccessItemRemoved{}

	if err = json.Unmarshal(bytes, &varAccessItemRemoved); err == nil {
		*o = AccessItemRemoved(varAccessItemRemoved)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accessItem")
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "dt")
		delete(additionalProperties, "governanceEvent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessItemRemoved struct {
	value *AccessItemRemoved
	isSet bool
}

func (v NullableAccessItemRemoved) Get() *AccessItemRemoved {
	return v.value
}

func (v *NullableAccessItemRemoved) Set(val *AccessItemRemoved) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessItemRemoved) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessItemRemoved) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessItemRemoved(val *AccessItemRemoved) *NullableAccessItemRemoved {
	return &NullableAccessItemRemoved{value: val, isSet: true}
}

func (v NullableAccessItemRemoved) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessItemRemoved) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

