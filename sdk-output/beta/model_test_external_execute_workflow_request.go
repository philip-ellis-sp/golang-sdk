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

// TestExternalExecuteWorkflowRequest struct for TestExternalExecuteWorkflowRequest
type TestExternalExecuteWorkflowRequest struct {
	// The test input for the workflow
	Input map[string]interface{} `json:"input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TestExternalExecuteWorkflowRequest TestExternalExecuteWorkflowRequest

// NewTestExternalExecuteWorkflowRequest instantiates a new TestExternalExecuteWorkflowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestExternalExecuteWorkflowRequest() *TestExternalExecuteWorkflowRequest {
	this := TestExternalExecuteWorkflowRequest{}
	return &this
}

// NewTestExternalExecuteWorkflowRequestWithDefaults instantiates a new TestExternalExecuteWorkflowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestExternalExecuteWorkflowRequestWithDefaults() *TestExternalExecuteWorkflowRequest {
	this := TestExternalExecuteWorkflowRequest{}
	return &this
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *TestExternalExecuteWorkflowRequest) GetInput() map[string]interface{} {
	if o == nil || isNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestExternalExecuteWorkflowRequest) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *TestExternalExecuteWorkflowRequest) HasInput() bool {
	if o != nil && !isNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *TestExternalExecuteWorkflowRequest) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o TestExternalExecuteWorkflowRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Input) {
		toSerialize["input"] = o.Input
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TestExternalExecuteWorkflowRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTestExternalExecuteWorkflowRequest := _TestExternalExecuteWorkflowRequest{}

	if err = json.Unmarshal(bytes, &varTestExternalExecuteWorkflowRequest); err == nil {
		*o = TestExternalExecuteWorkflowRequest(varTestExternalExecuteWorkflowRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestExternalExecuteWorkflowRequest struct {
	value *TestExternalExecuteWorkflowRequest
	isSet bool
}

func (v NullableTestExternalExecuteWorkflowRequest) Get() *TestExternalExecuteWorkflowRequest {
	return v.value
}

func (v *NullableTestExternalExecuteWorkflowRequest) Set(val *TestExternalExecuteWorkflowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTestExternalExecuteWorkflowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTestExternalExecuteWorkflowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestExternalExecuteWorkflowRequest(val *TestExternalExecuteWorkflowRequest) *NullableTestExternalExecuteWorkflowRequest {
	return &NullableTestExternalExecuteWorkflowRequest{value: val, isSet: true}
}

func (v NullableTestExternalExecuteWorkflowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestExternalExecuteWorkflowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


