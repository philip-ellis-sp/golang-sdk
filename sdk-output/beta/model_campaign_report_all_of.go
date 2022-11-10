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

// CampaignReportAllOf struct for CampaignReportAllOf
type CampaignReportAllOf struct {
	ReportType *ReportType `json:"reportType,omitempty"`
	// The most recent date and time this report was run
	LastRunAt *time.Time `json:"lastRunAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignReportAllOf CampaignReportAllOf

// NewCampaignReportAllOf instantiates a new CampaignReportAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignReportAllOf() *CampaignReportAllOf {
	this := CampaignReportAllOf{}
	return &this
}

// NewCampaignReportAllOfWithDefaults instantiates a new CampaignReportAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignReportAllOfWithDefaults() *CampaignReportAllOf {
	this := CampaignReportAllOf{}
	return &this
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *CampaignReportAllOf) GetReportType() ReportType {
	if o == nil || isNil(o.ReportType) {
		var ret ReportType
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportAllOf) GetReportTypeOk() (*ReportType, bool) {
	if o == nil || isNil(o.ReportType) {
		return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *CampaignReportAllOf) HasReportType() bool {
	if o != nil && !isNil(o.ReportType) {
		return true
	}

	return false
}

// SetReportType gets a reference to the given ReportType and assigns it to the ReportType field.
func (o *CampaignReportAllOf) SetReportType(v ReportType) {
	o.ReportType = &v
}

// GetLastRunAt returns the LastRunAt field value if set, zero value otherwise.
func (o *CampaignReportAllOf) GetLastRunAt() time.Time {
	if o == nil || isNil(o.LastRunAt) {
		var ret time.Time
		return ret
	}
	return *o.LastRunAt
}

// GetLastRunAtOk returns a tuple with the LastRunAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportAllOf) GetLastRunAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastRunAt) {
		return nil, false
	}
	return o.LastRunAt, true
}

// HasLastRunAt returns a boolean if a field has been set.
func (o *CampaignReportAllOf) HasLastRunAt() bool {
	if o != nil && !isNil(o.LastRunAt) {
		return true
	}

	return false
}

// SetLastRunAt gets a reference to the given time.Time and assigns it to the LastRunAt field.
func (o *CampaignReportAllOf) SetLastRunAt(v time.Time) {
	o.LastRunAt = &v
}

func (o CampaignReportAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReportType) {
		toSerialize["reportType"] = o.ReportType
	}
	if !isNil(o.LastRunAt) {
		toSerialize["lastRunAt"] = o.LastRunAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CampaignReportAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignReportAllOf := _CampaignReportAllOf{}

	if err = json.Unmarshal(bytes, &varCampaignReportAllOf); err == nil {
		*o = CampaignReportAllOf(varCampaignReportAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reportType")
		delete(additionalProperties, "lastRunAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignReportAllOf struct {
	value *CampaignReportAllOf
	isSet bool
}

func (v NullableCampaignReportAllOf) Get() *CampaignReportAllOf {
	return v.value
}

func (v *NullableCampaignReportAllOf) Set(val *CampaignReportAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignReportAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignReportAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignReportAllOf(val *CampaignReportAllOf) *NullableCampaignReportAllOf {
	return &NullableCampaignReportAllOf{value: val, isSet: true}
}

func (v NullableCampaignReportAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignReportAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

