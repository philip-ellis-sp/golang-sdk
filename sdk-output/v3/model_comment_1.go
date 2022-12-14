/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"time"
)

// Comment1 struct for Comment1
type Comment1 struct {
	// The comment text
	Comment *string `json:"comment,omitempty"`
	// The name of the commenter
	Commenter *string `json:"commenter,omitempty"`
	// A date-time in ISO-8601 format
	Date NullableTime `json:"date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Comment1 Comment1

// NewComment1 instantiates a new Comment1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComment1() *Comment1 {
	this := Comment1{}
	return &this
}

// NewComment1WithDefaults instantiates a new Comment1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComment1WithDefaults() *Comment1 {
	this := Comment1{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Comment1) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment1) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Comment1) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Comment1) SetComment(v string) {
	o.Comment = &v
}

// GetCommenter returns the Commenter field value if set, zero value otherwise.
func (o *Comment1) GetCommenter() string {
	if o == nil || isNil(o.Commenter) {
		var ret string
		return ret
	}
	return *o.Commenter
}

// GetCommenterOk returns a tuple with the Commenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment1) GetCommenterOk() (*string, bool) {
	if o == nil || isNil(o.Commenter) {
		return nil, false
	}
	return o.Commenter, true
}

// HasCommenter returns a boolean if a field has been set.
func (o *Comment1) HasCommenter() bool {
	if o != nil && !isNil(o.Commenter) {
		return true
	}

	return false
}

// SetCommenter gets a reference to the given string and assigns it to the Commenter field.
func (o *Comment1) SetCommenter(v string) {
	o.Commenter = &v
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Comment1) GetDate() time.Time {
	if o == nil || isNil(o.Date.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Comment1) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *Comment1) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableTime and assigns it to the Date field.
func (o *Comment1) SetDate(v time.Time) {
	o.Date.Set(&v)
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *Comment1) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *Comment1) UnsetDate() {
	o.Date.Unset()
}

func (o Comment1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.Commenter) {
		toSerialize["commenter"] = o.Commenter
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Comment1) UnmarshalJSON(bytes []byte) (err error) {
	varComment1 := _Comment1{}

	if err = json.Unmarshal(bytes, &varComment1); err == nil {
		*o = Comment1(varComment1)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "commenter")
		delete(additionalProperties, "date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComment1 struct {
	value *Comment1
	isSet bool
}

func (v NullableComment1) Get() *Comment1 {
	return v.value
}

func (v *NullableComment1) Set(val *Comment1) {
	v.value = val
	v.isSet = true
}

func (v NullableComment1) IsSet() bool {
	return v.isSet
}

func (v *NullableComment1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComment1(val *Comment1) *NullableComment1 {
	return &NullableComment1{value: val, isSet: true}
}

func (v NullableComment1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComment1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


