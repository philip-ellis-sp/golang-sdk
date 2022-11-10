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

// AccountAggregationCompleted struct for AccountAggregationCompleted
type AccountAggregationCompleted struct {
	Source TriggerInputAccountAggregationCompletedSource `json:"source"`
	// The overall status of the aggregation.
	Status map[string]interface{} `json:"status"`
	// The date and time when the account aggregation started.
	Started time.Time `json:"started"`
	// The date and time when the account aggregation finished.
	Completed time.Time `json:"completed"`
	// A list of errors that occurred during the aggregation.
	Errors []string `json:"errors"`
	// A list of warnings that occurred during the aggregation.
	Warnings []string `json:"warnings"`
	Stats TriggerInputAccountAggregationCompletedStats `json:"stats"`
	AdditionalProperties map[string]interface{}
}

type _AccountAggregationCompleted AccountAggregationCompleted

// NewAccountAggregationCompleted instantiates a new AccountAggregationCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAggregationCompleted(source TriggerInputAccountAggregationCompletedSource, status map[string]interface{}, started time.Time, completed time.Time, errors []string, warnings []string, stats TriggerInputAccountAggregationCompletedStats) *AccountAggregationCompleted {
	this := AccountAggregationCompleted{}
	this.Source = source
	this.Status = status
	this.Started = started
	this.Completed = completed
	this.Errors = errors
	this.Warnings = warnings
	this.Stats = stats
	return &this
}

// NewAccountAggregationCompletedWithDefaults instantiates a new AccountAggregationCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAggregationCompletedWithDefaults() *AccountAggregationCompleted {
	this := AccountAggregationCompleted{}
	return &this
}

// GetSource returns the Source field value
func (o *AccountAggregationCompleted) GetSource() TriggerInputAccountAggregationCompletedSource {
	if o == nil {
		var ret TriggerInputAccountAggregationCompletedSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AccountAggregationCompleted) GetSourceOk() (*TriggerInputAccountAggregationCompletedSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AccountAggregationCompleted) SetSource(v TriggerInputAccountAggregationCompletedSource) {
	o.Source = v
}

// GetStatus returns the Status field value
func (o *AccountAggregationCompleted) GetStatus() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccountAggregationCompleted) GetStatusOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *AccountAggregationCompleted) SetStatus(v map[string]interface{}) {
	o.Status = v
}

// GetStarted returns the Started field value
func (o *AccountAggregationCompleted) GetStarted() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Started
}

// GetStartedOk returns a tuple with the Started field value
// and a boolean to check if the value has been set.
func (o *AccountAggregationCompleted) GetStartedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Started, true
}

// SetStarted sets field value
func (o *AccountAggregationCompleted) SetStarted(v time.Time) {
	o.Started = v
}

// GetCompleted returns the Completed field value
func (o *AccountAggregationCompleted) GetCompleted() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value
// and a boolean to check if the value has been set.
func (o *AccountAggregationCompleted) GetCompletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completed, true
}

// SetCompleted sets field value
func (o *AccountAggregationCompleted) SetCompleted(v time.Time) {
	o.Completed = v
}

// GetErrors returns the Errors field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *AccountAggregationCompleted) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountAggregationCompleted) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *AccountAggregationCompleted) SetErrors(v []string) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *AccountAggregationCompleted) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountAggregationCompleted) GetWarningsOk() ([]string, bool) {
	if o == nil || isNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// SetWarnings sets field value
func (o *AccountAggregationCompleted) SetWarnings(v []string) {
	o.Warnings = v
}

// GetStats returns the Stats field value
func (o *AccountAggregationCompleted) GetStats() TriggerInputAccountAggregationCompletedStats {
	if o == nil {
		var ret TriggerInputAccountAggregationCompletedStats
		return ret
	}

	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value
// and a boolean to check if the value has been set.
func (o *AccountAggregationCompleted) GetStatsOk() (*TriggerInputAccountAggregationCompletedStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stats, true
}

// SetStats sets field value
func (o *AccountAggregationCompleted) SetStats(v TriggerInputAccountAggregationCompletedStats) {
	o.Stats = v
}

func (o AccountAggregationCompleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["started"] = o.Started
	}
	if true {
		toSerialize["completed"] = o.Completed
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if true {
		toSerialize["stats"] = o.Stats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountAggregationCompleted) UnmarshalJSON(bytes []byte) (err error) {
	varAccountAggregationCompleted := _AccountAggregationCompleted{}

	if err = json.Unmarshal(bytes, &varAccountAggregationCompleted); err == nil {
		*o = AccountAggregationCompleted(varAccountAggregationCompleted)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		delete(additionalProperties, "status")
		delete(additionalProperties, "started")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "warnings")
		delete(additionalProperties, "stats")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountAggregationCompleted struct {
	value *AccountAggregationCompleted
	isSet bool
}

func (v NullableAccountAggregationCompleted) Get() *AccountAggregationCompleted {
	return v.value
}

func (v *NullableAccountAggregationCompleted) Set(val *AccountAggregationCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAggregationCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAggregationCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAggregationCompleted(val *AccountAggregationCompleted) *NullableAccountAggregationCompleted {
	return &NullableAccountAggregationCompleted{value: val, isSet: true}
}

func (v NullableAccountAggregationCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAggregationCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


