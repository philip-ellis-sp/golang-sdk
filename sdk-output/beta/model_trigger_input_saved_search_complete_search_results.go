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

// TriggerInputSavedSearchCompleteSearchResults A preview of the search results for each object type. This includes a count as well as headers, and the first several rows of data, per object type.
type TriggerInputSavedSearchCompleteSearchResults struct {
	Account NullableTriggerInputSavedSearchCompleteSearchResultsAccount `json:"Account,omitempty"`
	Entitlement NullableTriggerInputSavedSearchCompleteSearchResultsEntitlement `json:"Entitlement,omitempty"`
	Identity NullableTriggerInputSavedSearchCompleteSearchResultsIdentity `json:"Identity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputSavedSearchCompleteSearchResults TriggerInputSavedSearchCompleteSearchResults

// NewTriggerInputSavedSearchCompleteSearchResults instantiates a new TriggerInputSavedSearchCompleteSearchResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputSavedSearchCompleteSearchResults() *TriggerInputSavedSearchCompleteSearchResults {
	this := TriggerInputSavedSearchCompleteSearchResults{}
	return &this
}

// NewTriggerInputSavedSearchCompleteSearchResultsWithDefaults instantiates a new TriggerInputSavedSearchCompleteSearchResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputSavedSearchCompleteSearchResultsWithDefaults() *TriggerInputSavedSearchCompleteSearchResults {
	this := TriggerInputSavedSearchCompleteSearchResults{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputSavedSearchCompleteSearchResults) GetAccount() TriggerInputSavedSearchCompleteSearchResultsAccount {
	if o == nil || isNil(o.Account.Get()) {
		var ret TriggerInputSavedSearchCompleteSearchResultsAccount
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputSavedSearchCompleteSearchResults) GetAccountOk() (*TriggerInputSavedSearchCompleteSearchResultsAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *TriggerInputSavedSearchCompleteSearchResults) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableTriggerInputSavedSearchCompleteSearchResultsAccount and assigns it to the Account field.
func (o *TriggerInputSavedSearchCompleteSearchResults) SetAccount(v TriggerInputSavedSearchCompleteSearchResultsAccount) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *TriggerInputSavedSearchCompleteSearchResults) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *TriggerInputSavedSearchCompleteSearchResults) UnsetAccount() {
	o.Account.Unset()
}

// GetEntitlement returns the Entitlement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputSavedSearchCompleteSearchResults) GetEntitlement() TriggerInputSavedSearchCompleteSearchResultsEntitlement {
	if o == nil || isNil(o.Entitlement.Get()) {
		var ret TriggerInputSavedSearchCompleteSearchResultsEntitlement
		return ret
	}
	return *o.Entitlement.Get()
}

// GetEntitlementOk returns a tuple with the Entitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputSavedSearchCompleteSearchResults) GetEntitlementOk() (*TriggerInputSavedSearchCompleteSearchResultsEntitlement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entitlement.Get(), o.Entitlement.IsSet()
}

// HasEntitlement returns a boolean if a field has been set.
func (o *TriggerInputSavedSearchCompleteSearchResults) HasEntitlement() bool {
	if o != nil && o.Entitlement.IsSet() {
		return true
	}

	return false
}

// SetEntitlement gets a reference to the given NullableTriggerInputSavedSearchCompleteSearchResultsEntitlement and assigns it to the Entitlement field.
func (o *TriggerInputSavedSearchCompleteSearchResults) SetEntitlement(v TriggerInputSavedSearchCompleteSearchResultsEntitlement) {
	o.Entitlement.Set(&v)
}
// SetEntitlementNil sets the value for Entitlement to be an explicit nil
func (o *TriggerInputSavedSearchCompleteSearchResults) SetEntitlementNil() {
	o.Entitlement.Set(nil)
}

// UnsetEntitlement ensures that no value is present for Entitlement, not even an explicit nil
func (o *TriggerInputSavedSearchCompleteSearchResults) UnsetEntitlement() {
	o.Entitlement.Unset()
}

// GetIdentity returns the Identity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerInputSavedSearchCompleteSearchResults) GetIdentity() TriggerInputSavedSearchCompleteSearchResultsIdentity {
	if o == nil || isNil(o.Identity.Get()) {
		var ret TriggerInputSavedSearchCompleteSearchResultsIdentity
		return ret
	}
	return *o.Identity.Get()
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerInputSavedSearchCompleteSearchResults) GetIdentityOk() (*TriggerInputSavedSearchCompleteSearchResultsIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identity.Get(), o.Identity.IsSet()
}

// HasIdentity returns a boolean if a field has been set.
func (o *TriggerInputSavedSearchCompleteSearchResults) HasIdentity() bool {
	if o != nil && o.Identity.IsSet() {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NullableTriggerInputSavedSearchCompleteSearchResultsIdentity and assigns it to the Identity field.
func (o *TriggerInputSavedSearchCompleteSearchResults) SetIdentity(v TriggerInputSavedSearchCompleteSearchResultsIdentity) {
	o.Identity.Set(&v)
}
// SetIdentityNil sets the value for Identity to be an explicit nil
func (o *TriggerInputSavedSearchCompleteSearchResults) SetIdentityNil() {
	o.Identity.Set(nil)
}

// UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
func (o *TriggerInputSavedSearchCompleteSearchResults) UnsetIdentity() {
	o.Identity.Unset()
}

func (o TriggerInputSavedSearchCompleteSearchResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}
	if o.Entitlement.IsSet() {
		toSerialize["Entitlement"] = o.Entitlement.Get()
	}
	if o.Identity.IsSet() {
		toSerialize["Identity"] = o.Identity.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputSavedSearchCompleteSearchResults) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputSavedSearchCompleteSearchResults := _TriggerInputSavedSearchCompleteSearchResults{}

	if err = json.Unmarshal(bytes, &varTriggerInputSavedSearchCompleteSearchResults); err == nil {
		*o = TriggerInputSavedSearchCompleteSearchResults(varTriggerInputSavedSearchCompleteSearchResults)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Entitlement")
		delete(additionalProperties, "Identity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputSavedSearchCompleteSearchResults struct {
	value *TriggerInputSavedSearchCompleteSearchResults
	isSet bool
}

func (v NullableTriggerInputSavedSearchCompleteSearchResults) Get() *TriggerInputSavedSearchCompleteSearchResults {
	return v.value
}

func (v *NullableTriggerInputSavedSearchCompleteSearchResults) Set(val *TriggerInputSavedSearchCompleteSearchResults) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputSavedSearchCompleteSearchResults) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputSavedSearchCompleteSearchResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputSavedSearchCompleteSearchResults(val *TriggerInputSavedSearchCompleteSearchResults) *NullableTriggerInputSavedSearchCompleteSearchResults {
	return &NullableTriggerInputSavedSearchCompleteSearchResults{value: val, isSet: true}
}

func (v NullableTriggerInputSavedSearchCompleteSearchResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputSavedSearchCompleteSearchResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

