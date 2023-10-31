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

type AtlasSearchApi interface {

	/*
		CreateAtlasSearchIndex Create One Atlas Search Index

		Creates one Atlas Search index on the specified collection. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. Only clusters running MongoDB v4.2 or later can use Atlas Search. To use this resource, the requesting API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection on which to create an Atlas Search index.
		@return CreateAtlasSearchIndexApiRequest
	*/
	CreateAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, clusterSearchIndex *ClusterSearchIndex) CreateAtlasSearchIndexApiRequest
	/*
		CreateAtlasSearchIndex Create One Atlas Search Index


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateAtlasSearchIndexApiParams - Parameters for the request
		@return CreateAtlasSearchIndexApiRequest
	*/
	CreateAtlasSearchIndexWithParams(ctx context.Context, args *CreateAtlasSearchIndexApiParams) CreateAtlasSearchIndexApiRequest

	// Interface only available internally
	createAtlasSearchIndexExecute(r CreateAtlasSearchIndexApiRequest) (*ClusterSearchIndex, *http.Response, error)

	/*
		DeleteAtlasSearchIndex Remove One Atlas Search Index

		Removes one Atlas Search index that you identified with its unique ID. To use this resource, the requesting API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the database and collection with one or more Application Search indexes.
		@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
		@return DeleteAtlasSearchIndexApiRequest
	*/
	DeleteAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string) DeleteAtlasSearchIndexApiRequest
	/*
		DeleteAtlasSearchIndex Remove One Atlas Search Index


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteAtlasSearchIndexApiParams - Parameters for the request
		@return DeleteAtlasSearchIndexApiRequest
	*/
	DeleteAtlasSearchIndexWithParams(ctx context.Context, args *DeleteAtlasSearchIndexApiParams) DeleteAtlasSearchIndexApiRequest

	// Interface only available internally
	deleteAtlasSearchIndexExecute(r DeleteAtlasSearchIndexApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		GetAtlasSearchIndex Return One Atlas Search Index

		Returns one Atlas Search index in the specified project. You identify this index using its unique ID. Atlas Search index contains the indexed fields and the analyzers used to create the index. To use this resource, the requesting API Key must have the Project Data Access Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
		@param indexId Unique 24-hexadecimal digit string that identifies the Application Search [index](https://docs.atlas.mongodb.com/reference/atlas-search/index-definitions/). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.
		@return GetAtlasSearchIndexApiRequest
	*/
	GetAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string) GetAtlasSearchIndexApiRequest
	/*
		GetAtlasSearchIndex Return One Atlas Search Index


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAtlasSearchIndexApiParams - Parameters for the request
		@return GetAtlasSearchIndexApiRequest
	*/
	GetAtlasSearchIndexWithParams(ctx context.Context, args *GetAtlasSearchIndexApiParams) GetAtlasSearchIndexApiRequest

	// Interface only available internally
	getAtlasSearchIndexExecute(r GetAtlasSearchIndexApiRequest) (*ClusterSearchIndex, *http.Response, error)

	/*
		ListAtlasSearchIndexes Return All Atlas Search Indexes for One Collection

		Returns all Atlas Search indexes on the specified collection. Atlas Search indexes contain the indexed fields and the analyzers used to create the indexes. To use this resource, the requesting API Key must have the Project Data Access Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
		@param collectionName Name of the collection that contains one or more Atlas Search indexes.
		@param databaseName Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.
		@return ListAtlasSearchIndexesApiRequest
	*/
	ListAtlasSearchIndexes(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string) ListAtlasSearchIndexesApiRequest
	/*
		ListAtlasSearchIndexes Return All Atlas Search Indexes for One Collection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAtlasSearchIndexesApiParams - Parameters for the request
		@return ListAtlasSearchIndexesApiRequest
	*/
	ListAtlasSearchIndexesWithParams(ctx context.Context, args *ListAtlasSearchIndexesApiParams) ListAtlasSearchIndexesApiRequest

	// Interface only available internally
	listAtlasSearchIndexesExecute(r ListAtlasSearchIndexesApiRequest) ([]ClusterSearchIndex, *http.Response, error)

	/*
		UpdateAtlasSearchIndex Update One Atlas Search Index

		Updates one Atlas Search index that you identified with its unique ID. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. To use this resource, the requesting API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection whose Atlas Search index to update.
		@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://docs.atlas.mongodb.com/reference/atlas-search/index-definitions/). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
		@return UpdateAtlasSearchIndexApiRequest
	*/
	UpdateAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string, clusterSearchIndex *ClusterSearchIndex) UpdateAtlasSearchIndexApiRequest
	/*
		UpdateAtlasSearchIndex Update One Atlas Search Index


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateAtlasSearchIndexApiParams - Parameters for the request
		@return UpdateAtlasSearchIndexApiRequest
	*/
	UpdateAtlasSearchIndexWithParams(ctx context.Context, args *UpdateAtlasSearchIndexApiParams) UpdateAtlasSearchIndexApiRequest

	// Interface only available internally
	updateAtlasSearchIndexExecute(r UpdateAtlasSearchIndexApiRequest) (*ClusterSearchIndex, *http.Response, error)
}

// AtlasSearchApiService AtlasSearchApi service
type AtlasSearchApiService service

type CreateAtlasSearchIndexApiRequest struct {
	ctx                context.Context
	ApiService         AtlasSearchApi
	groupId            string
	clusterName        string
	clusterSearchIndex *ClusterSearchIndex
}

type CreateAtlasSearchIndexApiParams struct {
	GroupId            string
	ClusterName        string
	ClusterSearchIndex *ClusterSearchIndex
}

func (a *AtlasSearchApiService) CreateAtlasSearchIndexWithParams(ctx context.Context, args *CreateAtlasSearchIndexApiParams) CreateAtlasSearchIndexApiRequest {
	return CreateAtlasSearchIndexApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            args.GroupId,
		clusterName:        args.ClusterName,
		clusterSearchIndex: args.ClusterSearchIndex,
	}
}

func (r CreateAtlasSearchIndexApiRequest) Execute() (*ClusterSearchIndex, *http.Response, error) {
	return r.ApiService.createAtlasSearchIndexExecute(r)
}

/*
CreateAtlasSearchIndex Create One Atlas Search Index

Creates one Atlas Search index on the specified collection. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. Only clusters running MongoDB v4.2 or later can use Atlas Search. To use this resource, the requesting API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection on which to create an Atlas Search index.
	@return CreateAtlasSearchIndexApiRequest
*/
func (a *AtlasSearchApiService) CreateAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, clusterSearchIndex *ClusterSearchIndex) CreateAtlasSearchIndexApiRequest {
	return CreateAtlasSearchIndexApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            groupId,
		clusterName:        clusterName,
		clusterSearchIndex: clusterSearchIndex,
	}
}

// Execute executes the request
//
//	@return ClusterSearchIndex
func (a *AtlasSearchApiService) createAtlasSearchIndexExecute(r CreateAtlasSearchIndexApiRequest) (*ClusterSearchIndex, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterSearchIndex
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.CreateAtlasSearchIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes"
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
	if r.clusterSearchIndex == nil {
		return localVarReturnValue, nil, reportError("clusterSearchIndex is required and must be specified")
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
	localVarPostBody = r.clusterSearchIndex
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

type DeleteAtlasSearchIndexApiRequest struct {
	ctx         context.Context
	ApiService  AtlasSearchApi
	groupId     string
	clusterName string
	indexId     string
}

type DeleteAtlasSearchIndexApiParams struct {
	GroupId     string
	ClusterName string
	IndexId     string
}

func (a *AtlasSearchApiService) DeleteAtlasSearchIndexWithParams(ctx context.Context, args *DeleteAtlasSearchIndexApiParams) DeleteAtlasSearchIndexApiRequest {
	return DeleteAtlasSearchIndexApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		indexId:     args.IndexId,
	}
}

func (r DeleteAtlasSearchIndexApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.deleteAtlasSearchIndexExecute(r)
}

/*
DeleteAtlasSearchIndex Remove One Atlas Search Index

Removes one Atlas Search index that you identified with its unique ID. To use this resource, the requesting API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the database and collection with one or more Application Search indexes.
	@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
	@return DeleteAtlasSearchIndexApiRequest
*/
func (a *AtlasSearchApiService) DeleteAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string) DeleteAtlasSearchIndexApiRequest {
	return DeleteAtlasSearchIndexApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		indexId:     indexId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AtlasSearchApiService) deleteAtlasSearchIndexExecute(r DeleteAtlasSearchIndexApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.DeleteAtlasSearchIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"indexId"+"}", url.PathEscape(parameterValueToString(r.indexId, "indexId")), -1)

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
	if strlen(r.indexId) < 24 {
		return localVarReturnValue, nil, reportError("indexId must have at least 24 elements")
	}
	if strlen(r.indexId) > 24 {
		return localVarReturnValue, nil, reportError("indexId must have less than 24 elements")
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

type GetAtlasSearchIndexApiRequest struct {
	ctx         context.Context
	ApiService  AtlasSearchApi
	groupId     string
	clusterName string
	indexId     string
}

type GetAtlasSearchIndexApiParams struct {
	GroupId     string
	ClusterName string
	IndexId     string
}

func (a *AtlasSearchApiService) GetAtlasSearchIndexWithParams(ctx context.Context, args *GetAtlasSearchIndexApiParams) GetAtlasSearchIndexApiRequest {
	return GetAtlasSearchIndexApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		indexId:     args.IndexId,
	}
}

func (r GetAtlasSearchIndexApiRequest) Execute() (*ClusterSearchIndex, *http.Response, error) {
	return r.ApiService.getAtlasSearchIndexExecute(r)
}

/*
GetAtlasSearchIndex Return One Atlas Search Index

Returns one Atlas Search index in the specified project. You identify this index using its unique ID. Atlas Search index contains the indexed fields and the analyzers used to create the index. To use this resource, the requesting API Key must have the Project Data Access Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
	@param indexId Unique 24-hexadecimal digit string that identifies the Application Search [index](https://docs.atlas.mongodb.com/reference/atlas-search/index-definitions/). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.
	@return GetAtlasSearchIndexApiRequest
*/
func (a *AtlasSearchApiService) GetAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string) GetAtlasSearchIndexApiRequest {
	return GetAtlasSearchIndexApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		indexId:     indexId,
	}
}

// Execute executes the request
//
//	@return ClusterSearchIndex
func (a *AtlasSearchApiService) getAtlasSearchIndexExecute(r GetAtlasSearchIndexApiRequest) (*ClusterSearchIndex, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterSearchIndex
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.GetAtlasSearchIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"indexId"+"}", url.PathEscape(parameterValueToString(r.indexId, "indexId")), -1)

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
	if strlen(r.indexId) < 24 {
		return localVarReturnValue, nil, reportError("indexId must have at least 24 elements")
	}
	if strlen(r.indexId) > 24 {
		return localVarReturnValue, nil, reportError("indexId must have less than 24 elements")
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

type ListAtlasSearchIndexesApiRequest struct {
	ctx            context.Context
	ApiService     AtlasSearchApi
	groupId        string
	clusterName    string
	collectionName string
	databaseName   string
}

type ListAtlasSearchIndexesApiParams struct {
	GroupId        string
	ClusterName    string
	CollectionName string
	DatabaseName   string
}

func (a *AtlasSearchApiService) ListAtlasSearchIndexesWithParams(ctx context.Context, args *ListAtlasSearchIndexesApiParams) ListAtlasSearchIndexesApiRequest {
	return ListAtlasSearchIndexesApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		clusterName:    args.ClusterName,
		collectionName: args.CollectionName,
		databaseName:   args.DatabaseName,
	}
}

func (r ListAtlasSearchIndexesApiRequest) Execute() ([]ClusterSearchIndex, *http.Response, error) {
	return r.ApiService.listAtlasSearchIndexesExecute(r)
}

/*
ListAtlasSearchIndexes Return All Atlas Search Indexes for One Collection

Returns all Atlas Search indexes on the specified collection. Atlas Search indexes contain the indexed fields and the analyzers used to create the indexes. To use this resource, the requesting API Key must have the Project Data Access Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
	@param collectionName Name of the collection that contains one or more Atlas Search indexes.
	@param databaseName Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.
	@return ListAtlasSearchIndexesApiRequest
*/
func (a *AtlasSearchApiService) ListAtlasSearchIndexes(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string) ListAtlasSearchIndexesApiRequest {
	return ListAtlasSearchIndexesApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		clusterName:    clusterName,
		collectionName: collectionName,
		databaseName:   databaseName,
	}
}

// Execute executes the request
//
//	@return []ClusterSearchIndex
func (a *AtlasSearchApiService) listAtlasSearchIndexesExecute(r ListAtlasSearchIndexesApiRequest) ([]ClusterSearchIndex, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ClusterSearchIndex
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.ListAtlasSearchIndexes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"collectionName"+"}", url.PathEscape(parameterValueToString(r.collectionName, "collectionName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"databaseName"+"}", url.PathEscape(parameterValueToString(r.databaseName, "databaseName")), -1)

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

type UpdateAtlasSearchIndexApiRequest struct {
	ctx                context.Context
	ApiService         AtlasSearchApi
	groupId            string
	clusterName        string
	indexId            string
	clusterSearchIndex *ClusterSearchIndex
}

type UpdateAtlasSearchIndexApiParams struct {
	GroupId            string
	ClusterName        string
	IndexId            string
	ClusterSearchIndex *ClusterSearchIndex
}

func (a *AtlasSearchApiService) UpdateAtlasSearchIndexWithParams(ctx context.Context, args *UpdateAtlasSearchIndexApiParams) UpdateAtlasSearchIndexApiRequest {
	return UpdateAtlasSearchIndexApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            args.GroupId,
		clusterName:        args.ClusterName,
		indexId:            args.IndexId,
		clusterSearchIndex: args.ClusterSearchIndex,
	}
}

func (r UpdateAtlasSearchIndexApiRequest) Execute() (*ClusterSearchIndex, *http.Response, error) {
	return r.ApiService.updateAtlasSearchIndexExecute(r)
}

/*
UpdateAtlasSearchIndex Update One Atlas Search Index

Updates one Atlas Search index that you identified with its unique ID. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. To use this resource, the requesting API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection whose Atlas Search index to update.
	@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://docs.atlas.mongodb.com/reference/atlas-search/index-definitions/). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
	@return UpdateAtlasSearchIndexApiRequest
*/
func (a *AtlasSearchApiService) UpdateAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string, clusterSearchIndex *ClusterSearchIndex) UpdateAtlasSearchIndexApiRequest {
	return UpdateAtlasSearchIndexApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            groupId,
		clusterName:        clusterName,
		indexId:            indexId,
		clusterSearchIndex: clusterSearchIndex,
	}
}

// Execute executes the request
//
//	@return ClusterSearchIndex
func (a *AtlasSearchApiService) updateAtlasSearchIndexExecute(r UpdateAtlasSearchIndexApiRequest) (*ClusterSearchIndex, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterSearchIndex
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.UpdateAtlasSearchIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"indexId"+"}", url.PathEscape(parameterValueToString(r.indexId, "indexId")), -1)

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
	if strlen(r.indexId) < 24 {
		return localVarReturnValue, nil, reportError("indexId must have at least 24 elements")
	}
	if strlen(r.indexId) > 24 {
		return localVarReturnValue, nil, reportError("indexId must have less than 24 elements")
	}
	if r.clusterSearchIndex == nil {
		return localVarReturnValue, nil, reportError("clusterSearchIndex is required and must be specified")
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
	localVarPostBody = r.clusterSearchIndex
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