/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointbetasdk

import (
	"encoding/json"
	"time"
)

// ApprovalForwardHistory struct for ApprovalForwardHistory
type ApprovalForwardHistory struct {
	// Display name of approver that forwarded the approval.
	OldApproverName *string `json:"oldApproverName,omitempty"`
	// Display name of approver to whom the approval was forwarded.
	NewApproverName *string `json:"newApproverName,omitempty"`
	// Comment made by old approver when forwarding.
	Comment *string `json:"comment,omitempty"`
	// Time at which approval was forwarded.
	Modified *time.Time `json:"modified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalForwardHistory ApprovalForwardHistory

// NewApprovalForwardHistory instantiates a new ApprovalForwardHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalForwardHistory() *ApprovalForwardHistory {
	this := ApprovalForwardHistory{}
	return &this
}

// NewApprovalForwardHistoryWithDefaults instantiates a new ApprovalForwardHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalForwardHistoryWithDefaults() *ApprovalForwardHistory {
	this := ApprovalForwardHistory{}
	return &this
}

// GetOldApproverName returns the OldApproverName field value if set, zero value otherwise.
func (o *ApprovalForwardHistory) GetOldApproverName() string {
	if o == nil || isNil(o.OldApproverName) {
		var ret string
		return ret
	}
	return *o.OldApproverName
}

// GetOldApproverNameOk returns a tuple with the OldApproverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalForwardHistory) GetOldApproverNameOk() (*string, bool) {
	if o == nil || isNil(o.OldApproverName) {
		return nil, false
	}
	return o.OldApproverName, true
}

// HasOldApproverName returns a boolean if a field has been set.
func (o *ApprovalForwardHistory) HasOldApproverName() bool {
	if o != nil && !isNil(o.OldApproverName) {
		return true
	}

	return false
}

// SetOldApproverName gets a reference to the given string and assigns it to the OldApproverName field.
func (o *ApprovalForwardHistory) SetOldApproverName(v string) {
	o.OldApproverName = &v
}

// GetNewApproverName returns the NewApproverName field value if set, zero value otherwise.
func (o *ApprovalForwardHistory) GetNewApproverName() string {
	if o == nil || isNil(o.NewApproverName) {
		var ret string
		return ret
	}
	return *o.NewApproverName
}

// GetNewApproverNameOk returns a tuple with the NewApproverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalForwardHistory) GetNewApproverNameOk() (*string, bool) {
	if o == nil || isNil(o.NewApproverName) {
		return nil, false
	}
	return o.NewApproverName, true
}

// HasNewApproverName returns a boolean if a field has been set.
func (o *ApprovalForwardHistory) HasNewApproverName() bool {
	if o != nil && !isNil(o.NewApproverName) {
		return true
	}

	return false
}

// SetNewApproverName gets a reference to the given string and assigns it to the NewApproverName field.
func (o *ApprovalForwardHistory) SetNewApproverName(v string) {
	o.NewApproverName = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApprovalForwardHistory) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalForwardHistory) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApprovalForwardHistory) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApprovalForwardHistory) SetComment(v string) {
	o.Comment = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *ApprovalForwardHistory) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalForwardHistory) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *ApprovalForwardHistory) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *ApprovalForwardHistory) SetModified(v time.Time) {
	o.Modified = &v
}

func (o ApprovalForwardHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OldApproverName) {
		toSerialize["oldApproverName"] = o.OldApproverName
	}
	if !isNil(o.NewApproverName) {
		toSerialize["newApproverName"] = o.NewApproverName
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApprovalForwardHistory) UnmarshalJSON(bytes []byte) (err error) {
	varApprovalForwardHistory := _ApprovalForwardHistory{}

	if err = json.Unmarshal(bytes, &varApprovalForwardHistory); err == nil {
		*o = ApprovalForwardHistory(varApprovalForwardHistory)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "oldApproverName")
		delete(additionalProperties, "newApproverName")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "modified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalForwardHistory struct {
	value *ApprovalForwardHistory
	isSet bool
}

func (v NullableApprovalForwardHistory) Get() *ApprovalForwardHistory {
	return v.value
}

func (v *NullableApprovalForwardHistory) Set(val *ApprovalForwardHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalForwardHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalForwardHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalForwardHistory(val *ApprovalForwardHistory) *NullableApprovalForwardHistory {
	return &NullableApprovalForwardHistory{value: val, isSet: true}
}

func (v NullableApprovalForwardHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalForwardHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


