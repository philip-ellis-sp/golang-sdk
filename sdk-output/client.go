/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpoint

import (

	"net/http"
	"regexp"

	sailpointsdk "github.com/philip-ellis-sp/golang-sdk/sdk-output/v3"
	sailpointbetasdk "github.com/philip-ellis-sp/golang-sdk/sdk-output/beta"
	sailpointccsdk "github.com/philip-ellis-sp/golang-sdk/sdk-output/cc"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// APIClient manages communication with the IdentityNow V3 API API v3.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *sailpointsdk.Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	V3 *sailpointsdk.APIClient
	Beta *sailpointbetasdk.APIClient
	CC *sailpointccsdk.APIClient
	token string
}

type service struct {
	client *sailpointsdk.APIClient
	betaClient *sailpointbetasdk.APIClient
	ccClient *sailpointccsdk.APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}

	
	c.V3 = sailpointsdk.NewAPIClient(sailpointsdk.NewConfiguration(cfg.Tenant))
	c.Beta = sailpointbetasdk.NewAPIClient(sailpointbetasdk.NewConfiguration(cfg.Tenant))
	c.CC = sailpointccsdk.NewAPIClient(sailpointccsdk.NewConfiguration(cfg.Tenant))

	// API Services


	return c
}
