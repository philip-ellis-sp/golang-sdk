# AccessRequestPostApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRequestId** | **string** | The unique ID of the access request. | 
**RequestedFor** | [**TriggerInputAccessRequestPostApprovalRequestedFor**](TriggerInputAccessRequestPostApprovalRequestedFor.md) |  | 
**RequestedItemsStatus** | [**[]TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner**](TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner.md) | Details on the outcome of each access item. | 
**RequestedBy** | [**TriggerInputAccessRequestPostApprovalRequestedBy**](TriggerInputAccessRequestPostApprovalRequestedBy.md) |  | 

## Methods

### NewAccessRequestPostApproval

`func NewAccessRequestPostApproval(accessRequestId string, requestedFor TriggerInputAccessRequestPostApprovalRequestedFor, requestedItemsStatus []TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner, requestedBy TriggerInputAccessRequestPostApprovalRequestedBy, ) *AccessRequestPostApproval`

NewAccessRequestPostApproval instantiates a new AccessRequestPostApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestPostApprovalWithDefaults

`func NewAccessRequestPostApprovalWithDefaults() *AccessRequestPostApproval`

NewAccessRequestPostApprovalWithDefaults instantiates a new AccessRequestPostApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRequestId

`func (o *AccessRequestPostApproval) GetAccessRequestId() string`

GetAccessRequestId returns the AccessRequestId field if non-nil, zero value otherwise.

### GetAccessRequestIdOk

`func (o *AccessRequestPostApproval) GetAccessRequestIdOk() (*string, bool)`

GetAccessRequestIdOk returns a tuple with the AccessRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestId

`func (o *AccessRequestPostApproval) SetAccessRequestId(v string)`

SetAccessRequestId sets AccessRequestId field to given value.


### GetRequestedFor

`func (o *AccessRequestPostApproval) GetRequestedFor() TriggerInputAccessRequestPostApprovalRequestedFor`

GetRequestedFor returns the RequestedFor field if non-nil, zero value otherwise.

### GetRequestedForOk

`func (o *AccessRequestPostApproval) GetRequestedForOk() (*TriggerInputAccessRequestPostApprovalRequestedFor, bool)`

GetRequestedForOk returns a tuple with the RequestedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFor

`func (o *AccessRequestPostApproval) SetRequestedFor(v TriggerInputAccessRequestPostApprovalRequestedFor)`

SetRequestedFor sets RequestedFor field to given value.


### GetRequestedItemsStatus

`func (o *AccessRequestPostApproval) GetRequestedItemsStatus() []TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner`

GetRequestedItemsStatus returns the RequestedItemsStatus field if non-nil, zero value otherwise.

### GetRequestedItemsStatusOk

`func (o *AccessRequestPostApproval) GetRequestedItemsStatusOk() (*[]TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner, bool)`

GetRequestedItemsStatusOk returns a tuple with the RequestedItemsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedItemsStatus

`func (o *AccessRequestPostApproval) SetRequestedItemsStatus(v []TriggerInputAccessRequestPostApprovalRequestedItemsStatusInner)`

SetRequestedItemsStatus sets RequestedItemsStatus field to given value.


### GetRequestedBy

`func (o *AccessRequestPostApproval) GetRequestedBy() TriggerInputAccessRequestPostApprovalRequestedBy`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *AccessRequestPostApproval) GetRequestedByOk() (*TriggerInputAccessRequestPostApprovalRequestedBy, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *AccessRequestPostApproval) SetRequestedBy(v TriggerInputAccessRequestPostApprovalRequestedBy)`

SetRequestedBy sets RequestedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


