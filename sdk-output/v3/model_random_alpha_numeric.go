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

// RandomAlphaNumeric struct for RandomAlphaNumeric
type RandomAlphaNumeric struct {
	// This is an integer value specifying the size/number of characters the random string must contain   * This value must be a positive number and cannot be blank   * If no length is provided, the transform will default to a value of `32`   * Due to identity attribute data constraints, the maximum allowable value is `450` characters 
	Length *string `json:"length,omitempty"`
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	// This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI.
	Input map[string]interface{} `json:"input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RandomAlphaNumeric RandomAlphaNumeric

// NewRandomAlphaNumeric instantiates a new RandomAlphaNumeric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRandomAlphaNumeric() *RandomAlphaNumeric {
	this := RandomAlphaNumeric{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// NewRandomAlphaNumericWithDefaults instantiates a new RandomAlphaNumeric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRandomAlphaNumericWithDefaults() *RandomAlphaNumeric {
	this := RandomAlphaNumeric{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *RandomAlphaNumeric) GetLength() string {
	if o == nil || isNil(o.Length) {
		var ret string
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RandomAlphaNumeric) GetLengthOk() (*string, bool) {
	if o == nil || isNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *RandomAlphaNumeric) HasLength() bool {
	if o != nil && !isNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given string and assigns it to the Length field.
func (o *RandomAlphaNumeric) SetLength(v string) {
	o.Length = &v
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *RandomAlphaNumeric) GetRequiresPeriodicRefresh() bool {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RandomAlphaNumeric) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *RandomAlphaNumeric) HasRequiresPeriodicRefresh() bool {
	if o != nil && !isNil(o.RequiresPeriodicRefresh) {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *RandomAlphaNumeric) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *RandomAlphaNumeric) GetInput() map[string]interface{} {
	if o == nil || isNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RandomAlphaNumeric) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *RandomAlphaNumeric) HasInput() bool {
	if o != nil && !isNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *RandomAlphaNumeric) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o RandomAlphaNumeric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !isNil(o.RequiresPeriodicRefresh) {
		toSerialize["requiresPeriodicRefresh"] = o.RequiresPeriodicRefresh
	}
	if !isNil(o.Input) {
		toSerialize["input"] = o.Input
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RandomAlphaNumeric) UnmarshalJSON(bytes []byte) (err error) {
	varRandomAlphaNumeric := _RandomAlphaNumeric{}

	if err = json.Unmarshal(bytes, &varRandomAlphaNumeric); err == nil {
		*o = RandomAlphaNumeric(varRandomAlphaNumeric)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "length")
		delete(additionalProperties, "requiresPeriodicRefresh")
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRandomAlphaNumeric struct {
	value *RandomAlphaNumeric
	isSet bool
}

func (v NullableRandomAlphaNumeric) Get() *RandomAlphaNumeric {
	return v.value
}

func (v *NullableRandomAlphaNumeric) Set(val *RandomAlphaNumeric) {
	v.value = val
	v.isSet = true
}

func (v NullableRandomAlphaNumeric) IsSet() bool {
	return v.isSet
}

func (v *NullableRandomAlphaNumeric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRandomAlphaNumeric(val *RandomAlphaNumeric) *NullableRandomAlphaNumeric {
	return &NullableRandomAlphaNumeric{value: val, isSet: true}
}

func (v NullableRandomAlphaNumeric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRandomAlphaNumeric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

