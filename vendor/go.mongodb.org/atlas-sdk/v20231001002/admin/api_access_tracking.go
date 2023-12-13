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

type AccessTrackingApi interface {

	/*
		ListAccessLogsByClusterName Return Database Access History for One Cluster using Its Cluster Name

		Returns the access logs of one cluster identified by the cluster's name. Access logs contain a list of authentication requests made against your cluster. You can't use this feature on tenant-tier clusters (M0, M2, M5). To use this resource, the requesting API Key must have the Project Monitoring Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return ListAccessLogsByClusterNameApiRequest
	*/
	ListAccessLogsByClusterName(ctx context.Context, groupId string, clusterName string) ListAccessLogsByClusterNameApiRequest
	/*
		ListAccessLogsByClusterName Return Database Access History for One Cluster using Its Cluster Name


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAccessLogsByClusterNameApiParams - Parameters for the request
		@return ListAccessLogsByClusterNameApiRequest
	*/
	ListAccessLogsByClusterNameWithParams(ctx context.Context, args *ListAccessLogsByClusterNameApiParams) ListAccessLogsByClusterNameApiRequest

	// Interface only available internally
	listAccessLogsByClusterNameExecute(r ListAccessLogsByClusterNameApiRequest) (*MongoDBAccessLogsList, *http.Response, error)

	/*
		ListAccessLogsByHostname Return Database Access History for One Cluster using Its Hostname

		Returns the access logs of one cluster identified by the cluster's hostname. Access logs contain a list of authentication requests made against your clusters. You can't use this feature on tenant-tier clusters (M0, M2, M5). To use this resource, the requesting API Key must have the Project Monitoring Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param hostname Fully qualified domain name or IP address of the MongoDB host that stores the log files that you want to download.
		@return ListAccessLogsByHostnameApiRequest
	*/
	ListAccessLogsByHostname(ctx context.Context, groupId string, hostname string) ListAccessLogsByHostnameApiRequest
	/*
		ListAccessLogsByHostname Return Database Access History for One Cluster using Its Hostname


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAccessLogsByHostnameApiParams - Parameters for the request
		@return ListAccessLogsByHostnameApiRequest
	*/
	ListAccessLogsByHostnameWithParams(ctx context.Context, args *ListAccessLogsByHostnameApiParams) ListAccessLogsByHostnameApiRequest

	// Interface only available internally
	listAccessLogsByHostnameExecute(r ListAccessLogsByHostnameApiRequest) (*MongoDBAccessLogsList, *http.Response, error)
}

// AccessTrackingApiService AccessTrackingApi service
type AccessTrackingApiService service

type ListAccessLogsByClusterNameApiRequest struct {
	ctx         context.Context
	ApiService  AccessTrackingApi
	groupId     string
	clusterName string
	authResult  *bool
	end         *int64
	ipAddress   *string
	nLogs       *int
	start       *int64
}

type ListAccessLogsByClusterNameApiParams struct {
	GroupId     string
	ClusterName string
	AuthResult  *bool
	End         *int64
	IpAddress   *string
	NLogs       *int
	Start       *int64
}

func (a *AccessTrackingApiService) ListAccessLogsByClusterNameWithParams(ctx context.Context, args *ListAccessLogsByClusterNameApiParams) ListAccessLogsByClusterNameApiRequest {
	return ListAccessLogsByClusterNameApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		authResult:  args.AuthResult,
		end:         args.End,
		ipAddress:   args.IpAddress,
		nLogs:       args.NLogs,
		start:       args.Start,
	}
}

// Flag that indicates whether the response returns the successful authentication attempts only.
func (r ListAccessLogsByClusterNameApiRequest) AuthResult(authResult bool) ListAccessLogsByClusterNameApiRequest {
	r.authResult = &authResult
	return r
}

// Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses UNIX epoch time in milliseconds.
func (r ListAccessLogsByClusterNameApiRequest) End(end int64) ListAccessLogsByClusterNameApiRequest {
	r.end = &end
	return r
}

// One Internet Protocol address that attempted to authenticate with the database.
func (r ListAccessLogsByClusterNameApiRequest) IpAddress(ipAddress string) ListAccessLogsByClusterNameApiRequest {
	r.ipAddress = &ipAddress
	return r
}

// Maximum number of lines from the log to return.
func (r ListAccessLogsByClusterNameApiRequest) NLogs(nLogs int) ListAccessLogsByClusterNameApiRequest {
	r.nLogs = &nLogs
	return r
}

// Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses UNIX epoch time in milliseconds.
func (r ListAccessLogsByClusterNameApiRequest) Start(start int64) ListAccessLogsByClusterNameApiRequest {
	r.start = &start
	return r
}

func (r ListAccessLogsByClusterNameApiRequest) Execute() (*MongoDBAccessLogsList, *http.Response, error) {
	return r.ApiService.listAccessLogsByClusterNameExecute(r)
}

/*
ListAccessLogsByClusterName Return Database Access History for One Cluster using Its Cluster Name

Returns the access logs of one cluster identified by the cluster's name. Access logs contain a list of authentication requests made against your cluster. You can't use this feature on tenant-tier clusters (M0, M2, M5). To use this resource, the requesting API Key must have the Project Monitoring Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return ListAccessLogsByClusterNameApiRequest
*/
func (a *AccessTrackingApiService) ListAccessLogsByClusterName(ctx context.Context, groupId string, clusterName string) ListAccessLogsByClusterNameApiRequest {
	return ListAccessLogsByClusterNameApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return MongoDBAccessLogsList
func (a *AccessTrackingApiService) listAccessLogsByClusterNameExecute(r ListAccessLogsByClusterNameApiRequest) (*MongoDBAccessLogsList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MongoDBAccessLogsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessTrackingApiService.ListAccessLogsByClusterName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/dbAccessHistory/clusters/{clusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}

	if r.authResult != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "authResult", r.authResult, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	if r.ipAddress != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ipAddress", r.ipAddress, "")
	}
	if r.nLogs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nLogs", r.nLogs, "")
	} else {
		var defaultValue int = 20000
		r.nLogs = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "nLogs", r.nLogs, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
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

type ListAccessLogsByHostnameApiRequest struct {
	ctx        context.Context
	ApiService AccessTrackingApi
	groupId    string
	hostname   string
	authResult *bool
	end        *int64
	ipAddress  *string
	nLogs      *int
	start      *int64
}

type ListAccessLogsByHostnameApiParams struct {
	GroupId    string
	Hostname   string
	AuthResult *bool
	End        *int64
	IpAddress  *string
	NLogs      *int
	Start      *int64
}

func (a *AccessTrackingApiService) ListAccessLogsByHostnameWithParams(ctx context.Context, args *ListAccessLogsByHostnameApiParams) ListAccessLogsByHostnameApiRequest {
	return ListAccessLogsByHostnameApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		hostname:   args.Hostname,
		authResult: args.AuthResult,
		end:        args.End,
		ipAddress:  args.IpAddress,
		nLogs:      args.NLogs,
		start:      args.Start,
	}
}

// Flag that indicates whether the response returns the successful authentication attempts only.
func (r ListAccessLogsByHostnameApiRequest) AuthResult(authResult bool) ListAccessLogsByHostnameApiRequest {
	r.authResult = &authResult
	return r
}

// Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses UNIX epoch time in milliseconds.
func (r ListAccessLogsByHostnameApiRequest) End(end int64) ListAccessLogsByHostnameApiRequest {
	r.end = &end
	return r
}

// One Internet Protocol address that attempted to authenticate with the database.
func (r ListAccessLogsByHostnameApiRequest) IpAddress(ipAddress string) ListAccessLogsByHostnameApiRequest {
	r.ipAddress = &ipAddress
	return r
}

// Maximum number of lines from the log to return.
func (r ListAccessLogsByHostnameApiRequest) NLogs(nLogs int) ListAccessLogsByHostnameApiRequest {
	r.nLogs = &nLogs
	return r
}

// Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses UNIX epoch time in milliseconds.
func (r ListAccessLogsByHostnameApiRequest) Start(start int64) ListAccessLogsByHostnameApiRequest {
	r.start = &start
	return r
}

func (r ListAccessLogsByHostnameApiRequest) Execute() (*MongoDBAccessLogsList, *http.Response, error) {
	return r.ApiService.listAccessLogsByHostnameExecute(r)
}

/*
ListAccessLogsByHostname Return Database Access History for One Cluster using Its Hostname

Returns the access logs of one cluster identified by the cluster's hostname. Access logs contain a list of authentication requests made against your clusters. You can't use this feature on tenant-tier clusters (M0, M2, M5). To use this resource, the requesting API Key must have the Project Monitoring Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param hostname Fully qualified domain name or IP address of the MongoDB host that stores the log files that you want to download.
	@return ListAccessLogsByHostnameApiRequest
*/
func (a *AccessTrackingApiService) ListAccessLogsByHostname(ctx context.Context, groupId string, hostname string) ListAccessLogsByHostnameApiRequest {
	return ListAccessLogsByHostnameApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		hostname:   hostname,
	}
}

// Execute executes the request
//
//	@return MongoDBAccessLogsList
func (a *AccessTrackingApiService) listAccessLogsByHostnameExecute(r ListAccessLogsByHostnameApiRequest) (*MongoDBAccessLogsList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MongoDBAccessLogsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessTrackingApiService.ListAccessLogsByHostname")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/dbAccessHistory/processes/{hostname}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"hostname"+"}", url.PathEscape(parameterValueToString(r.hostname, "hostname")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}

	if r.authResult != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "authResult", r.authResult, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	if r.ipAddress != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ipAddress", r.ipAddress, "")
	}
	if r.nLogs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nLogs", r.nLogs, "")
	} else {
		var defaultValue int = 20000
		r.nLogs = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "nLogs", r.nLogs, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
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
