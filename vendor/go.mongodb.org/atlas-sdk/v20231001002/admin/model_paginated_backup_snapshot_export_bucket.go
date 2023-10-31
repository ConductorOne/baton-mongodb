// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// PaginatedBackupSnapshotExportBucket struct for PaginatedBackupSnapshotExportBucket
type PaginatedBackupSnapshotExportBucket struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []DiskBackupSnapshotAWSExportBucket `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedBackupSnapshotExportBucket instantiates a new PaginatedBackupSnapshotExportBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedBackupSnapshotExportBucket() *PaginatedBackupSnapshotExportBucket {
	this := PaginatedBackupSnapshotExportBucket{}
	return &this
}

// NewPaginatedBackupSnapshotExportBucketWithDefaults instantiates a new PaginatedBackupSnapshotExportBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedBackupSnapshotExportBucketWithDefaults() *PaginatedBackupSnapshotExportBucket {
	this := PaginatedBackupSnapshotExportBucket{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *PaginatedBackupSnapshotExportBucket) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedBackupSnapshotExportBucket) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedBackupSnapshotExportBucket) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedBackupSnapshotExportBucket) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise
func (o *PaginatedBackupSnapshotExportBucket) GetResults() []DiskBackupSnapshotAWSExportBucket {
	if o == nil || IsNil(o.Results) {
		var ret []DiskBackupSnapshotAWSExportBucket
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedBackupSnapshotExportBucket) GetResultsOk() ([]DiskBackupSnapshotAWSExportBucket, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}

	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedBackupSnapshotExportBucket) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DiskBackupSnapshotAWSExportBucket and assigns it to the Results field.
func (o *PaginatedBackupSnapshotExportBucket) SetResults(v []DiskBackupSnapshotAWSExportBucket) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise
func (o *PaginatedBackupSnapshotExportBucket) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedBackupSnapshotExportBucket) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}

	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedBackupSnapshotExportBucket) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedBackupSnapshotExportBucket) SetTotalCount(v int) {
	o.TotalCount = &v
}

func (o PaginatedBackupSnapshotExportBucket) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedBackupSnapshotExportBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
