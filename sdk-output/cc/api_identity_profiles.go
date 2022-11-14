/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointccsdk

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// IdentityProfilesApiService IdentityProfilesApi service
type IdentityProfilesApiService service

type ApiDeleteLifecycleStateRequest struct {
	ctx context.Context
	ApiService *IdentityProfilesApiService
	id string
	xCSRFToken *string
}

func (r ApiDeleteLifecycleStateRequest) XCSRFToken(xCSRFToken string) ApiDeleteLifecycleStateRequest {
	r.xCSRFToken = &xCSRFToken
	return r
}

func (r ApiDeleteLifecycleStateRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteLifecycleStateExecute(r)
}

/*
DeleteLifecycleState LCS - Delete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiDeleteLifecycleStateRequest
*/
func (a *IdentityProfilesApiService) DeleteLifecycleState(ctx context.Context, id string) ApiDeleteLifecycleStateRequest {
	return ApiDeleteLifecycleStateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *IdentityProfilesApiService) DeleteLifecycleStateExecute(r ApiDeleteLifecycleStateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProfilesApiService.DeleteLifecycleState")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/lifecycleState/delete/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xCSRFToken != nil {
		localVarHeaderParams["X-CSRF-Token"] = parameterToString(*r.xCSRFToken, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetIdentityProfileRequest struct {
	ctx context.Context
	ApiService *IdentityProfilesApiService
	id string
}

func (r ApiGetIdentityProfileRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetIdentityProfileExecute(r)
}

/*
GetIdentityProfile Get Identity Profile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetIdentityProfileRequest
*/
func (a *IdentityProfilesApiService) GetIdentityProfile(ctx context.Context, id string) ApiGetIdentityProfileRequest {
	return ApiGetIdentityProfileRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *IdentityProfilesApiService) GetIdentityProfileExecute(r ApiGetIdentityProfileRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProfilesApiService.GetIdentityProfile")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cc/api/profile/get/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetLifecycleStateRequest struct {
	ctx context.Context
	ApiService *IdentityProfilesApiService
	id string
	authorization *string
}

func (r ApiGetLifecycleStateRequest) Authorization(authorization string) ApiGetLifecycleStateRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetLifecycleStateRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetLifecycleStateExecute(r)
}

/*
GetLifecycleState LCS - Get

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetLifecycleStateRequest
*/
func (a *IdentityProfilesApiService) GetLifecycleState(ctx context.Context, id string) ApiGetLifecycleStateRequest {
	return ApiGetLifecycleStateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *IdentityProfilesApiService) GetLifecycleStateExecute(r ApiGetLifecycleStateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProfilesApiService.GetLifecycleState")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cc/api/lifecycleState/get/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiListIdentityProfilesRequest struct {
	ctx context.Context
	ApiService *IdentityProfilesApiService
	authorization *string
}

func (r ApiListIdentityProfilesRequest) Authorization(authorization string) ApiListIdentityProfilesRequest {
	r.authorization = &authorization
	return r
}

func (r ApiListIdentityProfilesRequest) Execute() (*http.Response, error) {
	return r.ApiService.ListIdentityProfilesExecute(r)
}

/*
ListIdentityProfiles List Identity Profiles

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListIdentityProfilesRequest
*/
func (a *IdentityProfilesApiService) ListIdentityProfiles(ctx context.Context) ApiListIdentityProfilesRequest {
	return ApiListIdentityProfilesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *IdentityProfilesApiService) ListIdentityProfilesExecute(r ApiListIdentityProfilesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProfilesApiService.ListIdentityProfiles")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cc/api/profile/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateIdentityProfileRequest struct {
	ctx context.Context
	ApiService *IdentityProfilesApiService
	profileId string
	xCSRFToken *string
	priority *int32
}

func (r ApiUpdateIdentityProfileRequest) XCSRFToken(xCSRFToken string) ApiUpdateIdentityProfileRequest {
	r.xCSRFToken = &xCSRFToken
	return r
}

func (r ApiUpdateIdentityProfileRequest) Priority(priority int32) ApiUpdateIdentityProfileRequest {
	r.priority = &priority
	return r
}

func (r ApiUpdateIdentityProfileRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateIdentityProfileExecute(r)
}

/*
UpdateIdentityProfile Update Identity Profile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param profileId
 @return ApiUpdateIdentityProfileRequest
*/
func (a *IdentityProfilesApiService) UpdateIdentityProfile(ctx context.Context, profileId string) ApiUpdateIdentityProfileRequest {
	return ApiUpdateIdentityProfileRequest{
		ApiService: a,
		ctx: ctx,
		profileId: profileId,
	}
}

// Execute executes the request
func (a *IdentityProfilesApiService) UpdateIdentityProfileExecute(r ApiUpdateIdentityProfileRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProfilesApiService.UpdateIdentityProfile")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cc/api/profile/update/{profileId}"
	localVarPath = strings.Replace(localVarPath, "{"+"profileId"+"}", url.PathEscape(parameterToString(r.profileId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xCSRFToken != nil {
		localVarHeaderParams["X-CSRF-Token"] = parameterToString(*r.xCSRFToken, "")
	}
	if r.priority != nil {
		localVarFormParams.Add("priority", parameterToString(*r.priority, ""))
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateLifecycleStateRequest struct {
	ctx context.Context
	ApiService *IdentityProfilesApiService
	id string
	xCSRFToken *string
	contentType *string
	body *string
}

func (r ApiUpdateLifecycleStateRequest) XCSRFToken(xCSRFToken string) ApiUpdateLifecycleStateRequest {
	r.xCSRFToken = &xCSRFToken
	return r
}

func (r ApiUpdateLifecycleStateRequest) ContentType(contentType string) ApiUpdateLifecycleStateRequest {
	r.contentType = &contentType
	return r
}

func (r ApiUpdateLifecycleStateRequest) Body(body string) ApiUpdateLifecycleStateRequest {
	r.body = &body
	return r
}

func (r ApiUpdateLifecycleStateRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateLifecycleStateExecute(r)
}

/*
UpdateLifecycleState LCS - Update

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdateLifecycleStateRequest
*/
func (a *IdentityProfilesApiService) UpdateLifecycleState(ctx context.Context, id string) ApiUpdateLifecycleStateRequest {
	return ApiUpdateLifecycleStateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *IdentityProfilesApiService) UpdateLifecycleStateExecute(r ApiUpdateLifecycleStateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProfilesApiService.UpdateLifecycleState")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/lifecycleState/update/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xCSRFToken != nil {
		localVarHeaderParams["X-CSRF-Token"] = parameterToString(*r.xCSRFToken, "")
	}
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}