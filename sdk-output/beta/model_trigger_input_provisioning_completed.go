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

// TriggerInputProvisioningCompleted struct for TriggerInputProvisioningCompleted
type TriggerInputProvisioningCompleted struct {
	// The reference number of the provisioning request. Useful for tracking status in the Account Activity search interface.
	TrackingNumber string `json:"trackingNumber"`
	// One or more sources that the provisioning transaction(s) were done against.  Sources are comma separated.
	Sources string `json:"sources"`
	// Origin of where the provisioning request came from.
	Action NullableString `json:"action,omitempty"`
	// A list of any accumulated error messages that occurred during provisioning.
	Errors []string `json:"errors,omitempty"`
	// A list of any accumulated warning messages that occurred during provisioning.
	Warnings []string `json:"warnings,omitempty"`
	Recipient TriggerInputProvisioningCompletedRecipient `json:"recipient"`
	Requester NullableTriggerInputProvisioningCompletedRequester `json:"requester,omitempty"`
	// A list of provisioning instructions to perform on an account-by-account basis.
	AccountRequests []TriggerInputProvisioningCompletedAccountRequestsInner `json:"accountRequests"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputProvisioningCompleted TriggerInputProvisioningCompleted

// NewTriggerInputProvisioningCompleted instantiates a new TriggerInputProvisioningCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputProvisioningCompleted(trackingNumber string, sources string, recipient TriggerInputProvisioningCompletedRecipient, accountRequests []TriggerInputProvisioningCompletedAccountRequestsInner) *TriggerInputProvisioningCompleted {
	this := TriggerInputProvisioningCompleted{}
	this.TrackingNumber = trackingNumber
	this.Sources = sources
	this.Recipient = recipient
	this.AccountRequests = accountRequests
	return &this
}

// NewTriggerInputProvisioningCompletedWithDefaults instantiates a new TriggerInputProvisioningCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputProvisioningCompletedWithDefaults() *TriggerInputProvisioningCompleted {
	this := TriggerInputProvisioningCompleted{}
	return &this
}

// GetTrackingNumber returns the TrackingNumber field value
func (o *TriggerInputProvisioningCompleted) GetTrackingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value
// and a boolean to check if the value has been set.
func (o *TriggerInputProvisioningCompleted) GetTrackingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackingNumber, true
}

// SetTrackingNumber sets field value
func (o *TriggerInputProvisioningCompleted) SetTrackingNumber(v string) {
	o.TrackingNumber = v
}

// GetSources returns the Sources field value
func (o *TriggerInputProvisioningCompleted) GetSources() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *TriggerInputProvisioningCompleted) GetSourcesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value
func (o *TriggerInputProvisioningCompleted) SetSources(v string) {
	o.Sources = v
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputProvisioningCompleted) GetAction() string {
	if o == nil || isNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputProvisioningCompleted) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *TriggerInputProvisioningCompleted) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *TriggerInputProvisioningCompleted) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *TriggerInputProvisioningCompleted) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *TriggerInputProvisioningCompleted) UnsetAction() {
	o.Action.Unset()
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputProvisioningCompleted) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputProvisioningCompleted) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *TriggerInputProvisioningCompleted) HasErrors() bool {
	if o != nil && isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *TriggerInputProvisioningCompleted) SetErrors(v []string) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputProvisioningCompleted) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputProvisioningCompleted) GetWarningsOk() ([]string, bool) {
	if o == nil || isNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *TriggerInputProvisioningCompleted) HasWarnings() bool {
	if o != nil && isNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *TriggerInputProvisioningCompleted) SetWarnings(v []string) {
	o.Warnings = v
}

// GetRecipient returns the Recipient field value
func (o *TriggerInputProvisioningCompleted) GetRecipient() TriggerInputProvisioningCompletedRecipient {
	if o == nil {
		var ret TriggerInputProvisioningCompletedRecipient
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *TriggerInputProvisioningCompleted) GetRecipientOk() (*TriggerInputProvisioningCompletedRecipient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *TriggerInputProvisioningCompleted) SetRecipient(v TriggerInputProvisioningCompletedRecipient) {
	o.Recipient = v
}

// GetRequester returns the Requester field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputProvisioningCompleted) GetRequester() TriggerInputProvisioningCompletedRequester {
	if o == nil || isNil(o.Requester.Get()) {
		var ret TriggerInputProvisioningCompletedRequester
		return ret
	}
	return *o.Requester.Get()
}

// GetRequesterOk returns a tuple with the Requester field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputProvisioningCompleted) GetRequesterOk() (*TriggerInputProvisioningCompletedRequester, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requester.Get(), o.Requester.IsSet()
}

// HasRequester returns a boolean if a field has been set.
func (o *TriggerInputProvisioningCompleted) HasRequester() bool {
	if o != nil && o.Requester.IsSet() {
		return true
	}

	return false
}

// SetRequester gets a reference to the given NullableTriggerInputProvisioningCompletedRequester and assigns it to the Requester field.
func (o *TriggerInputProvisioningCompleted) SetRequester(v TriggerInputProvisioningCompletedRequester) {
	o.Requester.Set(&v)
}
// SetRequesterNil sets the value for Requester to be an explicit nil
func (o *TriggerInputProvisioningCompleted) SetRequesterNil() {
	o.Requester.Set(nil)
}

// UnsetRequester ensures that no value is present for Requester, not even an explicit nil
func (o *TriggerInputProvisioningCompleted) UnsetRequester() {
	o.Requester.Unset()
}

// GetAccountRequests returns the AccountRequests field value
func (o *TriggerInputProvisioningCompleted) GetAccountRequests() []TriggerInputProvisioningCompletedAccountRequestsInner {
	if o == nil {
		var ret []TriggerInputProvisioningCompletedAccountRequestsInner
		return ret
	}

	return o.AccountRequests
}

// GetAccountRequestsOk returns a tuple with the AccountRequests field value
// and a boolean to check if the value has been set.
func (o *TriggerInputProvisioningCompleted) GetAccountRequestsOk() ([]TriggerInputProvisioningCompletedAccountRequestsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountRequests, true
}

// SetAccountRequests sets field value
func (o *TriggerInputProvisioningCompleted) SetAccountRequests(v []TriggerInputProvisioningCompletedAccountRequestsInner) {
	o.AccountRequests = v
}

func (o TriggerInputProvisioningCompleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["trackingNumber"] = o.TrackingNumber
	}
	if true {
		toSerialize["sources"] = o.Sources
	}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if true {
		toSerialize["recipient"] = o.Recipient
	}
	if o.Requester.IsSet() {
		toSerialize["requester"] = o.Requester.Get()
	}
	if true {
		toSerialize["accountRequests"] = o.AccountRequests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputProvisioningCompleted) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputProvisioningCompleted := _TriggerInputProvisioningCompleted{}

	if err = json.Unmarshal(bytes, &varTriggerInputProvisioningCompleted); err == nil {
		*o = TriggerInputProvisioningCompleted(varTriggerInputProvisioningCompleted)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "trackingNumber")
		delete(additionalProperties, "sources")
		delete(additionalProperties, "action")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "warnings")
		delete(additionalProperties, "recipient")
		delete(additionalProperties, "requester")
		delete(additionalProperties, "accountRequests")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputProvisioningCompleted struct {
	value *TriggerInputProvisioningCompleted
	isSet bool
}

func (v NullableTriggerInputProvisioningCompleted) Get() *TriggerInputProvisioningCompleted {
	return v.value
}

func (v *NullableTriggerInputProvisioningCompleted) Set(val *TriggerInputProvisioningCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputProvisioningCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputProvisioningCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputProvisioningCompleted(val *TriggerInputProvisioningCompleted) *NullableTriggerInputProvisioningCompleted {
	return &NullableTriggerInputProvisioningCompleted{value: val, isSet: true}
}

func (v NullableTriggerInputProvisioningCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputProvisioningCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


