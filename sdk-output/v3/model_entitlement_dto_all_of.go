/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
)

// EntitlementDtoAllOf Entitlement object that represents entitlement
type EntitlementDtoAllOf struct {
	// Name of the entitlement attribute
	Attribute *string `json:"attribute,omitempty"`
	// Raw value of the entitlement
	Value *string `json:"value,omitempty"`
	// Entitlment description
	Description *string `json:"description,omitempty"`
	// Entitlement attributes
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Schema objectType on the given application that maps to an Account Group
	SourceSchemaObjectType *string `json:"sourceSchemaObjectType,omitempty"`
	// Determines if this Entitlement is privileged.
	Privileged *bool `json:"privileged,omitempty"`
	// Determines if this Entitlement is goverened in the cloud.
	CloudGoverned *bool `json:"cloudGoverned,omitempty"`
	Source *BaseReferenceDto `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementDtoAllOf EntitlementDtoAllOf

// NewEntitlementDtoAllOf instantiates a new EntitlementDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementDtoAllOf() *EntitlementDtoAllOf {
	this := EntitlementDtoAllOf{}
	return &this
}

// NewEntitlementDtoAllOfWithDefaults instantiates a new EntitlementDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementDtoAllOfWithDefaults() *EntitlementDtoAllOf {
	this := EntitlementDtoAllOf{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *EntitlementDtoAllOf) GetAttribute() string {
	if o == nil || isNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDtoAllOf) GetAttributeOk() (*string, bool) {
	if o == nil || isNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *EntitlementDtoAllOf) HasAttribute() bool {
	if o != nil && !isNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *EntitlementDtoAllOf) SetAttribute(v string) {
	o.Attribute = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EntitlementDtoAllOf) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDtoAllOf) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EntitlementDtoAllOf) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EntitlementDtoAllOf) SetValue(v string) {
	o.Value = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementDtoAllOf) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDtoAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementDtoAllOf) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementDtoAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EntitlementDtoAllOf) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDtoAllOf) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EntitlementDtoAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *EntitlementDtoAllOf) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetSourceSchemaObjectType returns the SourceSchemaObjectType field value if set, zero value otherwise.
func (o *EntitlementDtoAllOf) GetSourceSchemaObjectType() string {
	if o == nil || isNil(o.SourceSchemaObjectType) {
		var ret string
		return ret
	}
	return *o.SourceSchemaObjectType
}

// GetSourceSchemaObjectTypeOk returns a tuple with the SourceSchemaObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDtoAllOf) GetSourceSchemaObjectTypeOk() (*string, bool) {
	if o == nil || isNil(o.SourceSchemaObjectType) {
		return nil, false
	}
	return o.SourceSchemaObjectType, true
}

// HasSourceSchemaObjectType returns a boolean if a field has been set.
func (o *EntitlementDtoAllOf) HasSourceSchemaObjectType() bool {
	if o != nil && !isNil(o.SourceSchemaObjectType) {
		return true
	}

	return false
}

// SetSourceSchemaObjectType gets a reference to the given string and assigns it to the SourceSchemaObjectType field.
func (o *EntitlementDtoAllOf) SetSourceSchemaObjectType(v string) {
	o.SourceSchemaObjectType = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *EntitlementDtoAllOf) GetPrivileged() bool {
	if o == nil || isNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDtoAllOf) GetPrivilegedOk() (*bool, bool) {
	if o == nil || isNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *EntitlementDtoAllOf) HasPrivileged() bool {
	if o != nil && !isNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *EntitlementDtoAllOf) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetCloudGoverned returns the CloudGoverned field value if set, zero value otherwise.
func (o *EntitlementDtoAllOf) GetCloudGoverned() bool {
	if o == nil || isNil(o.CloudGoverned) {
		var ret bool
		return ret
	}
	return *o.CloudGoverned
}

// GetCloudGovernedOk returns a tuple with the CloudGoverned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDtoAllOf) GetCloudGovernedOk() (*bool, bool) {
	if o == nil || isNil(o.CloudGoverned) {
		return nil, false
	}
	return o.CloudGoverned, true
}

// HasCloudGoverned returns a boolean if a field has been set.
func (o *EntitlementDtoAllOf) HasCloudGoverned() bool {
	if o != nil && !isNil(o.CloudGoverned) {
		return true
	}

	return false
}

// SetCloudGoverned gets a reference to the given bool and assigns it to the CloudGoverned field.
func (o *EntitlementDtoAllOf) SetCloudGoverned(v bool) {
	o.CloudGoverned = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *EntitlementDtoAllOf) GetSource() BaseReferenceDto {
	if o == nil || isNil(o.Source) {
		var ret BaseReferenceDto
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDtoAllOf) GetSourceOk() (*BaseReferenceDto, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *EntitlementDtoAllOf) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given BaseReferenceDto and assigns it to the Source field.
func (o *EntitlementDtoAllOf) SetSource(v BaseReferenceDto) {
	o.Source = &v
}

func (o EntitlementDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.SourceSchemaObjectType) {
		toSerialize["sourceSchemaObjectType"] = o.SourceSchemaObjectType
	}
	if !isNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !isNil(o.CloudGoverned) {
		toSerialize["cloudGoverned"] = o.CloudGoverned
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementDtoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementDtoAllOf := _EntitlementDtoAllOf{}

	if err = json.Unmarshal(bytes, &varEntitlementDtoAllOf); err == nil {
		*o = EntitlementDtoAllOf(varEntitlementDtoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "value")
		delete(additionalProperties, "description")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "sourceSchemaObjectType")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "cloudGoverned")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementDtoAllOf struct {
	value *EntitlementDtoAllOf
	isSet bool
}

func (v NullableEntitlementDtoAllOf) Get() *EntitlementDtoAllOf {
	return v.value
}

func (v *NullableEntitlementDtoAllOf) Set(val *EntitlementDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDtoAllOf(val *EntitlementDtoAllOf) *NullableEntitlementDtoAllOf {
	return &NullableEntitlementDtoAllOf{value: val, isSet: true}
}

func (v NullableEntitlementDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


