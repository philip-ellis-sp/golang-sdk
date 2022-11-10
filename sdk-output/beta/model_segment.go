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

// Segment struct for Segment
type Segment struct {
	Id *string `json:"id,omitempty"`
	// Segment Business Name
	Name *string `json:"name,omitempty"`
	// The time when this Segment is created
	Created *time.Time `json:"created,omitempty"`
	// The time when this Segment is modified
	Modified *time.Time `json:"modified,omitempty"`
	// Optional description of the Segment
	Description *string `json:"description,omitempty"`
	Owner *OwnerReference `json:"owner,omitempty"`
	VisibilityCriteria *VisibilityCriteria `json:"visibilityCriteria,omitempty"`
	// Whether the Segment is currently active. Inactive segments have no effect.
	Active *bool `json:"active,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Segment Segment

// NewSegment instantiates a new Segment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegment() *Segment {
	this := Segment{}
	return &this
}

// NewSegmentWithDefaults instantiates a new Segment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentWithDefaults() *Segment {
	this := Segment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Segment) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Segment) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Segment) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Segment) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Segment) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Segment) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Segment) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Segment) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Segment) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *Segment) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *Segment) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *Segment) SetModified(v time.Time) {
	o.Modified = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Segment) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Segment) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Segment) SetDescription(v string) {
	o.Description = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Segment) GetOwner() OwnerReference {
	if o == nil || isNil(o.Owner) {
		var ret OwnerReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetOwnerOk() (*OwnerReference, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Segment) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given OwnerReference and assigns it to the Owner field.
func (o *Segment) SetOwner(v OwnerReference) {
	o.Owner = &v
}

// GetVisibilityCriteria returns the VisibilityCriteria field value if set, zero value otherwise.
func (o *Segment) GetVisibilityCriteria() VisibilityCriteria {
	if o == nil || isNil(o.VisibilityCriteria) {
		var ret VisibilityCriteria
		return ret
	}
	return *o.VisibilityCriteria
}

// GetVisibilityCriteriaOk returns a tuple with the VisibilityCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetVisibilityCriteriaOk() (*VisibilityCriteria, bool) {
	if o == nil || isNil(o.VisibilityCriteria) {
		return nil, false
	}
	return o.VisibilityCriteria, true
}

// HasVisibilityCriteria returns a boolean if a field has been set.
func (o *Segment) HasVisibilityCriteria() bool {
	if o != nil && !isNil(o.VisibilityCriteria) {
		return true
	}

	return false
}

// SetVisibilityCriteria gets a reference to the given VisibilityCriteria and assigns it to the VisibilityCriteria field.
func (o *Segment) SetVisibilityCriteria(v VisibilityCriteria) {
	o.VisibilityCriteria = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Segment) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Segment) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Segment) SetActive(v bool) {
	o.Active = &v
}

func (o Segment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.VisibilityCriteria) {
		toSerialize["visibilityCriteria"] = o.VisibilityCriteria
	}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Segment) UnmarshalJSON(bytes []byte) (err error) {
	varSegment := _Segment{}

	if err = json.Unmarshal(bytes, &varSegment); err == nil {
		*o = Segment(varSegment)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "description")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "visibilityCriteria")
		delete(additionalProperties, "active")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSegment struct {
	value *Segment
	isSet bool
}

func (v NullableSegment) Get() *Segment {
	return v.value
}

func (v *NullableSegment) Set(val *Segment) {
	v.value = val
	v.isSet = true
}

func (v NullableSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegment(val *Segment) *NullableSegment {
	return &NullableSegment{value: val, isSet: true}
}

func (v NullableSegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


