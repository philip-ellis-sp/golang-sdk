/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"time"
	"encoding/json"
)

// Campaign struct for Campaign
type Campaign struct {
	// Id of the campaign
	Id *string `json:"id,omitempty"`
	// The campaign name. If this object is part of a template, special formatting applies; see the `/campaign-templates/{id}/generate` endpoint documentation for details.
	Name string `json:"name"`
	// The campaign description. If this object is part of a template, special formatting applies; see the `/campaign-templates/{id}/generate` endpoint documentation for details.
	Description string `json:"description"`
	// The campaign's completion deadline.
	Deadline *time.Time `json:"deadline,omitempty"`
	// The type of campaign. Could be extended in the future.
	Type string `json:"type"`
	// Enables email notification for this campaign
	EmailNotificationEnabled *bool `json:"emailNotificationEnabled,omitempty"`
	// Allows auto revoke for this campaign
	AutoRevokeAllowed *bool `json:"autoRevokeAllowed,omitempty"`
	// Enables IAI for this campaign. Accepts true even if the IAI product feature is off. If IAI is turned off then campaigns generated from this template will indicate false. The real value will then be returned if IAI is ever enabled for the org in the future.
	RecommendationsEnabled *bool `json:"recommendationsEnabled,omitempty"`
	// The campaign's current status.
	Status *string `json:"status,omitempty"`
	// The correlatedStatus of the campaign. Only SOURCE_OWNER campaigns can be Uncorrelated. An Uncorrelated certification campaign only includes Uncorrelated identities (An identity is uncorrelated if it has no accounts on an authoritative source).
	CorrelatedStatus map[string]interface{} `json:"correlatedStatus,omitempty"`
	// Created time of the campaign
	Created *time.Time `json:"created,omitempty"`
	// Modified time of the campaign
	Modified *time.Time `json:"modified,omitempty"`
	Filter *FullcampaignAllOfFilter `json:"filter,omitempty"`
	// Determines if comments on sunset date changes are required.
	SunsetCommentsRequired *bool `json:"sunsetCommentsRequired,omitempty"`
	SourceOwnerCampaignInfo *FullcampaignAllOfSourceOwnerCampaignInfo `json:"sourceOwnerCampaignInfo,omitempty"`
	SearchCampaignInfo *FullcampaignAllOfSearchCampaignInfo `json:"searchCampaignInfo,omitempty"`
	RoleCompositionCampaignInfo *FullcampaignAllOfRoleCompositionCampaignInfo `json:"roleCompositionCampaignInfo,omitempty"`
	// A list of errors and warnings that have accumulated.
	Alerts []CampaignAlert `json:"alerts,omitempty"`
	// The total number of certifications in this campaign.
	TotalCertifications *int32 `json:"totalCertifications,omitempty"`
	// The number of completed certifications in this campaign.
	CompletedCertifications *int32 `json:"completedCertifications,omitempty"`
	// A list of sources in the campaign that contain \\\"orphan entitlements\\\" (entitlements without a corresponding Managed Attribute). An empty list indicates the campaign has no orphan entitlements. Null indicates there may be unknown orphan entitlements in the campaign (the campaign was created before this feature was implemented).
	SourcesWithOrphanEntitlements []FullcampaignAllOfSourcesWithOrphanEntitlements `json:"sourcesWithOrphanEntitlements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Campaign Campaign

// NewCampaign instantiates a new Campaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaign(name string, description string, type_ string) *Campaign {
	this := Campaign{}
	this.Name = name
	this.Description = description
	this.Type = type_
	var sunsetCommentsRequired bool = true
	this.SunsetCommentsRequired = &sunsetCommentsRequired
	return &this
}

// NewCampaignWithDefaults instantiates a new Campaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignWithDefaults() *Campaign {
	this := Campaign{}
	var sunsetCommentsRequired bool = true
	this.SunsetCommentsRequired = &sunsetCommentsRequired
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Campaign) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Campaign) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Campaign) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Campaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Campaign) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Campaign) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Campaign) SetDescription(v string) {
	o.Description = v
}

// GetDeadline returns the Deadline field value if set, zero value otherwise.
func (o *Campaign) GetDeadline() time.Time {
	if o == nil || isNil(o.Deadline) {
		var ret time.Time
		return ret
	}
	return *o.Deadline
}

// GetDeadlineOk returns a tuple with the Deadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetDeadlineOk() (*time.Time, bool) {
	if o == nil || isNil(o.Deadline) {
		return nil, false
	}
	return o.Deadline, true
}

// HasDeadline returns a boolean if a field has been set.
func (o *Campaign) HasDeadline() bool {
	if o != nil && !isNil(o.Deadline) {
		return true
	}

	return false
}

// SetDeadline gets a reference to the given time.Time and assigns it to the Deadline field.
func (o *Campaign) SetDeadline(v time.Time) {
	o.Deadline = &v
}

// GetType returns the Type field value
func (o *Campaign) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Campaign) SetType(v string) {
	o.Type = v
}

// GetEmailNotificationEnabled returns the EmailNotificationEnabled field value if set, zero value otherwise.
func (o *Campaign) GetEmailNotificationEnabled() bool {
	if o == nil || isNil(o.EmailNotificationEnabled) {
		var ret bool
		return ret
	}
	return *o.EmailNotificationEnabled
}

// GetEmailNotificationEnabledOk returns a tuple with the EmailNotificationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetEmailNotificationEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.EmailNotificationEnabled) {
		return nil, false
	}
	return o.EmailNotificationEnabled, true
}

// HasEmailNotificationEnabled returns a boolean if a field has been set.
func (o *Campaign) HasEmailNotificationEnabled() bool {
	if o != nil && !isNil(o.EmailNotificationEnabled) {
		return true
	}

	return false
}

// SetEmailNotificationEnabled gets a reference to the given bool and assigns it to the EmailNotificationEnabled field.
func (o *Campaign) SetEmailNotificationEnabled(v bool) {
	o.EmailNotificationEnabled = &v
}

// GetAutoRevokeAllowed returns the AutoRevokeAllowed field value if set, zero value otherwise.
func (o *Campaign) GetAutoRevokeAllowed() bool {
	if o == nil || isNil(o.AutoRevokeAllowed) {
		var ret bool
		return ret
	}
	return *o.AutoRevokeAllowed
}

// GetAutoRevokeAllowedOk returns a tuple with the AutoRevokeAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetAutoRevokeAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.AutoRevokeAllowed) {
		return nil, false
	}
	return o.AutoRevokeAllowed, true
}

// HasAutoRevokeAllowed returns a boolean if a field has been set.
func (o *Campaign) HasAutoRevokeAllowed() bool {
	if o != nil && !isNil(o.AutoRevokeAllowed) {
		return true
	}

	return false
}

// SetAutoRevokeAllowed gets a reference to the given bool and assigns it to the AutoRevokeAllowed field.
func (o *Campaign) SetAutoRevokeAllowed(v bool) {
	o.AutoRevokeAllowed = &v
}

// GetRecommendationsEnabled returns the RecommendationsEnabled field value if set, zero value otherwise.
func (o *Campaign) GetRecommendationsEnabled() bool {
	if o == nil || isNil(o.RecommendationsEnabled) {
		var ret bool
		return ret
	}
	return *o.RecommendationsEnabled
}

// GetRecommendationsEnabledOk returns a tuple with the RecommendationsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetRecommendationsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RecommendationsEnabled) {
		return nil, false
	}
	return o.RecommendationsEnabled, true
}

// HasRecommendationsEnabled returns a boolean if a field has been set.
func (o *Campaign) HasRecommendationsEnabled() bool {
	if o != nil && !isNil(o.RecommendationsEnabled) {
		return true
	}

	return false
}

// SetRecommendationsEnabled gets a reference to the given bool and assigns it to the RecommendationsEnabled field.
func (o *Campaign) SetRecommendationsEnabled(v bool) {
	o.RecommendationsEnabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Campaign) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Campaign) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Campaign) SetStatus(v string) {
	o.Status = &v
}

// GetCorrelatedStatus returns the CorrelatedStatus field value if set, zero value otherwise.
func (o *Campaign) GetCorrelatedStatus() map[string]interface{} {
	if o == nil || isNil(o.CorrelatedStatus) {
		var ret map[string]interface{}
		return ret
	}
	return o.CorrelatedStatus
}

// GetCorrelatedStatusOk returns a tuple with the CorrelatedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetCorrelatedStatusOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.CorrelatedStatus) {
		return map[string]interface{}{}, false
	}
	return o.CorrelatedStatus, true
}

// HasCorrelatedStatus returns a boolean if a field has been set.
func (o *Campaign) HasCorrelatedStatus() bool {
	if o != nil && !isNil(o.CorrelatedStatus) {
		return true
	}

	return false
}

// SetCorrelatedStatus gets a reference to the given map[string]interface{} and assigns it to the CorrelatedStatus field.
func (o *Campaign) SetCorrelatedStatus(v map[string]interface{}) {
	o.CorrelatedStatus = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Campaign) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Campaign) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Campaign) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *Campaign) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *Campaign) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *Campaign) SetModified(v time.Time) {
	o.Modified = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *Campaign) GetFilter() FullcampaignAllOfFilter {
	if o == nil || isNil(o.Filter) {
		var ret FullcampaignAllOfFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetFilterOk() (*FullcampaignAllOfFilter, bool) {
	if o == nil || isNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *Campaign) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given FullcampaignAllOfFilter and assigns it to the Filter field.
func (o *Campaign) SetFilter(v FullcampaignAllOfFilter) {
	o.Filter = &v
}

// GetSunsetCommentsRequired returns the SunsetCommentsRequired field value if set, zero value otherwise.
func (o *Campaign) GetSunsetCommentsRequired() bool {
	if o == nil || isNil(o.SunsetCommentsRequired) {
		var ret bool
		return ret
	}
	return *o.SunsetCommentsRequired
}

// GetSunsetCommentsRequiredOk returns a tuple with the SunsetCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetSunsetCommentsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.SunsetCommentsRequired) {
		return nil, false
	}
	return o.SunsetCommentsRequired, true
}

// HasSunsetCommentsRequired returns a boolean if a field has been set.
func (o *Campaign) HasSunsetCommentsRequired() bool {
	if o != nil && !isNil(o.SunsetCommentsRequired) {
		return true
	}

	return false
}

// SetSunsetCommentsRequired gets a reference to the given bool and assigns it to the SunsetCommentsRequired field.
func (o *Campaign) SetSunsetCommentsRequired(v bool) {
	o.SunsetCommentsRequired = &v
}

// GetSourceOwnerCampaignInfo returns the SourceOwnerCampaignInfo field value if set, zero value otherwise.
func (o *Campaign) GetSourceOwnerCampaignInfo() FullcampaignAllOfSourceOwnerCampaignInfo {
	if o == nil || isNil(o.SourceOwnerCampaignInfo) {
		var ret FullcampaignAllOfSourceOwnerCampaignInfo
		return ret
	}
	return *o.SourceOwnerCampaignInfo
}

// GetSourceOwnerCampaignInfoOk returns a tuple with the SourceOwnerCampaignInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetSourceOwnerCampaignInfoOk() (*FullcampaignAllOfSourceOwnerCampaignInfo, bool) {
	if o == nil || isNil(o.SourceOwnerCampaignInfo) {
		return nil, false
	}
	return o.SourceOwnerCampaignInfo, true
}

// HasSourceOwnerCampaignInfo returns a boolean if a field has been set.
func (o *Campaign) HasSourceOwnerCampaignInfo() bool {
	if o != nil && !isNil(o.SourceOwnerCampaignInfo) {
		return true
	}

	return false
}

// SetSourceOwnerCampaignInfo gets a reference to the given FullcampaignAllOfSourceOwnerCampaignInfo and assigns it to the SourceOwnerCampaignInfo field.
func (o *Campaign) SetSourceOwnerCampaignInfo(v FullcampaignAllOfSourceOwnerCampaignInfo) {
	o.SourceOwnerCampaignInfo = &v
}

// GetSearchCampaignInfo returns the SearchCampaignInfo field value if set, zero value otherwise.
func (o *Campaign) GetSearchCampaignInfo() FullcampaignAllOfSearchCampaignInfo {
	if o == nil || isNil(o.SearchCampaignInfo) {
		var ret FullcampaignAllOfSearchCampaignInfo
		return ret
	}
	return *o.SearchCampaignInfo
}

// GetSearchCampaignInfoOk returns a tuple with the SearchCampaignInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetSearchCampaignInfoOk() (*FullcampaignAllOfSearchCampaignInfo, bool) {
	if o == nil || isNil(o.SearchCampaignInfo) {
		return nil, false
	}
	return o.SearchCampaignInfo, true
}

// HasSearchCampaignInfo returns a boolean if a field has been set.
func (o *Campaign) HasSearchCampaignInfo() bool {
	if o != nil && !isNil(o.SearchCampaignInfo) {
		return true
	}

	return false
}

// SetSearchCampaignInfo gets a reference to the given FullcampaignAllOfSearchCampaignInfo and assigns it to the SearchCampaignInfo field.
func (o *Campaign) SetSearchCampaignInfo(v FullcampaignAllOfSearchCampaignInfo) {
	o.SearchCampaignInfo = &v
}

// GetRoleCompositionCampaignInfo returns the RoleCompositionCampaignInfo field value if set, zero value otherwise.
func (o *Campaign) GetRoleCompositionCampaignInfo() FullcampaignAllOfRoleCompositionCampaignInfo {
	if o == nil || isNil(o.RoleCompositionCampaignInfo) {
		var ret FullcampaignAllOfRoleCompositionCampaignInfo
		return ret
	}
	return *o.RoleCompositionCampaignInfo
}

// GetRoleCompositionCampaignInfoOk returns a tuple with the RoleCompositionCampaignInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetRoleCompositionCampaignInfoOk() (*FullcampaignAllOfRoleCompositionCampaignInfo, bool) {
	if o == nil || isNil(o.RoleCompositionCampaignInfo) {
		return nil, false
	}
	return o.RoleCompositionCampaignInfo, true
}

// HasRoleCompositionCampaignInfo returns a boolean if a field has been set.
func (o *Campaign) HasRoleCompositionCampaignInfo() bool {
	if o != nil && !isNil(o.RoleCompositionCampaignInfo) {
		return true
	}

	return false
}

// SetRoleCompositionCampaignInfo gets a reference to the given FullcampaignAllOfRoleCompositionCampaignInfo and assigns it to the RoleCompositionCampaignInfo field.
func (o *Campaign) SetRoleCompositionCampaignInfo(v FullcampaignAllOfRoleCompositionCampaignInfo) {
	o.RoleCompositionCampaignInfo = &v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *Campaign) GetAlerts() []CampaignAlert {
	if o == nil || isNil(o.Alerts) {
		var ret []CampaignAlert
		return ret
	}
	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetAlertsOk() ([]CampaignAlert, bool) {
	if o == nil || isNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *Campaign) HasAlerts() bool {
	if o != nil && !isNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []CampaignAlert and assigns it to the Alerts field.
func (o *Campaign) SetAlerts(v []CampaignAlert) {
	o.Alerts = v
}

// GetTotalCertifications returns the TotalCertifications field value if set, zero value otherwise.
func (o *Campaign) GetTotalCertifications() int32 {
	if o == nil || isNil(o.TotalCertifications) {
		var ret int32
		return ret
	}
	return *o.TotalCertifications
}

// GetTotalCertificationsOk returns a tuple with the TotalCertifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetTotalCertificationsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalCertifications) {
		return nil, false
	}
	return o.TotalCertifications, true
}

// HasTotalCertifications returns a boolean if a field has been set.
func (o *Campaign) HasTotalCertifications() bool {
	if o != nil && !isNil(o.TotalCertifications) {
		return true
	}

	return false
}

// SetTotalCertifications gets a reference to the given int32 and assigns it to the TotalCertifications field.
func (o *Campaign) SetTotalCertifications(v int32) {
	o.TotalCertifications = &v
}

// GetCompletedCertifications returns the CompletedCertifications field value if set, zero value otherwise.
func (o *Campaign) GetCompletedCertifications() int32 {
	if o == nil || isNil(o.CompletedCertifications) {
		var ret int32
		return ret
	}
	return *o.CompletedCertifications
}

// GetCompletedCertificationsOk returns a tuple with the CompletedCertifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetCompletedCertificationsOk() (*int32, bool) {
	if o == nil || isNil(o.CompletedCertifications) {
		return nil, false
	}
	return o.CompletedCertifications, true
}

// HasCompletedCertifications returns a boolean if a field has been set.
func (o *Campaign) HasCompletedCertifications() bool {
	if o != nil && !isNil(o.CompletedCertifications) {
		return true
	}

	return false
}

// SetCompletedCertifications gets a reference to the given int32 and assigns it to the CompletedCertifications field.
func (o *Campaign) SetCompletedCertifications(v int32) {
	o.CompletedCertifications = &v
}

// GetSourcesWithOrphanEntitlements returns the SourcesWithOrphanEntitlements field value if set, zero value otherwise.
func (o *Campaign) GetSourcesWithOrphanEntitlements() []FullcampaignAllOfSourcesWithOrphanEntitlements {
	if o == nil || isNil(o.SourcesWithOrphanEntitlements) {
		var ret []FullcampaignAllOfSourcesWithOrphanEntitlements
		return ret
	}
	return o.SourcesWithOrphanEntitlements
}

// GetSourcesWithOrphanEntitlementsOk returns a tuple with the SourcesWithOrphanEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetSourcesWithOrphanEntitlementsOk() ([]FullcampaignAllOfSourcesWithOrphanEntitlements, bool) {
	if o == nil || isNil(o.SourcesWithOrphanEntitlements) {
		return nil, false
	}
	return o.SourcesWithOrphanEntitlements, true
}

// HasSourcesWithOrphanEntitlements returns a boolean if a field has been set.
func (o *Campaign) HasSourcesWithOrphanEntitlements() bool {
	if o != nil && !isNil(o.SourcesWithOrphanEntitlements) {
		return true
	}

	return false
}

// SetSourcesWithOrphanEntitlements gets a reference to the given []FullcampaignAllOfSourcesWithOrphanEntitlements and assigns it to the SourcesWithOrphanEntitlements field.
func (o *Campaign) SetSourcesWithOrphanEntitlements(v []FullcampaignAllOfSourcesWithOrphanEntitlements) {
	o.SourcesWithOrphanEntitlements = v
}

func (o Campaign) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Deadline) {
		toSerialize["deadline"] = o.Deadline
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.EmailNotificationEnabled) {
		toSerialize["emailNotificationEnabled"] = o.EmailNotificationEnabled
	}
	if !isNil(o.AutoRevokeAllowed) {
		toSerialize["autoRevokeAllowed"] = o.AutoRevokeAllowed
	}
	if !isNil(o.RecommendationsEnabled) {
		toSerialize["recommendationsEnabled"] = o.RecommendationsEnabled
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.CorrelatedStatus) {
		toSerialize["correlatedStatus"] = o.CorrelatedStatus
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !isNil(o.SunsetCommentsRequired) {
		toSerialize["sunsetCommentsRequired"] = o.SunsetCommentsRequired
	}
	if !isNil(o.SourceOwnerCampaignInfo) {
		toSerialize["sourceOwnerCampaignInfo"] = o.SourceOwnerCampaignInfo
	}
	if !isNil(o.SearchCampaignInfo) {
		toSerialize["searchCampaignInfo"] = o.SearchCampaignInfo
	}
	if !isNil(o.RoleCompositionCampaignInfo) {
		toSerialize["roleCompositionCampaignInfo"] = o.RoleCompositionCampaignInfo
	}
	if !isNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !isNil(o.TotalCertifications) {
		toSerialize["totalCertifications"] = o.TotalCertifications
	}
	if !isNil(o.CompletedCertifications) {
		toSerialize["completedCertifications"] = o.CompletedCertifications
	}
	if !isNil(o.SourcesWithOrphanEntitlements) {
		toSerialize["sourcesWithOrphanEntitlements"] = o.SourcesWithOrphanEntitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Campaign) UnmarshalJSON(bytes []byte) (err error) {
	varCampaign := _Campaign{}

	if err = json.Unmarshal(bytes, &varCampaign); err == nil {
		*o = Campaign(varCampaign)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "deadline")
		delete(additionalProperties, "type")
		delete(additionalProperties, "emailNotificationEnabled")
		delete(additionalProperties, "autoRevokeAllowed")
		delete(additionalProperties, "recommendationsEnabled")
		delete(additionalProperties, "status")
		delete(additionalProperties, "correlatedStatus")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "sunsetCommentsRequired")
		delete(additionalProperties, "sourceOwnerCampaignInfo")
		delete(additionalProperties, "searchCampaignInfo")
		delete(additionalProperties, "roleCompositionCampaignInfo")
		delete(additionalProperties, "alerts")
		delete(additionalProperties, "totalCertifications")
		delete(additionalProperties, "completedCertifications")
		delete(additionalProperties, "sourcesWithOrphanEntitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaign struct {
	value *Campaign
	isSet bool
}

func (v NullableCampaign) Get() *Campaign {
	return v.value
}

func (v *NullableCampaign) Set(val *Campaign) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaign(val *Campaign) *NullableCampaign {
	return &NullableCampaign{value: val, isSet: true}
}

func (v NullableCampaign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


