// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AlertsApi interface {

	/*
		AcknowledgeAlert Acknowledge One Alert from One Project

		Confirms receipt of one existing alert. This alert applies to any component in one project. Acknowledging an alert prevents successive notifications. You receive an alert when a monitored component meets or exceeds a value you set until you acknowledge the alert. To use this resource, the requesting API Key must have the Organization Owner or Project Owner role.

	This resource remains under revision and may change.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param alertId Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access.
		@return AcknowledgeAlertApiRequest
	*/
	AcknowledgeAlert(ctx context.Context, groupId string, alertId string, alertViewForNdsGroup *AlertViewForNdsGroup) AcknowledgeAlertApiRequest
	/*
		AcknowledgeAlert Acknowledge One Alert from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AcknowledgeAlertApiParams - Parameters for the request
		@return AcknowledgeAlertApiRequest
	*/
	AcknowledgeAlertWithParams(ctx context.Context, args *AcknowledgeAlertApiParams) AcknowledgeAlertApiRequest

	// Interface only available internally
	acknowledgeAlertExecute(r AcknowledgeAlertApiRequest) (*AlertViewForNdsGroup, *http.Response, error)

	/*
		GetAlert Return One Alert from One Project

		Returns one alert. This alert applies to any component in one project. You receive an alert when a monitored component meets or exceeds a value you set. To use this resource, the requesting API Key must have the Project Read Only role.

	This resource remains under revision and may change.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param alertId Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access.
		@return GetAlertApiRequest
	*/
	GetAlert(ctx context.Context, groupId string, alertId string) GetAlertApiRequest
	/*
		GetAlert Return One Alert from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAlertApiParams - Parameters for the request
		@return GetAlertApiRequest
	*/
	GetAlertWithParams(ctx context.Context, args *GetAlertApiParams) GetAlertApiRequest

	// Interface only available internally
	getAlertExecute(r GetAlertApiRequest) (*AlertViewForNdsGroup, *http.Response, error)

	/*
		ListAlerts Return All Alerts from One Project

		Returns all alerts. These alerts apply to all components in one project. You receive an alert when a monitored component meets or exceeds a value you set. To use this resource, the requesting API Key must have the Project Read Only role.

	This resource remains under revision and may change.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListAlertsApiRequest
	*/
	ListAlerts(ctx context.Context, groupId string) ListAlertsApiRequest
	/*
		ListAlerts Return All Alerts from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAlertsApiParams - Parameters for the request
		@return ListAlertsApiRequest
	*/
	ListAlertsWithParams(ctx context.Context, args *ListAlertsApiParams) ListAlertsApiRequest

	// Interface only available internally
	listAlertsExecute(r ListAlertsApiRequest) (*PaginatedAlert, *http.Response, error)

	/*
		ListAlertsByAlertConfigurationId Return All Open Alerts for Alert Configuration

		[experimental] Returns all open alerts that the specified alert configuration triggers. These alert configurations apply to the specified project only. Alert configurations define the triggers and notification methods for alerts. Open alerts have been triggered but remain unacknowledged. To use this resource, the requesting API Key must have the Project Read Only role.

	This resource remains under revision and may change.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param alertConfigId Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.
		@return ListAlertsByAlertConfigurationIdApiRequest
	*/
	ListAlertsByAlertConfigurationId(ctx context.Context, groupId string, alertConfigId string) ListAlertsByAlertConfigurationIdApiRequest
	/*
		ListAlertsByAlertConfigurationId Return All Open Alerts for Alert Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAlertsByAlertConfigurationIdApiParams - Parameters for the request
		@return ListAlertsByAlertConfigurationIdApiRequest
	*/
	ListAlertsByAlertConfigurationIdWithParams(ctx context.Context, args *ListAlertsByAlertConfigurationIdApiParams) ListAlertsByAlertConfigurationIdApiRequest

	// Interface only available internally
	listAlertsByAlertConfigurationIdExecute(r ListAlertsByAlertConfigurationIdApiRequest) (*PaginatedAlert, *http.Response, error)
}

// AlertsApiService AlertsApi service
type AlertsApiService service

type AcknowledgeAlertApiRequest struct {
	ctx                  context.Context
	ApiService           AlertsApi
	groupId              string
	alertId              string
	alertViewForNdsGroup *AlertViewForNdsGroup
}

type AcknowledgeAlertApiParams struct {
	GroupId              string
	AlertId              string
	AlertViewForNdsGroup *AlertViewForNdsGroup
}

func (a *AlertsApiService) AcknowledgeAlertWithParams(ctx context.Context, args *AcknowledgeAlertApiParams) AcknowledgeAlertApiRequest {
	return AcknowledgeAlertApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		groupId:              args.GroupId,
		alertId:              args.AlertId,
		alertViewForNdsGroup: args.AlertViewForNdsGroup,
	}
}

func (r AcknowledgeAlertApiRequest) Execute() (*AlertViewForNdsGroup, *http.Response, error) {
	return r.ApiService.acknowledgeAlertExecute(r)
}

/*
AcknowledgeAlert Acknowledge One Alert from One Project

Confirms receipt of one existing alert. This alert applies to any component in one project. Acknowledging an alert prevents successive notifications. You receive an alert when a monitored component meets or exceeds a value you set until you acknowledge the alert. To use this resource, the requesting API Key must have the Organization Owner or Project Owner role.

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param alertId Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access.
	@return AcknowledgeAlertApiRequest
*/
func (a *AlertsApiService) AcknowledgeAlert(ctx context.Context, groupId string, alertId string, alertViewForNdsGroup *AlertViewForNdsGroup) AcknowledgeAlertApiRequest {
	return AcknowledgeAlertApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		groupId:              groupId,
		alertId:              alertId,
		alertViewForNdsGroup: alertViewForNdsGroup,
	}
}

// Execute executes the request
//
//	@return AlertViewForNdsGroup
func (a *AlertsApiService) acknowledgeAlertExecute(r AcknowledgeAlertApiRequest) (*AlertViewForNdsGroup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlertViewForNdsGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertsApiService.AcknowledgeAlert")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/alerts/{alertId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"alertId"+"}", url.PathEscape(parameterValueToString(r.alertId, "alertId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.alertId) < 24 {
		return localVarReturnValue, nil, reportError("alertId must have at least 24 elements")
	}
	if strlen(r.alertId) > 24 {
		return localVarReturnValue, nil, reportError("alertId must have less than 24 elements")
	}
	if r.alertViewForNdsGroup == nil {
		return localVarReturnValue, nil, reportError("alertViewForNdsGroup is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.alertViewForNdsGroup
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
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
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

type GetAlertApiRequest struct {
	ctx        context.Context
	ApiService AlertsApi
	groupId    string
	alertId    string
}

type GetAlertApiParams struct {
	GroupId string
	AlertId string
}

func (a *AlertsApiService) GetAlertWithParams(ctx context.Context, args *GetAlertApiParams) GetAlertApiRequest {
	return GetAlertApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		alertId:    args.AlertId,
	}
}

func (r GetAlertApiRequest) Execute() (*AlertViewForNdsGroup, *http.Response, error) {
	return r.ApiService.getAlertExecute(r)
}

/*
GetAlert Return One Alert from One Project

Returns one alert. This alert applies to any component in one project. You receive an alert when a monitored component meets or exceeds a value you set. To use this resource, the requesting API Key must have the Project Read Only role.

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param alertId Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access.
	@return GetAlertApiRequest
*/
func (a *AlertsApiService) GetAlert(ctx context.Context, groupId string, alertId string) GetAlertApiRequest {
	return GetAlertApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		alertId:    alertId,
	}
}

// Execute executes the request
//
//	@return AlertViewForNdsGroup
func (a *AlertsApiService) getAlertExecute(r GetAlertApiRequest) (*AlertViewForNdsGroup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlertViewForNdsGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertsApiService.GetAlert")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/alerts/{alertId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"alertId"+"}", url.PathEscape(parameterValueToString(r.alertId, "alertId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.alertId) < 24 {
		return localVarReturnValue, nil, reportError("alertId must have at least 24 elements")
	}
	if strlen(r.alertId) > 24 {
		return localVarReturnValue, nil, reportError("alertId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
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

type ListAlertsApiRequest struct {
	ctx          context.Context
	ApiService   AlertsApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
	status       *string
}

type ListAlertsApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
	Status       *string
}

func (a *AlertsApiService) ListAlertsWithParams(ctx context.Context, args *ListAlertsApiParams) ListAlertsApiRequest {
	return ListAlertsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
		status:       args.Status,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListAlertsApiRequest) IncludeCount(includeCount bool) ListAlertsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListAlertsApiRequest) ItemsPerPage(itemsPerPage int) ListAlertsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListAlertsApiRequest) PageNum(pageNum int) ListAlertsApiRequest {
	r.pageNum = &pageNum
	return r
}

// Status of the alerts to return. Omit to return all alerts in all statuses.
func (r ListAlertsApiRequest) Status(status string) ListAlertsApiRequest {
	r.status = &status
	return r
}

func (r ListAlertsApiRequest) Execute() (*PaginatedAlert, *http.Response, error) {
	return r.ApiService.listAlertsExecute(r)
}

/*
ListAlerts Return All Alerts from One Project

Returns all alerts. These alerts apply to all components in one project. You receive an alert when a monitored component meets or exceeds a value you set. To use this resource, the requesting API Key must have the Project Read Only role.

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListAlertsApiRequest
*/
func (a *AlertsApiService) ListAlerts(ctx context.Context, groupId string) ListAlertsApiRequest {
	return ListAlertsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return PaginatedAlert
func (a *AlertsApiService) listAlertsExecute(r ListAlertsApiRequest) (*PaginatedAlert, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedAlert
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertsApiService.ListAlerts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/alerts"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
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

type ListAlertsByAlertConfigurationIdApiRequest struct {
	ctx           context.Context
	ApiService    AlertsApi
	groupId       string
	alertConfigId string
	includeCount  *bool
	itemsPerPage  *int
	pageNum       *int
}

type ListAlertsByAlertConfigurationIdApiParams struct {
	GroupId       string
	AlertConfigId string
	IncludeCount  *bool
	ItemsPerPage  *int
	PageNum       *int
}

func (a *AlertsApiService) ListAlertsByAlertConfigurationIdWithParams(ctx context.Context, args *ListAlertsByAlertConfigurationIdApiParams) ListAlertsByAlertConfigurationIdApiRequest {
	return ListAlertsByAlertConfigurationIdApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		alertConfigId: args.AlertConfigId,
		includeCount:  args.IncludeCount,
		itemsPerPage:  args.ItemsPerPage,
		pageNum:       args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListAlertsByAlertConfigurationIdApiRequest) IncludeCount(includeCount bool) ListAlertsByAlertConfigurationIdApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListAlertsByAlertConfigurationIdApiRequest) ItemsPerPage(itemsPerPage int) ListAlertsByAlertConfigurationIdApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListAlertsByAlertConfigurationIdApiRequest) PageNum(pageNum int) ListAlertsByAlertConfigurationIdApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListAlertsByAlertConfigurationIdApiRequest) Execute() (*PaginatedAlert, *http.Response, error) {
	return r.ApiService.listAlertsByAlertConfigurationIdExecute(r)
}

/*
ListAlertsByAlertConfigurationId Return All Open Alerts for Alert Configuration

[experimental] Returns all open alerts that the specified alert configuration triggers. These alert configurations apply to the specified project only. Alert configurations define the triggers and notification methods for alerts. Open alerts have been triggered but remain unacknowledged. To use this resource, the requesting API Key must have the Project Read Only role.

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param alertConfigId Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.
	@return ListAlertsByAlertConfigurationIdApiRequest
*/
func (a *AlertsApiService) ListAlertsByAlertConfigurationId(ctx context.Context, groupId string, alertConfigId string) ListAlertsByAlertConfigurationIdApiRequest {
	return ListAlertsByAlertConfigurationIdApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		alertConfigId: alertConfigId,
	}
}

// Execute executes the request
//
//	@return PaginatedAlert
func (a *AlertsApiService) listAlertsByAlertConfigurationIdExecute(r ListAlertsByAlertConfigurationIdApiRequest) (*PaginatedAlert, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedAlert
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertsApiService.ListAlertsByAlertConfigurationId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId}/alerts"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"alertConfigId"+"}", url.PathEscape(parameterValueToString(r.alertConfigId, "alertConfigId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.alertConfigId) < 24 {
		return localVarReturnValue, nil, reportError("alertConfigId must have at least 24 elements")
	}
	if strlen(r.alertConfigId) > 24 {
		return localVarReturnValue, nil, reportError("alertConfigId must have less than 24 elements")
	}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
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
