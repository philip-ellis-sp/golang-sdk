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

// IdentityPreviewRequest struct for IdentityPreviewRequest
type IdentityPreviewRequest struct {
	IdentityId *string `json:"identityId,omitempty"`
	IdentityAttributeConfig *IdentityAttributeConfig `json:"identityAttributeConfig,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityPreviewRequest IdentityPreviewRequest

// NewIdentityPreviewRequest instantiates a new IdentityPreviewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityPreviewRequest() *IdentityPreviewRequest {
	this := IdentityPreviewRequest{}
	return &this
}

// NewIdentityPreviewRequestWithDefaults instantiates a new IdentityPreviewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityPreviewRequestWithDefaults() *IdentityPreviewRequest {
	this := IdentityPreviewRequest{}
	return &this
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *IdentityPreviewRequest) GetIdentityId() string {
	if o == nil || isNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityPreviewRequest) GetIdentityIdOk() (*string, bool) {
	if o == nil || isNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *IdentityPreviewRequest) HasIdentityId() bool {
	if o != nil && !isNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *IdentityPreviewRequest) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetIdentityAttributeConfig returns the IdentityAttributeConfig field value if set, zero value otherwise.
func (o *IdentityPreviewRequest) GetIdentityAttributeConfig() IdentityAttributeConfig {
	if o == nil || isNil(o.IdentityAttributeConfig) {
		var ret IdentityAttributeConfig
		return ret
	}
	return *o.IdentityAttributeConfig
}

// GetIdentityAttributeConfigOk returns a tuple with the IdentityAttributeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityPreviewRequest) GetIdentityAttributeConfigOk() (*IdentityAttributeConfig, bool) {
	if o == nil || isNil(o.IdentityAttributeConfig) {
		return nil, false
	}
	return o.IdentityAttributeConfig, true
}

// HasIdentityAttributeConfig returns a boolean if a field has been set.
func (o *IdentityPreviewRequest) HasIdentityAttributeConfig() bool {
	if o != nil && !isNil(o.IdentityAttributeConfig) {
		return true
	}

	return false
}

// SetIdentityAttributeConfig gets a reference to the given IdentityAttributeConfig and assigns it to the IdentityAttributeConfig field.
func (o *IdentityPreviewRequest) SetIdentityAttributeConfig(v IdentityAttributeConfig) {
	o.IdentityAttributeConfig = &v
}

func (o IdentityPreviewRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdentityId) {
		toSerialize["identityId"] = o.IdentityId
	}
	if !isNil(o.IdentityAttributeConfig) {
		toSerialize["identityAttributeConfig"] = o.IdentityAttributeConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityPreviewRequest) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityPreviewRequest := _IdentityPreviewRequest{}

	if err = json.Unmarshal(bytes, &varIdentityPreviewRequest); err == nil {
		*o = IdentityPreviewRequest(varIdentityPreviewRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identityId")
		delete(additionalProperties, "identityAttributeConfig")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityPreviewRequest struct {
	value *IdentityPreviewRequest
	isSet bool
}

func (v NullableIdentityPreviewRequest) Get() *IdentityPreviewRequest {
	return v.value
}

func (v *NullableIdentityPreviewRequest) Set(val *IdentityPreviewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityPreviewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityPreviewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityPreviewRequest(val *IdentityPreviewRequest) *NullableIdentityPreviewRequest {
	return &NullableIdentityPreviewRequest{value: val, isSet: true}
}

func (v NullableIdentityPreviewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityPreviewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


