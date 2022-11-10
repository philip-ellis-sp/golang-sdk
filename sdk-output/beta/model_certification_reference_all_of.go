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

// CertificationReferenceAllOf struct for CertificationReferenceAllOf
type CertificationReferenceAllOf struct {
	Reviewer *Reviewer `json:"reviewer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificationReferenceAllOf CertificationReferenceAllOf

// NewCertificationReferenceAllOf instantiates a new CertificationReferenceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificationReferenceAllOf() *CertificationReferenceAllOf {
	this := CertificationReferenceAllOf{}
	return &this
}

// NewCertificationReferenceAllOfWithDefaults instantiates a new CertificationReferenceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificationReferenceAllOfWithDefaults() *CertificationReferenceAllOf {
	this := CertificationReferenceAllOf{}
	return &this
}

// GetReviewer returns the Reviewer field value if set, zero value otherwise.
func (o *CertificationReferenceAllOf) GetReviewer() Reviewer {
	if o == nil || isNil(o.Reviewer) {
		var ret Reviewer
		return ret
	}
	return *o.Reviewer
}

// GetReviewerOk returns a tuple with the Reviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificationReferenceAllOf) GetReviewerOk() (*Reviewer, bool) {
	if o == nil || isNil(o.Reviewer) {
		return nil, false
	}
	return o.Reviewer, true
}

// HasReviewer returns a boolean if a field has been set.
func (o *CertificationReferenceAllOf) HasReviewer() bool {
	if o != nil && !isNil(o.Reviewer) {
		return true
	}

	return false
}

// SetReviewer gets a reference to the given Reviewer and assigns it to the Reviewer field.
func (o *CertificationReferenceAllOf) SetReviewer(v Reviewer) {
	o.Reviewer = &v
}

func (o CertificationReferenceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Reviewer) {
		toSerialize["reviewer"] = o.Reviewer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CertificationReferenceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCertificationReferenceAllOf := _CertificationReferenceAllOf{}

	if err = json.Unmarshal(bytes, &varCertificationReferenceAllOf); err == nil {
		*o = CertificationReferenceAllOf(varCertificationReferenceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reviewer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificationReferenceAllOf struct {
	value *CertificationReferenceAllOf
	isSet bool
}

func (v NullableCertificationReferenceAllOf) Get() *CertificationReferenceAllOf {
	return v.value
}

func (v *NullableCertificationReferenceAllOf) Set(val *CertificationReferenceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificationReferenceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificationReferenceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificationReferenceAllOf(val *CertificationReferenceAllOf) *NullableCertificationReferenceAllOf {
	return &NullableCertificationReferenceAllOf{value: val, isSet: true}
}

func (v NullableCertificationReferenceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificationReferenceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

