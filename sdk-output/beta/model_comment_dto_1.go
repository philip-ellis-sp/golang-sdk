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

// CommentDto1 struct for CommentDto1
type CommentDto1 struct {
	// Content of the comment
	Comment *string `json:"comment,omitempty"`
	Author *CommentDto1Author `json:"author,omitempty"`
	// Date and time comment was created
	Created *time.Time `json:"created,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommentDto1 CommentDto1

// NewCommentDto1 instantiates a new CommentDto1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentDto1() *CommentDto1 {
	this := CommentDto1{}
	return &this
}

// NewCommentDto1WithDefaults instantiates a new CommentDto1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentDto1WithDefaults() *CommentDto1 {
	this := CommentDto1{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CommentDto1) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentDto1) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CommentDto1) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CommentDto1) SetComment(v string) {
	o.Comment = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *CommentDto1) GetAuthor() CommentDto1Author {
	if o == nil || isNil(o.Author) {
		var ret CommentDto1Author
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentDto1) GetAuthorOk() (*CommentDto1Author, bool) {
	if o == nil || isNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *CommentDto1) HasAuthor() bool {
	if o != nil && !isNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given CommentDto1Author and assigns it to the Author field.
func (o *CommentDto1) SetAuthor(v CommentDto1Author) {
	o.Author = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CommentDto1) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentDto1) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CommentDto1) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CommentDto1) SetCreated(v time.Time) {
	o.Created = &v
}

func (o CommentDto1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CommentDto1) UnmarshalJSON(bytes []byte) (err error) {
	varCommentDto1 := _CommentDto1{}

	if err = json.Unmarshal(bytes, &varCommentDto1); err == nil {
		*o = CommentDto1(varCommentDto1)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "author")
		delete(additionalProperties, "created")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommentDto1 struct {
	value *CommentDto1
	isSet bool
}

func (v NullableCommentDto1) Get() *CommentDto1 {
	return v.value
}

func (v *NullableCommentDto1) Set(val *CommentDto1) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentDto1) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentDto1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentDto1(val *CommentDto1) *NullableCommentDto1 {
	return &NullableCommentDto1{value: val, isSet: true}
}

func (v NullableCommentDto1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentDto1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


