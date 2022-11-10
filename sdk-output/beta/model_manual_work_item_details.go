/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"time"
)

// ManualWorkItemDetails struct for ManualWorkItemDetails
type ManualWorkItemDetails struct {
	// True if the request for this item was forwarded from one owner to another.
	Forwarded *bool `json:"forwarded,omitempty"`
	OriginalOwner *BaseReferenceDto `json:"originalOwner,omitempty"`
	CurrentOwner *BaseReferenceDto `json:"currentOwner,omitempty"`
	// Time at which item was modified.
	Modified *time.Time `json:"modified,omitempty"`
	Status *ManualWorkItemState `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManualWorkItemDetails ManualWorkItemDetails

// NewManualWorkItemDetails instantiates a new ManualWorkItemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualWorkItemDetails() *ManualWorkItemDetails {
	this := ManualWorkItemDetails{}
	return &this
}

// NewManualWorkItemDetailsWithDefaults instantiates a new ManualWorkItemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualWorkItemDetailsWithDefaults() *ManualWorkItemDetails {
	this := ManualWorkItemDetails{}
	return &this
}

// GetForwarded returns the Forwarded field value if set, zero value otherwise.
func (o *ManualWorkItemDetails) GetForwarded() bool {
	if o == nil || isNil(o.Forwarded) {
		var ret bool
		return ret
	}
	return *o.Forwarded
}

// GetForwardedOk returns a tuple with the Forwarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualWorkItemDetails) GetForwardedOk() (*bool, bool) {
	if o == nil || isNil(o.Forwarded) {
		return nil, false
	}
	return o.Forwarded, true
}

// HasForwarded returns a boolean if a field has been set.
func (o *ManualWorkItemDetails) HasForwarded() bool {
	if o != nil && !isNil(o.Forwarded) {
		return true
	}

	return false
}

// SetForwarded gets a reference to the given bool and assigns it to the Forwarded field.
func (o *ManualWorkItemDetails) SetForwarded(v bool) {
	o.Forwarded = &v
}

// GetOriginalOwner returns the OriginalOwner field value if set, zero value otherwise.
func (o *ManualWorkItemDetails) GetOriginalOwner() BaseReferenceDto {
	if o == nil || isNil(o.OriginalOwner) {
		var ret BaseReferenceDto
		return ret
	}
	return *o.OriginalOwner
}

// GetOriginalOwnerOk returns a tuple with the OriginalOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualWorkItemDetails) GetOriginalOwnerOk() (*BaseReferenceDto, bool) {
	if o == nil || isNil(o.OriginalOwner) {
		return nil, false
	}
	return o.OriginalOwner, true
}

// HasOriginalOwner returns a boolean if a field has been set.
func (o *ManualWorkItemDetails) HasOriginalOwner() bool {
	if o != nil && !isNil(o.OriginalOwner) {
		return true
	}

	return false
}

// SetOriginalOwner gets a reference to the given BaseReferenceDto and assigns it to the OriginalOwner field.
func (o *ManualWorkItemDetails) SetOriginalOwner(v BaseReferenceDto) {
	o.OriginalOwner = &v
}

// GetCurrentOwner returns the CurrentOwner field value if set, zero value otherwise.
func (o *ManualWorkItemDetails) GetCurrentOwner() BaseReferenceDto {
	if o == nil || isNil(o.CurrentOwner) {
		var ret BaseReferenceDto
		return ret
	}
	return *o.CurrentOwner
}

// GetCurrentOwnerOk returns a tuple with the CurrentOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualWorkItemDetails) GetCurrentOwnerOk() (*BaseReferenceDto, bool) {
	if o == nil || isNil(o.CurrentOwner) {
		return nil, false
	}
	return o.CurrentOwner, true
}

// HasCurrentOwner returns a boolean if a field has been set.
func (o *ManualWorkItemDetails) HasCurrentOwner() bool {
	if o != nil && !isNil(o.CurrentOwner) {
		return true
	}

	return false
}

// SetCurrentOwner gets a reference to the given BaseReferenceDto and assigns it to the CurrentOwner field.
func (o *ManualWorkItemDetails) SetCurrentOwner(v BaseReferenceDto) {
	o.CurrentOwner = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *ManualWorkItemDetails) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualWorkItemDetails) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *ManualWorkItemDetails) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *ManualWorkItemDetails) SetModified(v time.Time) {
	o.Modified = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ManualWorkItemDetails) GetStatus() ManualWorkItemState {
	if o == nil || isNil(o.Status) {
		var ret ManualWorkItemState
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualWorkItemDetails) GetStatusOk() (*ManualWorkItemState, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ManualWorkItemDetails) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ManualWorkItemState and assigns it to the Status field.
func (o *ManualWorkItemDetails) SetStatus(v ManualWorkItemState) {
	o.Status = &v
}

func (o ManualWorkItemDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Forwarded) {
		toSerialize["forwarded"] = o.Forwarded
	}
	if !isNil(o.OriginalOwner) {
		toSerialize["originalOwner"] = o.OriginalOwner
	}
	if !isNil(o.CurrentOwner) {
		toSerialize["currentOwner"] = o.CurrentOwner
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ManualWorkItemDetails) UnmarshalJSON(bytes []byte) (err error) {
	varManualWorkItemDetails := _ManualWorkItemDetails{}

	if err = json.Unmarshal(bytes, &varManualWorkItemDetails); err == nil {
		*o = ManualWorkItemDetails(varManualWorkItemDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "forwarded")
		delete(additionalProperties, "originalOwner")
		delete(additionalProperties, "currentOwner")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManualWorkItemDetails struct {
	value *ManualWorkItemDetails
	isSet bool
}

func (v NullableManualWorkItemDetails) Get() *ManualWorkItemDetails {
	return v.value
}

func (v *NullableManualWorkItemDetails) Set(val *ManualWorkItemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableManualWorkItemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableManualWorkItemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualWorkItemDetails(val *ManualWorkItemDetails) *NullableManualWorkItemDetails {
	return &NullableManualWorkItemDetails{value: val, isSet: true}
}

func (v NullableManualWorkItemDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualWorkItemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


