/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"fmt"
)

// SearchDocument - struct for SearchDocument
type SearchDocument struct {
	AccessProfile *AccessProfile
	Account1 *Account1
	AccountActivity1 *AccountActivity1
	Aggregation *Aggregation
	Entitlement *Entitlement
	Event *Event
	Identity *Identity
	Role *Role
}

// AccessProfileAsSearchDocument is a convenience function that returns AccessProfile wrapped in SearchDocument
func AccessProfileAsSearchDocument(v *AccessProfile) SearchDocument {
	return SearchDocument{
		AccessProfile: v,
	}
}

// Account1AsSearchDocument is a convenience function that returns Account1 wrapped in SearchDocument
func Account1AsSearchDocument(v *Account1) SearchDocument {
	return SearchDocument{
		Account1: v,
	}
}

// AccountActivity1AsSearchDocument is a convenience function that returns AccountActivity1 wrapped in SearchDocument
func AccountActivity1AsSearchDocument(v *AccountActivity1) SearchDocument {
	return SearchDocument{
		AccountActivity1: v,
	}
}

// AggregationAsSearchDocument is a convenience function that returns Aggregation wrapped in SearchDocument
func AggregationAsSearchDocument(v *Aggregation) SearchDocument {
	return SearchDocument{
		Aggregation: v,
	}
}

// EntitlementAsSearchDocument is a convenience function that returns Entitlement wrapped in SearchDocument
func EntitlementAsSearchDocument(v *Entitlement) SearchDocument {
	return SearchDocument{
		Entitlement: v,
	}
}

// EventAsSearchDocument is a convenience function that returns Event wrapped in SearchDocument
func EventAsSearchDocument(v *Event) SearchDocument {
	return SearchDocument{
		Event: v,
	}
}

// IdentityAsSearchDocument is a convenience function that returns Identity wrapped in SearchDocument
func IdentityAsSearchDocument(v *Identity) SearchDocument {
	return SearchDocument{
		Identity: v,
	}
}

// RoleAsSearchDocument is a convenience function that returns Role wrapped in SearchDocument
func RoleAsSearchDocument(v *Role) SearchDocument {
	return SearchDocument{
		Role: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SearchDocument) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AccessProfile
	err = newStrictDecoder(data).Decode(&dst.AccessProfile)
	if err == nil {
		jsonAccessProfile, _ := json.Marshal(dst.AccessProfile)
		if string(jsonAccessProfile) == "{}" { // empty struct
			dst.AccessProfile = nil
		} else {
			match++
		}
	} else {
		dst.AccessProfile = nil
	}

	// try to unmarshal data into Account1
	err = newStrictDecoder(data).Decode(&dst.Account1)
	if err == nil {
		jsonAccount1, _ := json.Marshal(dst.Account1)
		if string(jsonAccount1) == "{}" { // empty struct
			dst.Account1 = nil
		} else {
			match++
		}
	} else {
		dst.Account1 = nil
	}

	// try to unmarshal data into AccountActivity1
	err = newStrictDecoder(data).Decode(&dst.AccountActivity1)
	if err == nil {
		jsonAccountActivity1, _ := json.Marshal(dst.AccountActivity1)
		if string(jsonAccountActivity1) == "{}" { // empty struct
			dst.AccountActivity1 = nil
		} else {
			match++
		}
	} else {
		dst.AccountActivity1 = nil
	}

	// try to unmarshal data into Aggregation
	err = newStrictDecoder(data).Decode(&dst.Aggregation)
	if err == nil {
		jsonAggregation, _ := json.Marshal(dst.Aggregation)
		if string(jsonAggregation) == "{}" { // empty struct
			dst.Aggregation = nil
		} else {
			match++
		}
	} else {
		dst.Aggregation = nil
	}

	// try to unmarshal data into Entitlement
	err = newStrictDecoder(data).Decode(&dst.Entitlement)
	if err == nil {
		jsonEntitlement, _ := json.Marshal(dst.Entitlement)
		if string(jsonEntitlement) == "{}" { // empty struct
			dst.Entitlement = nil
		} else {
			match++
		}
	} else {
		dst.Entitlement = nil
	}

	// try to unmarshal data into Event
	err = newStrictDecoder(data).Decode(&dst.Event)
	if err == nil {
		jsonEvent, _ := json.Marshal(dst.Event)
		if string(jsonEvent) == "{}" { // empty struct
			dst.Event = nil
		} else {
			match++
		}
	} else {
		dst.Event = nil
	}

	// try to unmarshal data into Identity
	err = newStrictDecoder(data).Decode(&dst.Identity)
	if err == nil {
		jsonIdentity, _ := json.Marshal(dst.Identity)
		if string(jsonIdentity) == "{}" { // empty struct
			dst.Identity = nil
		} else {
			match++
		}
	} else {
		dst.Identity = nil
	}

	// try to unmarshal data into Role
	err = newStrictDecoder(data).Decode(&dst.Role)
	if err == nil {
		jsonRole, _ := json.Marshal(dst.Role)
		if string(jsonRole) == "{}" { // empty struct
			dst.Role = nil
		} else {
			match++
		}
	} else {
		dst.Role = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AccessProfile = nil
		dst.Account1 = nil
		dst.AccountActivity1 = nil
		dst.Aggregation = nil
		dst.Entitlement = nil
		dst.Event = nil
		dst.Identity = nil
		dst.Role = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SearchDocument)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SearchDocument)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SearchDocument) MarshalJSON() ([]byte, error) {
	if src.AccessProfile != nil {
		return json.Marshal(&src.AccessProfile)
	}

	if src.Account1 != nil {
		return json.Marshal(&src.Account1)
	}

	if src.AccountActivity1 != nil {
		return json.Marshal(&src.AccountActivity1)
	}

	if src.Aggregation != nil {
		return json.Marshal(&src.Aggregation)
	}

	if src.Entitlement != nil {
		return json.Marshal(&src.Entitlement)
	}

	if src.Event != nil {
		return json.Marshal(&src.Event)
	}

	if src.Identity != nil {
		return json.Marshal(&src.Identity)
	}

	if src.Role != nil {
		return json.Marshal(&src.Role)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SearchDocument) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AccessProfile != nil {
		return obj.AccessProfile
	}

	if obj.Account1 != nil {
		return obj.Account1
	}

	if obj.AccountActivity1 != nil {
		return obj.AccountActivity1
	}

	if obj.Aggregation != nil {
		return obj.Aggregation
	}

	if obj.Entitlement != nil {
		return obj.Entitlement
	}

	if obj.Event != nil {
		return obj.Event
	}

	if obj.Identity != nil {
		return obj.Identity
	}

	if obj.Role != nil {
		return obj.Role
	}

	// all schemas are nil
	return nil
}

type NullableSearchDocument struct {
	value *SearchDocument
	isSet bool
}

func (v NullableSearchDocument) Get() *SearchDocument {
	return v.value
}

func (v *NullableSearchDocument) Set(val *SearchDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchDocument(val *SearchDocument) *NullableSearchDocument {
	return &NullableSearchDocument{value: val, isSet: true}
}

func (v NullableSearchDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


