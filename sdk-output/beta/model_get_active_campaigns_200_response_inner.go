/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"fmt"
)

// GetActiveCampaigns200ResponseInner - struct for GetActiveCampaigns200ResponseInner
type GetActiveCampaigns200ResponseInner struct {
	Campaign *Campaign
	Slimcampaign *Slimcampaign
}

// CampaignAsGetActiveCampaigns200ResponseInner is a convenience function that returns Campaign wrapped in GetActiveCampaigns200ResponseInner
func CampaignAsGetActiveCampaigns200ResponseInner(v *Campaign) GetActiveCampaigns200ResponseInner {
	return GetActiveCampaigns200ResponseInner{
		Campaign: v,
	}
}

// SlimcampaignAsGetActiveCampaigns200ResponseInner is a convenience function that returns Slimcampaign wrapped in GetActiveCampaigns200ResponseInner
func SlimcampaignAsGetActiveCampaigns200ResponseInner(v *Slimcampaign) GetActiveCampaigns200ResponseInner {
	return GetActiveCampaigns200ResponseInner{
		Slimcampaign: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetActiveCampaigns200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Campaign
	err = newStrictDecoder(data).Decode(&dst.Campaign)
	if err == nil {
		jsonCampaign, _ := json.Marshal(dst.Campaign)
		if string(jsonCampaign) == "{}" { // empty struct
			dst.Campaign = nil
		} else {
			match++
		}
	} else {
		dst.Campaign = nil
	}

	// try to unmarshal data into Slimcampaign
	err = newStrictDecoder(data).Decode(&dst.Slimcampaign)
	if err == nil {
		jsonSlimcampaign, _ := json.Marshal(dst.Slimcampaign)
		if string(jsonSlimcampaign) == "{}" { // empty struct
			dst.Slimcampaign = nil
		} else {
			match++
		}
	} else {
		dst.Slimcampaign = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Campaign = nil
		dst.Slimcampaign = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetActiveCampaigns200ResponseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetActiveCampaigns200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetActiveCampaigns200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.Campaign != nil {
		return json.Marshal(&src.Campaign)
	}

	if src.Slimcampaign != nil {
		return json.Marshal(&src.Slimcampaign)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetActiveCampaigns200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Campaign != nil {
		return obj.Campaign
	}

	if obj.Slimcampaign != nil {
		return obj.Slimcampaign
	}

	// all schemas are nil
	return nil
}

type NullableGetActiveCampaigns200ResponseInner struct {
	value *GetActiveCampaigns200ResponseInner
	isSet bool
}

func (v NullableGetActiveCampaigns200ResponseInner) Get() *GetActiveCampaigns200ResponseInner {
	return v.value
}

func (v *NullableGetActiveCampaigns200ResponseInner) Set(val *GetActiveCampaigns200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetActiveCampaigns200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetActiveCampaigns200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetActiveCampaigns200ResponseInner(val *GetActiveCampaigns200ResponseInner) *NullableGetActiveCampaigns200ResponseInner {
	return &NullableGetActiveCampaigns200ResponseInner{value: val, isSet: true}
}

func (v NullableGetActiveCampaigns200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetActiveCampaigns200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


