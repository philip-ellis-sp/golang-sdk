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

// PasswordChangeResponse struct for PasswordChangeResponse
type PasswordChangeResponse struct {
	// The password change request ID
	RequestId NullableString `json:"requestId,omitempty"`
	// Password change state
	State *string `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordChangeResponse PasswordChangeResponse

// NewPasswordChangeResponse instantiates a new PasswordChangeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordChangeResponse() *PasswordChangeResponse {
	this := PasswordChangeResponse{}
	return &this
}

// NewPasswordChangeResponseWithDefaults instantiates a new PasswordChangeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordChangeResponseWithDefaults() *PasswordChangeResponse {
	this := PasswordChangeResponse{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordChangeResponse) GetRequestId() string {
	if o == nil || isNil(o.RequestId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestId.Get()
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordChangeResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestId.Get(), o.RequestId.IsSet()
}

// HasRequestId returns a boolean if a field has been set.
func (o *PasswordChangeResponse) HasRequestId() bool {
	if o != nil && o.RequestId.IsSet() {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given NullableString and assigns it to the RequestId field.
func (o *PasswordChangeResponse) SetRequestId(v string) {
	o.RequestId.Set(&v)
}
// SetRequestIdNil sets the value for RequestId to be an explicit nil
func (o *PasswordChangeResponse) SetRequestIdNil() {
	o.RequestId.Set(nil)
}

// UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
func (o *PasswordChangeResponse) UnsetRequestId() {
	o.RequestId.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PasswordChangeResponse) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordChangeResponse) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PasswordChangeResponse) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PasswordChangeResponse) SetState(v string) {
	o.State = &v
}

func (o PasswordChangeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId.IsSet() {
		toSerialize["requestId"] = o.RequestId.Get()
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordChangeResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordChangeResponse := _PasswordChangeResponse{}

	if err = json.Unmarshal(bytes, &varPasswordChangeResponse); err == nil {
		*o = PasswordChangeResponse(varPasswordChangeResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordChangeResponse struct {
	value *PasswordChangeResponse
	isSet bool
}

func (v NullablePasswordChangeResponse) Get() *PasswordChangeResponse {
	return v.value
}

func (v *NullablePasswordChangeResponse) Set(val *PasswordChangeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordChangeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordChangeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordChangeResponse(val *PasswordChangeResponse) *NullablePasswordChangeResponse {
	return &NullablePasswordChangeResponse{value: val, isSet: true}
}

func (v NullablePasswordChangeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordChangeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


