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

// AccessRequestItem struct for AccessRequestItem
type AccessRequestItem struct {
	// The type of the item being requested.
	Type string `json:"type"`
	// ID of Role, Access Profile or Entitlement being requested.
	Id string `json:"id"`
	// Comment provided by requester. * Comment is required when the request is of type Revoke Access. 
	Comment *string `json:"comment,omitempty"`
	// Arbitrary key-value pairs. They will never be processed by the IdentityNow system but will be returned on associated APIs such as /account-activities and /access-request-status.
	ClientMetadata *map[string]string `json:"clientMetadata,omitempty"`
	// The date the role or access profile is no longer assigned to the specified identity. * Specify a date in the future. * The current SLA for the deprovisioning is 24 hours. * This date can be modified to either extend or decrease the duration of access item assignments for the specified identity. * Currently it is not supported for entitlements. 
	RemoveDate *time.Time `json:"removeDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequestItem AccessRequestItem

// NewAccessRequestItem instantiates a new AccessRequestItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestItem(type_ string, id string) *AccessRequestItem {
	this := AccessRequestItem{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAccessRequestItemWithDefaults instantiates a new AccessRequestItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestItemWithDefaults() *AccessRequestItem {
	this := AccessRequestItem{}
	return &this
}

// GetType returns the Type field value
func (o *AccessRequestItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessRequestItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccessRequestItem) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AccessRequestItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessRequestItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessRequestItem) SetId(v string) {
	o.Id = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *AccessRequestItem) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItem) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *AccessRequestItem) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *AccessRequestItem) SetComment(v string) {
	o.Comment = &v
}

// GetClientMetadata returns the ClientMetadata field value if set, zero value otherwise.
func (o *AccessRequestItem) GetClientMetadata() map[string]string {
	if o == nil || isNil(o.ClientMetadata) {
		var ret map[string]string
		return ret
	}
	return *o.ClientMetadata
}

// GetClientMetadataOk returns a tuple with the ClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItem) GetClientMetadataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ClientMetadata) {
		return nil, false
	}
	return o.ClientMetadata, true
}

// HasClientMetadata returns a boolean if a field has been set.
func (o *AccessRequestItem) HasClientMetadata() bool {
	if o != nil && !isNil(o.ClientMetadata) {
		return true
	}

	return false
}

// SetClientMetadata gets a reference to the given map[string]string and assigns it to the ClientMetadata field.
func (o *AccessRequestItem) SetClientMetadata(v map[string]string) {
	o.ClientMetadata = &v
}

// GetRemoveDate returns the RemoveDate field value if set, zero value otherwise.
func (o *AccessRequestItem) GetRemoveDate() time.Time {
	if o == nil || isNil(o.RemoveDate) {
		var ret time.Time
		return ret
	}
	return *o.RemoveDate
}

// GetRemoveDateOk returns a tuple with the RemoveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestItem) GetRemoveDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.RemoveDate) {
		return nil, false
	}
	return o.RemoveDate, true
}

// HasRemoveDate returns a boolean if a field has been set.
func (o *AccessRequestItem) HasRemoveDate() bool {
	if o != nil && !isNil(o.RemoveDate) {
		return true
	}

	return false
}

// SetRemoveDate gets a reference to the given time.Time and assigns it to the RemoveDate field.
func (o *AccessRequestItem) SetRemoveDate(v time.Time) {
	o.RemoveDate = &v
}

func (o AccessRequestItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.ClientMetadata) {
		toSerialize["clientMetadata"] = o.ClientMetadata
	}
	if !isNil(o.RemoveDate) {
		toSerialize["removeDate"] = o.RemoveDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessRequestItem) UnmarshalJSON(bytes []byte) (err error) {
	varAccessRequestItem := _AccessRequestItem{}

	if err = json.Unmarshal(bytes, &varAccessRequestItem); err == nil {
		*o = AccessRequestItem(varAccessRequestItem)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "clientMetadata")
		delete(additionalProperties, "removeDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequestItem struct {
	value *AccessRequestItem
	isSet bool
}

func (v NullableAccessRequestItem) Get() *AccessRequestItem {
	return v.value
}

func (v *NullableAccessRequestItem) Set(val *AccessRequestItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestItem(val *AccessRequestItem) *NullableAccessRequestItem {
	return &NullableAccessRequestItem{value: val, isSet: true}
}

func (v NullableAccessRequestItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


