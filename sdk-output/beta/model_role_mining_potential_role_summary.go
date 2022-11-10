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

// RoleMiningPotentialRoleSummary struct for RoleMiningPotentialRoleSummary
type RoleMiningPotentialRoleSummary struct {
	PotentialRoleRef *RoleMiningPotentialRoleRef `json:"potentialRoleRef,omitempty"`
	// The number of identities in a potential role.
	IdentityCount *int32 `json:"identityCount,omitempty"`
	// The number of entitlements in a potential role.
	EntitlementCount *int32 `json:"entitlementCount,omitempty"`
	// The status for this identity group which can be \"REQUESTED\" or \"OBTAINED\"
	IdentityGroupStatus *string `json:"identityGroupStatus,omitempty"`
	ProvisionState *RoleMiningPotentialRoleProvisionState `json:"provisionState,omitempty"`
	// ID of the provisioned role in IIQ or IDN.  Null if this potential role has not been provisioned.
	RoleId *string `json:"roleId,omitempty"`
	// The density metric (0-100) of this potential role. Higher density values indicate higher similarity amongst the identities.
	Density *int32 `json:"density,omitempty"`
	// The freshness metric (0-100) of this potential role. Higher freshness values indicate this potential role is more distinctive compared to existing roles.
	Freshness *int32 `json:"freshness,omitempty"`
	// The quality metric (0-100) of this potential role. Higher quality values indicate this potential role has high density and freshness.
	Quality *int32 `json:"quality,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleMiningPotentialRoleSummary RoleMiningPotentialRoleSummary

// NewRoleMiningPotentialRoleSummary instantiates a new RoleMiningPotentialRoleSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMiningPotentialRoleSummary() *RoleMiningPotentialRoleSummary {
	this := RoleMiningPotentialRoleSummary{}
	return &this
}

// NewRoleMiningPotentialRoleSummaryWithDefaults instantiates a new RoleMiningPotentialRoleSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMiningPotentialRoleSummaryWithDefaults() *RoleMiningPotentialRoleSummary {
	this := RoleMiningPotentialRoleSummary{}
	return &this
}

// GetPotentialRoleRef returns the PotentialRoleRef field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleSummary) GetPotentialRoleRef() RoleMiningPotentialRoleRef {
	if o == nil || isNil(o.PotentialRoleRef) {
		var ret RoleMiningPotentialRoleRef
		return ret
	}
	return *o.PotentialRoleRef
}

// GetPotentialRoleRefOk returns a tuple with the PotentialRoleRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleSummary) GetPotentialRoleRefOk() (*RoleMiningPotentialRoleRef, bool) {
	if o == nil || isNil(o.PotentialRoleRef) {
		return nil, false
	}
	return o.PotentialRoleRef, true
}

// HasPotentialRoleRef returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleSummary) HasPotentialRoleRef() bool {
	if o != nil && !isNil(o.PotentialRoleRef) {
		return true
	}

	return false
}

// SetPotentialRoleRef gets a reference to the given RoleMiningPotentialRoleRef and assigns it to the PotentialRoleRef field.
func (o *RoleMiningPotentialRoleSummary) SetPotentialRoleRef(v RoleMiningPotentialRoleRef) {
	o.PotentialRoleRef = &v
}

// GetIdentityCount returns the IdentityCount field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleSummary) GetIdentityCount() int32 {
	if o == nil || isNil(o.IdentityCount) {
		var ret int32
		return ret
	}
	return *o.IdentityCount
}

// GetIdentityCountOk returns a tuple with the IdentityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleSummary) GetIdentityCountOk() (*int32, bool) {
	if o == nil || isNil(o.IdentityCount) {
		return nil, false
	}
	return o.IdentityCount, true
}

// HasIdentityCount returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleSummary) HasIdentityCount() bool {
	if o != nil && !isNil(o.IdentityCount) {
		return true
	}

	return false
}

// SetIdentityCount gets a reference to the given int32 and assigns it to the IdentityCount field.
func (o *RoleMiningPotentialRoleSummary) SetIdentityCount(v int32) {
	o.IdentityCount = &v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleSummary) GetEntitlementCount() int32 {
	if o == nil || isNil(o.EntitlementCount) {
		var ret int32
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleSummary) GetEntitlementCountOk() (*int32, bool) {
	if o == nil || isNil(o.EntitlementCount) {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleSummary) HasEntitlementCount() bool {
	if o != nil && !isNil(o.EntitlementCount) {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given int32 and assigns it to the EntitlementCount field.
func (o *RoleMiningPotentialRoleSummary) SetEntitlementCount(v int32) {
	o.EntitlementCount = &v
}

// GetIdentityGroupStatus returns the IdentityGroupStatus field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleSummary) GetIdentityGroupStatus() string {
	if o == nil || isNil(o.IdentityGroupStatus) {
		var ret string
		return ret
	}
	return *o.IdentityGroupStatus
}

// GetIdentityGroupStatusOk returns a tuple with the IdentityGroupStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleSummary) GetIdentityGroupStatusOk() (*string, bool) {
	if o == nil || isNil(o.IdentityGroupStatus) {
		return nil, false
	}
	return o.IdentityGroupStatus, true
}

// HasIdentityGroupStatus returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleSummary) HasIdentityGroupStatus() bool {
	if o != nil && !isNil(o.IdentityGroupStatus) {
		return true
	}

	return false
}

// SetIdentityGroupStatus gets a reference to the given string and assigns it to the IdentityGroupStatus field.
func (o *RoleMiningPotentialRoleSummary) SetIdentityGroupStatus(v string) {
	o.IdentityGroupStatus = &v
}

// GetProvisionState returns the ProvisionState field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleSummary) GetProvisionState() RoleMiningPotentialRoleProvisionState {
	if o == nil || isNil(o.ProvisionState) {
		var ret RoleMiningPotentialRoleProvisionState
		return ret
	}
	return *o.ProvisionState
}

// GetProvisionStateOk returns a tuple with the ProvisionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleSummary) GetProvisionStateOk() (*RoleMiningPotentialRoleProvisionState, bool) {
	if o == nil || isNil(o.ProvisionState) {
		return nil, false
	}
	return o.ProvisionState, true
}

// HasProvisionState returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleSummary) HasProvisionState() bool {
	if o != nil && !isNil(o.ProvisionState) {
		return true
	}

	return false
}

// SetProvisionState gets a reference to the given RoleMiningPotentialRoleProvisionState and assigns it to the ProvisionState field.
func (o *RoleMiningPotentialRoleSummary) SetProvisionState(v RoleMiningPotentialRoleProvisionState) {
	o.ProvisionState = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleSummary) GetRoleId() string {
	if o == nil || isNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleSummary) GetRoleIdOk() (*string, bool) {
	if o == nil || isNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleSummary) HasRoleId() bool {
	if o != nil && !isNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *RoleMiningPotentialRoleSummary) SetRoleId(v string) {
	o.RoleId = &v
}

// GetDensity returns the Density field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleSummary) GetDensity() int32 {
	if o == nil || isNil(o.Density) {
		var ret int32
		return ret
	}
	return *o.Density
}

// GetDensityOk returns a tuple with the Density field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleSummary) GetDensityOk() (*int32, bool) {
	if o == nil || isNil(o.Density) {
		return nil, false
	}
	return o.Density, true
}

// HasDensity returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleSummary) HasDensity() bool {
	if o != nil && !isNil(o.Density) {
		return true
	}

	return false
}

// SetDensity gets a reference to the given int32 and assigns it to the Density field.
func (o *RoleMiningPotentialRoleSummary) SetDensity(v int32) {
	o.Density = &v
}

// GetFreshness returns the Freshness field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleSummary) GetFreshness() int32 {
	if o == nil || isNil(o.Freshness) {
		var ret int32
		return ret
	}
	return *o.Freshness
}

// GetFreshnessOk returns a tuple with the Freshness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleSummary) GetFreshnessOk() (*int32, bool) {
	if o == nil || isNil(o.Freshness) {
		return nil, false
	}
	return o.Freshness, true
}

// HasFreshness returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleSummary) HasFreshness() bool {
	if o != nil && !isNil(o.Freshness) {
		return true
	}

	return false
}

// SetFreshness gets a reference to the given int32 and assigns it to the Freshness field.
func (o *RoleMiningPotentialRoleSummary) SetFreshness(v int32) {
	o.Freshness = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *RoleMiningPotentialRoleSummary) GetQuality() int32 {
	if o == nil || isNil(o.Quality) {
		var ret int32
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMiningPotentialRoleSummary) GetQualityOk() (*int32, bool) {
	if o == nil || isNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *RoleMiningPotentialRoleSummary) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given int32 and assigns it to the Quality field.
func (o *RoleMiningPotentialRoleSummary) SetQuality(v int32) {
	o.Quality = &v
}

func (o RoleMiningPotentialRoleSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PotentialRoleRef) {
		toSerialize["potentialRoleRef"] = o.PotentialRoleRef
	}
	if !isNil(o.IdentityCount) {
		toSerialize["identityCount"] = o.IdentityCount
	}
	if !isNil(o.EntitlementCount) {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}
	if !isNil(o.IdentityGroupStatus) {
		toSerialize["identityGroupStatus"] = o.IdentityGroupStatus
	}
	if !isNil(o.ProvisionState) {
		toSerialize["provisionState"] = o.ProvisionState
	}
	if !isNil(o.RoleId) {
		toSerialize["roleId"] = o.RoleId
	}
	if !isNil(o.Density) {
		toSerialize["density"] = o.Density
	}
	if !isNil(o.Freshness) {
		toSerialize["freshness"] = o.Freshness
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleMiningPotentialRoleSummary) UnmarshalJSON(bytes []byte) (err error) {
	varRoleMiningPotentialRoleSummary := _RoleMiningPotentialRoleSummary{}

	if err = json.Unmarshal(bytes, &varRoleMiningPotentialRoleSummary); err == nil {
		*o = RoleMiningPotentialRoleSummary(varRoleMiningPotentialRoleSummary)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "potentialRoleRef")
		delete(additionalProperties, "identityCount")
		delete(additionalProperties, "entitlementCount")
		delete(additionalProperties, "identityGroupStatus")
		delete(additionalProperties, "provisionState")
		delete(additionalProperties, "roleId")
		delete(additionalProperties, "density")
		delete(additionalProperties, "freshness")
		delete(additionalProperties, "quality")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleMiningPotentialRoleSummary struct {
	value *RoleMiningPotentialRoleSummary
	isSet bool
}

func (v NullableRoleMiningPotentialRoleSummary) Get() *RoleMiningPotentialRoleSummary {
	return v.value
}

func (v *NullableRoleMiningPotentialRoleSummary) Set(val *RoleMiningPotentialRoleSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningPotentialRoleSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningPotentialRoleSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningPotentialRoleSummary(val *RoleMiningPotentialRoleSummary) *NullableRoleMiningPotentialRoleSummary {
	return &NullableRoleMiningPotentialRoleSummary{value: val, isSet: true}
}

func (v NullableRoleMiningPotentialRoleSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningPotentialRoleSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


