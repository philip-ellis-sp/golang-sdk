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

// ManagedClientStatus Managed Client Status
type ManagedClientStatus struct {
	// ManagedClientStatus body information
	Body map[string]interface{} `json:"body"`
	Status ManagedClientStatusEnum `json:"status"`
	Type ManagedClientType `json:"type"`
	// timestamp on the Client Status update
	Timestamp time.Time `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _ManagedClientStatus ManagedClientStatus

// NewManagedClientStatus instantiates a new ManagedClientStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedClientStatus(body map[string]interface{}, status ManagedClientStatusEnum, type_ ManagedClientType, timestamp time.Time) *ManagedClientStatus {
	this := ManagedClientStatus{}
	this.Body = body
	this.Status = status
	this.Type = type_
	this.Timestamp = timestamp
	return &this
}

// NewManagedClientStatusWithDefaults instantiates a new ManagedClientStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedClientStatusWithDefaults() *ManagedClientStatus {
	this := ManagedClientStatus{}
	return &this
}

// GetBody returns the Body field value
func (o *ManagedClientStatus) GetBody() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *ManagedClientStatus) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Body, true
}

// SetBody sets field value
func (o *ManagedClientStatus) SetBody(v map[string]interface{}) {
	o.Body = v
}

// GetStatus returns the Status field value
func (o *ManagedClientStatus) GetStatus() ManagedClientStatusEnum {
	if o == nil {
		var ret ManagedClientStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ManagedClientStatus) GetStatusOk() (*ManagedClientStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ManagedClientStatus) SetStatus(v ManagedClientStatusEnum) {
	o.Status = v
}

// GetType returns the Type field value
func (o *ManagedClientStatus) GetType() ManagedClientType {
	if o == nil {
		var ret ManagedClientType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ManagedClientStatus) GetTypeOk() (*ManagedClientType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ManagedClientStatus) SetType(v ManagedClientType) {
	o.Type = v
}

// GetTimestamp returns the Timestamp field value
func (o *ManagedClientStatus) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ManagedClientStatus) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ManagedClientStatus) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o ManagedClientStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["body"] = o.Body
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ManagedClientStatus) UnmarshalJSON(bytes []byte) (err error) {
	varManagedClientStatus := _ManagedClientStatus{}

	if err = json.Unmarshal(bytes, &varManagedClientStatus); err == nil {
		*o = ManagedClientStatus(varManagedClientStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "body")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagedClientStatus struct {
	value *ManagedClientStatus
	isSet bool
}

func (v NullableManagedClientStatus) Get() *ManagedClientStatus {
	return v.value
}

func (v *NullableManagedClientStatus) Set(val *ManagedClientStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedClientStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedClientStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedClientStatus(val *ManagedClientStatus) *NullableManagedClientStatus {
	return &NullableManagedClientStatus{value: val, isSet: true}
}

func (v NullableManagedClientStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedClientStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


