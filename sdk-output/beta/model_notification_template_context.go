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

// NotificationTemplateContext struct for NotificationTemplateContext
type NotificationTemplateContext struct {
	// A JSON object that stores the context.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// When the global context was created
	Created *time.Time `json:"created,omitempty"`
	// When the global context was last modified
	Modified *time.Time `json:"modified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationTemplateContext NotificationTemplateContext

// NewNotificationTemplateContext instantiates a new NotificationTemplateContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationTemplateContext() *NotificationTemplateContext {
	this := NotificationTemplateContext{}
	return &this
}

// NewNotificationTemplateContextWithDefaults instantiates a new NotificationTemplateContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationTemplateContextWithDefaults() *NotificationTemplateContext {
	this := NotificationTemplateContext{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NotificationTemplateContext) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTemplateContext) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NotificationTemplateContext) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *NotificationTemplateContext) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NotificationTemplateContext) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTemplateContext) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NotificationTemplateContext) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NotificationTemplateContext) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *NotificationTemplateContext) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTemplateContext) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *NotificationTemplateContext) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *NotificationTemplateContext) SetModified(v time.Time) {
	o.Modified = &v
}

func (o NotificationTemplateContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NotificationTemplateContext) UnmarshalJSON(bytes []byte) (err error) {
	varNotificationTemplateContext := _NotificationTemplateContext{}

	if err = json.Unmarshal(bytes, &varNotificationTemplateContext); err == nil {
		*o = NotificationTemplateContext(varNotificationTemplateContext)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationTemplateContext struct {
	value *NotificationTemplateContext
	isSet bool
}

func (v NullableNotificationTemplateContext) Get() *NotificationTemplateContext {
	return v.value
}

func (v *NullableNotificationTemplateContext) Set(val *NotificationTemplateContext) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTemplateContext) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTemplateContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTemplateContext(val *NotificationTemplateContext) *NullableNotificationTemplateContext {
	return &NullableNotificationTemplateContext{value: val, isSet: true}
}

func (v NullableNotificationTemplateContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTemplateContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


