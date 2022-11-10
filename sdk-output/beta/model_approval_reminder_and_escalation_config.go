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

// ApprovalReminderAndEscalationConfig struct for ApprovalReminderAndEscalationConfig
type ApprovalReminderAndEscalationConfig struct {
	// Number of days to wait before the first reminder. If no reminders are configured, then this is the number of days to wait before escalation.
	DaysUntilEscalation *int32 `json:"daysUntilEscalation,omitempty"`
	// Number of days to wait between reminder notifications.
	DaysBetweenReminders *int32 `json:"daysBetweenReminders,omitempty"`
	// Maximum number of reminder notification to send to the reviewer before approval escalation.
	MaxReminders *int32 `json:"maxReminders,omitempty"`
	FallbackApproverRef NullableIdentityReferenceWithNameAndEmail `json:"fallbackApproverRef,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalReminderAndEscalationConfig ApprovalReminderAndEscalationConfig

// NewApprovalReminderAndEscalationConfig instantiates a new ApprovalReminderAndEscalationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalReminderAndEscalationConfig() *ApprovalReminderAndEscalationConfig {
	this := ApprovalReminderAndEscalationConfig{}
	return &this
}

// NewApprovalReminderAndEscalationConfigWithDefaults instantiates a new ApprovalReminderAndEscalationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalReminderAndEscalationConfigWithDefaults() *ApprovalReminderAndEscalationConfig {
	this := ApprovalReminderAndEscalationConfig{}
	return &this
}

// GetDaysUntilEscalation returns the DaysUntilEscalation field value if set, zero value otherwise.
func (o *ApprovalReminderAndEscalationConfig) GetDaysUntilEscalation() int32 {
	if o == nil || isNil(o.DaysUntilEscalation) {
		var ret int32
		return ret
	}
	return *o.DaysUntilEscalation
}

// GetDaysUntilEscalationOk returns a tuple with the DaysUntilEscalation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalReminderAndEscalationConfig) GetDaysUntilEscalationOk() (*int32, bool) {
	if o == nil || isNil(o.DaysUntilEscalation) {
		return nil, false
	}
	return o.DaysUntilEscalation, true
}

// HasDaysUntilEscalation returns a boolean if a field has been set.
func (o *ApprovalReminderAndEscalationConfig) HasDaysUntilEscalation() bool {
	if o != nil && !isNil(o.DaysUntilEscalation) {
		return true
	}

	return false
}

// SetDaysUntilEscalation gets a reference to the given int32 and assigns it to the DaysUntilEscalation field.
func (o *ApprovalReminderAndEscalationConfig) SetDaysUntilEscalation(v int32) {
	o.DaysUntilEscalation = &v
}

// GetDaysBetweenReminders returns the DaysBetweenReminders field value if set, zero value otherwise.
func (o *ApprovalReminderAndEscalationConfig) GetDaysBetweenReminders() int32 {
	if o == nil || isNil(o.DaysBetweenReminders) {
		var ret int32
		return ret
	}
	return *o.DaysBetweenReminders
}

// GetDaysBetweenRemindersOk returns a tuple with the DaysBetweenReminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalReminderAndEscalationConfig) GetDaysBetweenRemindersOk() (*int32, bool) {
	if o == nil || isNil(o.DaysBetweenReminders) {
		return nil, false
	}
	return o.DaysBetweenReminders, true
}

// HasDaysBetweenReminders returns a boolean if a field has been set.
func (o *ApprovalReminderAndEscalationConfig) HasDaysBetweenReminders() bool {
	if o != nil && !isNil(o.DaysBetweenReminders) {
		return true
	}

	return false
}

// SetDaysBetweenReminders gets a reference to the given int32 and assigns it to the DaysBetweenReminders field.
func (o *ApprovalReminderAndEscalationConfig) SetDaysBetweenReminders(v int32) {
	o.DaysBetweenReminders = &v
}

// GetMaxReminders returns the MaxReminders field value if set, zero value otherwise.
func (o *ApprovalReminderAndEscalationConfig) GetMaxReminders() int32 {
	if o == nil || isNil(o.MaxReminders) {
		var ret int32
		return ret
	}
	return *o.MaxReminders
}

// GetMaxRemindersOk returns a tuple with the MaxReminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalReminderAndEscalationConfig) GetMaxRemindersOk() (*int32, bool) {
	if o == nil || isNil(o.MaxReminders) {
		return nil, false
	}
	return o.MaxReminders, true
}

// HasMaxReminders returns a boolean if a field has been set.
func (o *ApprovalReminderAndEscalationConfig) HasMaxReminders() bool {
	if o != nil && !isNil(o.MaxReminders) {
		return true
	}

	return false
}

// SetMaxReminders gets a reference to the given int32 and assigns it to the MaxReminders field.
func (o *ApprovalReminderAndEscalationConfig) SetMaxReminders(v int32) {
	o.MaxReminders = &v
}

// GetFallbackApproverRef returns the FallbackApproverRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalReminderAndEscalationConfig) GetFallbackApproverRef() IdentityReferenceWithNameAndEmail {
	if o == nil || isNil(o.FallbackApproverRef.Get()) {
		var ret IdentityReferenceWithNameAndEmail
		return ret
	}
	return *o.FallbackApproverRef.Get()
}

// GetFallbackApproverRefOk returns a tuple with the FallbackApproverRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalReminderAndEscalationConfig) GetFallbackApproverRefOk() (*IdentityReferenceWithNameAndEmail, bool) {
	if o == nil {
		return nil, false
	}
	return o.FallbackApproverRef.Get(), o.FallbackApproverRef.IsSet()
}

// HasFallbackApproverRef returns a boolean if a field has been set.
func (o *ApprovalReminderAndEscalationConfig) HasFallbackApproverRef() bool {
	if o != nil && o.FallbackApproverRef.IsSet() {
		return true
	}

	return false
}

// SetFallbackApproverRef gets a reference to the given NullableIdentityReferenceWithNameAndEmail and assigns it to the FallbackApproverRef field.
func (o *ApprovalReminderAndEscalationConfig) SetFallbackApproverRef(v IdentityReferenceWithNameAndEmail) {
	o.FallbackApproverRef.Set(&v)
}
// SetFallbackApproverRefNil sets the value for FallbackApproverRef to be an explicit nil
func (o *ApprovalReminderAndEscalationConfig) SetFallbackApproverRefNil() {
	o.FallbackApproverRef.Set(nil)
}

// UnsetFallbackApproverRef ensures that no value is present for FallbackApproverRef, not even an explicit nil
func (o *ApprovalReminderAndEscalationConfig) UnsetFallbackApproverRef() {
	o.FallbackApproverRef.Unset()
}

func (o ApprovalReminderAndEscalationConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DaysUntilEscalation) {
		toSerialize["daysUntilEscalation"] = o.DaysUntilEscalation
	}
	if !isNil(o.DaysBetweenReminders) {
		toSerialize["daysBetweenReminders"] = o.DaysBetweenReminders
	}
	if !isNil(o.MaxReminders) {
		toSerialize["maxReminders"] = o.MaxReminders
	}
	if o.FallbackApproverRef.IsSet() {
		toSerialize["fallbackApproverRef"] = o.FallbackApproverRef.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApprovalReminderAndEscalationConfig) UnmarshalJSON(bytes []byte) (err error) {
	varApprovalReminderAndEscalationConfig := _ApprovalReminderAndEscalationConfig{}

	if err = json.Unmarshal(bytes, &varApprovalReminderAndEscalationConfig); err == nil {
		*o = ApprovalReminderAndEscalationConfig(varApprovalReminderAndEscalationConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "daysUntilEscalation")
		delete(additionalProperties, "daysBetweenReminders")
		delete(additionalProperties, "maxReminders")
		delete(additionalProperties, "fallbackApproverRef")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalReminderAndEscalationConfig struct {
	value *ApprovalReminderAndEscalationConfig
	isSet bool
}

func (v NullableApprovalReminderAndEscalationConfig) Get() *ApprovalReminderAndEscalationConfig {
	return v.value
}

func (v *NullableApprovalReminderAndEscalationConfig) Set(val *ApprovalReminderAndEscalationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalReminderAndEscalationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalReminderAndEscalationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalReminderAndEscalationConfig(val *ApprovalReminderAndEscalationConfig) *NullableApprovalReminderAndEscalationConfig {
	return &NullableApprovalReminderAndEscalationConfig{value: val, isSet: true}
}

func (v NullableApprovalReminderAndEscalationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalReminderAndEscalationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


