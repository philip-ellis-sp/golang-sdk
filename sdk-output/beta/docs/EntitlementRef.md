# EntitlementRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the Entitlement | [optional] 
**Type** | Pointer to **string** | The type of the Entitlement, will always be ENTITLEMENT | [optional] 
**Name** | Pointer to **string** | The display name of the Entitlement | [optional] 

## Methods

### NewEntitlementRef

`func NewEntitlementRef() *EntitlementRef`

NewEntitlementRef instantiates a new EntitlementRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementRefWithDefaults

`func NewEntitlementRefWithDefaults() *EntitlementRef`

NewEntitlementRefWithDefaults instantiates a new EntitlementRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementRef) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementRef) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *EntitlementRef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitlementRef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitlementRef) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntitlementRef) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *EntitlementRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntitlementRef) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


