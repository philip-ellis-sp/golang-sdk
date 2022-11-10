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

// RoleInsightsResponse struct for RoleInsightsResponse
type RoleInsightsResponse struct {
	// Request Id for a role insight generation request
	Id *string `json:"id,omitempty"`
	// The date-time role insights request was created.
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// The date-time role insights request was completed.
	LastGenerated *time.Time `json:"lastGenerated,omitempty"`
	// Total number of updates for this request. Starts with 0 and will have correct number when request is COMPLETED.
	NumberOfUpdates *int32 `json:"numberOfUpdates,omitempty"`
	// The role IDs that are in this request.
	RoleIds []string `json:"roleIds,omitempty"`
	// Request status
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleInsightsResponse RoleInsightsResponse

// NewRoleInsightsResponse instantiates a new RoleInsightsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleInsightsResponse() *RoleInsightsResponse {
	this := RoleInsightsResponse{}
	return &this
}

// NewRoleInsightsResponseWithDefaults instantiates a new RoleInsightsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleInsightsResponseWithDefaults() *RoleInsightsResponse {
	this := RoleInsightsResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleInsightsResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleInsightsResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleInsightsResponse) SetId(v string) {
	o.Id = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *RoleInsightsResponse) GetCreatedDate() time.Time {
	if o == nil || isNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsResponse) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *RoleInsightsResponse) HasCreatedDate() bool {
	if o != nil && !isNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *RoleInsightsResponse) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetLastGenerated returns the LastGenerated field value if set, zero value otherwise.
func (o *RoleInsightsResponse) GetLastGenerated() time.Time {
	if o == nil || isNil(o.LastGenerated) {
		var ret time.Time
		return ret
	}
	return *o.LastGenerated
}

// GetLastGeneratedOk returns a tuple with the LastGenerated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsResponse) GetLastGeneratedOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastGenerated) {
		return nil, false
	}
	return o.LastGenerated, true
}

// HasLastGenerated returns a boolean if a field has been set.
func (o *RoleInsightsResponse) HasLastGenerated() bool {
	if o != nil && !isNil(o.LastGenerated) {
		return true
	}

	return false
}

// SetLastGenerated gets a reference to the given time.Time and assigns it to the LastGenerated field.
func (o *RoleInsightsResponse) SetLastGenerated(v time.Time) {
	o.LastGenerated = &v
}

// GetNumberOfUpdates returns the NumberOfUpdates field value if set, zero value otherwise.
func (o *RoleInsightsResponse) GetNumberOfUpdates() int32 {
	if o == nil || isNil(o.NumberOfUpdates) {
		var ret int32
		return ret
	}
	return *o.NumberOfUpdates
}

// GetNumberOfUpdatesOk returns a tuple with the NumberOfUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsResponse) GetNumberOfUpdatesOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfUpdates) {
		return nil, false
	}
	return o.NumberOfUpdates, true
}

// HasNumberOfUpdates returns a boolean if a field has been set.
func (o *RoleInsightsResponse) HasNumberOfUpdates() bool {
	if o != nil && !isNil(o.NumberOfUpdates) {
		return true
	}

	return false
}

// SetNumberOfUpdates gets a reference to the given int32 and assigns it to the NumberOfUpdates field.
func (o *RoleInsightsResponse) SetNumberOfUpdates(v int32) {
	o.NumberOfUpdates = &v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *RoleInsightsResponse) GetRoleIds() []string {
	if o == nil || isNil(o.RoleIds) {
		var ret []string
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsResponse) GetRoleIdsOk() ([]string, bool) {
	if o == nil || isNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *RoleInsightsResponse) HasRoleIds() bool {
	if o != nil && !isNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *RoleInsightsResponse) SetRoleIds(v []string) {
	o.RoleIds = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RoleInsightsResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleInsightsResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RoleInsightsResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RoleInsightsResponse) SetStatus(v string) {
	o.Status = &v
}

func (o RoleInsightsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !isNil(o.LastGenerated) {
		toSerialize["lastGenerated"] = o.LastGenerated
	}
	if !isNil(o.NumberOfUpdates) {
		toSerialize["numberOfUpdates"] = o.NumberOfUpdates
	}
	if !isNil(o.RoleIds) {
		toSerialize["roleIds"] = o.RoleIds
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleInsightsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varRoleInsightsResponse := _RoleInsightsResponse{}

	if err = json.Unmarshal(bytes, &varRoleInsightsResponse); err == nil {
		*o = RoleInsightsResponse(varRoleInsightsResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "lastGenerated")
		delete(additionalProperties, "numberOfUpdates")
		delete(additionalProperties, "roleIds")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleInsightsResponse struct {
	value *RoleInsightsResponse
	isSet bool
}

func (v NullableRoleInsightsResponse) Get() *RoleInsightsResponse {
	return v.value
}

func (v *NullableRoleInsightsResponse) Set(val *RoleInsightsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleInsightsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleInsightsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleInsightsResponse(val *RoleInsightsResponse) *NullableRoleInsightsResponse {
	return &NullableRoleInsightsResponse{value: val, isSet: true}
}

func (v NullableRoleInsightsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleInsightsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


