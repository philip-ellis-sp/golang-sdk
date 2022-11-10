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

// AccessItemAccountResponse struct for AccessItemAccountResponse
type AccessItemAccountResponse struct {
	// the access item type. account in this case
	AccessType *string `json:"accessType,omitempty"`
	// the access item id
	Id *string `json:"id,omitempty"`
	// the native identifier used to uniquely identify an acccount
	NativeIdentity *string `json:"nativeIdentity,omitempty"`
	// the name of the source
	SourceName *string `json:"sourceName,omitempty"`
	// the id of the source
	SourceId *string `json:"sourceId,omitempty"`
	// the number of entitlements the account will create
	EntitlementCount *string `json:"entitlementCount,omitempty"`
	// the display name of the identity
	DisplayName *string `json:"displayName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessItemAccountResponse AccessItemAccountResponse

// NewAccessItemAccountResponse instantiates a new AccessItemAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessItemAccountResponse() *AccessItemAccountResponse {
	this := AccessItemAccountResponse{}
	return &this
}

// NewAccessItemAccountResponseWithDefaults instantiates a new AccessItemAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessItemAccountResponseWithDefaults() *AccessItemAccountResponse {
	this := AccessItemAccountResponse{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *AccessItemAccountResponse) GetAccessType() string {
	if o == nil || isNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccountResponse) GetAccessTypeOk() (*string, bool) {
	if o == nil || isNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *AccessItemAccountResponse) HasAccessType() bool {
	if o != nil && !isNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *AccessItemAccountResponse) SetAccessType(v string) {
	o.AccessType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessItemAccountResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccountResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessItemAccountResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessItemAccountResponse) SetId(v string) {
	o.Id = &v
}

// GetNativeIdentity returns the NativeIdentity field value if set, zero value otherwise.
func (o *AccessItemAccountResponse) GetNativeIdentity() string {
	if o == nil || isNil(o.NativeIdentity) {
		var ret string
		return ret
	}
	return *o.NativeIdentity
}

// GetNativeIdentityOk returns a tuple with the NativeIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccountResponse) GetNativeIdentityOk() (*string, bool) {
	if o == nil || isNil(o.NativeIdentity) {
		return nil, false
	}
	return o.NativeIdentity, true
}

// HasNativeIdentity returns a boolean if a field has been set.
func (o *AccessItemAccountResponse) HasNativeIdentity() bool {
	if o != nil && !isNil(o.NativeIdentity) {
		return true
	}

	return false
}

// SetNativeIdentity gets a reference to the given string and assigns it to the NativeIdentity field.
func (o *AccessItemAccountResponse) SetNativeIdentity(v string) {
	o.NativeIdentity = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *AccessItemAccountResponse) GetSourceName() string {
	if o == nil || isNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccountResponse) GetSourceNameOk() (*string, bool) {
	if o == nil || isNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *AccessItemAccountResponse) HasSourceName() bool {
	if o != nil && !isNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *AccessItemAccountResponse) SetSourceName(v string) {
	o.SourceName = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *AccessItemAccountResponse) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccountResponse) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *AccessItemAccountResponse) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *AccessItemAccountResponse) SetSourceId(v string) {
	o.SourceId = &v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *AccessItemAccountResponse) GetEntitlementCount() string {
	if o == nil || isNil(o.EntitlementCount) {
		var ret string
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccountResponse) GetEntitlementCountOk() (*string, bool) {
	if o == nil || isNil(o.EntitlementCount) {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *AccessItemAccountResponse) HasEntitlementCount() bool {
	if o != nil && !isNil(o.EntitlementCount) {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given string and assigns it to the EntitlementCount field.
func (o *AccessItemAccountResponse) SetEntitlementCount(v string) {
	o.EntitlementCount = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AccessItemAccountResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessItemAccountResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AccessItemAccountResponse) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AccessItemAccountResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o AccessItemAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.NativeIdentity) {
		toSerialize["nativeIdentity"] = o.NativeIdentity
	}
	if !isNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !isNil(o.EntitlementCount) {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessItemAccountResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAccessItemAccountResponse := _AccessItemAccountResponse{}

	if err = json.Unmarshal(bytes, &varAccessItemAccountResponse); err == nil {
		*o = AccessItemAccountResponse(varAccessItemAccountResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accessType")
		delete(additionalProperties, "id")
		delete(additionalProperties, "nativeIdentity")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "entitlementCount")
		delete(additionalProperties, "displayName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessItemAccountResponse struct {
	value *AccessItemAccountResponse
	isSet bool
}

func (v NullableAccessItemAccountResponse) Get() *AccessItemAccountResponse {
	return v.value
}

func (v *NullableAccessItemAccountResponse) Set(val *AccessItemAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessItemAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessItemAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessItemAccountResponse(val *AccessItemAccountResponse) *NullableAccessItemAccountResponse {
	return &NullableAccessItemAccountResponse{value: val, isSet: true}
}

func (v NullableAccessItemAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessItemAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

