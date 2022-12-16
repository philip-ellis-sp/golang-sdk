/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"time"
	"encoding/json"
)

// AccessProfileDocument This is more of a complete representation of an access profile.  
type AccessProfileDocument struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type DocumentType `json:"_type"`
	// The description of the access item
	Description *string `json:"description,omitempty"`
	// A date-time in ISO-8601 format
	Created NullableTime `json:"created,omitempty"`
	// A date-time in ISO-8601 format
	Modified NullableTime `json:"modified,omitempty"`
	// A date-time in ISO-8601 format
	Synced NullableTime `json:"synced,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// Indicates if the access can be requested
	Requestable *bool `json:"requestable,omitempty"`
	// Indicates if comments are required when requesting access
	RequestCommentsRequired *bool `json:"requestCommentsRequired,omitempty"`
	Owner *Owner `json:"owner,omitempty"`
	Source *Reference `json:"source,omitempty"`
	Entitlements []BaseEntitlement `json:"entitlements,omitempty"`
	EntitlementCount *int32 `json:"entitlementCount,omitempty"`
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessProfileDocument AccessProfileDocument

// NewAccessProfileDocument instantiates a new AccessProfileDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessProfileDocument(id string, name string, type_ DocumentType) *AccessProfileDocument {
	this := AccessProfileDocument{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewAccessProfileDocumentWithDefaults instantiates a new AccessProfileDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessProfileDocumentWithDefaults() *AccessProfileDocument {
	this := AccessProfileDocument{}
	return &this
}

// GetId returns the Id field value
func (o *AccessProfileDocument) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessProfileDocument) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessProfileDocument) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AccessProfileDocument) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccessProfileDocument) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccessProfileDocument) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *AccessProfileDocument) GetType() DocumentType {
	if o == nil {
		var ret DocumentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessProfileDocument) GetTypeOk() (*DocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccessProfileDocument) SetType(v DocumentType) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccessProfileDocument) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDocument) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessProfileDocument) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccessProfileDocument) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfileDocument) GetCreated() time.Time {
	if o == nil || isNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfileDocument) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *AccessProfileDocument) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *AccessProfileDocument) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *AccessProfileDocument) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *AccessProfileDocument) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfileDocument) GetModified() time.Time {
	if o == nil || isNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfileDocument) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *AccessProfileDocument) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *AccessProfileDocument) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *AccessProfileDocument) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *AccessProfileDocument) UnsetModified() {
	o.Modified.Unset()
}

// GetSynced returns the Synced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessProfileDocument) GetSynced() time.Time {
	if o == nil || isNil(o.Synced.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Synced.Get()
}

// GetSyncedOk returns a tuple with the Synced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessProfileDocument) GetSyncedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Synced.Get(), o.Synced.IsSet()
}

// HasSynced returns a boolean if a field has been set.
func (o *AccessProfileDocument) HasSynced() bool {
	if o != nil && o.Synced.IsSet() {
		return true
	}

	return false
}

// SetSynced gets a reference to the given NullableTime and assigns it to the Synced field.
func (o *AccessProfileDocument) SetSynced(v time.Time) {
	o.Synced.Set(&v)
}
// SetSyncedNil sets the value for Synced to be an explicit nil
func (o *AccessProfileDocument) SetSyncedNil() {
	o.Synced.Set(nil)
}

// UnsetSynced ensures that no value is present for Synced, not even an explicit nil
func (o *AccessProfileDocument) UnsetSynced() {
	o.Synced.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AccessProfileDocument) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDocument) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AccessProfileDocument) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AccessProfileDocument) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRequestable returns the Requestable field value if set, zero value otherwise.
func (o *AccessProfileDocument) GetRequestable() bool {
	if o == nil || isNil(o.Requestable) {
		var ret bool
		return ret
	}
	return *o.Requestable
}

// GetRequestableOk returns a tuple with the Requestable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDocument) GetRequestableOk() (*bool, bool) {
	if o == nil || isNil(o.Requestable) {
		return nil, false
	}
	return o.Requestable, true
}

// HasRequestable returns a boolean if a field has been set.
func (o *AccessProfileDocument) HasRequestable() bool {
	if o != nil && !isNil(o.Requestable) {
		return true
	}

	return false
}

// SetRequestable gets a reference to the given bool and assigns it to the Requestable field.
func (o *AccessProfileDocument) SetRequestable(v bool) {
	o.Requestable = &v
}

// GetRequestCommentsRequired returns the RequestCommentsRequired field value if set, zero value otherwise.
func (o *AccessProfileDocument) GetRequestCommentsRequired() bool {
	if o == nil || isNil(o.RequestCommentsRequired) {
		var ret bool
		return ret
	}
	return *o.RequestCommentsRequired
}

// GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDocument) GetRequestCommentsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.RequestCommentsRequired) {
		return nil, false
	}
	return o.RequestCommentsRequired, true
}

// HasRequestCommentsRequired returns a boolean if a field has been set.
func (o *AccessProfileDocument) HasRequestCommentsRequired() bool {
	if o != nil && !isNil(o.RequestCommentsRequired) {
		return true
	}

	return false
}

// SetRequestCommentsRequired gets a reference to the given bool and assigns it to the RequestCommentsRequired field.
func (o *AccessProfileDocument) SetRequestCommentsRequired(v bool) {
	o.RequestCommentsRequired = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *AccessProfileDocument) GetOwner() Owner {
	if o == nil || isNil(o.Owner) {
		var ret Owner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDocument) GetOwnerOk() (*Owner, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *AccessProfileDocument) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Owner and assigns it to the Owner field.
func (o *AccessProfileDocument) SetOwner(v Owner) {
	o.Owner = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AccessProfileDocument) GetSource() Reference {
	if o == nil || isNil(o.Source) {
		var ret Reference
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDocument) GetSourceOk() (*Reference, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AccessProfileDocument) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Reference and assigns it to the Source field.
func (o *AccessProfileDocument) SetSource(v Reference) {
	o.Source = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *AccessProfileDocument) GetEntitlements() []BaseEntitlement {
	if o == nil || isNil(o.Entitlements) {
		var ret []BaseEntitlement
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDocument) GetEntitlementsOk() ([]BaseEntitlement, bool) {
	if o == nil || isNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *AccessProfileDocument) HasEntitlements() bool {
	if o != nil && !isNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []BaseEntitlement and assigns it to the Entitlements field.
func (o *AccessProfileDocument) SetEntitlements(v []BaseEntitlement) {
	o.Entitlements = v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *AccessProfileDocument) GetEntitlementCount() int32 {
	if o == nil || isNil(o.EntitlementCount) {
		var ret int32
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDocument) GetEntitlementCountOk() (*int32, bool) {
	if o == nil || isNil(o.EntitlementCount) {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *AccessProfileDocument) HasEntitlementCount() bool {
	if o != nil && !isNil(o.EntitlementCount) {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given int32 and assigns it to the EntitlementCount field.
func (o *AccessProfileDocument) SetEntitlementCount(v int32) {
	o.EntitlementCount = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AccessProfileDocument) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessProfileDocument) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AccessProfileDocument) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AccessProfileDocument) SetTags(v []string) {
	o.Tags = v
}

func (o AccessProfileDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["_type"] = o.Type
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if o.Synced.IsSet() {
		toSerialize["synced"] = o.Synced.Get()
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Requestable) {
		toSerialize["requestable"] = o.Requestable
	}
	if !isNil(o.RequestCommentsRequired) {
		toSerialize["requestCommentsRequired"] = o.RequestCommentsRequired
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !isNil(o.EntitlementCount) {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessProfileDocument) UnmarshalJSON(bytes []byte) (err error) {
	varAccessProfileDocument := _AccessProfileDocument{}

	if err = json.Unmarshal(bytes, &varAccessProfileDocument); err == nil {
		*o = AccessProfileDocument(varAccessProfileDocument)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "synced")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "requestable")
		delete(additionalProperties, "requestCommentsRequired")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "source")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "entitlementCount")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessProfileDocument struct {
	value *AccessProfileDocument
	isSet bool
}

func (v NullableAccessProfileDocument) Get() *AccessProfileDocument {
	return v.value
}

func (v *NullableAccessProfileDocument) Set(val *AccessProfileDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessProfileDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessProfileDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessProfileDocument(val *AccessProfileDocument) *NullableAccessProfileDocument {
	return &NullableAccessProfileDocument{value: val, isSet: true}
}

func (v NullableAccessProfileDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessProfileDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


