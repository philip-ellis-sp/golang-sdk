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

// WorkflowLibraryOperator struct for WorkflowLibraryOperator
type WorkflowLibraryOperator struct {
	// Operator ID.
	Id *string `json:"id,omitempty"`
	// Operator friendly name
	Name *string `json:"name,omitempty"`
	// Description of the operator
	Description *string `json:"description,omitempty"`
	// One or more inputs that the operator accepts
	FormFields []WorkflowLibraryFormFields `json:"formFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowLibraryOperator WorkflowLibraryOperator

// NewWorkflowLibraryOperator instantiates a new WorkflowLibraryOperator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowLibraryOperator() *WorkflowLibraryOperator {
	this := WorkflowLibraryOperator{}
	return &this
}

// NewWorkflowLibraryOperatorWithDefaults instantiates a new WorkflowLibraryOperator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowLibraryOperatorWithDefaults() *WorkflowLibraryOperator {
	this := WorkflowLibraryOperator{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowLibraryOperator) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryOperator) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowLibraryOperator) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowLibraryOperator) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowLibraryOperator) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryOperator) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowLibraryOperator) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowLibraryOperator) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowLibraryOperator) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryOperator) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowLibraryOperator) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowLibraryOperator) SetDescription(v string) {
	o.Description = &v
}

// GetFormFields returns the FormFields field value if set, zero value otherwise.
func (o *WorkflowLibraryOperator) GetFormFields() []WorkflowLibraryFormFields {
	if o == nil || isNil(o.FormFields) {
		var ret []WorkflowLibraryFormFields
		return ret
	}
	return o.FormFields
}

// GetFormFieldsOk returns a tuple with the FormFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryOperator) GetFormFieldsOk() ([]WorkflowLibraryFormFields, bool) {
	if o == nil || isNil(o.FormFields) {
		return nil, false
	}
	return o.FormFields, true
}

// HasFormFields returns a boolean if a field has been set.
func (o *WorkflowLibraryOperator) HasFormFields() bool {
	if o != nil && !isNil(o.FormFields) {
		return true
	}

	return false
}

// SetFormFields gets a reference to the given []WorkflowLibraryFormFields and assigns it to the FormFields field.
func (o *WorkflowLibraryOperator) SetFormFields(v []WorkflowLibraryFormFields) {
	o.FormFields = v
}

func (o WorkflowLibraryOperator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.FormFields) {
		toSerialize["formFields"] = o.FormFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowLibraryOperator) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowLibraryOperator := _WorkflowLibraryOperator{}

	if err = json.Unmarshal(bytes, &varWorkflowLibraryOperator); err == nil {
		*o = WorkflowLibraryOperator(varWorkflowLibraryOperator)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "formFields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowLibraryOperator struct {
	value *WorkflowLibraryOperator
	isSet bool
}

func (v NullableWorkflowLibraryOperator) Get() *WorkflowLibraryOperator {
	return v.value
}

func (v *NullableWorkflowLibraryOperator) Set(val *WorkflowLibraryOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowLibraryOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowLibraryOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowLibraryOperator(val *WorkflowLibraryOperator) *NullableWorkflowLibraryOperator {
	return &NullableWorkflowLibraryOperator{value: val, isSet: true}
}

func (v NullableWorkflowLibraryOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowLibraryOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


