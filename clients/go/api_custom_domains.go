/*
Fermyon Cloud API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloud_openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CustomDomainsApiService CustomDomainsApi service
type CustomDomainsApiService service

type ApiApiCustomDomainsNameGetRequest struct {
	ctx context.Context
	ApiService *CustomDomainsApiService
	name string
	domainName *string
	apiVersion *string
}

func (r ApiApiCustomDomainsNameGetRequest) DomainName(domainName string) ApiApiCustomDomainsNameGetRequest {
	r.domainName = &domainName
	return r
}

// The requested API version
func (r ApiApiCustomDomainsNameGetRequest) ApiVersion(apiVersion string) ApiApiCustomDomainsNameGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiCustomDomainsNameGetRequest) Execute() (*DomainItem, *http.Response, error) {
	return r.ApiService.ApiCustomDomainsNameGetExecute(r)
}

/*
ApiCustomDomainsNameGet Method for ApiCustomDomainsNameGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ApiApiCustomDomainsNameGetRequest
*/
func (a *CustomDomainsApiService) ApiCustomDomainsNameGet(ctx context.Context, name string) ApiApiCustomDomainsNameGetRequest {
	return ApiApiCustomDomainsNameGetRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return DomainItem
func (a *CustomDomainsApiService) ApiCustomDomainsNameGetExecute(r ApiApiCustomDomainsNameGetRequest) (*DomainItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DomainItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDomainsApiService.ApiCustomDomainsNameGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/custom-domains/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.domainName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "domainName", r.domainName, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.apiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Api-Version", r.apiVersion, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
