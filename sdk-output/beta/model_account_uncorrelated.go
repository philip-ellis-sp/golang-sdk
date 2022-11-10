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

// AccountUncorrelated struct for AccountUncorrelated
type AccountUncorrelated struct {
	Identity TriggerInputAccountUncorrelatedIdentity `json:"identity"`
	Source TriggerInputAccountUncorrelatedSource `json:"source"`
	Account TriggerInputAccountUncorrelatedAccount `json:"account"`
	// The number of entitlements associated with this account.
	EntitlementCount *int32 `json:"entitlementCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountUncorrelated AccountUncorrelated

// NewAccountUncorrelated instantiates a new AccountUncorrelated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUncorrelated(identity TriggerInputAccountUncorrelatedIdentity, source TriggerInputAccountUncorrelatedSource, account TriggerInputAccountUncorrelatedAccount) *AccountUncorrelated {
	this := AccountUncorrelated{}
	this.Identity = identity
	this.Source = source
	this.Account = account
	return &this
}

// NewAccountUncorrelatedWithDefaults instantiates a new AccountUncorrelated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUncorrelatedWithDefaults() *AccountUncorrelated {
	this := AccountUncorrelated{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *AccountUncorrelated) GetIdentity() TriggerInputAccountUncorrelatedIdentity {
	if o == nil {
		var ret TriggerInputAccountUncorrelatedIdentity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *AccountUncorrelated) GetIdentityOk() (*TriggerInputAccountUncorrelatedIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *AccountUncorrelated) SetIdentity(v TriggerInputAccountUncorrelatedIdentity) {
	o.Identity = v
}

// GetSource returns the Source field value
func (o *AccountUncorrelated) GetSource() TriggerInputAccountUncorrelatedSource {
	if o == nil {
		var ret TriggerInputAccountUncorrelatedSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AccountUncorrelated) GetSourceOk() (*TriggerInputAccountUncorrelatedSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AccountUncorrelated) SetSource(v TriggerInputAccountUncorrelatedSource) {
	o.Source = v
}

// GetAccount returns the Account field value
func (o *AccountUncorrelated) GetAccount() TriggerInputAccountUncorrelatedAccount {
	if o == nil {
		var ret TriggerInputAccountUncorrelatedAccount
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AccountUncorrelated) GetAccountOk() (*TriggerInputAccountUncorrelatedAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AccountUncorrelated) SetAccount(v TriggerInputAccountUncorrelatedAccount) {
	o.Account = v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *AccountUncorrelated) GetEntitlementCount() int32 {
	if o == nil || isNil(o.EntitlementCount) {
		var ret int32
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUncorrelated) GetEntitlementCountOk() (*int32, bool) {
	if o == nil || isNil(o.EntitlementCount) {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *AccountUncorrelated) HasEntitlementCount() bool {
	if o != nil && !isNil(o.EntitlementCount) {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given int32 and assigns it to the EntitlementCount field.
func (o *AccountUncorrelated) SetEntitlementCount(v int32) {
	o.EntitlementCount = &v
}

func (o AccountUncorrelated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identity"] = o.Identity
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if !isNil(o.EntitlementCount) {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountUncorrelated) UnmarshalJSON(bytes []byte) (err error) {
	varAccountUncorrelated := _AccountUncorrelated{}

	if err = json.Unmarshal(bytes, &varAccountUncorrelated); err == nil {
		*o = AccountUncorrelated(varAccountUncorrelated)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "source")
		delete(additionalProperties, "account")
		delete(additionalProperties, "entitlementCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUncorrelated struct {
	value *AccountUncorrelated
	isSet bool
}

func (v NullableAccountUncorrelated) Get() *AccountUncorrelated {
	return v.value
}

func (v *NullableAccountUncorrelated) Set(val *AccountUncorrelated) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUncorrelated) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUncorrelated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUncorrelated(val *AccountUncorrelated) *NullableAccountUncorrelated {
	return &NullableAccountUncorrelated{value: val, isSet: true}
}

func (v NullableAccountUncorrelated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUncorrelated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


