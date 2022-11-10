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

// ResourceObject Representation of the object which is returned from source connectors.
type ResourceObject struct {
	// Identifier of the specific instance where this object resides.
	Instance *string `json:"instance,omitempty"`
	// Native identity of the object in the Source.
	Identity *string `json:"identity,omitempty"`
	// Universal unique identifier of the object in the Source.
	Uuid *string `json:"uuid,omitempty"`
	// Native identity that the object has previously.
	PreviousIdentity *string `json:"previousIdentity,omitempty"`
	// Display name for this object.
	Name *string `json:"name,omitempty"`
	// Type of object.
	ObjectType *string `json:"objectType,omitempty"`
	// A flag indicating that this is an incomplete object. Used in special cases where the connector has to return account information in several phases and the objects might not have a complete set of all account attributes. The attributes in this object will replace the corresponding attributes in the Link, but no other Link attributes will be changed.
	Incomplete *bool `json:"incomplete,omitempty"`
	// A flag indicating that this is an incremental change object. This is similar to incomplete but it also means that the values of any multi-valued attributes in this object should be merged with the existing values in the Link rather than replacing the existing Link value.
	Incremental *bool `json:"incremental,omitempty"`
	// A flag indicating that this object has been deleted. This is set only when doing delta aggregation and the connector supports detection of native deletes.
	Delete *bool `json:"delete,omitempty"`
	// A flag set indicating that the values in the attributes represent things to remove rather than things to add. Setting this implies incremental. The values which are always for multi-valued attributes are removed from the current values.
	Remove *bool `json:"remove,omitempty"`
	// A list of attribute names that are not included in this object. This is only used with SMConnector and will only contain \"groups\".
	Missing []string `json:"missing,omitempty"`
	// Attributes of this ResourceObject.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// In Aggregation, for sparse object the count for total accounts scanned identities updated is not incremented.
	FinalUpdate *bool `json:"finalUpdate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceObject ResourceObject

// NewResourceObject instantiates a new ResourceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceObject() *ResourceObject {
	this := ResourceObject{}
	return &this
}

// NewResourceObjectWithDefaults instantiates a new ResourceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceObjectWithDefaults() *ResourceObject {
	this := ResourceObject{}
	return &this
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ResourceObject) GetInstance() string {
	if o == nil || isNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObject) GetInstanceOk() (*string, bool) {
	if o == nil || isNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ResourceObject) HasInstance() bool {
	if o != nil && !isNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *ResourceObject) SetInstance(v string) {
	o.Instance = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *ResourceObject) GetIdentity() string {
	if o == nil || isNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObject) GetIdentityOk() (*string, bool) {
	if o == nil || isNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *ResourceObject) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *ResourceObject) SetIdentity(v string) {
	o.Identity = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ResourceObject) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObject) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ResourceObject) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ResourceObject) SetUuid(v string) {
	o.Uuid = &v
}

// GetPreviousIdentity returns the PreviousIdentity field value if set, zero value otherwise.
func (o *ResourceObject) GetPreviousIdentity() string {
	if o == nil || isNil(o.PreviousIdentity) {
		var ret string
		return ret
	}
	return *o.PreviousIdentity
}

// GetPreviousIdentityOk returns a tuple with the PreviousIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObject) GetPreviousIdentityOk() (*string, bool) {
	if o == nil || isNil(o.PreviousIdentity) {
		return nil, false
	}
	return o.PreviousIdentity, true
}

// HasPreviousIdentity returns a boolean if a field has been set.
func (o *ResourceObject) HasPreviousIdentity() bool {
	if o != nil && !isNil(o.PreviousIdentity) {
		return true
	}

	return false
}

// SetPreviousIdentity gets a reference to the given string and assigns it to the PreviousIdentity field.
func (o *ResourceObject) SetPreviousIdentity(v string) {
	o.PreviousIdentity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceObject) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObject) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceObject) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceObject) SetName(v string) {
	o.Name = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *ResourceObject) GetObjectType() string {
	if o == nil || isNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObject) GetObjectTypeOk() (*string, bool) {
	if o == nil || isNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *ResourceObject) HasObjectType() bool {
	if o != nil && !isNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *ResourceObject) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetIncomplete returns the Incomplete field value if set, zero value otherwise.
func (o *ResourceObject) GetIncomplete() bool {
	if o == nil || isNil(o.Incomplete) {
		var ret bool
		return ret
	}
	return *o.Incomplete
}

// GetIncompleteOk returns a tuple with the Incomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObject) GetIncompleteOk() (*bool, bool) {
	if o == nil || isNil(o.Incomplete) {
		return nil, false
	}
	return o.Incomplete, true
}

// HasIncomplete returns a boolean if a field has been set.
func (o *ResourceObject) HasIncomplete() bool {
	if o != nil && !isNil(o.Incomplete) {
		return true
	}

	return false
}

// SetIncomplete gets a reference to the given bool and assigns it to the Incomplete field.
func (o *ResourceObject) SetIncomplete(v bool) {
	o.Incomplete = &v
}

// GetIncremental returns the Incremental field value if set, zero value otherwise.
func (o *ResourceObject) GetIncremental() bool {
	if o == nil || isNil(o.Incremental) {
		var ret bool
		return ret
	}
	return *o.Incremental
}

// GetIncrementalOk returns a tuple with the Incremental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObject) GetIncrementalOk() (*bool, bool) {
	if o == nil || isNil(o.Incremental) {
		return nil, false
	}
	return o.Incremental, true
}

// HasIncremental returns a boolean if a field has been set.
func (o *ResourceObject) HasIncremental() bool {
	if o != nil && !isNil(o.Incremental) {
		return true
	}

	return false
}

// SetIncremental gets a reference to the given bool and assigns it to the Incremental field.
func (o *ResourceObject) SetIncremental(v bool) {
	o.Incremental = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *ResourceObject) GetDelete() bool {
	if o == nil || isNil(o.Delete) {
		var ret bool
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObject) GetDeleteOk() (*bool, bool) {
	if o == nil || isNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *ResourceObject) HasDelete() bool {
	if o != nil && !isNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given bool and assigns it to the Delete field.
func (o *ResourceObject) SetDelete(v bool) {
	o.Delete = &v
}

// GetRemove returns the Remove field value if set, zero value otherwise.
func (o *ResourceObject) GetRemove() bool {
	if o == nil || isNil(o.Remove) {
		var ret bool
		return ret
	}
	return *o.Remove
}

// GetRemoveOk returns a tuple with the Remove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObject) GetRemoveOk() (*bool, bool) {
	if o == nil || isNil(o.Remove) {
		return nil, false
	}
	return o.Remove, true
}

// HasRemove returns a boolean if a field has been set.
func (o *ResourceObject) HasRemove() bool {
	if o != nil && !isNil(o.Remove) {
		return true
	}

	return false
}

// SetRemove gets a reference to the given bool and assigns it to the Remove field.
func (o *ResourceObject) SetRemove(v bool) {
	o.Remove = &v
}

// GetMissing returns the Missing field value if set, zero value otherwise.
func (o *ResourceObject) GetMissing() []string {
	if o == nil || isNil(o.Missing) {
		var ret []string
		return ret
	}
	return o.Missing
}

// GetMissingOk returns a tuple with the Missing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObject) GetMissingOk() ([]string, bool) {
	if o == nil || isNil(o.Missing) {
		return nil, false
	}
	return o.Missing, true
}

// HasMissing returns a boolean if a field has been set.
func (o *ResourceObject) HasMissing() bool {
	if o != nil && !isNil(o.Missing) {
		return true
	}

	return false
}

// SetMissing gets a reference to the given []string and assigns it to the Missing field.
func (o *ResourceObject) SetMissing(v []string) {
	o.Missing = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ResourceObject) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObject) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ResourceObject) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *ResourceObject) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetFinalUpdate returns the FinalUpdate field value if set, zero value otherwise.
func (o *ResourceObject) GetFinalUpdate() bool {
	if o == nil || isNil(o.FinalUpdate) {
		var ret bool
		return ret
	}
	return *o.FinalUpdate
}

// GetFinalUpdateOk returns a tuple with the FinalUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObject) GetFinalUpdateOk() (*bool, bool) {
	if o == nil || isNil(o.FinalUpdate) {
		return nil, false
	}
	return o.FinalUpdate, true
}

// HasFinalUpdate returns a boolean if a field has been set.
func (o *ResourceObject) HasFinalUpdate() bool {
	if o != nil && !isNil(o.FinalUpdate) {
		return true
	}

	return false
}

// SetFinalUpdate gets a reference to the given bool and assigns it to the FinalUpdate field.
func (o *ResourceObject) SetFinalUpdate(v bool) {
	o.FinalUpdate = &v
}

func (o ResourceObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !isNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !isNil(o.PreviousIdentity) {
		toSerialize["previousIdentity"] = o.PreviousIdentity
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !isNil(o.Incomplete) {
		toSerialize["incomplete"] = o.Incomplete
	}
	if !isNil(o.Incremental) {
		toSerialize["incremental"] = o.Incremental
	}
	if !isNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	if !isNil(o.Remove) {
		toSerialize["remove"] = o.Remove
	}
	if !isNil(o.Missing) {
		toSerialize["missing"] = o.Missing
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.FinalUpdate) {
		toSerialize["finalUpdate"] = o.FinalUpdate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceObject) UnmarshalJSON(bytes []byte) (err error) {
	varResourceObject := _ResourceObject{}

	if err = json.Unmarshal(bytes, &varResourceObject); err == nil {
		*o = ResourceObject(varResourceObject)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "instance")
		delete(additionalProperties, "identity")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "previousIdentity")
		delete(additionalProperties, "name")
		delete(additionalProperties, "objectType")
		delete(additionalProperties, "incomplete")
		delete(additionalProperties, "incremental")
		delete(additionalProperties, "delete")
		delete(additionalProperties, "remove")
		delete(additionalProperties, "missing")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "finalUpdate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceObject struct {
	value *ResourceObject
	isSet bool
}

func (v NullableResourceObject) Get() *ResourceObject {
	return v.value
}

func (v *NullableResourceObject) Set(val *ResourceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceObject(val *ResourceObject) *NullableResourceObject {
	return &NullableResourceObject{value: val, isSet: true}
}

func (v NullableResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


