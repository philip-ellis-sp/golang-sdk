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

// EntitlementRequestConfig struct for EntitlementRequestConfig
type EntitlementRequestConfig struct {
	// Flag for allowing entitlement request.
	AllowEntitlementRequest *bool `json:"allowEntitlementRequest,omitempty"`
	// Flag for requiring comments while submitting an entitlement request.
	RequestCommentsRequired *bool `json:"requestCommentsRequired,omitempty"`
	// Flag for requiring comments while rejecting an entitlement request.
	DeniedCommentsRequired *bool `json:"deniedCommentsRequired,omitempty"`
	// Approval schemes for granting entitlement request. This can be empty if no approval is needed. Multiple schemes must be comma-separated. The valid schemes are \"sourceOwner\", \"manager\" and \"workgroup:{id}\". Multiple workgroups (governance groups) can be used. 
	GrantRequestApprovalSchemes *string `json:"grantRequestApprovalSchemes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementRequestConfig EntitlementRequestConfig

// NewEntitlementRequestConfig instantiates a new EntitlementRequestConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementRequestConfig() *EntitlementRequestConfig {
	this := EntitlementRequestConfig{}
	var requestCommentsRequired bool = false
	this.RequestCommentsRequired = &requestCommentsRequired
	var deniedCommentsRequired bool = false
	this.DeniedCommentsRequired = &deniedCommentsRequired
	var grantRequestApprovalSchemes string = "sourceOwner"
	this.GrantRequestApprovalSchemes = &grantRequestApprovalSchemes
	return &this
}

// NewEntitlementRequestConfigWithDefaults instantiates a new EntitlementRequestConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementRequestConfigWithDefaults() *EntitlementRequestConfig {
	this := EntitlementRequestConfig{}
	var requestCommentsRequired bool = false
	this.RequestCommentsRequired = &requestCommentsRequired
	var deniedCommentsRequired bool = false
	this.DeniedCommentsRequired = &deniedCommentsRequired
	var grantRequestApprovalSchemes string = "sourceOwner"
	this.GrantRequestApprovalSchemes = &grantRequestApprovalSchemes
	return &this
}

// GetAllowEntitlementRequest returns the AllowEntitlementRequest field value if set, zero value otherwise.
func (o *EntitlementRequestConfig) GetAllowEntitlementRequest() bool {
	if o == nil || isNil(o.AllowEntitlementRequest) {
		var ret bool
		return ret
	}
	return *o.AllowEntitlementRequest
}

// GetAllowEntitlementRequestOk returns a tuple with the AllowEntitlementRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementRequestConfig) GetAllowEntitlementRequestOk() (*bool, bool) {
	if o == nil || isNil(o.AllowEntitlementRequest) {
		return nil, false
	}
	return o.AllowEntitlementRequest, true
}

// HasAllowEntitlementRequest returns a boolean if a field has been set.
func (o *EntitlementRequestConfig) HasAllowEntitlementRequest() bool {
	if o != nil && !isNil(o.AllowEntitlementRequest) {
		return true
	}

	return false
}

// SetAllowEntitlementRequest gets a reference to the given bool and assigns it to the AllowEntitlementRequest field.
func (o *EntitlementRequestConfig) SetAllowEntitlementRequest(v bool) {
	o.AllowEntitlementRequest = &v
}

// GetRequestCommentsRequired returns the RequestCommentsRequired field value if set, zero value otherwise.
func (o *EntitlementRequestConfig) GetRequestCommentsRequired() bool {
	if o == nil || isNil(o.RequestCommentsRequired) {
		var ret bool
		return ret
	}
	return *o.RequestCommentsRequired
}

// GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementRequestConfig) GetRequestCommentsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.RequestCommentsRequired) {
		return nil, false
	}
	return o.RequestCommentsRequired, true
}

// HasRequestCommentsRequired returns a boolean if a field has been set.
func (o *EntitlementRequestConfig) HasRequestCommentsRequired() bool {
	if o != nil && !isNil(o.RequestCommentsRequired) {
		return true
	}

	return false
}

// SetRequestCommentsRequired gets a reference to the given bool and assigns it to the RequestCommentsRequired field.
func (o *EntitlementRequestConfig) SetRequestCommentsRequired(v bool) {
	o.RequestCommentsRequired = &v
}

// GetDeniedCommentsRequired returns the DeniedCommentsRequired field value if set, zero value otherwise.
func (o *EntitlementRequestConfig) GetDeniedCommentsRequired() bool {
	if o == nil || isNil(o.DeniedCommentsRequired) {
		var ret bool
		return ret
	}
	return *o.DeniedCommentsRequired
}

// GetDeniedCommentsRequiredOk returns a tuple with the DeniedCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementRequestConfig) GetDeniedCommentsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.DeniedCommentsRequired) {
		return nil, false
	}
	return o.DeniedCommentsRequired, true
}

// HasDeniedCommentsRequired returns a boolean if a field has been set.
func (o *EntitlementRequestConfig) HasDeniedCommentsRequired() bool {
	if o != nil && !isNil(o.DeniedCommentsRequired) {
		return true
	}

	return false
}

// SetDeniedCommentsRequired gets a reference to the given bool and assigns it to the DeniedCommentsRequired field.
func (o *EntitlementRequestConfig) SetDeniedCommentsRequired(v bool) {
	o.DeniedCommentsRequired = &v
}

// GetGrantRequestApprovalSchemes returns the GrantRequestApprovalSchemes field value if set, zero value otherwise.
func (o *EntitlementRequestConfig) GetGrantRequestApprovalSchemes() string {
	if o == nil || isNil(o.GrantRequestApprovalSchemes) {
		var ret string
		return ret
	}
	return *o.GrantRequestApprovalSchemes
}

// GetGrantRequestApprovalSchemesOk returns a tuple with the GrantRequestApprovalSchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementRequestConfig) GetGrantRequestApprovalSchemesOk() (*string, bool) {
	if o == nil || isNil(o.GrantRequestApprovalSchemes) {
		return nil, false
	}
	return o.GrantRequestApprovalSchemes, true
}

// HasGrantRequestApprovalSchemes returns a boolean if a field has been set.
func (o *EntitlementRequestConfig) HasGrantRequestApprovalSchemes() bool {
	if o != nil && !isNil(o.GrantRequestApprovalSchemes) {
		return true
	}

	return false
}

// SetGrantRequestApprovalSchemes gets a reference to the given string and assigns it to the GrantRequestApprovalSchemes field.
func (o *EntitlementRequestConfig) SetGrantRequestApprovalSchemes(v string) {
	o.GrantRequestApprovalSchemes = &v
}

func (o EntitlementRequestConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllowEntitlementRequest) {
		toSerialize["allowEntitlementRequest"] = o.AllowEntitlementRequest
	}
	if !isNil(o.RequestCommentsRequired) {
		toSerialize["requestCommentsRequired"] = o.RequestCommentsRequired
	}
	if !isNil(o.DeniedCommentsRequired) {
		toSerialize["deniedCommentsRequired"] = o.DeniedCommentsRequired
	}
	if !isNil(o.GrantRequestApprovalSchemes) {
		toSerialize["grantRequestApprovalSchemes"] = o.GrantRequestApprovalSchemes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementRequestConfig) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementRequestConfig := _EntitlementRequestConfig{}

	if err = json.Unmarshal(bytes, &varEntitlementRequestConfig); err == nil {
		*o = EntitlementRequestConfig(varEntitlementRequestConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "allowEntitlementRequest")
		delete(additionalProperties, "requestCommentsRequired")
		delete(additionalProperties, "deniedCommentsRequired")
		delete(additionalProperties, "grantRequestApprovalSchemes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementRequestConfig struct {
	value *EntitlementRequestConfig
	isSet bool
}

func (v NullableEntitlementRequestConfig) Get() *EntitlementRequestConfig {
	return v.value
}

func (v *NullableEntitlementRequestConfig) Set(val *EntitlementRequestConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementRequestConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementRequestConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementRequestConfig(val *EntitlementRequestConfig) *NullableEntitlementRequestConfig {
	return &NullableEntitlementRequestConfig{value: val, isSet: true}
}

func (v NullableEntitlementRequestConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementRequestConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


