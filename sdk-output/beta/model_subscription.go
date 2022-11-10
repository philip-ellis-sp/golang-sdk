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

// Subscription struct for Subscription
type Subscription struct {
	// Subscription ID.
	Id string `json:"id"`
	// Subscription name.
	Name string `json:"name"`
	// Subscription description.
	Description *string `json:"description,omitempty"`
	// ID of trigger subscribed to.
	TriggerId string `json:"triggerId"`
	// Trigger name of trigger subscribed to.
	TriggerName string `json:"triggerName"`
	Type SubscriptionType `json:"type"`
	// Deadline for completing REQUEST_RESPONSE trigger invocation, represented in ISO-8601 duration format.
	ResponseDeadline string `json:"responseDeadline"`
	HttpConfig *HttpConfig `json:"httpConfig,omitempty"`
	EventBridgeConfig *EventBridgeConfig `json:"eventBridgeConfig,omitempty"`
	// Whether subscription should receive real-time trigger invocations or not. Test trigger invocations are always enabled regardless of this option.
	Enabled bool `json:"enabled"`
	// JSONPath filter to conditionally invoke trigger when expression evaluates to true.
	Filter *string `json:"filter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Subscription Subscription

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription(id string, name string, triggerId string, triggerName string, type_ SubscriptionType, responseDeadline string, enabled bool) *Subscription {
	this := Subscription{}
	this.Id = id
	this.Name = name
	this.TriggerId = triggerId
	this.TriggerName = triggerName
	this.Type = type_
	this.ResponseDeadline = responseDeadline
	this.Enabled = enabled
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	var responseDeadline string = "PT1H"
	this.ResponseDeadline = responseDeadline
	var enabled bool = true
	this.Enabled = enabled
	return &this
}

// GetId returns the Id field value
func (o *Subscription) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Subscription) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Subscription) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Subscription) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Subscription) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Subscription) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Subscription) SetDescription(v string) {
	o.Description = &v
}

// GetTriggerId returns the TriggerId field value
func (o *Subscription) GetTriggerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerId
}

// GetTriggerIdOk returns a tuple with the TriggerId field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetTriggerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerId, true
}

// SetTriggerId sets field value
func (o *Subscription) SetTriggerId(v string) {
	o.TriggerId = v
}

// GetTriggerName returns the TriggerName field value
func (o *Subscription) GetTriggerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerName
}

// GetTriggerNameOk returns a tuple with the TriggerName field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetTriggerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerName, true
}

// SetTriggerName sets field value
func (o *Subscription) SetTriggerName(v string) {
	o.TriggerName = v
}

// GetType returns the Type field value
func (o *Subscription) GetType() SubscriptionType {
	if o == nil {
		var ret SubscriptionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetTypeOk() (*SubscriptionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Subscription) SetType(v SubscriptionType) {
	o.Type = v
}

// GetResponseDeadline returns the ResponseDeadline field value
func (o *Subscription) GetResponseDeadline() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseDeadline
}

// GetResponseDeadlineOk returns a tuple with the ResponseDeadline field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetResponseDeadlineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseDeadline, true
}

// SetResponseDeadline sets field value
func (o *Subscription) SetResponseDeadline(v string) {
	o.ResponseDeadline = v
}

// GetHttpConfig returns the HttpConfig field value if set, zero value otherwise.
func (o *Subscription) GetHttpConfig() HttpConfig {
	if o == nil || isNil(o.HttpConfig) {
		var ret HttpConfig
		return ret
	}
	return *o.HttpConfig
}

// GetHttpConfigOk returns a tuple with the HttpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetHttpConfigOk() (*HttpConfig, bool) {
	if o == nil || isNil(o.HttpConfig) {
		return nil, false
	}
	return o.HttpConfig, true
}

// HasHttpConfig returns a boolean if a field has been set.
func (o *Subscription) HasHttpConfig() bool {
	if o != nil && !isNil(o.HttpConfig) {
		return true
	}

	return false
}

// SetHttpConfig gets a reference to the given HttpConfig and assigns it to the HttpConfig field.
func (o *Subscription) SetHttpConfig(v HttpConfig) {
	o.HttpConfig = &v
}

// GetEventBridgeConfig returns the EventBridgeConfig field value if set, zero value otherwise.
func (o *Subscription) GetEventBridgeConfig() EventBridgeConfig {
	if o == nil || isNil(o.EventBridgeConfig) {
		var ret EventBridgeConfig
		return ret
	}
	return *o.EventBridgeConfig
}

// GetEventBridgeConfigOk returns a tuple with the EventBridgeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetEventBridgeConfigOk() (*EventBridgeConfig, bool) {
	if o == nil || isNil(o.EventBridgeConfig) {
		return nil, false
	}
	return o.EventBridgeConfig, true
}

// HasEventBridgeConfig returns a boolean if a field has been set.
func (o *Subscription) HasEventBridgeConfig() bool {
	if o != nil && !isNil(o.EventBridgeConfig) {
		return true
	}

	return false
}

// SetEventBridgeConfig gets a reference to the given EventBridgeConfig and assigns it to the EventBridgeConfig field.
func (o *Subscription) SetEventBridgeConfig(v EventBridgeConfig) {
	o.EventBridgeConfig = &v
}

// GetEnabled returns the Enabled field value
func (o *Subscription) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Subscription) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *Subscription) GetFilter() string {
	if o == nil || isNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetFilterOk() (*string, bool) {
	if o == nil || isNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *Subscription) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *Subscription) SetFilter(v string) {
	o.Filter = &v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["triggerId"] = o.TriggerId
	}
	if true {
		toSerialize["triggerName"] = o.TriggerName
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["responseDeadline"] = o.ResponseDeadline
	}
	if !isNil(o.HttpConfig) {
		toSerialize["httpConfig"] = o.HttpConfig
	}
	if !isNil(o.EventBridgeConfig) {
		toSerialize["eventBridgeConfig"] = o.EventBridgeConfig
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Subscription) UnmarshalJSON(bytes []byte) (err error) {
	varSubscription := _Subscription{}

	if err = json.Unmarshal(bytes, &varSubscription); err == nil {
		*o = Subscription(varSubscription)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "triggerId")
		delete(additionalProperties, "triggerName")
		delete(additionalProperties, "type")
		delete(additionalProperties, "responseDeadline")
		delete(additionalProperties, "httpConfig")
		delete(additionalProperties, "eventBridgeConfig")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "filter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


