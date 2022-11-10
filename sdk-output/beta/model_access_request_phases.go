/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"time"
)

// AccessRequestPhases Provides additional details about this access request phase.
type AccessRequestPhases struct {
	// The time that this phase started.
	Started *time.Time `json:"started,omitempty"`
	// The time that this phase finished.
	Finished *time.Time `json:"finished,omitempty"`
	// The name of this phase.
	Name *string `json:"name,omitempty"`
	// The state of this phase.
	State *string `json:"state,omitempty"`
	// The state of this phase.
	Result *string `json:"result,omitempty"`
	// A reference to another object on the RequestedItemStatus that contains more details about the phase. Note that for the Provisioning phase, this will be empty if there are no manual work items.
	PhaseReference *string `json:"phaseReference,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestPhases AccessRequestPhases

// NewAccessRequestPhases instantiates a new AccessRequestPhases object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestPhases() *AccessRequestPhases {
	this := AccessRequestPhases{}
	return &this
}

// NewAccessRequestPhasesWithDefaults instantiates a new AccessRequestPhases object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestPhasesWithDefaults() *AccessRequestPhases {
	this := AccessRequestPhases{}
	return &this
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *AccessRequestPhases) GetStarted() time.Time {
	if o == nil || isNil(o.Started) {
		var ret time.Time
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestPhases) GetStartedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Started) {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *AccessRequestPhases) HasStarted() bool {
	if o != nil && !isNil(o.Started) {
		return true
	}

	return false
}

// SetStarted gets a reference to the given time.Time and assigns it to the Started field.
func (o *AccessRequestPhases) SetStarted(v time.Time) {
	o.Started = &v
}

// GetFinished returns the Finished field value if set, zero value otherwise.
func (o *AccessRequestPhases) GetFinished() time.Time {
	if o == nil || isNil(o.Finished) {
		var ret time.Time
		return ret
	}
	return *o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestPhases) GetFinishedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Finished) {
		return nil, false
	}
	return o.Finished, true
}

// HasFinished returns a boolean if a field has been set.
func (o *AccessRequestPhases) HasFinished() bool {
	if o != nil && !isNil(o.Finished) {
		return true
	}

	return false
}

// SetFinished gets a reference to the given time.Time and assigns it to the Finished field.
func (o *AccessRequestPhases) SetFinished(v time.Time) {
	o.Finished = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessRequestPhases) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestPhases) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessRequestPhases) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessRequestPhases) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AccessRequestPhases) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestPhases) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AccessRequestPhases) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AccessRequestPhases) SetState(v string) {
	o.State = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AccessRequestPhases) GetResult() string {
	if o == nil || isNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestPhases) GetResultOk() (*string, bool) {
	if o == nil || isNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AccessRequestPhases) HasResult() bool {
	if o != nil && !isNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *AccessRequestPhases) SetResult(v string) {
	o.Result = &v
}

// GetPhaseReference returns the PhaseReference field value if set, zero value otherwise.
func (o *AccessRequestPhases) GetPhaseReference() string {
	if o == nil || isNil(o.PhaseReference) {
		var ret string
		return ret
	}
	return *o.PhaseReference
}

// GetPhaseReferenceOk returns a tuple with the PhaseReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestPhases) GetPhaseReferenceOk() (*string, bool) {
	if o == nil || isNil(o.PhaseReference) {
		return nil, false
	}
	return o.PhaseReference, true
}

// HasPhaseReference returns a boolean if a field has been set.
func (o *AccessRequestPhases) HasPhaseReference() bool {
	if o != nil && !isNil(o.PhaseReference) {
		return true
	}

	return false
}

// SetPhaseReference gets a reference to the given string and assigns it to the PhaseReference field.
func (o *AccessRequestPhases) SetPhaseReference(v string) {
	o.PhaseReference = &v
}

func (o AccessRequestPhases) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Started) {
		toSerialize["started"] = o.Started
	}
	if !isNil(o.Finished) {
		toSerialize["finished"] = o.Finished
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !isNil(o.PhaseReference) {
		toSerialize["phaseReference"] = o.PhaseReference
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessRequestPhases) UnmarshalJSON(bytes []byte) (err error) {
	varAccessRequestPhases := _AccessRequestPhases{}

	if err = json.Unmarshal(bytes, &varAccessRequestPhases); err == nil {
		*o = AccessRequestPhases(varAccessRequestPhases)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "started")
		delete(additionalProperties, "finished")
		delete(additionalProperties, "name")
		delete(additionalProperties, "state")
		delete(additionalProperties, "result")
		delete(additionalProperties, "phaseReference")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestPhases struct {
	value *AccessRequestPhases
	isSet bool
}

func (v NullableAccessRequestPhases) Get() *AccessRequestPhases {
	return v.value
}

func (v *NullableAccessRequestPhases) Set(val *AccessRequestPhases) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestPhases) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestPhases) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestPhases(val *AccessRequestPhases) *NullableAccessRequestPhases {
	return &NullableAccessRequestPhases{value: val, isSet: true}
}

func (v NullableAccessRequestPhases) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestPhases) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


