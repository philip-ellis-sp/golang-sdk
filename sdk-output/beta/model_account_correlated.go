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

// AccountCorrelated struct for AccountCorrelated
type AccountCorrelated struct {
	Identity TriggerInputAccountCorrelatedIdentity `json:"identity"`
	Source TriggerInputAccountCorrelatedSource `json:"source"`
	Account TriggerInputAccountCorrelatedAccount `json:"account"`
	// The attributes associated with the account.  Attributes are unique per source.
	Attributes map[string]interface{} `json:"attributes"`
	// The number of entitlements associated with this account.
	EntitlementCount *int32 `json:"entitlementCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountCorrelated AccountCorrelated

// NewAccountCorrelated instantiates a new AccountCorrelated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCorrelated(identity TriggerInputAccountCorrelatedIdentity, source TriggerInputAccountCorrelatedSource, account TriggerInputAccountCorrelatedAccount, attributes map[string]interface{}) *AccountCorrelated {
	this := AccountCorrelated{}
	this.Identity = identity
	this.Source = source
	this.Account = account
	this.Attributes = attributes
	return &this
}

// NewAccountCorrelatedWithDefaults instantiates a new AccountCorrelated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCorrelatedWithDefaults() *AccountCorrelated {
	this := AccountCorrelated{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *AccountCorrelated) GetIdentity() TriggerInputAccountCorrelatedIdentity {
	if o == nil {
		var ret TriggerInputAccountCorrelatedIdentity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *AccountCorrelated) GetIdentityOk() (*TriggerInputAccountCorrelatedIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *AccountCorrelated) SetIdentity(v TriggerInputAccountCorrelatedIdentity) {
	o.Identity = v
}

// GetSource returns the Source field value
func (o *AccountCorrelated) GetSource() TriggerInputAccountCorrelatedSource {
	if o == nil {
		var ret TriggerInputAccountCorrelatedSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AccountCorrelated) GetSourceOk() (*TriggerInputAccountCorrelatedSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AccountCorrelated) SetSource(v TriggerInputAccountCorrelatedSource) {
	o.Source = v
}

// GetAccount returns the Account field value
func (o *AccountCorrelated) GetAccount() TriggerInputAccountCorrelatedAccount {
	if o == nil {
		var ret TriggerInputAccountCorrelatedAccount
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AccountCorrelated) GetAccountOk() (*TriggerInputAccountCorrelatedAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AccountCorrelated) SetAccount(v TriggerInputAccountCorrelatedAccount) {
	o.Account = v
}

// GetAttributes returns the Attributes field value
func (o *AccountCorrelated) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AccountCorrelated) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *AccountCorrelated) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *AccountCorrelated) GetEntitlementCount() int32 {
	if o == nil || isNil(o.EntitlementCount) {
		var ret int32
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCorrelated) GetEntitlementCountOk() (*int32, bool) {
	if o == nil || isNil(o.EntitlementCount) {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *AccountCorrelated) HasEntitlementCount() bool {
	if o != nil && !isNil(o.EntitlementCount) {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given int32 and assigns it to the EntitlementCount field.
func (o *AccountCorrelated) SetEntitlementCount(v int32) {
	o.EntitlementCount = &v
}

func (o AccountCorrelated) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.EntitlementCount) {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountCorrelated) UnmarshalJSON(bytes []byte) (err error) {
	varAccountCorrelated := _AccountCorrelated{}

	if err = json.Unmarshal(bytes, &varAccountCorrelated); err == nil {
		*o = AccountCorrelated(varAccountCorrelated)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "source")
		delete(additionalProperties, "account")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "entitlementCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountCorrelated struct {
	value *AccountCorrelated
	isSet bool
}

func (v NullableAccountCorrelated) Get() *AccountCorrelated {
	return v.value
}

func (v *NullableAccountCorrelated) Set(val *AccountCorrelated) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCorrelated) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCorrelated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCorrelated(val *AccountCorrelated) *NullableAccountCorrelated {
	return &NullableAccountCorrelated{value: val, isSet: true}
}

func (v NullableAccountCorrelated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCorrelated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


