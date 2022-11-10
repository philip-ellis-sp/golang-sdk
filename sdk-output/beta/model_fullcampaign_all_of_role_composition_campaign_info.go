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

// FullcampaignAllOfRoleCompositionCampaignInfo Optional configuration options for role composition campaigns.
type FullcampaignAllOfRoleCompositionCampaignInfo struct {
	// If specified, this identity or governance group will be the reviewer for all certifications in this campaign. The allowed DTO types are IDENTITY and GOVERNANCE_GROUP
	Reviewer *BaseReferenceDto `json:"reviewer,omitempty"`
	// Optional list of roles to include in this campaign. Only one of `roleIds` and `query` may be set; if neither are set, all roles are included.
	RoleIds []string `json:"roleIds,omitempty"`
	RemediatorRef FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef `json:"remediatorRef"`
	// Optional search query to scope this campaign to a set of roles. Only one of `roleIds` and `query` may be set; if neither are set, all roles are included.
	Query *string `json:"query,omitempty"`
	// Describes this role composition campaign. Intended for storing the query used, and possibly the number of roles selected/available.
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FullcampaignAllOfRoleCompositionCampaignInfo FullcampaignAllOfRoleCompositionCampaignInfo

// NewFullcampaignAllOfRoleCompositionCampaignInfo instantiates a new FullcampaignAllOfRoleCompositionCampaignInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullcampaignAllOfRoleCompositionCampaignInfo(remediatorRef FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) *FullcampaignAllOfRoleCompositionCampaignInfo {
	this := FullcampaignAllOfRoleCompositionCampaignInfo{}
	this.RemediatorRef = remediatorRef
	return &this
}

// NewFullcampaignAllOfRoleCompositionCampaignInfoWithDefaults instantiates a new FullcampaignAllOfRoleCompositionCampaignInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullcampaignAllOfRoleCompositionCampaignInfoWithDefaults() *FullcampaignAllOfRoleCompositionCampaignInfo {
	this := FullcampaignAllOfRoleCompositionCampaignInfo{}
	return &this
}

// GetReviewer returns the Reviewer field value if set, zero value otherwise.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) GetReviewer() BaseReferenceDto {
	if o == nil || isNil(o.Reviewer) {
		var ret BaseReferenceDto
		return ret
	}
	return *o.Reviewer
}

// GetReviewerOk returns a tuple with the Reviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) GetReviewerOk() (*BaseReferenceDto, bool) {
	if o == nil || isNil(o.Reviewer) {
		return nil, false
	}
	return o.Reviewer, true
}

// HasReviewer returns a boolean if a field has been set.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) HasReviewer() bool {
	if o != nil && !isNil(o.Reviewer) {
		return true
	}

	return false
}

// SetReviewer gets a reference to the given BaseReferenceDto and assigns it to the Reviewer field.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) SetReviewer(v BaseReferenceDto) {
	o.Reviewer = &v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) GetRoleIds() []string {
	if o == nil || isNil(o.RoleIds) {
		var ret []string
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) GetRoleIdsOk() ([]string, bool) {
	if o == nil || isNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) HasRoleIds() bool {
	if o != nil && !isNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) SetRoleIds(v []string) {
	o.RoleIds = v
}

// GetRemediatorRef returns the RemediatorRef field value
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) GetRemediatorRef() FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef {
	if o == nil {
		var ret FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef
		return ret
	}

	return o.RemediatorRef
}

// GetRemediatorRefOk returns a tuple with the RemediatorRef field value
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) GetRemediatorRefOk() (*FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediatorRef, true
}

// SetRemediatorRef sets field value
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) SetRemediatorRef(v FullcampaignAllOfRoleCompositionCampaignInfoRemediatorRef) {
	o.RemediatorRef = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) GetQuery() string {
	if o == nil || isNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) GetQueryOk() (*string, bool) {
	if o == nil || isNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) SetQuery(v string) {
	o.Query = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FullcampaignAllOfRoleCompositionCampaignInfo) SetDescription(v string) {
	o.Description = &v
}

func (o FullcampaignAllOfRoleCompositionCampaignInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Reviewer) {
		toSerialize["reviewer"] = o.Reviewer
	}
	if !isNil(o.RoleIds) {
		toSerialize["roleIds"] = o.RoleIds
	}
	if true {
		toSerialize["remediatorRef"] = o.RemediatorRef
	}
	if !isNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FullcampaignAllOfRoleCompositionCampaignInfo) UnmarshalJSON(bytes []byte) (err error) {
	varFullcampaignAllOfRoleCompositionCampaignInfo := _FullcampaignAllOfRoleCompositionCampaignInfo{}

	if err = json.Unmarshal(bytes, &varFullcampaignAllOfRoleCompositionCampaignInfo); err == nil {
		*o = FullcampaignAllOfRoleCompositionCampaignInfo(varFullcampaignAllOfRoleCompositionCampaignInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reviewer")
		delete(additionalProperties, "roleIds")
		delete(additionalProperties, "remediatorRef")
		delete(additionalProperties, "query")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFullcampaignAllOfRoleCompositionCampaignInfo struct {
	value *FullcampaignAllOfRoleCompositionCampaignInfo
	isSet bool
}

func (v NullableFullcampaignAllOfRoleCompositionCampaignInfo) Get() *FullcampaignAllOfRoleCompositionCampaignInfo {
	return v.value
}

func (v *NullableFullcampaignAllOfRoleCompositionCampaignInfo) Set(val *FullcampaignAllOfRoleCompositionCampaignInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFullcampaignAllOfRoleCompositionCampaignInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFullcampaignAllOfRoleCompositionCampaignInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullcampaignAllOfRoleCompositionCampaignInfo(val *FullcampaignAllOfRoleCompositionCampaignInfo) *NullableFullcampaignAllOfRoleCompositionCampaignInfo {
	return &NullableFullcampaignAllOfRoleCompositionCampaignInfo{value: val, isSet: true}
}

func (v NullableFullcampaignAllOfRoleCompositionCampaignInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullcampaignAllOfRoleCompositionCampaignInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


