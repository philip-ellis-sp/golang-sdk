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

// TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner struct for TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner
type TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner struct {
	// The name of the attribute being provisioned.
	AttributeName string `json:"attributeName"`
	// The value of the attribute being provisioned.
	AttributeValue NullableString `json:"attributeValue,omitempty"`
	// The operation to handle the attribute.
	Operation map[string]interface{} `json:"operation"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner

// NewTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner instantiates a new TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner(attributeName string, operation map[string]interface{}) *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner {
	this := TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner{}
	this.AttributeName = attributeName
	this.Operation = operation
	return &this
}

// NewTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInnerWithDefaults instantiates a new TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInnerWithDefaults() *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner {
	this := TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner{}
	return &this
}

// GetAttributeName returns the AttributeName field value
func (o *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) GetAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value
// and a boolean to check if the value has been set.
func (o *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) GetAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeName, true
}

// SetAttributeName sets field value
func (o *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) SetAttributeName(v string) {
	o.AttributeName = v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) GetAttributeValue() string {
	if o == nil || isNil(o.AttributeValue.Get()) {
		var ret string
		return ret
	}
	return *o.AttributeValue.Get()
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) GetAttributeValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeValue.Get(), o.AttributeValue.IsSet()
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) HasAttributeValue() bool {
	if o != nil && o.AttributeValue.IsSet() {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given NullableString and assigns it to the AttributeValue field.
func (o *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) SetAttributeValue(v string) {
	o.AttributeValue.Set(&v)
}
// SetAttributeValueNil sets the value for AttributeValue to be an explicit nil
func (o *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) SetAttributeValueNil() {
	o.AttributeValue.Set(nil)
}

// UnsetAttributeValue ensures that no value is present for AttributeValue, not even an explicit nil
func (o *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) UnsetAttributeValue() {
	o.AttributeValue.Unset()
}

// GetOperation returns the Operation field value
func (o *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) GetOperation() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) GetOperationOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Operation, true
}

// SetOperation sets field value
func (o *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) SetOperation(v map[string]interface{}) {
	o.Operation = v
}

func (o TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributeName"] = o.AttributeName
	}
	if o.AttributeValue.IsSet() {
		toSerialize["attributeValue"] = o.AttributeValue.Get()
	}
	if true {
		toSerialize["operation"] = o.Operation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner := _TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner{}

	if err = json.Unmarshal(bytes, &varTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner); err == nil {
		*o = TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner(varTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "attributeName")
		delete(additionalProperties, "attributeValue")
		delete(additionalProperties, "operation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner struct {
	value *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner
	isSet bool
}

func (v NullableTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) Get() *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner {
	return v.value
}

func (v *NullableTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) Set(val *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner(val *TriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) *NullableTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner {
	return &NullableTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner{value: val, isSet: true}
}

func (v NullableTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputProvisioningCompletedAccountRequestsInnerAttributeRequestsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


