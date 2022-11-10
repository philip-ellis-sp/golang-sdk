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

// ApprovalItems struct for ApprovalItems
type ApprovalItems struct {
	// ID of the approval item
	Id *string `json:"id,omitempty"`
	// The account referenced by the approval item
	Account *string `json:"account,omitempty"`
	// The name the application/source
	Application *string `json:"application,omitempty"`
	// The name of the attribute
	AttributeName *string `json:"attributeName,omitempty"`
	// The operation of the attribute
	AttributeOperation *string `json:"attributeOperation,omitempty"`
	// The value of the attribute
	AttributeValue *string `json:"attributeValue,omitempty"`
	State *WorkItemState `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalItems ApprovalItems

// NewApprovalItems instantiates a new ApprovalItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalItems() *ApprovalItems {
	this := ApprovalItems{}
	return &this
}

// NewApprovalItemsWithDefaults instantiates a new ApprovalItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalItemsWithDefaults() *ApprovalItems {
	this := ApprovalItems{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApprovalItems) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApprovalItems) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApprovalItems) SetId(v string) {
	o.Id = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApprovalItems) GetAccount() string {
	if o == nil || isNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetAccountOk() (*string, bool) {
	if o == nil || isNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApprovalItems) HasAccount() bool {
	if o != nil && !isNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *ApprovalItems) SetAccount(v string) {
	o.Account = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ApprovalItems) GetApplication() string {
	if o == nil || isNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetApplicationOk() (*string, bool) {
	if o == nil || isNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ApprovalItems) HasApplication() bool {
	if o != nil && !isNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *ApprovalItems) SetApplication(v string) {
	o.Application = &v
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *ApprovalItems) GetAttributeName() string {
	if o == nil || isNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetAttributeNameOk() (*string, bool) {
	if o == nil || isNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *ApprovalItems) HasAttributeName() bool {
	if o != nil && !isNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *ApprovalItems) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetAttributeOperation returns the AttributeOperation field value if set, zero value otherwise.
func (o *ApprovalItems) GetAttributeOperation() string {
	if o == nil || isNil(o.AttributeOperation) {
		var ret string
		return ret
	}
	return *o.AttributeOperation
}

// GetAttributeOperationOk returns a tuple with the AttributeOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetAttributeOperationOk() (*string, bool) {
	if o == nil || isNil(o.AttributeOperation) {
		return nil, false
	}
	return o.AttributeOperation, true
}

// HasAttributeOperation returns a boolean if a field has been set.
func (o *ApprovalItems) HasAttributeOperation() bool {
	if o != nil && !isNil(o.AttributeOperation) {
		return true
	}

	return false
}

// SetAttributeOperation gets a reference to the given string and assigns it to the AttributeOperation field.
func (o *ApprovalItems) SetAttributeOperation(v string) {
	o.AttributeOperation = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *ApprovalItems) GetAttributeValue() string {
	if o == nil || isNil(o.AttributeValue) {
		var ret string
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetAttributeValueOk() (*string, bool) {
	if o == nil || isNil(o.AttributeValue) {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *ApprovalItems) HasAttributeValue() bool {
	if o != nil && !isNil(o.AttributeValue) {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given string and assigns it to the AttributeValue field.
func (o *ApprovalItems) SetAttributeValue(v string) {
	o.AttributeValue = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ApprovalItems) GetState() WorkItemState {
	if o == nil || isNil(o.State) {
		var ret WorkItemState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItems) GetStateOk() (*WorkItemState, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApprovalItems) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given WorkItemState and assigns it to the State field.
func (o *ApprovalItems) SetState(v WorkItemState) {
	o.State = &v
}

func (o ApprovalItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !isNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !isNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !isNil(o.AttributeOperation) {
		toSerialize["attributeOperation"] = o.AttributeOperation
	}
	if !isNil(o.AttributeValue) {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApprovalItems) UnmarshalJSON(bytes []byte) (err error) {
	varApprovalItems := _ApprovalItems{}

	if err = json.Unmarshal(bytes, &varApprovalItems); err == nil {
		*o = ApprovalItems(varApprovalItems)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "account")
		delete(additionalProperties, "application")
		delete(additionalProperties, "attributeName")
		delete(additionalProperties, "attributeOperation")
		delete(additionalProperties, "attributeValue")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalItems struct {
	value *ApprovalItems
	isSet bool
}

func (v NullableApprovalItems) Get() *ApprovalItems {
	return v.value
}

func (v *NullableApprovalItems) Set(val *ApprovalItems) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalItems) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalItems(val *ApprovalItems) *NullableApprovalItems {
	return &NullableApprovalItems{value: val, isSet: true}
}

func (v NullableApprovalItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


