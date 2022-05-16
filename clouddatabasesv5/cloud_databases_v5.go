/**
 * (C) Copyright IBM Corp. 2022.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.45.1-632ec580-20220210-190638
 */

// Package clouddatabasesv5 : Operations and models for the CloudDatabasesV5 service
package clouddatabasesv5

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	common "github.com/IBM/cloud-databases-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
)

// CloudDatabasesV5 : The IBM Cloud Databases API enables interaction between applications and Cloud Databases database
// deployments.
//
// Access to the API requires an IAM Bearer Token or an IAM API Key to be presented through bearer authentication.
//
// Deployment IDs are CRNs on the IBM Cloud Databases v5 API platform. No lookup or translation the Compose style UUIDs
// is needed. The Deployment ID is a traditional UUID on the Compose v5 API platform.
//
// When you use CRNs, remember to URL escape the CRN value as they can include the forward-slash (/) character.
//
// API Version: 5.0.0
type CloudDatabasesV5 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.databases.cloud.ibm.com/v5/ibm"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "cloud_databases"

const ParameterizedServiceURL = "https://api.{region}.databases.cloud.ibm.com/v5/{platform}"

var defaultUrlVariables = map[string]string{
	"platform": "ibm",
	"region": "us-south",
}

// CloudDatabasesV5Options : Service options
type CloudDatabasesV5Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewCloudDatabasesV5UsingExternalConfig : constructs an instance of CloudDatabasesV5 with passed in options and external configuration.
func NewCloudDatabasesV5UsingExternalConfig(options *CloudDatabasesV5Options) (cloudDatabases *CloudDatabasesV5, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	cloudDatabases, err = NewCloudDatabasesV5(options)
	if err != nil {
		return
	}

	err = cloudDatabases.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = cloudDatabases.Service.SetServiceURL(options.URL)
	}
	return
}

// NewCloudDatabasesV5 : constructs an instance of CloudDatabasesV5 with passed in options.
func NewCloudDatabasesV5(options *CloudDatabasesV5Options) (service *CloudDatabasesV5, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &CloudDatabasesV5{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "cloudDatabases" suitable for processing requests.
func (cloudDatabases *CloudDatabasesV5) Clone() *CloudDatabasesV5 {
	if core.IsNil(cloudDatabases) {
		return nil
	}
	clone := *cloudDatabases
	clone.Service = cloudDatabases.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (cloudDatabases *CloudDatabasesV5) SetServiceURL(url string) error {
	return cloudDatabases.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (cloudDatabases *CloudDatabasesV5) GetServiceURL() string {
	return cloudDatabases.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (cloudDatabases *CloudDatabasesV5) SetDefaultHeaders(headers http.Header) {
	cloudDatabases.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (cloudDatabases *CloudDatabasesV5) SetEnableGzipCompression(enableGzip bool) {
	cloudDatabases.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (cloudDatabases *CloudDatabasesV5) GetEnableGzipCompression() bool {
	return cloudDatabases.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (cloudDatabases *CloudDatabasesV5) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	cloudDatabases.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (cloudDatabases *CloudDatabasesV5) DisableRetries() {
	cloudDatabases.Service.DisableRetries()
}

// ListDeployables : List all deployable databases
// Returns a list of all the types and associated major versions of database deployments that can be provisioned.
func (cloudDatabases *CloudDatabasesV5) ListDeployables(listDeployablesOptions *ListDeployablesOptions) (result *ListDeployablesResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.ListDeployablesWithContext(context.Background(), listDeployablesOptions)
}

// ListDeployablesWithContext is an alternate form of the ListDeployables method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) ListDeployablesWithContext(ctx context.Context, listDeployablesOptions *ListDeployablesOptions) (result *ListDeployablesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDeployablesOptions, "listDeployablesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployables`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDeployablesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "ListDeployables")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListDeployablesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListRegions : List all deployable regions
// Returns a list of all the regions that deployments can be provisioned into from the current region. Used to determine
// region availability for read-only replicas.
func (cloudDatabases *CloudDatabasesV5) ListRegions(listRegionsOptions *ListRegionsOptions) (result *ListRegionsResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.ListRegionsWithContext(context.Background(), listRegionsOptions)
}

// ListRegionsWithContext is an alternate form of the ListRegions method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) ListRegionsWithContext(ctx context.Context, listRegionsOptions *ListRegionsOptions) (result *ListRegionsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listRegionsOptions, "listRegionsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/regions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listRegionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "ListRegions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListRegionsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDeploymentInfo : Get deployment information
// Gets the full data that is associated with a deployment. This data includes the ID, name, database type, and version.
func (cloudDatabases *CloudDatabasesV5) GetDeploymentInfo(getDeploymentInfoOptions *GetDeploymentInfoOptions) (result *GetDeploymentInfoResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetDeploymentInfoWithContext(context.Background(), getDeploymentInfoOptions)
}

// GetDeploymentInfoWithContext is an alternate form of the GetDeploymentInfo method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetDeploymentInfoWithContext(ctx context.Context, getDeploymentInfoOptions *GetDeploymentInfoOptions) (result *GetDeploymentInfoResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDeploymentInfoOptions, "getDeploymentInfoOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDeploymentInfoOptions, "getDeploymentInfoOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getDeploymentInfoOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDeploymentInfoOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetDeploymentInfo")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetDeploymentInfoResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateDatabaseUser : Creates a user based on user type
// Creates a user in the database that can access the database through a connection.
func (cloudDatabases *CloudDatabasesV5) CreateDatabaseUser(createDatabaseUserOptions *CreateDatabaseUserOptions) (result *CreateDatabaseUserResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.CreateDatabaseUserWithContext(context.Background(), createDatabaseUserOptions)
}

// CreateDatabaseUserWithContext is an alternate form of the CreateDatabaseUser method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) CreateDatabaseUserWithContext(ctx context.Context, createDatabaseUserOptions *CreateDatabaseUserOptions) (result *CreateDatabaseUserResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDatabaseUserOptions, "createDatabaseUserOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDatabaseUserOptions, "createDatabaseUserOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *createDatabaseUserOptions.ID,
		"user_type": *createDatabaseUserOptions.UserType,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/users/{user_type}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDatabaseUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "CreateDatabaseUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createDatabaseUserOptions.User != nil {
		body["user"] = createDatabaseUserOptions.User
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateDatabaseUserResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ChangeUserPassword : Set specified user's password
// Sets the password of a specified user.
func (cloudDatabases *CloudDatabasesV5) ChangeUserPassword(changeUserPasswordOptions *ChangeUserPasswordOptions) (result *ChangeUserPasswordResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.ChangeUserPasswordWithContext(context.Background(), changeUserPasswordOptions)
}

// ChangeUserPasswordWithContext is an alternate form of the ChangeUserPassword method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) ChangeUserPasswordWithContext(ctx context.Context, changeUserPasswordOptions *ChangeUserPasswordOptions) (result *ChangeUserPasswordResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(changeUserPasswordOptions, "changeUserPasswordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(changeUserPasswordOptions, "changeUserPasswordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *changeUserPasswordOptions.ID,
		"user_type": *changeUserPasswordOptions.UserType,
		"username": *changeUserPasswordOptions.Username,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/users/{user_type}/{username}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range changeUserPasswordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "ChangeUserPassword")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if changeUserPasswordOptions.User != nil {
		body["user"] = changeUserPasswordOptions.User
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalChangeUserPasswordResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDatabaseUser : Deletes a user based on user type
// Removes a user from the deployment.
func (cloudDatabases *CloudDatabasesV5) DeleteDatabaseUser(deleteDatabaseUserOptions *DeleteDatabaseUserOptions) (result *DeleteDatabaseUserResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.DeleteDatabaseUserWithContext(context.Background(), deleteDatabaseUserOptions)
}

// DeleteDatabaseUserWithContext is an alternate form of the DeleteDatabaseUser method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) DeleteDatabaseUserWithContext(ctx context.Context, deleteDatabaseUserOptions *DeleteDatabaseUserOptions) (result *DeleteDatabaseUserResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDatabaseUserOptions, "deleteDatabaseUserOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDatabaseUserOptions, "deleteDatabaseUserOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteDatabaseUserOptions.ID,
		"user_type": *deleteDatabaseUserOptions.UserType,
		"username": *deleteDatabaseUserOptions.Username,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/users/{user_type}/{username}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDatabaseUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "DeleteDatabaseUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteDatabaseUserResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateDatabaseConfiguration : Change your database configuration
// Change your database configuration. Available for PostgreSQL, EnterpriseDB, MySQL, and Redis ONLY.
func (cloudDatabases *CloudDatabasesV5) UpdateDatabaseConfiguration(updateDatabaseConfigurationOptions *UpdateDatabaseConfigurationOptions) (result *UpdateDatabaseConfigurationResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.UpdateDatabaseConfigurationWithContext(context.Background(), updateDatabaseConfigurationOptions)
}

// UpdateDatabaseConfigurationWithContext is an alternate form of the UpdateDatabaseConfiguration method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) UpdateDatabaseConfigurationWithContext(ctx context.Context, updateDatabaseConfigurationOptions *UpdateDatabaseConfigurationOptions) (result *UpdateDatabaseConfigurationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDatabaseConfigurationOptions, "updateDatabaseConfigurationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDatabaseConfigurationOptions, "updateDatabaseConfigurationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateDatabaseConfigurationOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/configuration`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDatabaseConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "UpdateDatabaseConfiguration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateDatabaseConfigurationOptions.Configuration != nil {
		body["configuration"] = updateDatabaseConfigurationOptions.Configuration
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUpdateDatabaseConfigurationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListRemotes : List read-only replica information
// Get the read-only replicas associated with a deployment. Available for PostgreSQL and EnterpriseDB ONLY.
func (cloudDatabases *CloudDatabasesV5) ListRemotes(listRemotesOptions *ListRemotesOptions) (result *ListRemotesResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.ListRemotesWithContext(context.Background(), listRemotesOptions)
}

// ListRemotesWithContext is an alternate form of the ListRemotes method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) ListRemotesWithContext(ctx context.Context, listRemotesOptions *ListRemotesOptions) (result *ListRemotesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listRemotesOptions, "listRemotesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listRemotesOptions, "listRemotesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *listRemotesOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/remotes`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listRemotesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "ListRemotes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListRemotesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ResyncReplica : Resync read-only replica
// Reinitialize a read-only replica. Available for PostgreSQL and EnterpriseDB ONLY.
func (cloudDatabases *CloudDatabasesV5) ResyncReplica(resyncReplicaOptions *ResyncReplicaOptions) (result *ResyncReplicaResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.ResyncReplicaWithContext(context.Background(), resyncReplicaOptions)
}

// ResyncReplicaWithContext is an alternate form of the ResyncReplica method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) ResyncReplicaWithContext(ctx context.Context, resyncReplicaOptions *ResyncReplicaOptions) (result *ResyncReplicaResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(resyncReplicaOptions, "resyncReplicaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(resyncReplicaOptions, "resyncReplicaOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *resyncReplicaOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/remotes/resync`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range resyncReplicaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "ResyncReplica")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResyncReplicaResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PromoteReadOnlyReplica : Promote read-only replica to a full deployment
// Promote a read-only replica or upgrade and promote a read-only replica. Available for PostgreSQL and EnterpriseDB
// ONLY.
func (cloudDatabases *CloudDatabasesV5) PromoteReadOnlyReplica(promoteReadOnlyReplicaOptions *PromoteReadOnlyReplicaOptions) (result *PromoteReadOnlyReplicaResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.PromoteReadOnlyReplicaWithContext(context.Background(), promoteReadOnlyReplicaOptions)
}

// PromoteReadOnlyReplicaWithContext is an alternate form of the PromoteReadOnlyReplica method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) PromoteReadOnlyReplicaWithContext(ctx context.Context, promoteReadOnlyReplicaOptions *PromoteReadOnlyReplicaOptions) (result *PromoteReadOnlyReplicaResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(promoteReadOnlyReplicaOptions, "promoteReadOnlyReplicaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(promoteReadOnlyReplicaOptions, "promoteReadOnlyReplicaOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *promoteReadOnlyReplicaOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/remotes/promotion`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range promoteReadOnlyReplicaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "PromoteReadOnlyReplica")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if promoteReadOnlyReplicaOptions.Promotion != nil {
		body["promotion"] = promoteReadOnlyReplicaOptions.Promotion
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPromoteReadOnlyReplicaResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDeploymentTasks : List currently running tasks on a deployment
// Obtain a list of tasks currently running or recently run on a deployment. Tasks are ephemeral. Records of successful
// tasks are shown for 24-48 hours, and unsuccessful tasks are shown for 7-8 days.
func (cloudDatabases *CloudDatabasesV5) ListDeploymentTasks(listDeploymentTasksOptions *ListDeploymentTasksOptions) (result *Tasks, response *core.DetailedResponse, err error) {
	return cloudDatabases.ListDeploymentTasksWithContext(context.Background(), listDeploymentTasksOptions)
}

// ListDeploymentTasksWithContext is an alternate form of the ListDeploymentTasks method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) ListDeploymentTasksWithContext(ctx context.Context, listDeploymentTasksOptions *ListDeploymentTasksOptions) (result *Tasks, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDeploymentTasksOptions, "listDeploymentTasksOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listDeploymentTasksOptions, "listDeploymentTasksOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *listDeploymentTasksOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/tasks`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDeploymentTasksOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "ListDeploymentTasks")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTasks)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetTask : Get information about a task
// Get information about a task and its status. Tasks themselves are persistent so old tasks can be consulted as well as
// running tasks.
func (cloudDatabases *CloudDatabasesV5) GetTask(getTaskOptions *GetTaskOptions) (result *GetTaskResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetTaskWithContext(context.Background(), getTaskOptions)
}

// GetTaskWithContext is an alternate form of the GetTask method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetTaskWithContext(ctx context.Context, getTaskOptions *GetTaskOptions) (result *GetTaskResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTaskOptions, "getTaskOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTaskOptions, "getTaskOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getTaskOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/tasks/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTaskOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetTask")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetTaskResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetBackupInfo : Get information about a backup
// Get information about a backup, such as creation date.
func (cloudDatabases *CloudDatabasesV5) GetBackupInfo(getBackupInfoOptions *GetBackupInfoOptions) (result *GetBackupInfoResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetBackupInfoWithContext(context.Background(), getBackupInfoOptions)
}

// GetBackupInfoWithContext is an alternate form of the GetBackupInfo method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetBackupInfoWithContext(ctx context.Context, getBackupInfoOptions *GetBackupInfoOptions) (result *GetBackupInfoResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBackupInfoOptions, "getBackupInfoOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBackupInfoOptions, "getBackupInfoOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"backup_id": *getBackupInfoOptions.BackupID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/backups/{backup_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBackupInfoOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetBackupInfo")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetBackupInfoResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDeploymentBackups : List currently available backups from a deployment
// Get details of all currently available backups from a deployment.
func (cloudDatabases *CloudDatabasesV5) ListDeploymentBackups(listDeploymentBackupsOptions *ListDeploymentBackupsOptions) (result *Backups, response *core.DetailedResponse, err error) {
	return cloudDatabases.ListDeploymentBackupsWithContext(context.Background(), listDeploymentBackupsOptions)
}

// ListDeploymentBackupsWithContext is an alternate form of the ListDeploymentBackups method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) ListDeploymentBackupsWithContext(ctx context.Context, listDeploymentBackupsOptions *ListDeploymentBackupsOptions) (result *Backups, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDeploymentBackupsOptions, "listDeploymentBackupsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listDeploymentBackupsOptions, "listDeploymentBackupsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *listDeploymentBackupsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/backups`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDeploymentBackupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "ListDeploymentBackups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBackups)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// StartOndemandBackup : Initiate an on-demand backup
// Signal the platform to create an on-demand backup for the specified deployment. The returned task can be polled to
// track progress of the backup as it takes place.
func (cloudDatabases *CloudDatabasesV5) StartOndemandBackup(startOndemandBackupOptions *StartOndemandBackupOptions) (result *StartOndemandBackupResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.StartOndemandBackupWithContext(context.Background(), startOndemandBackupOptions)
}

// StartOndemandBackupWithContext is an alternate form of the StartOndemandBackup method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) StartOndemandBackupWithContext(ctx context.Context, startOndemandBackupOptions *StartOndemandBackupOptions) (result *StartOndemandBackupResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(startOndemandBackupOptions, "startOndemandBackupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(startOndemandBackupOptions, "startOndemandBackupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *startOndemandBackupOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/backups`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range startOndemandBackupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "StartOndemandBackup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalStartOndemandBackupResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetPitrData : Get earliest point-in-time-recovery timestamp
// Returns the earliest available time for point-in-time-recovery in ISO8601 UTC format. PostgreSQL and EnterpriseDB
// only.
func (cloudDatabases *CloudDatabasesV5) GetPitrData(getPitrDataOptions *GetPitrDataOptions) (result *GetPitrDataResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetPitrDataWithContext(context.Background(), getPitrDataOptions)
}

// GetPitrDataWithContext is an alternate form of the GetPitrData method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetPitrDataWithContext(ctx context.Context, getPitrDataOptions *GetPitrDataOptions) (result *GetPitrDataResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPitrDataOptions, "getPitrDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPitrDataOptions, "getPitrDataOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getPitrDataOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/point_in_time_recovery_data`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPitrDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetPitrData")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetPitrDataResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetConnection : Discover connection information for a deployment for a user with an endpoint type
// Discover connection information for a deployment for a user with an endpoint type.
func (cloudDatabases *CloudDatabasesV5) GetConnection(getConnectionOptions *GetConnectionOptions) (result *GetConnectionResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetConnectionWithContext(context.Background(), getConnectionOptions)
}

// GetConnectionWithContext is an alternate form of the GetConnection method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetConnectionWithContext(ctx context.Context, getConnectionOptions *GetConnectionOptions) (result *GetConnectionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getConnectionOptions, "getConnectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getConnectionOptions, "getConnectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getConnectionOptions.ID,
		"user_type": *getConnectionOptions.UserType,
		"user_id": *getConnectionOptions.UserID,
		"endpoint_type": *getConnectionOptions.EndpointType,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/users/{user_type}/{user_id}/connections/{endpoint_type}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConnectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetConnection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getConnectionOptions.CertificateRoot != nil {
		builder.AddQuery("certificate_root", fmt.Sprint(*getConnectionOptions.CertificateRoot))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetConnectionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CompleteConnection : Discover connection information for a deployment for a user with substitutions and an endpoint type
// Discover connection information for a deployment for a user. Behaves the same as the GET method but substitutes the
// provided password parameter into the returned connection information.
func (cloudDatabases *CloudDatabasesV5) CompleteConnection(completeConnectionOptions *CompleteConnectionOptions) (result *CompleteConnectionResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.CompleteConnectionWithContext(context.Background(), completeConnectionOptions)
}

// CompleteConnectionWithContext is an alternate form of the CompleteConnection method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) CompleteConnectionWithContext(ctx context.Context, completeConnectionOptions *CompleteConnectionOptions) (result *CompleteConnectionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(completeConnectionOptions, "completeConnectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(completeConnectionOptions, "completeConnectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *completeConnectionOptions.ID,
		"user_type": *completeConnectionOptions.UserType,
		"user_id": *completeConnectionOptions.UserID,
		"endpoint_type": *completeConnectionOptions.EndpointType,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/users/{user_type}/{user_id}/connections/{endpoint_type}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range completeConnectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "CompleteConnection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if completeConnectionOptions.Password != nil {
		body["password"] = completeConnectionOptions.Password
	}
	if completeConnectionOptions.CertificateRoot != nil {
		body["certificate_root"] = completeConnectionOptions.CertificateRoot
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCompleteConnectionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDeploymentScalingGroups : List currently available scaling groups from a deployment
// Scaling groups represent the various resources that are allocated to a deployment. This command allows for the
// retrieval of all of the groups for a particular deployment.
func (cloudDatabases *CloudDatabasesV5) ListDeploymentScalingGroups(listDeploymentScalingGroupsOptions *ListDeploymentScalingGroupsOptions) (result *ListDeploymentScalingGroupsResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.ListDeploymentScalingGroupsWithContext(context.Background(), listDeploymentScalingGroupsOptions)
}

// ListDeploymentScalingGroupsWithContext is an alternate form of the ListDeploymentScalingGroups method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) ListDeploymentScalingGroupsWithContext(ctx context.Context, listDeploymentScalingGroupsOptions *ListDeploymentScalingGroupsOptions) (result *ListDeploymentScalingGroupsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDeploymentScalingGroupsOptions, "listDeploymentScalingGroupsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listDeploymentScalingGroupsOptions, "listDeploymentScalingGroupsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *listDeploymentScalingGroupsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/groups`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDeploymentScalingGroupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "ListDeploymentScalingGroups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListDeploymentScalingGroupsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDefaultScalingGroups : Get default scaling groups for a new deployment
// Scaling groups represent the various resources allocated to a deployment. When a new deployment is created, there are
// a set of defaults for each database type. This endpoint returns them for a particular database.
func (cloudDatabases *CloudDatabasesV5) GetDefaultScalingGroups(getDefaultScalingGroupsOptions *GetDefaultScalingGroupsOptions) (result *GetDefaultScalingGroupsResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetDefaultScalingGroupsWithContext(context.Background(), getDefaultScalingGroupsOptions)
}

// GetDefaultScalingGroupsWithContext is an alternate form of the GetDefaultScalingGroups method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetDefaultScalingGroupsWithContext(ctx context.Context, getDefaultScalingGroupsOptions *GetDefaultScalingGroupsOptions) (result *GetDefaultScalingGroupsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDefaultScalingGroupsOptions, "getDefaultScalingGroupsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDefaultScalingGroupsOptions, "getDefaultScalingGroupsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"type": *getDefaultScalingGroupsOptions.Type,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployables/{type}/groups`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDefaultScalingGroupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetDefaultScalingGroups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetDefaultScalingGroupsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetDeploymentScalingGroup : Set scaling values on a specified group
// Set scaling value on a specified group. Can only be performed on is_adjustable=true groups. Values set are for the
// group as a whole and resources are distributed amongst the group. Values must be greater than or equal to the minimum
// size and must be a multiple of the step size.
func (cloudDatabases *CloudDatabasesV5) SetDeploymentScalingGroup(setDeploymentScalingGroupOptions *SetDeploymentScalingGroupOptions) (result *SetDeploymentScalingGroupResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.SetDeploymentScalingGroupWithContext(context.Background(), setDeploymentScalingGroupOptions)
}

// SetDeploymentScalingGroupWithContext is an alternate form of the SetDeploymentScalingGroup method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) SetDeploymentScalingGroupWithContext(ctx context.Context, setDeploymentScalingGroupOptions *SetDeploymentScalingGroupOptions) (result *SetDeploymentScalingGroupResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setDeploymentScalingGroupOptions, "setDeploymentScalingGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setDeploymentScalingGroupOptions, "setDeploymentScalingGroupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *setDeploymentScalingGroupOptions.ID,
		"group_id": *setDeploymentScalingGroupOptions.GroupID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/groups/{group_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setDeploymentScalingGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "SetDeploymentScalingGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setDeploymentScalingGroupOptions.Group != nil {
		body["group"] = setDeploymentScalingGroupOptions.Group
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSetDeploymentScalingGroupResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAutoscalingConditions : Get the autoscaling configuration from a deployment
// The Autoscaling configuration represents the various conditions that control autoscaling for a deployment. This
// command allows for the retrieval of all autoscaling conditions for a particular deployment.
func (cloudDatabases *CloudDatabasesV5) GetAutoscalingConditions(getAutoscalingConditionsOptions *GetAutoscalingConditionsOptions) (result *AutoscalingGroup, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetAutoscalingConditionsWithContext(context.Background(), getAutoscalingConditionsOptions)
}

// GetAutoscalingConditionsWithContext is an alternate form of the GetAutoscalingConditions method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetAutoscalingConditionsWithContext(ctx context.Context, getAutoscalingConditionsOptions *GetAutoscalingConditionsOptions) (result *AutoscalingGroup, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAutoscalingConditionsOptions, "getAutoscalingConditionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAutoscalingConditionsOptions, "getAutoscalingConditionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getAutoscalingConditionsOptions.ID,
		"group_id": *getAutoscalingConditionsOptions.GroupID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/groups/{group_id}/autoscaling`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAutoscalingConditionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetAutoscalingConditions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAutoscalingGroup)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetAutoscalingConditions : Set the autoscaling configuration from a deployment
// Enable, disable, or set the conditions for autoscaling on your deployment. Memory, disk, and CPU (if available) can
// be set separately and are not all required.
func (cloudDatabases *CloudDatabasesV5) SetAutoscalingConditions(setAutoscalingConditionsOptions *SetAutoscalingConditionsOptions) (result *SetAutoscalingConditionsResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.SetAutoscalingConditionsWithContext(context.Background(), setAutoscalingConditionsOptions)
}

// SetAutoscalingConditionsWithContext is an alternate form of the SetAutoscalingConditions method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) SetAutoscalingConditionsWithContext(ctx context.Context, setAutoscalingConditionsOptions *SetAutoscalingConditionsOptions) (result *SetAutoscalingConditionsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setAutoscalingConditionsOptions, "setAutoscalingConditionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setAutoscalingConditionsOptions, "setAutoscalingConditionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *setAutoscalingConditionsOptions.ID,
		"group_id": *setAutoscalingConditionsOptions.GroupID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/groups/{group_id}/autoscaling`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setAutoscalingConditionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "SetAutoscalingConditions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setAutoscalingConditionsOptions.Autoscaling != nil {
		body["autoscaling"] = setAutoscalingConditionsOptions.Autoscaling
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSetAutoscalingConditionsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// KillConnections : Kill connections to a PostgreSQL or EnterpriseDB deployment
// Closes all the connections on a deployment. Available for PostgreSQL and EnterpriseDB ONLY.
func (cloudDatabases *CloudDatabasesV5) KillConnections(killConnectionsOptions *KillConnectionsOptions) (result *KillConnectionsResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.KillConnectionsWithContext(context.Background(), killConnectionsOptions)
}

// KillConnectionsWithContext is an alternate form of the KillConnections method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) KillConnectionsWithContext(ctx context.Context, killConnectionsOptions *KillConnectionsOptions) (result *KillConnectionsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(killConnectionsOptions, "killConnectionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(killConnectionsOptions, "killConnectionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *killConnectionsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/management/database_connections`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range killConnectionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "KillConnections")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalKillConnectionsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAllowlist : Retrieve the allowlisted addresses and ranges for a deployment
// Retrieve the allowlisted addresses and ranges for a deployment.
func (cloudDatabases *CloudDatabasesV5) GetAllowlist(getAllowlistOptions *GetAllowlistOptions) (result *GetAllowlistResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetAllowlistWithContext(context.Background(), getAllowlistOptions)
}

// GetAllowlistWithContext is an alternate form of the GetAllowlist method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetAllowlistWithContext(ctx context.Context, getAllowlistOptions *GetAllowlistOptions) (result *GetAllowlistResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAllowlistOptions, "getAllowlistOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAllowlistOptions, "getAllowlistOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getAllowlistOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/allowlists/ip_addresses`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAllowlistOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetAllowlist")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetAllowlistResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetAllowlist : Set the allowlist for a deployment
// Set the allowlist for a deployment. This action overwrites all existing entries, so when you modify the allowlist via
// a GET/update/PUT, provide the GET response's ETag header value in this endpoint's If-Match header to ensure that
// changes that are made by other clients are not accidentally overwritten.
func (cloudDatabases *CloudDatabasesV5) SetAllowlist(setAllowlistOptions *SetAllowlistOptions) (result *SetAllowlistResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.SetAllowlistWithContext(context.Background(), setAllowlistOptions)
}

// SetAllowlistWithContext is an alternate form of the SetAllowlist method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) SetAllowlistWithContext(ctx context.Context, setAllowlistOptions *SetAllowlistOptions) (result *SetAllowlistResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setAllowlistOptions, "setAllowlistOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setAllowlistOptions, "setAllowlistOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *setAllowlistOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/allowlists/ip_addresses`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setAllowlistOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "SetAllowlist")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if setAllowlistOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*setAllowlistOptions.IfMatch))
	}

	body := make(map[string]interface{})
	if setAllowlistOptions.IPAddresses != nil {
		body["ip_addresses"] = setAllowlistOptions.IPAddresses
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSetAllowlistResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddAllowlistEntry : Add an address or range to the allowlist for a deployment
// Add an address or range to the allowlist for a deployment.
func (cloudDatabases *CloudDatabasesV5) AddAllowlistEntry(addAllowlistEntryOptions *AddAllowlistEntryOptions) (result *AddAllowlistEntryResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.AddAllowlistEntryWithContext(context.Background(), addAllowlistEntryOptions)
}

// AddAllowlistEntryWithContext is an alternate form of the AddAllowlistEntry method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) AddAllowlistEntryWithContext(ctx context.Context, addAllowlistEntryOptions *AddAllowlistEntryOptions) (result *AddAllowlistEntryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addAllowlistEntryOptions, "addAllowlistEntryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addAllowlistEntryOptions, "addAllowlistEntryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *addAllowlistEntryOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/allowlists/ip_addresses`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addAllowlistEntryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "AddAllowlistEntry")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addAllowlistEntryOptions.IPAddress != nil {
		body["ip_address"] = addAllowlistEntryOptions.IPAddress
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAddAllowlistEntryResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteAllowlistEntry : Delete an address or range from the allowlist of a deployment
// Delete an address or range from the allowlist of a deployment.
func (cloudDatabases *CloudDatabasesV5) DeleteAllowlistEntry(deleteAllowlistEntryOptions *DeleteAllowlistEntryOptions) (result *DeleteAllowlistEntryResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.DeleteAllowlistEntryWithContext(context.Background(), deleteAllowlistEntryOptions)
}

// DeleteAllowlistEntryWithContext is an alternate form of the DeleteAllowlistEntry method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) DeleteAllowlistEntryWithContext(ctx context.Context, deleteAllowlistEntryOptions *DeleteAllowlistEntryOptions) (result *DeleteAllowlistEntryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAllowlistEntryOptions, "deleteAllowlistEntryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAllowlistEntryOptions, "deleteAllowlistEntryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteAllowlistEntryOptions.ID,
		"ipaddress": *deleteAllowlistEntryOptions.Ipaddress,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/allowlists/ip_addresses/{ipaddress}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAllowlistEntryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "DeleteAllowlistEntry")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteAllowlistEntryResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAutoscalingCapability : Discover autoscaling capability information from a deployment
// Returns autoscaling capability information based on the parameters provided. Also returns request parameters or
// defaults if request parameters are blank.
func (cloudDatabases *CloudDatabasesV5) GetAutoscalingCapability(getAutoscalingCapabilityOptions *GetAutoscalingCapabilityOptions) (result *GetAutoscalingCapabilityResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetAutoscalingCapabilityWithContext(context.Background(), getAutoscalingCapabilityOptions)
}

// GetAutoscalingCapabilityWithContext is an alternate form of the GetAutoscalingCapability method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetAutoscalingCapabilityWithContext(ctx context.Context, getAutoscalingCapabilityOptions *GetAutoscalingCapabilityOptions) (result *GetAutoscalingCapabilityResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAutoscalingCapabilityOptions, "getAutoscalingCapabilityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAutoscalingCapabilityOptions, "getAutoscalingCapabilityOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getAutoscalingCapabilityOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/capability/autoscaling`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAutoscalingCapabilityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetAutoscalingCapability")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if getAutoscalingCapabilityOptions.Deployment != nil {
		body["deployment"] = getAutoscalingCapabilityOptions.Deployment
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetAutoscalingCapabilityResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetEncryptionCapability : Discover encryption capability information from a deployment
// Returns encryption capability information based on the parameters provided. Also returns request parameters or
// defaults if request parameters are blank.
func (cloudDatabases *CloudDatabasesV5) GetEncryptionCapability(getEncryptionCapabilityOptions *GetEncryptionCapabilityOptions) (result *GetEncryptionCapabilityResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetEncryptionCapabilityWithContext(context.Background(), getEncryptionCapabilityOptions)
}

// GetEncryptionCapabilityWithContext is an alternate form of the GetEncryptionCapability method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetEncryptionCapabilityWithContext(ctx context.Context, getEncryptionCapabilityOptions *GetEncryptionCapabilityOptions) (result *GetEncryptionCapabilityResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getEncryptionCapabilityOptions, "getEncryptionCapabilityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getEncryptionCapabilityOptions, "getEncryptionCapabilityOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getEncryptionCapabilityOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/capability/encryption`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getEncryptionCapabilityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetEncryptionCapability")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if getEncryptionCapabilityOptions.Deployment != nil {
		body["deployment"] = getEncryptionCapabilityOptions.Deployment
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetEncryptionCapabilityResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetEndpointsCapability : Discover endpoints capability information from a deployment
// Returns endpoints capability information based on the parameters provided. Also returns request parameters or
// defaults if request parameters are blank.
func (cloudDatabases *CloudDatabasesV5) GetEndpointsCapability(getEndpointsCapabilityOptions *GetEndpointsCapabilityOptions) (result *GetEndpointsCapabilityResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetEndpointsCapabilityWithContext(context.Background(), getEndpointsCapabilityOptions)
}

// GetEndpointsCapabilityWithContext is an alternate form of the GetEndpointsCapability method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetEndpointsCapabilityWithContext(ctx context.Context, getEndpointsCapabilityOptions *GetEndpointsCapabilityOptions) (result *GetEndpointsCapabilityResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getEndpointsCapabilityOptions, "getEndpointsCapabilityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getEndpointsCapabilityOptions, "getEndpointsCapabilityOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getEndpointsCapabilityOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/capability/endpoints`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getEndpointsCapabilityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetEndpointsCapability")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if getEndpointsCapabilityOptions.Deployment != nil {
		body["deployment"] = getEndpointsCapabilityOptions.Deployment
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetEndpointsCapabilityResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetGroupsCapability : Discover groups capability information from a deployment
// Returns groups capability information based on the parameters provided. Also returns request parameters or defaults
// if request parameters are blank.
func (cloudDatabases *CloudDatabasesV5) GetGroupsCapability(getGroupsCapabilityOptions *GetGroupsCapabilityOptions) (result *GetGroupsCapabilityResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetGroupsCapabilityWithContext(context.Background(), getGroupsCapabilityOptions)
}

// GetGroupsCapabilityWithContext is an alternate form of the GetGroupsCapability method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetGroupsCapabilityWithContext(ctx context.Context, getGroupsCapabilityOptions *GetGroupsCapabilityOptions) (result *GetGroupsCapabilityResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getGroupsCapabilityOptions, "getGroupsCapabilityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getGroupsCapabilityOptions, "getGroupsCapabilityOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getGroupsCapabilityOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/capability/groups`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getGroupsCapabilityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetGroupsCapability")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if getGroupsCapabilityOptions.Deployment != nil {
		body["deployment"] = getGroupsCapabilityOptions.Deployment
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetGroupsCapabilityResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetRegionsCapability : Discover regions capability information for a deployment
// Returns regions capability information based on the parameters provided. Also returns request parameters or defaults
// if request parameters are blank.
func (cloudDatabases *CloudDatabasesV5) GetRegionsCapability(getRegionsCapabilityOptions *GetRegionsCapabilityOptions) (result *GetRegionsCapabilityResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetRegionsCapabilityWithContext(context.Background(), getRegionsCapabilityOptions)
}

// GetRegionsCapabilityWithContext is an alternate form of the GetRegionsCapability method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetRegionsCapabilityWithContext(ctx context.Context, getRegionsCapabilityOptions *GetRegionsCapabilityOptions) (result *GetRegionsCapabilityResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRegionsCapabilityOptions, "getRegionsCapabilityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRegionsCapabilityOptions, "getRegionsCapabilityOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getRegionsCapabilityOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/capability/regions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRegionsCapabilityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetRegionsCapability")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if getRegionsCapabilityOptions.Deployment != nil {
		body["deployment"] = getRegionsCapabilityOptions.Deployment
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetRegionsCapabilityResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetRemotesCapability : Discover remotes capability information for a deployment
// Returns remotes capability information based on the parameters provided. Also returns request parameters or defaults
// if request parameters are blank.
func (cloudDatabases *CloudDatabasesV5) GetRemotesCapability(getRemotesCapabilityOptions *GetRemotesCapabilityOptions) (result *GetRemotesCapabilityResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetRemotesCapabilityWithContext(context.Background(), getRemotesCapabilityOptions)
}

// GetRemotesCapabilityWithContext is an alternate form of the GetRemotesCapability method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetRemotesCapabilityWithContext(ctx context.Context, getRemotesCapabilityOptions *GetRemotesCapabilityOptions) (result *GetRemotesCapabilityResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRemotesCapabilityOptions, "getRemotesCapabilityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRemotesCapabilityOptions, "getRemotesCapabilityOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getRemotesCapabilityOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/capability/remotes`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRemotesCapabilityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetRemotesCapability")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if getRemotesCapabilityOptions.Deployment != nil {
		body["deployment"] = getRemotesCapabilityOptions.Deployment
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetRemotesCapabilityResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetVersionsCapability : Discover versions capability information for a deployment
// Returns versions capability information based on the parameters provided. Also returns request parameters or defaults
// if request parameters are blank.
func (cloudDatabases *CloudDatabasesV5) GetVersionsCapability(getVersionsCapabilityOptions *GetVersionsCapabilityOptions) (result *GetVersionsCapabilityResponse, response *core.DetailedResponse, err error) {
	return cloudDatabases.GetVersionsCapabilityWithContext(context.Background(), getVersionsCapabilityOptions)
}

// GetVersionsCapabilityWithContext is an alternate form of the GetVersionsCapability method which supports a Context parameter
func (cloudDatabases *CloudDatabasesV5) GetVersionsCapabilityWithContext(ctx context.Context, getVersionsCapabilityOptions *GetVersionsCapabilityOptions) (result *GetVersionsCapabilityResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVersionsCapabilityOptions, "getVersionsCapabilityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVersionsCapabilityOptions, "getVersionsCapabilityOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getVersionsCapabilityOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = cloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(cloudDatabases.Service.Options.URL, `/deployments/{id}/capability/versions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVersionsCapabilityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cloud_databases", "V5", "GetVersionsCapability")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if getVersionsCapabilityOptions.Deployment != nil {
		body["deployment"] = getVersionsCapabilityOptions.Deployment
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetVersionsCapabilityResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// APasswordSettingUser : APasswordSettingUser struct
type APasswordSettingUser struct {
	Password *string `json:"password,omitempty"`
}

// UnmarshalAPasswordSettingUser unmarshals an instance of APasswordSettingUser from the specified map of raw messages.
func UnmarshalAPasswordSettingUser(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(APasswordSettingUser)
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddAllowlistEntryOptions : The AddAllowlistEntry options.
type AddAllowlistEntryOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	IPAddress *AllowlistEntry `json:"ip_address,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddAllowlistEntryOptions : Instantiate AddAllowlistEntryOptions
func (*CloudDatabasesV5) NewAddAllowlistEntryOptions(id string) *AddAllowlistEntryOptions {
	return &AddAllowlistEntryOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *AddAllowlistEntryOptions) SetID(id string) *AddAllowlistEntryOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIPAddress : Allow user to set IPAddress
func (_options *AddAllowlistEntryOptions) SetIPAddress(ipAddress *AllowlistEntry) *AddAllowlistEntryOptions {
	_options.IPAddress = ipAddress
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddAllowlistEntryOptions) SetHeaders(param map[string]string) *AddAllowlistEntryOptions {
	options.Headers = param
	return options
}

// AddAllowlistEntryResponse : AddAllowlistEntryResponse struct
type AddAllowlistEntryResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalAddAllowlistEntryResponse unmarshals an instance of AddAllowlistEntryResponse from the specified map of raw messages.
func UnmarshalAddAllowlistEntryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddAllowlistEntryResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AllowlistEntry : AllowlistEntry struct
type AllowlistEntry struct {
	// An IPv4 address or a CIDR range (netmasked IPv4 address).
	Address *string `json:"address,omitempty"`

	// A human readable description of the address or range for identification purposes.
	Description *string `json:"description,omitempty"`
}

// UnmarshalAllowlistEntry unmarshals an instance of AllowlistEntry from the specified map of raw messages.
func UnmarshalAllowlistEntry(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AllowlistEntry)
	err = core.UnmarshalPrimitive(m, "address", &obj.Address)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingCPUGroupCPU : AutoscalingCPUGroupCPU struct
type AutoscalingCPUGroupCPU struct {
	Scalers interface{} `json:"scalers,omitempty"`

	Rate *AutoscalingCPUGroupCPURate `json:"rate,omitempty"`
}

// UnmarshalAutoscalingCPUGroupCPU unmarshals an instance of AutoscalingCPUGroupCPU from the specified map of raw messages.
func UnmarshalAutoscalingCPUGroupCPU(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingCPUGroupCPU)
	err = core.UnmarshalPrimitive(m, "scalers", &obj.Scalers)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rate", &obj.Rate, UnmarshalAutoscalingCPUGroupCPURate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingCPUGroupCPURate : AutoscalingCPUGroupCPURate struct
type AutoscalingCPUGroupCPURate struct {
	IncreasePercent *float64 `json:"increase_percent,omitempty"`

	PeriodSeconds *int64 `json:"period_seconds,omitempty"`

	LimitCountPerMember *int64 `json:"limit_count_per_member,omitempty"`

	Units *string `json:"units,omitempty"`
}

// UnmarshalAutoscalingCPUGroupCPURate unmarshals an instance of AutoscalingCPUGroupCPURate from the specified map of raw messages.
func UnmarshalAutoscalingCPUGroupCPURate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingCPUGroupCPURate)
	err = core.UnmarshalPrimitive(m, "increase_percent", &obj.IncreasePercent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "period_seconds", &obj.PeriodSeconds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit_count_per_member", &obj.LimitCountPerMember)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingDiskGroupDisk : AutoscalingDiskGroupDisk struct
type AutoscalingDiskGroupDisk struct {
	Scalers *AutoscalingDiskGroupDiskScalers `json:"scalers,omitempty"`

	Rate *AutoscalingDiskGroupDiskRate `json:"rate,omitempty"`
}

// UnmarshalAutoscalingDiskGroupDisk unmarshals an instance of AutoscalingDiskGroupDisk from the specified map of raw messages.
func UnmarshalAutoscalingDiskGroupDisk(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingDiskGroupDisk)
	err = core.UnmarshalModel(m, "scalers", &obj.Scalers, UnmarshalAutoscalingDiskGroupDiskScalers)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rate", &obj.Rate, UnmarshalAutoscalingDiskGroupDiskRate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingDiskGroupDiskRate : AutoscalingDiskGroupDiskRate struct
type AutoscalingDiskGroupDiskRate struct {
	IncreasePercent *float64 `json:"increase_percent,omitempty"`

	PeriodSeconds *int64 `json:"period_seconds,omitempty"`

	LimitMbPerMember *float64 `json:"limit_mb_per_member,omitempty"`

	Units *string `json:"units,omitempty"`
}

// UnmarshalAutoscalingDiskGroupDiskRate unmarshals an instance of AutoscalingDiskGroupDiskRate from the specified map of raw messages.
func UnmarshalAutoscalingDiskGroupDiskRate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingDiskGroupDiskRate)
	err = core.UnmarshalPrimitive(m, "increase_percent", &obj.IncreasePercent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "period_seconds", &obj.PeriodSeconds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit_mb_per_member", &obj.LimitMbPerMember)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingDiskGroupDiskScalers : AutoscalingDiskGroupDiskScalers struct
type AutoscalingDiskGroupDiskScalers struct {
	Capacity *AutoscalingDiskGroupDiskScalersCapacity `json:"capacity,omitempty"`

	IoUtilization *AutoscalingDiskGroupDiskScalersIoUtilization `json:"io_utilization,omitempty"`
}

// UnmarshalAutoscalingDiskGroupDiskScalers unmarshals an instance of AutoscalingDiskGroupDiskScalers from the specified map of raw messages.
func UnmarshalAutoscalingDiskGroupDiskScalers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingDiskGroupDiskScalers)
	err = core.UnmarshalModel(m, "capacity", &obj.Capacity, UnmarshalAutoscalingDiskGroupDiskScalersCapacity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "io_utilization", &obj.IoUtilization, UnmarshalAutoscalingDiskGroupDiskScalersIoUtilization)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingDiskGroupDiskScalersCapacity : AutoscalingDiskGroupDiskScalersCapacity struct
type AutoscalingDiskGroupDiskScalersCapacity struct {
	Enabled *bool `json:"enabled,omitempty"`

	FreeSpaceLessThanPercent *int64 `json:"free_space_less_than_percent,omitempty"`
}

// UnmarshalAutoscalingDiskGroupDiskScalersCapacity unmarshals an instance of AutoscalingDiskGroupDiskScalersCapacity from the specified map of raw messages.
func UnmarshalAutoscalingDiskGroupDiskScalersCapacity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingDiskGroupDiskScalersCapacity)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "free_space_less_than_percent", &obj.FreeSpaceLessThanPercent)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingDiskGroupDiskScalersIoUtilization : AutoscalingDiskGroupDiskScalersIoUtilization struct
type AutoscalingDiskGroupDiskScalersIoUtilization struct {
	Enabled *bool `json:"enabled,omitempty"`

	OverPeriod *string `json:"over_period,omitempty"`

	AbovePercent *int64 `json:"above_percent,omitempty"`
}

// UnmarshalAutoscalingDiskGroupDiskScalersIoUtilization unmarshals an instance of AutoscalingDiskGroupDiskScalersIoUtilization from the specified map of raw messages.
func UnmarshalAutoscalingDiskGroupDiskScalersIoUtilization(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingDiskGroupDiskScalersIoUtilization)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "over_period", &obj.OverPeriod)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "above_percent", &obj.AbovePercent)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingGroup : AutoscalingGroup struct
type AutoscalingGroup struct {
	Autoscaling *AutoscalingGroupAutoscaling `json:"autoscaling" validate:"required"`
}

// UnmarshalAutoscalingGroup unmarshals an instance of AutoscalingGroup from the specified map of raw messages.
func UnmarshalAutoscalingGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingGroup)
	err = core.UnmarshalModel(m, "autoscaling", &obj.Autoscaling, UnmarshalAutoscalingGroupAutoscaling)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingGroupAutoscaling : AutoscalingGroupAutoscaling struct
type AutoscalingGroupAutoscaling struct {
	Disk *AutoscalingDiskGroupDisk `json:"disk,omitempty"`

	Memory *AutoscalingMemoryGroupMemory `json:"memory,omitempty"`

	CPU *AutoscalingCPUGroupCPU `json:"cpu,omitempty"`
}

// UnmarshalAutoscalingGroupAutoscaling unmarshals an instance of AutoscalingGroupAutoscaling from the specified map of raw messages.
func UnmarshalAutoscalingGroupAutoscaling(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingGroupAutoscaling)
	err = core.UnmarshalModel(m, "disk", &obj.Disk, UnmarshalAutoscalingDiskGroupDisk)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "memory", &obj.Memory, UnmarshalAutoscalingMemoryGroupMemory)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cpu", &obj.CPU, UnmarshalAutoscalingCPUGroupCPU)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingMemoryGroupMemory : AutoscalingMemoryGroupMemory struct
type AutoscalingMemoryGroupMemory struct {
	Scalers *AutoscalingMemoryGroupMemoryScalers `json:"scalers,omitempty"`

	Rate *AutoscalingMemoryGroupMemoryRate `json:"rate,omitempty"`
}

// UnmarshalAutoscalingMemoryGroupMemory unmarshals an instance of AutoscalingMemoryGroupMemory from the specified map of raw messages.
func UnmarshalAutoscalingMemoryGroupMemory(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingMemoryGroupMemory)
	err = core.UnmarshalModel(m, "scalers", &obj.Scalers, UnmarshalAutoscalingMemoryGroupMemoryScalers)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rate", &obj.Rate, UnmarshalAutoscalingMemoryGroupMemoryRate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingMemoryGroupMemoryRate : AutoscalingMemoryGroupMemoryRate struct
type AutoscalingMemoryGroupMemoryRate struct {
	IncreasePercent *float64 `json:"increase_percent,omitempty"`

	PeriodSeconds *int64 `json:"period_seconds,omitempty"`

	LimitMbPerMember *float64 `json:"limit_mb_per_member,omitempty"`

	Units *string `json:"units,omitempty"`
}

// UnmarshalAutoscalingMemoryGroupMemoryRate unmarshals an instance of AutoscalingMemoryGroupMemoryRate from the specified map of raw messages.
func UnmarshalAutoscalingMemoryGroupMemoryRate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingMemoryGroupMemoryRate)
	err = core.UnmarshalPrimitive(m, "increase_percent", &obj.IncreasePercent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "period_seconds", &obj.PeriodSeconds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit_mb_per_member", &obj.LimitMbPerMember)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingMemoryGroupMemoryScalers : AutoscalingMemoryGroupMemoryScalers struct
type AutoscalingMemoryGroupMemoryScalers struct {
	IoUtilization *AutoscalingMemoryGroupMemoryScalersIoUtilization `json:"io_utilization,omitempty"`
}

// UnmarshalAutoscalingMemoryGroupMemoryScalers unmarshals an instance of AutoscalingMemoryGroupMemoryScalers from the specified map of raw messages.
func UnmarshalAutoscalingMemoryGroupMemoryScalers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingMemoryGroupMemoryScalers)
	err = core.UnmarshalModel(m, "io_utilization", &obj.IoUtilization, UnmarshalAutoscalingMemoryGroupMemoryScalersIoUtilization)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingMemoryGroupMemoryScalersIoUtilization : AutoscalingMemoryGroupMemoryScalersIoUtilization struct
type AutoscalingMemoryGroupMemoryScalersIoUtilization struct {
	Enabled *bool `json:"enabled,omitempty"`

	OverPeriod *string `json:"over_period,omitempty"`

	AbovePercent *int64 `json:"above_percent,omitempty"`
}

// UnmarshalAutoscalingMemoryGroupMemoryScalersIoUtilization unmarshals an instance of AutoscalingMemoryGroupMemoryScalersIoUtilization from the specified map of raw messages.
func UnmarshalAutoscalingMemoryGroupMemoryScalersIoUtilization(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingMemoryGroupMemoryScalersIoUtilization)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "over_period", &obj.OverPeriod)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "above_percent", &obj.AbovePercent)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingSetGroupAutoscaling : AutoscalingSetGroupAutoscaling struct
// Models which "extend" this model:
// - AutoscalingSetGroupAutoscalingAutoscalingDiskGroup
// - AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup
// - AutoscalingSetGroupAutoscalingAutoscalingCPUGroup
type AutoscalingSetGroupAutoscaling struct {
	Disk *AutoscalingDiskGroupDisk `json:"disk,omitempty"`

	Memory *AutoscalingMemoryGroupMemory `json:"memory,omitempty"`

	CPU *AutoscalingCPUGroupCPU `json:"cpu,omitempty"`
}
func (*AutoscalingSetGroupAutoscaling) isaAutoscalingSetGroupAutoscaling() bool {
	return true
}

type AutoscalingSetGroupAutoscalingIntf interface {
	isaAutoscalingSetGroupAutoscaling() bool
}

// UnmarshalAutoscalingSetGroupAutoscaling unmarshals an instance of AutoscalingSetGroupAutoscaling from the specified map of raw messages.
func UnmarshalAutoscalingSetGroupAutoscaling(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingSetGroupAutoscaling)
	err = core.UnmarshalModel(m, "disk", &obj.Disk, UnmarshalAutoscalingDiskGroupDisk)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "memory", &obj.Memory, UnmarshalAutoscalingMemoryGroupMemory)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cpu", &obj.CPU, UnmarshalAutoscalingCPUGroupCPU)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Backup : Backup struct
type Backup struct {
	// ID of this backup.
	ID *string `json:"id,omitempty"`

	// ID of the deployment this backup relates to.
	DeploymentID *string `json:"deployment_id,omitempty"`

	// The type of backup.
	Type *string `json:"type,omitempty"`

	// The status of this backup.
	Status *string `json:"status,omitempty"`

	// Is this backup available to download?.
	IsDownloadable *bool `json:"is_downloadable,omitempty"`

	// Can this backup be used to restore an instance?.
	IsRestorable *bool `json:"is_restorable,omitempty"`

	// URI which is currently available for file downloading.
	DownloadLink *string `json:"download_link,omitempty"`

	// Date and time when this backup was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`
}

// Constants associated with the Backup.Type property.
// The type of backup.
const (
	BackupTypeOnDemandConst = "on_demand"
	BackupTypeScheduledConst = "scheduled"
)

// Constants associated with the Backup.Status property.
// The status of this backup.
const (
	BackupStatusCompletedConst = "completed"
	BackupStatusFailedConst = "failed"
	BackupStatusRunningConst = "running"
)

// UnmarshalBackup unmarshals an instance of Backup from the specified map of raw messages.
func UnmarshalBackup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Backup)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployment_id", &obj.DeploymentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_downloadable", &obj.IsDownloadable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_restorable", &obj.IsRestorable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "download_link", &obj.DownloadLink)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Backups : Backups struct
type Backups struct {
	// An array of backups.
	Backups []Backup `json:"backups,omitempty"`
}

// UnmarshalBackups unmarshals an instance of Backups from the specified map of raw messages.
func UnmarshalBackups(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Backups)
	err = core.UnmarshalModel(m, "backups", &obj.Backups, UnmarshalBackup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Capability : Capability struct
type Capability struct {
	// Discover autoscaling capability.
	Autoscaling *string `json:"autoscaling,omitempty"`

	// Discover encryption capability.
	Encryption *bool `json:"encryption,omitempty"`

	// Discover endpoints capability.
	Endpoints *bool `json:"endpoints,omitempty"`

	// Discover currently available scaling groups. This command allows for the retrieval of all groups for a particular
	// deployment type. It also returns request parameters or defaults if request parameters are blank.
	Groups *string `json:"groups,omitempty"`

	// Discover all the regions into which deployments can be provisioned. Also returns request parameters or defaults if
	// request parameters are blank.
	Regions *string `json:"regions,omitempty"`

	// Discover read-only replica capability and supported regions. Also returns request parameters or defaults if request
	// parameters are blank.
	Remotes *string `json:"remotes,omitempty"`

	// Returns version capability, including upgrade paths. Also returns request parameters or defaults if request
	// parameters are blank.
	Versions *string `json:"versions,omitempty"`
}

// Constants associated with the Capability.Groups property.
// Discover currently available scaling groups. This command allows for the retrieval of all groups for a particular
// deployment type. It also returns request parameters or defaults if request parameters are blank.
const (
	CapabilityGroupsDatastaxConst = "datastax"
	CapabilityGroupsElasticsearchConst = "elasticsearch"
	CapabilityGroupsEnterprisedbConst = "enterprisedb"
	CapabilityGroupsEtcdConst = "etcd"
	CapabilityGroupsMongodbConst = "mongodb"
	CapabilityGroupsMysqlConst = "mysql"
	CapabilityGroupsPostgresqlConst = "postgresql"
	CapabilityGroupsRabbitmqConst = "rabbitmq"
	CapabilityGroupsRedisConst = "redis"
)

// Constants associated with the Capability.Remotes property.
// Discover read-only replica capability and supported regions. Also returns request parameters or defaults if request
// parameters are blank.
const (
	CapabilityRemotesDatastaxConst = "datastax"
	CapabilityRemotesElasticsearchConst = "elasticsearch"
	CapabilityRemotesEnterprisedbConst = "enterprisedb"
	CapabilityRemotesEtcdConst = "etcd"
	CapabilityRemotesMongodbConst = "mongodb"
	CapabilityRemotesMysqlConst = "mysql"
	CapabilityRemotesPostgresqlConst = "postgresql"
	CapabilityRemotesRabbitmqConst = "rabbitmq"
	CapabilityRemotesRedisConst = "redis"
)

// UnmarshalCapability unmarshals an instance of Capability from the specified map of raw messages.
func UnmarshalCapability(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Capability)
	err = core.UnmarshalPrimitive(m, "autoscaling", &obj.Autoscaling)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "encryption", &obj.Encryption)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoints", &obj.Endpoints)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "regions", &obj.Regions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "remotes", &obj.Remotes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "versions", &obj.Versions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChangeUserPasswordOptions : The ChangeUserPassword options.
type ChangeUserPasswordOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// User type.
	UserType *string `json:"user_type" validate:"required,ne="`

	// User ID.
	Username *string `json:"username" validate:"required,ne="`

	User *APasswordSettingUser `json:"user,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewChangeUserPasswordOptions : Instantiate ChangeUserPasswordOptions
func (*CloudDatabasesV5) NewChangeUserPasswordOptions(id string, userType string, username string) *ChangeUserPasswordOptions {
	return &ChangeUserPasswordOptions{
		ID: core.StringPtr(id),
		UserType: core.StringPtr(userType),
		Username: core.StringPtr(username),
	}
}

// SetID : Allow user to set ID
func (_options *ChangeUserPasswordOptions) SetID(id string) *ChangeUserPasswordOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetUserType : Allow user to set UserType
func (_options *ChangeUserPasswordOptions) SetUserType(userType string) *ChangeUserPasswordOptions {
	_options.UserType = core.StringPtr(userType)
	return _options
}

// SetUsername : Allow user to set Username
func (_options *ChangeUserPasswordOptions) SetUsername(username string) *ChangeUserPasswordOptions {
	_options.Username = core.StringPtr(username)
	return _options
}

// SetUser : Allow user to set User
func (_options *ChangeUserPasswordOptions) SetUser(user *APasswordSettingUser) *ChangeUserPasswordOptions {
	_options.User = user
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ChangeUserPasswordOptions) SetHeaders(param map[string]string) *ChangeUserPasswordOptions {
	options.Headers = param
	return options
}

// ChangeUserPasswordResponse : ChangeUserPasswordResponse struct
type ChangeUserPasswordResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalChangeUserPasswordResponse unmarshals an instance of ChangeUserPasswordResponse from the specified map of raw messages.
func UnmarshalChangeUserPasswordResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChangeUserPasswordResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CompleteConnectionOptions : The CompleteConnection options.
type CompleteConnectionOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// User type of `database` is the only currently supported value.
	UserType *string `json:"user_type" validate:"required,ne="`

	// User ID.
	UserID *string `json:"user_id" validate:"required,ne="`

	// Endpoint Type. The select endpoint must be enabled on the deployment before its connection information can be
	// fetched.
	EndpointType *string `json:"endpoint_type" validate:"required,ne="`

	// Password to be substituted into the response.
	Password *string `json:"password,omitempty"`

	// Optional certificate root path to prepend certificate names. Certificates would be stored in this directory for use
	// by other commands.
	CertificateRoot *string `json:"certificate_root,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CompleteConnectionOptions.EndpointType property.
// Endpoint Type. The select endpoint must be enabled on the deployment before its connection information can be
// fetched.
const (
	CompleteConnectionOptionsEndpointTypePrivateConst = "private"
	CompleteConnectionOptionsEndpointTypePublicConst = "public"
)

// NewCompleteConnectionOptions : Instantiate CompleteConnectionOptions
func (*CloudDatabasesV5) NewCompleteConnectionOptions(id string, userType string, userID string, endpointType string) *CompleteConnectionOptions {
	return &CompleteConnectionOptions{
		ID: core.StringPtr(id),
		UserType: core.StringPtr(userType),
		UserID: core.StringPtr(userID),
		EndpointType: core.StringPtr(endpointType),
	}
}

// SetID : Allow user to set ID
func (_options *CompleteConnectionOptions) SetID(id string) *CompleteConnectionOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetUserType : Allow user to set UserType
func (_options *CompleteConnectionOptions) SetUserType(userType string) *CompleteConnectionOptions {
	_options.UserType = core.StringPtr(userType)
	return _options
}

// SetUserID : Allow user to set UserID
func (_options *CompleteConnectionOptions) SetUserID(userID string) *CompleteConnectionOptions {
	_options.UserID = core.StringPtr(userID)
	return _options
}

// SetEndpointType : Allow user to set EndpointType
func (_options *CompleteConnectionOptions) SetEndpointType(endpointType string) *CompleteConnectionOptions {
	_options.EndpointType = core.StringPtr(endpointType)
	return _options
}

// SetPassword : Allow user to set Password
func (_options *CompleteConnectionOptions) SetPassword(password string) *CompleteConnectionOptions {
	_options.Password = core.StringPtr(password)
	return _options
}

// SetCertificateRoot : Allow user to set CertificateRoot
func (_options *CompleteConnectionOptions) SetCertificateRoot(certificateRoot string) *CompleteConnectionOptions {
	_options.CertificateRoot = core.StringPtr(certificateRoot)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CompleteConnectionOptions) SetHeaders(param map[string]string) *CompleteConnectionOptions {
	options.Headers = param
	return options
}

// CompleteConnectionResponse : CompleteConnectionResponse struct
type CompleteConnectionResponse struct {
	Connection ConnectionIntf `json:"connection,omitempty"`
}

// UnmarshalCompleteConnectionResponse unmarshals an instance of CompleteConnectionResponse from the specified map of raw messages.
func UnmarshalCompleteConnectionResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CompleteConnectionResponse)
	err = core.UnmarshalModel(m, "connection", &obj.Connection, UnmarshalConnection)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Configuration : Configuration struct
// Models which "extend" this model:
// - ConfigurationPgConfiguration
// - ConfigurationRedisConfiguration
// - ConfigurationMySQLConfiguration
type Configuration struct {
	// Maximum connections allowed.
	MaxConnections *int64 `json:"max_connections,omitempty"`

	// Max number of transactions that can be in the "prepared" state simultaneously.
	MaxPreparedTransactions *int64 `json:"max_prepared_transactions,omitempty"`

	// Deadlock timeout in ms. The time to wait on a lock before checking for deadlock.  Also the duration where lock waits
	// will be logged.
	DeadlockTimeout *int64 `json:"deadlock_timeout,omitempty"`

	// Number of simultaneous requests that can be handled efficiently by the disk subsystem.
	EffectiveIoConcurrency *int64 `json:"effective_io_concurrency,omitempty"`

	// Maximum number of simultaneously defined replication slots.
	MaxReplicationSlots *int64 `json:"max_replication_slots,omitempty"`

	// Maximum number of simultaneously running WAL sender processes.
	MaxWalSenders *int64 `json:"max_wal_senders,omitempty"`

	// The number of 8kB shared memory buffers used by the server.  Set to 1/4 of memory.  Setting too high will cause
	// crashes or prevent the database from starting.
	SharedBuffers *int64 `json:"shared_buffers,omitempty"`

	// Sets the current transaction's synchronization level.  Off can result in data loss.  remote_write with enable
	// synchronous replication which will impact performance and availabilty.
	SynchronousCommit *string `json:"synchronous_commit,omitempty"`

	// WAL level.  Set to logical to use logical decoding or logical replication.
	WalLevel *string `json:"wal_level,omitempty"`

	// The number of seconds to wait before forces a switch to the next WAL file if a new file has not been started.
	ArchiveTimeout *int64 `json:"archive_timeout,omitempty"`

	// The minimum number of milliseconds for execution time above which statements will be logged.
	LogMinDurationStatement *int64 `json:"log_min_duration_statement,omitempty"`

	// The maximum memory Redis should use, as bytes.
	MaxmemoryRedis *int64 `json:"maxmemory-redis,omitempty"`

	// The policy with which Redis evicts keys when maximum memory is reached.
	MaxmemoryPolicy *string `json:"maxmemory-policy,omitempty"`

	// If set to yes this will enable AOF persistence.
	Appendonly *string `json:"appendonly,omitempty"`

	// The maximum memory Redis should use, as bytes.
	MaxmemorySamples *int64 `json:"maxmemory-samples,omitempty"`

	// Whether or not to stop accepting writes when background persistence actions fail.
	StopWritesOnBgsaveError *string `json:"stop-writes-on-bgsave-error,omitempty"`

	// Maximum age for a binlog in seconds. If a binlog is older, it's archived.
	MysqlMaxBinlogAgeSec *int64 `json:"mysql_max_binlog_age_sec,omitempty"`

	// Determines which authentication plugin the server assigns to new accounts created by CREATE USER and GRANT
	// statements that do not explicitly specify an authentication plugin.
	MysqlDefaultAuthenticationPlugin *string `json:"mysql_default_authentication_plugin,omitempty"`
}

// Constants associated with the Configuration.SynchronousCommit property.
// Sets the current transaction's synchronization level.  Off can result in data loss.  remote_write with enable
// synchronous replication which will impact performance and availabilty.
const (
	ConfigurationSynchronousCommitLocalConst = "local"
	ConfigurationSynchronousCommitOffConst = "off"
)

// Constants associated with the Configuration.WalLevel property.
// WAL level.  Set to logical to use logical decoding or logical replication.
const (
	ConfigurationWalLevelHotStandbyConst = "hot_standby"
	ConfigurationWalLevelLogicalConst = "logical"
)

// Constants associated with the Configuration.MaxmemoryPolicy property.
// The policy with which Redis evicts keys when maximum memory is reached.
const (
	ConfigurationMaxmemoryPolicyAllkeysLruConst = "allkeys-lru"
	ConfigurationMaxmemoryPolicyAllkeysRandomConst = "allkeys-random"
	ConfigurationMaxmemoryPolicyNoevictionConst = "noeviction"
	ConfigurationMaxmemoryPolicyVolatileLruConst = "volatile-lru"
	ConfigurationMaxmemoryPolicyVolatileRandomConst = "volatile-random"
	ConfigurationMaxmemoryPolicyVolatileTTLConst = "volatile-ttl"
)

// Constants associated with the Configuration.Appendonly property.
// If set to yes this will enable AOF persistence.
const (
	ConfigurationAppendonlyNoConst = "no"
	ConfigurationAppendonlyYesConst = "yes"
)

// Constants associated with the Configuration.StopWritesOnBgsaveError property.
// Whether or not to stop accepting writes when background persistence actions fail.
const (
	ConfigurationStopWritesOnBgsaveErrorNoConst = "no"
	ConfigurationStopWritesOnBgsaveErrorYesConst = "yes"
)

// Constants associated with the Configuration.MysqlDefaultAuthenticationPlugin property.
// Determines which authentication plugin the server assigns to new accounts created by CREATE USER and GRANT statements
// that do not explicitly specify an authentication plugin.
const (
	ConfigurationMysqlDefaultAuthenticationPluginMysqlNativePasswordConst = "mysql_native_password"
	ConfigurationMysqlDefaultAuthenticationPluginSha256PasswordConst = "sha256_password"
)
func (*Configuration) isaConfiguration() bool {
	return true
}

type ConfigurationIntf interface {
	isaConfiguration() bool
}

// UnmarshalConfiguration unmarshals an instance of Configuration from the specified map of raw messages.
func UnmarshalConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Configuration)
	err = core.UnmarshalPrimitive(m, "max_connections", &obj.MaxConnections)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_prepared_transactions", &obj.MaxPreparedTransactions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deadlock_timeout", &obj.DeadlockTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "effective_io_concurrency", &obj.EffectiveIoConcurrency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_replication_slots", &obj.MaxReplicationSlots)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_wal_senders", &obj.MaxWalSenders)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "shared_buffers", &obj.SharedBuffers)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "synchronous_commit", &obj.SynchronousCommit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "wal_level", &obj.WalLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "archive_timeout", &obj.ArchiveTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "log_min_duration_statement", &obj.LogMinDurationStatement)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxmemory-redis", &obj.MaxmemoryRedis)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxmemory-policy", &obj.MaxmemoryPolicy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "appendonly", &obj.Appendonly)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxmemory-samples", &obj.MaxmemorySamples)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "stop-writes-on-bgsave-error", &obj.StopWritesOnBgsaveError)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "mysql_max_binlog_age_sec", &obj.MysqlMaxBinlogAgeSec)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "mysql_default_authentication_plugin", &obj.MysqlDefaultAuthenticationPlugin)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Connection : Connection struct
// Models which "extend" this model:
// - ConnectionPostgreSQLConnection
// - ConnectionRedisConnection
// - ConnectionElasticsearchConnection
// - ConnectionRabbitMqConnection
// - ConnectionEtcdConnection
// - ConnectionMongoDbConnection
// - ConnectionMongoDbeeConnection
// - ConnectionMongoDbeeOpsManagerConnection
// - ConnectionMySQLConnection
// - ConnectionDataStaxConnection
// - ConnectionEnterpriseDbConnection
type Connection struct {
	Postgres *PostgreSQLConnectionURI `json:"postgres,omitempty"`

	// CLI Connection.
	Cli *ConnectionCli `json:"cli,omitempty"`

	Rediss *RedisConnectionURI `json:"rediss,omitempty"`

	HTTPS *ConnectionURI `json:"https,omitempty"`

	Amqps *ConnectionURI `json:"amqps,omitempty"`

	Mqtts *ConnectionURI `json:"mqtts,omitempty"`

	StompSsl *ConnectionURI `json:"stomp_ssl,omitempty"`

	Grpc *ConnectionURI `json:"grpc,omitempty"`

	Mongodb *MongoDbConnectionURI `json:"mongodb,omitempty"`

	BiConnector *ConnectionURI `json:"bi_connector,omitempty"`

	Analytics *ConnectionURI `json:"analytics,omitempty"`

	OpsManager *ConnectionURI `json:"ops_manager,omitempty"`

	Mysql *MySQLConnectionURI `json:"mysql,omitempty"`

	Secure *DataStaxConnectionURI `json:"secure,omitempty"`

	Emp *ConnectionURI `json:"emp,omitempty"`
}
func (*Connection) isaConnection() bool {
	return true
}

type ConnectionIntf interface {
	isaConnection() bool
}

// UnmarshalConnection unmarshals an instance of Connection from the specified map of raw messages.
func UnmarshalConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Connection)
	err = core.UnmarshalModel(m, "postgres", &obj.Postgres, UnmarshalPostgreSQLConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCli)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rediss", &obj.Rediss, UnmarshalRedisConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "https", &obj.HTTPS, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "amqps", &obj.Amqps, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "mqtts", &obj.Mqtts, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "stomp_ssl", &obj.StompSsl, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "grpc", &obj.Grpc, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "mongodb", &obj.Mongodb, UnmarshalMongoDbConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "bi_connector", &obj.BiConnector, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "analytics", &obj.Analytics, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "ops_manager", &obj.OpsManager, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "mysql", &obj.Mysql, UnmarshalMySQLConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "secure", &obj.Secure, UnmarshalDataStaxConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "emp", &obj.Emp, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionAuthentication : Authentication data for Connection String.
type ConnectionAuthentication struct {
	// Authentication method for this credential.
	Method *string `json:"method,omitempty"`

	// Username part of credential.
	Username *string `json:"username,omitempty"`

	// Password part of credential.
	Password *string `json:"password,omitempty"`
}

// UnmarshalConnectionAuthentication unmarshals an instance of ConnectionAuthentication from the specified map of raw messages.
func UnmarshalConnectionAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionBundle : ConnectionBundle struct
type ConnectionBundle struct {
	// Name associated with the certificate.
	Name *string `json:"name,omitempty"`

	// Base64 encoded version of the certificate bundle.
	BundleBase64 *string `json:"bundle_base64,omitempty"`
}

// UnmarshalConnectionBundle unmarshals an instance of ConnectionBundle from the specified map of raw messages.
func UnmarshalConnectionBundle(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionBundle)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bundle_base64", &obj.BundleBase64)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionCli : CLI Connection.
type ConnectionCli struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// A map of environment variables for a CLI connection.
	Environment map[string]interface{} `json:"environment,omitempty"`

	// The name of the executable the CLI should run.
	Bin *string `json:"bin,omitempty"`

	// Sets of arguments to call the executable with. The outer array corresponds to a possible way to call the CLI; the
	// inner array is the set of arguments to use with that call.
	Arguments [][]string `json:"arguments,omitempty"`

	Certificate *ConnectionCertificate `json:"certificate,omitempty"`
}

// UnmarshalConnectionCli unmarshals an instance of ConnectionCli from the specified map of raw messages.
func UnmarshalConnectionCli(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionCli)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment", &obj.Environment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bin", &obj.Bin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "arguments", &obj.Arguments)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalConnectionCertificate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionCertificate : ConnectionCertificate struct
type ConnectionCertificate struct {
	// Name associated with the certificate.
	Name *string `json:"name,omitempty"`

	// Base64 encoded version of the certificate.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`
}

// UnmarshalConnectionCertificate unmarshals an instance of ConnectionCertificate from the specified map of raw messages.
func UnmarshalConnectionCertificate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionCertificate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionHost : Connection hostname and port.
type ConnectionHost struct {
	// Hostname for connection.
	Hostname *string `json:"hostname,omitempty"`

	// Port number for connection.
	Port *int64 `json:"port,omitempty"`
}

// UnmarshalConnectionHost unmarshals an instance of ConnectionHost from the specified map of raw messages.
func UnmarshalConnectionHost(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionHost)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionURI : ConnectionURI struct
type ConnectionURI struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// Scheme/protocol for URI connection.
	Scheme *string `json:"scheme,omitempty"`

	Hosts []ConnectionHost `json:"hosts,omitempty"`

	// Path for URI connection.
	Path *string `json:"path,omitempty"`

	// Query options to add to the URI connection.
	QueryOptions map[string]interface{} `json:"query_options,omitempty"`

	// Authentication data for Connection String.
	Authentication *ConnectionAuthentication `json:"authentication,omitempty"`

	Certificate *ConnectionCertificate `json:"certificate,omitempty"`

	// Indicates ssl is required for the connection.
	Ssl *bool `json:"ssl,omitempty"`

	// Indicates the address is accessible by browser.
	BrowserAccessible *bool `json:"browser_accessible,omitempty"`
}

// UnmarshalConnectionURI unmarshals an instance of ConnectionURI from the specified map of raw messages.
func UnmarshalConnectionURI(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionURI)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalConnectionHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_options", &obj.QueryOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalConnectionAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalConnectionCertificate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "browser_accessible", &obj.BrowserAccessible)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateDatabaseUserOptions : The CreateDatabaseUser options.
type CreateDatabaseUserOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// User type.
	UserType *string `json:"user_type" validate:"required,ne="`

	User *User `json:"user,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateDatabaseUserOptions : Instantiate CreateDatabaseUserOptions
func (*CloudDatabasesV5) NewCreateDatabaseUserOptions(id string, userType string) *CreateDatabaseUserOptions {
	return &CreateDatabaseUserOptions{
		ID: core.StringPtr(id),
		UserType: core.StringPtr(userType),
	}
}

// SetID : Allow user to set ID
func (_options *CreateDatabaseUserOptions) SetID(id string) *CreateDatabaseUserOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetUserType : Allow user to set UserType
func (_options *CreateDatabaseUserOptions) SetUserType(userType string) *CreateDatabaseUserOptions {
	_options.UserType = core.StringPtr(userType)
	return _options
}

// SetUser : Allow user to set User
func (_options *CreateDatabaseUserOptions) SetUser(user *User) *CreateDatabaseUserOptions {
	_options.User = user
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDatabaseUserOptions) SetHeaders(param map[string]string) *CreateDatabaseUserOptions {
	options.Headers = param
	return options
}

// CreateDatabaseUserResponse : CreateDatabaseUserResponse struct
type CreateDatabaseUserResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalCreateDatabaseUserResponse unmarshals an instance of CreateDatabaseUserResponse from the specified map of raw messages.
func UnmarshalCreateDatabaseUserResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateDatabaseUserResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataStaxConnectionURI : DataStaxConnectionURI struct
type DataStaxConnectionURI struct {
	Hosts []ConnectionHost `json:"hosts,omitempty"`

	// Authentication data for Connection String.
	Authentication *ConnectionAuthentication `json:"authentication,omitempty"`

	Bundle *ConnectionBundle `json:"bundle,omitempty"`
}

// UnmarshalDataStaxConnectionURI unmarshals an instance of DataStaxConnectionURI from the specified map of raw messages.
func UnmarshalDataStaxConnectionURI(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataStaxConnectionURI)
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalConnectionHost)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalConnectionAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "bundle", &obj.Bundle, UnmarshalConnectionBundle)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteAllowlistEntryOptions : The DeleteAllowlistEntry options.
type DeleteAllowlistEntryOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// An IPv4 address or a CIDR range (netmasked IPv4 address).
	Ipaddress *string `json:"ipaddress" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAllowlistEntryOptions : Instantiate DeleteAllowlistEntryOptions
func (*CloudDatabasesV5) NewDeleteAllowlistEntryOptions(id string, ipaddress string) *DeleteAllowlistEntryOptions {
	return &DeleteAllowlistEntryOptions{
		ID: core.StringPtr(id),
		Ipaddress: core.StringPtr(ipaddress),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteAllowlistEntryOptions) SetID(id string) *DeleteAllowlistEntryOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIpaddress : Allow user to set Ipaddress
func (_options *DeleteAllowlistEntryOptions) SetIpaddress(ipaddress string) *DeleteAllowlistEntryOptions {
	_options.Ipaddress = core.StringPtr(ipaddress)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAllowlistEntryOptions) SetHeaders(param map[string]string) *DeleteAllowlistEntryOptions {
	options.Headers = param
	return options
}

// DeleteAllowlistEntryResponse : DeleteAllowlistEntryResponse struct
type DeleteAllowlistEntryResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalDeleteAllowlistEntryResponse unmarshals an instance of DeleteAllowlistEntryResponse from the specified map of raw messages.
func UnmarshalDeleteAllowlistEntryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteAllowlistEntryResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteDatabaseUserOptions : The DeleteDatabaseUser options.
type DeleteDatabaseUserOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// User type.
	UserType *string `json:"user_type" validate:"required,ne="`

	// Username.
	Username *string `json:"username" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDatabaseUserOptions : Instantiate DeleteDatabaseUserOptions
func (*CloudDatabasesV5) NewDeleteDatabaseUserOptions(id string, userType string, username string) *DeleteDatabaseUserOptions {
	return &DeleteDatabaseUserOptions{
		ID: core.StringPtr(id),
		UserType: core.StringPtr(userType),
		Username: core.StringPtr(username),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteDatabaseUserOptions) SetID(id string) *DeleteDatabaseUserOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetUserType : Allow user to set UserType
func (_options *DeleteDatabaseUserOptions) SetUserType(userType string) *DeleteDatabaseUserOptions {
	_options.UserType = core.StringPtr(userType)
	return _options
}

// SetUsername : Allow user to set Username
func (_options *DeleteDatabaseUserOptions) SetUsername(username string) *DeleteDatabaseUserOptions {
	_options.Username = core.StringPtr(username)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDatabaseUserOptions) SetHeaders(param map[string]string) *DeleteDatabaseUserOptions {
	options.Headers = param
	return options
}

// DeleteDatabaseUserResponse : DeleteDatabaseUserResponse struct
type DeleteDatabaseUserResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalDeleteDatabaseUserResponse unmarshals an instance of DeleteDatabaseUserResponse from the specified map of raw messages.
func UnmarshalDeleteDatabaseUserResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteDatabaseUserResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Deployables : Deployable databases with their version information.
type Deployables struct {
	// Deployment type - typically the name of the database.
	Type *string `json:"type,omitempty"`

	// An array of versions of the database, their status, preferedness, and transitions.
	Versions []DeployablesVersionsItem `json:"versions,omitempty"`
}

// UnmarshalDeployables unmarshals an instance of Deployables from the specified map of raw messages.
func UnmarshalDeployables(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Deployables)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "versions", &obj.Versions, UnmarshalDeployablesVersionsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeployablesVersionsItem : DeployablesVersionsItem struct
type DeployablesVersionsItem struct {
	// The version number.
	Version *string `json:"version,omitempty"`

	// The status of this version: To be finalized.
	Status *string `json:"status,omitempty"`

	// Should this version be preferred over others?.
	IsPreferred *bool `json:"is_preferred,omitempty"`

	// versions that this version can be upgraded to.
	Transitions []DeployablesVersionsItemTransitionsItem `json:"transitions,omitempty"`
}

// Constants associated with the DeployablesVersionsItem.Status property.
// The status of this version: To be finalized.
const (
	DeployablesVersionsItemStatusBetaConst = "beta"
	DeployablesVersionsItemStatusDeprecatedConst = "deprecated"
	DeployablesVersionsItemStatusStableConst = "stable"
)

// UnmarshalDeployablesVersionsItem unmarshals an instance of DeployablesVersionsItem from the specified map of raw messages.
func UnmarshalDeployablesVersionsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeployablesVersionsItem)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_preferred", &obj.IsPreferred)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "transitions", &obj.Transitions, UnmarshalDeployablesVersionsItemTransitionsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeployablesVersionsItemTransitionsItem : DeployablesVersionsItemTransitionsItem struct
type DeployablesVersionsItemTransitionsItem struct {
	// The database type.
	Application *string `json:"application,omitempty"`

	// method of going from from_version to to_version.
	Method *string `json:"method,omitempty"`

	// The version the transition in from.
	FromVersion *string `json:"from_version,omitempty"`

	// The version the transition is to.
	ToVersion *string `json:"to_version,omitempty"`
}

// UnmarshalDeployablesVersionsItemTransitionsItem unmarshals an instance of DeployablesVersionsItemTransitionsItem from the specified map of raw messages.
func UnmarshalDeployablesVersionsItemTransitionsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeployablesVersionsItemTransitionsItem)
	err = core.UnmarshalPrimitive(m, "application", &obj.Application)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "from_version", &obj.FromVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "to_version", &obj.ToVersion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Deployment : Deployment struct
type Deployment struct {
	// ID of this deployment.
	ID *string `json:"id,omitempty"`

	// Readable name of this deployment.
	Name *string `json:"name,omitempty"`

	// Database type within this deployment.
	Type *string `json:"type,omitempty"`

	// Platform for this deployment.
	Platform *string `json:"platform,omitempty"`

	// Platform-specific options for this deployment.
	PlatformOptions map[string]interface{} `json:"platform_options,omitempty"`

	// Version number of the database.
	Version *string `json:"version,omitempty"`

	// Login name of administration level user.
	AdminUsernames map[string]string `json:"admin_usernames,omitempty"`

	// Whether access to this deployment is enabled from the public internet. This property can be modified by updating
	// this service instance through the Resource Controller API.
	EnablePublicEndpoints *bool `json:"enable_public_endpoints,omitempty"`

	// Whether access to this deployment is enabled from IBM Cloud via the IBM Cloud backbone network. This property can be
	// modified by updating this service instance through the Resource Controller API.
	EnablePrivateEndpoints *bool `json:"enable_private_endpoints,omitempty"`
}

// UnmarshalDeployment unmarshals an instance of Deployment from the specified map of raw messages.
func UnmarshalDeployment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Deployment)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "platform", &obj.Platform)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "platform_options", &obj.PlatformOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "admin_usernames", &obj.AdminUsernames)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enable_public_endpoints", &obj.EnablePublicEndpoints)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enable_private_endpoints", &obj.EnablePrivateEndpoints)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAllowlistOptions : The GetAllowlist options.
type GetAllowlistOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAllowlistOptions : Instantiate GetAllowlistOptions
func (*CloudDatabasesV5) NewGetAllowlistOptions(id string) *GetAllowlistOptions {
	return &GetAllowlistOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetAllowlistOptions) SetID(id string) *GetAllowlistOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAllowlistOptions) SetHeaders(param map[string]string) *GetAllowlistOptions {
	options.Headers = param
	return options
}

// GetAllowlistResponse : GetAllowlistResponse struct
type GetAllowlistResponse struct {
	IPAddresses []AllowlistEntry `json:"ip_addresses,omitempty"`
}

// UnmarshalGetAllowlistResponse unmarshals an instance of GetAllowlistResponse from the specified map of raw messages.
func UnmarshalGetAllowlistResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetAllowlistResponse)
	err = core.UnmarshalModel(m, "ip_addresses", &obj.IPAddresses, UnmarshalAllowlistEntry)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAutoscalingCapabilityOptions : The GetAutoscalingCapability options.
type GetAutoscalingCapabilityOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	Deployment *GetAutoscalingCapabilityRequestDeployment `json:"deployment,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAutoscalingCapabilityOptions : Instantiate GetAutoscalingCapabilityOptions
func (*CloudDatabasesV5) NewGetAutoscalingCapabilityOptions(id string) *GetAutoscalingCapabilityOptions {
	return &GetAutoscalingCapabilityOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetAutoscalingCapabilityOptions) SetID(id string) *GetAutoscalingCapabilityOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetDeployment : Allow user to set Deployment
func (_options *GetAutoscalingCapabilityOptions) SetDeployment(deployment *GetAutoscalingCapabilityRequestDeployment) *GetAutoscalingCapabilityOptions {
	_options.Deployment = deployment
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAutoscalingCapabilityOptions) SetHeaders(param map[string]string) *GetAutoscalingCapabilityOptions {
	options.Headers = param
	return options
}

// GetAutoscalingCapabilityRequestDeployment : GetAutoscalingCapabilityRequestDeployment struct
type GetAutoscalingCapabilityRequestDeployment struct {
	// Type of database.
	DeploymentType *string `json:"deployment_type,omitempty"`

	// database version.
	Version *string `json:"version,omitempty"`

	// classic or satellite.
	Platform *string `json:"platform,omitempty"`
}

// UnmarshalGetAutoscalingCapabilityRequestDeployment unmarshals an instance of GetAutoscalingCapabilityRequestDeployment from the specified map of raw messages.
func UnmarshalGetAutoscalingCapabilityRequestDeployment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetAutoscalingCapabilityRequestDeployment)
	err = core.UnmarshalPrimitive(m, "deployment_type", &obj.DeploymentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "platform", &obj.Platform)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAutoscalingCapabilityResponse : GetAutoscalingCapabilityResponse struct
type GetAutoscalingCapabilityResponse struct {
	Capability []Capability `json:"capability,omitempty"`
}

// UnmarshalGetAutoscalingCapabilityResponse unmarshals an instance of GetAutoscalingCapabilityResponse from the specified map of raw messages.
func UnmarshalGetAutoscalingCapabilityResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetAutoscalingCapabilityResponse)
	err = core.UnmarshalModel(m, "capability", &obj.Capability, UnmarshalCapability)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAutoscalingConditionsOptions : The GetAutoscalingConditions options.
type GetAutoscalingConditionsOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Group ID.
	GroupID *string `json:"group_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAutoscalingConditionsOptions : Instantiate GetAutoscalingConditionsOptions
func (*CloudDatabasesV5) NewGetAutoscalingConditionsOptions(id string, groupID string) *GetAutoscalingConditionsOptions {
	return &GetAutoscalingConditionsOptions{
		ID: core.StringPtr(id),
		GroupID: core.StringPtr(groupID),
	}
}

// SetID : Allow user to set ID
func (_options *GetAutoscalingConditionsOptions) SetID(id string) *GetAutoscalingConditionsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetGroupID : Allow user to set GroupID
func (_options *GetAutoscalingConditionsOptions) SetGroupID(groupID string) *GetAutoscalingConditionsOptions {
	_options.GroupID = core.StringPtr(groupID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAutoscalingConditionsOptions) SetHeaders(param map[string]string) *GetAutoscalingConditionsOptions {
	options.Headers = param
	return options
}

// GetBackupInfoOptions : The GetBackupInfo options.
type GetBackupInfoOptions struct {
	// Backup ID.
	BackupID *string `json:"backup_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBackupInfoOptions : Instantiate GetBackupInfoOptions
func (*CloudDatabasesV5) NewGetBackupInfoOptions(backupID string) *GetBackupInfoOptions {
	return &GetBackupInfoOptions{
		BackupID: core.StringPtr(backupID),
	}
}

// SetBackupID : Allow user to set BackupID
func (_options *GetBackupInfoOptions) SetBackupID(backupID string) *GetBackupInfoOptions {
	_options.BackupID = core.StringPtr(backupID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBackupInfoOptions) SetHeaders(param map[string]string) *GetBackupInfoOptions {
	options.Headers = param
	return options
}

// GetBackupInfoResponse : GetBackupInfoResponse struct
type GetBackupInfoResponse struct {
	Backup *Backup `json:"backup,omitempty"`
}

// UnmarshalGetBackupInfoResponse unmarshals an instance of GetBackupInfoResponse from the specified map of raw messages.
func UnmarshalGetBackupInfoResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetBackupInfoResponse)
	err = core.UnmarshalModel(m, "backup", &obj.Backup, UnmarshalBackup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetConnectionOptions : The GetConnection options.
type GetConnectionOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// User type.
	UserType *string `json:"user_type" validate:"required,ne="`

	// User ID.
	UserID *string `json:"user_id" validate:"required,ne="`

	// Endpoint Type. The endpoint must be enabled on the deployment before its connection information can be fetched.
	EndpointType *string `json:"endpoint_type" validate:"required,ne="`

	// Optional certificate root path to prepend certificate names. Certificates would be stored in this directory for use
	// by other commands.
	CertificateRoot *string `json:"certificate_root,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetConnectionOptions.EndpointType property.
// Endpoint Type. The endpoint must be enabled on the deployment before its connection information can be fetched.
const (
	GetConnectionOptionsEndpointTypePrivateConst = "private"
	GetConnectionOptionsEndpointTypePublicConst = "public"
)

// NewGetConnectionOptions : Instantiate GetConnectionOptions
func (*CloudDatabasesV5) NewGetConnectionOptions(id string, userType string, userID string, endpointType string) *GetConnectionOptions {
	return &GetConnectionOptions{
		ID: core.StringPtr(id),
		UserType: core.StringPtr(userType),
		UserID: core.StringPtr(userID),
		EndpointType: core.StringPtr(endpointType),
	}
}

// SetID : Allow user to set ID
func (_options *GetConnectionOptions) SetID(id string) *GetConnectionOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetUserType : Allow user to set UserType
func (_options *GetConnectionOptions) SetUserType(userType string) *GetConnectionOptions {
	_options.UserType = core.StringPtr(userType)
	return _options
}

// SetUserID : Allow user to set UserID
func (_options *GetConnectionOptions) SetUserID(userID string) *GetConnectionOptions {
	_options.UserID = core.StringPtr(userID)
	return _options
}

// SetEndpointType : Allow user to set EndpointType
func (_options *GetConnectionOptions) SetEndpointType(endpointType string) *GetConnectionOptions {
	_options.EndpointType = core.StringPtr(endpointType)
	return _options
}

// SetCertificateRoot : Allow user to set CertificateRoot
func (_options *GetConnectionOptions) SetCertificateRoot(certificateRoot string) *GetConnectionOptions {
	_options.CertificateRoot = core.StringPtr(certificateRoot)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetConnectionOptions) SetHeaders(param map[string]string) *GetConnectionOptions {
	options.Headers = param
	return options
}

// GetConnectionResponse : GetConnectionResponse struct
type GetConnectionResponse struct {
	Connection ConnectionIntf `json:"connection,omitempty"`
}

// UnmarshalGetConnectionResponse unmarshals an instance of GetConnectionResponse from the specified map of raw messages.
func UnmarshalGetConnectionResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetConnectionResponse)
	err = core.UnmarshalModel(m, "connection", &obj.Connection, UnmarshalConnection)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetDefaultScalingGroupsOptions : The GetDefaultScalingGroups options.
type GetDefaultScalingGroupsOptions struct {
	// Database type name.
	Type *string `json:"type" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetDefaultScalingGroupsOptions.Type property.
// Database type name.
const (
	GetDefaultScalingGroupsOptionsTypeEtcdConst = "etcd"
	GetDefaultScalingGroupsOptionsTypePostgresqlConst = "postgresql"
)

// NewGetDefaultScalingGroupsOptions : Instantiate GetDefaultScalingGroupsOptions
func (*CloudDatabasesV5) NewGetDefaultScalingGroupsOptions(typeVar string) *GetDefaultScalingGroupsOptions {
	return &GetDefaultScalingGroupsOptions{
		Type: core.StringPtr(typeVar),
	}
}

// SetType : Allow user to set Type
func (_options *GetDefaultScalingGroupsOptions) SetType(typeVar string) *GetDefaultScalingGroupsOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDefaultScalingGroupsOptions) SetHeaders(param map[string]string) *GetDefaultScalingGroupsOptions {
	options.Headers = param
	return options
}

// GetDefaultScalingGroupsResponse : GetDefaultScalingGroupsResponse struct
type GetDefaultScalingGroupsResponse struct {
	Groups []Group `json:"groups,omitempty"`
}

// UnmarshalGetDefaultScalingGroupsResponse unmarshals an instance of GetDefaultScalingGroupsResponse from the specified map of raw messages.
func UnmarshalGetDefaultScalingGroupsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetDefaultScalingGroupsResponse)
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalGroup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetDeploymentInfoOptions : The GetDeploymentInfo options.
type GetDeploymentInfoOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDeploymentInfoOptions : Instantiate GetDeploymentInfoOptions
func (*CloudDatabasesV5) NewGetDeploymentInfoOptions(id string) *GetDeploymentInfoOptions {
	return &GetDeploymentInfoOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetDeploymentInfoOptions) SetID(id string) *GetDeploymentInfoOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDeploymentInfoOptions) SetHeaders(param map[string]string) *GetDeploymentInfoOptions {
	options.Headers = param
	return options
}

// GetDeploymentInfoResponse : GetDeploymentInfoResponse struct
type GetDeploymentInfoResponse struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

// UnmarshalGetDeploymentInfoResponse unmarshals an instance of GetDeploymentInfoResponse from the specified map of raw messages.
func UnmarshalGetDeploymentInfoResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetDeploymentInfoResponse)
	err = core.UnmarshalModel(m, "deployment", &obj.Deployment, UnmarshalDeployment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetEncryptionCapabilityOptions : The GetEncryptionCapability options.
type GetEncryptionCapabilityOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	Deployment *GetEncryptionCapabilityRequestDeployment `json:"deployment,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetEncryptionCapabilityOptions : Instantiate GetEncryptionCapabilityOptions
func (*CloudDatabasesV5) NewGetEncryptionCapabilityOptions(id string) *GetEncryptionCapabilityOptions {
	return &GetEncryptionCapabilityOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetEncryptionCapabilityOptions) SetID(id string) *GetEncryptionCapabilityOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetDeployment : Allow user to set Deployment
func (_options *GetEncryptionCapabilityOptions) SetDeployment(deployment *GetEncryptionCapabilityRequestDeployment) *GetEncryptionCapabilityOptions {
	_options.Deployment = deployment
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetEncryptionCapabilityOptions) SetHeaders(param map[string]string) *GetEncryptionCapabilityOptions {
	options.Headers = param
	return options
}

// GetEncryptionCapabilityRequestDeployment : GetEncryptionCapabilityRequestDeployment struct
type GetEncryptionCapabilityRequestDeployment struct {
	// Type of database.
	DeploymentType *string `json:"deployment_type,omitempty"`
}

// UnmarshalGetEncryptionCapabilityRequestDeployment unmarshals an instance of GetEncryptionCapabilityRequestDeployment from the specified map of raw messages.
func UnmarshalGetEncryptionCapabilityRequestDeployment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetEncryptionCapabilityRequestDeployment)
	err = core.UnmarshalPrimitive(m, "deployment_type", &obj.DeploymentType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetEncryptionCapabilityResponse : GetEncryptionCapabilityResponse struct
type GetEncryptionCapabilityResponse struct {
	Capability []Capability `json:"capability,omitempty"`
}

// UnmarshalGetEncryptionCapabilityResponse unmarshals an instance of GetEncryptionCapabilityResponse from the specified map of raw messages.
func UnmarshalGetEncryptionCapabilityResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetEncryptionCapabilityResponse)
	err = core.UnmarshalModel(m, "capability", &obj.Capability, UnmarshalCapability)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetEndpointsCapabilityOptions : The GetEndpointsCapability options.
type GetEndpointsCapabilityOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	Deployment *GetEndpointsCapabilityRequestDeployment `json:"deployment,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetEndpointsCapabilityOptions : Instantiate GetEndpointsCapabilityOptions
func (*CloudDatabasesV5) NewGetEndpointsCapabilityOptions(id string) *GetEndpointsCapabilityOptions {
	return &GetEndpointsCapabilityOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetEndpointsCapabilityOptions) SetID(id string) *GetEndpointsCapabilityOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetDeployment : Allow user to set Deployment
func (_options *GetEndpointsCapabilityOptions) SetDeployment(deployment *GetEndpointsCapabilityRequestDeployment) *GetEndpointsCapabilityOptions {
	_options.Deployment = deployment
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetEndpointsCapabilityOptions) SetHeaders(param map[string]string) *GetEndpointsCapabilityOptions {
	options.Headers = param
	return options
}

// GetEndpointsCapabilityRequestDeployment : GetEndpointsCapabilityRequestDeployment struct
type GetEndpointsCapabilityRequestDeployment struct {
	// Type of database.
	DeploymentType *string `json:"deployment_type,omitempty"`
}

// UnmarshalGetEndpointsCapabilityRequestDeployment unmarshals an instance of GetEndpointsCapabilityRequestDeployment from the specified map of raw messages.
func UnmarshalGetEndpointsCapabilityRequestDeployment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetEndpointsCapabilityRequestDeployment)
	err = core.UnmarshalPrimitive(m, "deployment_type", &obj.DeploymentType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetEndpointsCapabilityResponse : GetEndpointsCapabilityResponse struct
type GetEndpointsCapabilityResponse struct {
	Capability []Capability `json:"capability,omitempty"`
}

// UnmarshalGetEndpointsCapabilityResponse unmarshals an instance of GetEndpointsCapabilityResponse from the specified map of raw messages.
func UnmarshalGetEndpointsCapabilityResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetEndpointsCapabilityResponse)
	err = core.UnmarshalModel(m, "capability", &obj.Capability, UnmarshalCapability)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetGroupsCapabilityOptions : The GetGroupsCapability options.
type GetGroupsCapabilityOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	Deployment *GetGroupsCapabilityRequestDeployment `json:"deployment,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetGroupsCapabilityOptions : Instantiate GetGroupsCapabilityOptions
func (*CloudDatabasesV5) NewGetGroupsCapabilityOptions(id string) *GetGroupsCapabilityOptions {
	return &GetGroupsCapabilityOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetGroupsCapabilityOptions) SetID(id string) *GetGroupsCapabilityOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetDeployment : Allow user to set Deployment
func (_options *GetGroupsCapabilityOptions) SetDeployment(deployment *GetGroupsCapabilityRequestDeployment) *GetGroupsCapabilityOptions {
	_options.Deployment = deployment
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetGroupsCapabilityOptions) SetHeaders(param map[string]string) *GetGroupsCapabilityOptions {
	options.Headers = param
	return options
}

// GetGroupsCapabilityRequestDeployment : GetGroupsCapabilityRequestDeployment struct
type GetGroupsCapabilityRequestDeployment struct {
	// Type of database.
	DeploymentType *string `json:"deployment_type,omitempty"`
}

// UnmarshalGetGroupsCapabilityRequestDeployment unmarshals an instance of GetGroupsCapabilityRequestDeployment from the specified map of raw messages.
func UnmarshalGetGroupsCapabilityRequestDeployment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetGroupsCapabilityRequestDeployment)
	err = core.UnmarshalPrimitive(m, "deployment_type", &obj.DeploymentType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetGroupsCapabilityResponse : GetGroupsCapabilityResponse struct
type GetGroupsCapabilityResponse struct {
	Capability []Capability `json:"capability,omitempty"`
}

// UnmarshalGetGroupsCapabilityResponse unmarshals an instance of GetGroupsCapabilityResponse from the specified map of raw messages.
func UnmarshalGetGroupsCapabilityResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetGroupsCapabilityResponse)
	err = core.UnmarshalModel(m, "capability", &obj.Capability, UnmarshalCapability)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetPitrDataOptions : The GetPitrData options.
type GetPitrDataOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPitrDataOptions : Instantiate GetPitrDataOptions
func (*CloudDatabasesV5) NewGetPitrDataOptions(id string) *GetPitrDataOptions {
	return &GetPitrDataOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetPitrDataOptions) SetID(id string) *GetPitrDataOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPitrDataOptions) SetHeaders(param map[string]string) *GetPitrDataOptions {
	options.Headers = param
	return options
}

// GetPitrDataResponse : GetPitrDataResponse struct
type GetPitrDataResponse struct {
	PointInTimeRecoveryData *PointInTimeRecoveryData `json:"point_in_time_recovery_data,omitempty"`
}

// UnmarshalGetPitrDataResponse unmarshals an instance of GetPitrDataResponse from the specified map of raw messages.
func UnmarshalGetPitrDataResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetPitrDataResponse)
	err = core.UnmarshalModel(m, "point_in_time_recovery_data", &obj.PointInTimeRecoveryData, UnmarshalPointInTimeRecoveryData)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetRegionsCapabilityOptions : The GetRegionsCapability options.
type GetRegionsCapabilityOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	Deployment *GetRegionsCapabilityRequestDeployment `json:"deployment,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRegionsCapabilityOptions : Instantiate GetRegionsCapabilityOptions
func (*CloudDatabasesV5) NewGetRegionsCapabilityOptions(id string) *GetRegionsCapabilityOptions {
	return &GetRegionsCapabilityOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetRegionsCapabilityOptions) SetID(id string) *GetRegionsCapabilityOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetDeployment : Allow user to set Deployment
func (_options *GetRegionsCapabilityOptions) SetDeployment(deployment *GetRegionsCapabilityRequestDeployment) *GetRegionsCapabilityOptions {
	_options.Deployment = deployment
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetRegionsCapabilityOptions) SetHeaders(param map[string]string) *GetRegionsCapabilityOptions {
	options.Headers = param
	return options
}

// GetRegionsCapabilityRequestDeployment : GetRegionsCapabilityRequestDeployment struct
type GetRegionsCapabilityRequestDeployment struct {
	// Type of database.
	DeploymentType *string `json:"deployment_type,omitempty"`

	// database version.
	Version *string `json:"version,omitempty"`

	// classic or satellite.
	Platform *string `json:"platform,omitempty"`
}

// UnmarshalGetRegionsCapabilityRequestDeployment unmarshals an instance of GetRegionsCapabilityRequestDeployment from the specified map of raw messages.
func UnmarshalGetRegionsCapabilityRequestDeployment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetRegionsCapabilityRequestDeployment)
	err = core.UnmarshalPrimitive(m, "deployment_type", &obj.DeploymentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "platform", &obj.Platform)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetRegionsCapabilityResponse : GetRegionsCapabilityResponse struct
type GetRegionsCapabilityResponse struct {
	Capability []Capability `json:"capability,omitempty"`
}

// UnmarshalGetRegionsCapabilityResponse unmarshals an instance of GetRegionsCapabilityResponse from the specified map of raw messages.
func UnmarshalGetRegionsCapabilityResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetRegionsCapabilityResponse)
	err = core.UnmarshalModel(m, "capability", &obj.Capability, UnmarshalCapability)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetRemotesCapabilityOptions : The GetRemotesCapability options.
type GetRemotesCapabilityOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	Deployment *GetRemotesCapabilityRequestDeployment `json:"deployment,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRemotesCapabilityOptions : Instantiate GetRemotesCapabilityOptions
func (*CloudDatabasesV5) NewGetRemotesCapabilityOptions(id string) *GetRemotesCapabilityOptions {
	return &GetRemotesCapabilityOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetRemotesCapabilityOptions) SetID(id string) *GetRemotesCapabilityOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetDeployment : Allow user to set Deployment
func (_options *GetRemotesCapabilityOptions) SetDeployment(deployment *GetRemotesCapabilityRequestDeployment) *GetRemotesCapabilityOptions {
	_options.Deployment = deployment
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetRemotesCapabilityOptions) SetHeaders(param map[string]string) *GetRemotesCapabilityOptions {
	options.Headers = param
	return options
}

// GetRemotesCapabilityRequestDeployment : GetRemotesCapabilityRequestDeployment struct
type GetRemotesCapabilityRequestDeployment struct {
	// Type of database.
	DeploymentType *string `json:"deployment_type,omitempty"`
}

// UnmarshalGetRemotesCapabilityRequestDeployment unmarshals an instance of GetRemotesCapabilityRequestDeployment from the specified map of raw messages.
func UnmarshalGetRemotesCapabilityRequestDeployment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetRemotesCapabilityRequestDeployment)
	err = core.UnmarshalPrimitive(m, "deployment_type", &obj.DeploymentType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetRemotesCapabilityResponse : GetRemotesCapabilityResponse struct
type GetRemotesCapabilityResponse struct {
	Capability []Capability `json:"capability,omitempty"`
}

// UnmarshalGetRemotesCapabilityResponse unmarshals an instance of GetRemotesCapabilityResponse from the specified map of raw messages.
func UnmarshalGetRemotesCapabilityResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetRemotesCapabilityResponse)
	err = core.UnmarshalModel(m, "capability", &obj.Capability, UnmarshalCapability)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetTaskOptions : The GetTask options.
type GetTaskOptions struct {
	// Task ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTaskOptions : Instantiate GetTaskOptions
func (*CloudDatabasesV5) NewGetTaskOptions(id string) *GetTaskOptions {
	return &GetTaskOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetTaskOptions) SetID(id string) *GetTaskOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTaskOptions) SetHeaders(param map[string]string) *GetTaskOptions {
	options.Headers = param
	return options
}

// GetTaskResponse : GetTaskResponse struct
type GetTaskResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalGetTaskResponse unmarshals an instance of GetTaskResponse from the specified map of raw messages.
func UnmarshalGetTaskResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetTaskResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetVersionsCapabilityOptions : The GetVersionsCapability options.
type GetVersionsCapabilityOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	Deployment *GetVersionsCapabilityRequestDeployment `json:"deployment,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetVersionsCapabilityOptions : Instantiate GetVersionsCapabilityOptions
func (*CloudDatabasesV5) NewGetVersionsCapabilityOptions(id string) *GetVersionsCapabilityOptions {
	return &GetVersionsCapabilityOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetVersionsCapabilityOptions) SetID(id string) *GetVersionsCapabilityOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetDeployment : Allow user to set Deployment
func (_options *GetVersionsCapabilityOptions) SetDeployment(deployment *GetVersionsCapabilityRequestDeployment) *GetVersionsCapabilityOptions {
	_options.Deployment = deployment
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetVersionsCapabilityOptions) SetHeaders(param map[string]string) *GetVersionsCapabilityOptions {
	options.Headers = param
	return options
}

// GetVersionsCapabilityRequestDeployment : GetVersionsCapabilityRequestDeployment struct
type GetVersionsCapabilityRequestDeployment struct {
	// Type of database.
	DeploymentType *string `json:"deployment_type,omitempty"`
}

// UnmarshalGetVersionsCapabilityRequestDeployment unmarshals an instance of GetVersionsCapabilityRequestDeployment from the specified map of raw messages.
func UnmarshalGetVersionsCapabilityRequestDeployment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetVersionsCapabilityRequestDeployment)
	err = core.UnmarshalPrimitive(m, "deployment_type", &obj.DeploymentType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetVersionsCapabilityResponse : GetVersionsCapabilityResponse struct
type GetVersionsCapabilityResponse struct {
	Capability []Capability `json:"capability,omitempty"`
}

// UnmarshalGetVersionsCapabilityResponse unmarshals an instance of GetVersionsCapabilityResponse from the specified map of raw messages.
func UnmarshalGetVersionsCapabilityResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetVersionsCapabilityResponse)
	err = core.UnmarshalModel(m, "capability", &obj.Capability, UnmarshalCapability)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Group : Group struct
type Group struct {
	// Id/name for group.
	ID *string `json:"id,omitempty"`

	// Number of entities in the group.
	Count *int64 `json:"count,omitempty"`

	Members *GroupMembers `json:"members,omitempty"`

	Memory *GroupMemory `json:"memory,omitempty"`

	CPU *GroupCPU `json:"cpu,omitempty"`

	Disk *GroupDisk `json:"disk,omitempty"`
}

// UnmarshalGroup unmarshals an instance of Group from the specified map of raw messages.
func UnmarshalGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Group)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "members", &obj.Members, UnmarshalGroupMembers)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "memory", &obj.Memory, UnmarshalGroupMemory)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cpu", &obj.CPU, UnmarshalGroupCPU)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disk", &obj.Disk, UnmarshalGroupDisk)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupCPU : GroupCPU struct
type GroupCPU struct {
	// Units used for scaling cpu - count means the value is the number of the unit(s) available.
	Units *string `json:"units,omitempty"`

	// Number of allocated CPUs.
	AllocationCount *int64 `json:"allocation_count,omitempty"`

	// Minimum number of CPUs.
	MinimumCount *int64 `json:"minimum_count,omitempty"`

	// Maximum number of CPUs.
	MaximumCount *int64 `json:"maximum_count,omitempty"`

	// Step size CPUs can be adjusted.
	StepSizeCount *int64 `json:"step_size_count,omitempty"`

	// Is this group's CPU count adjustable.
	IsAdjustable *bool `json:"is_adjustable,omitempty"`

	// Is this group's CPU optional?.
	IsOptional *bool `json:"is_optional,omitempty"`

	// Can this group's CPU scale down?.
	CanScaleDown *bool `json:"can_scale_down,omitempty"`
}

// UnmarshalGroupCPU unmarshals an instance of GroupCPU from the specified map of raw messages.
func UnmarshalGroupCPU(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupCPU)
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allocation_count", &obj.AllocationCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_count", &obj.MinimumCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_count", &obj.MaximumCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "step_size_count", &obj.StepSizeCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_adjustable", &obj.IsAdjustable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_optional", &obj.IsOptional)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "can_scale_down", &obj.CanScaleDown)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupDisk : GroupDisk struct
type GroupDisk struct {
	// Units used for scaling storage.
	Units *string `json:"units,omitempty"`

	// Allocated storage in MB.
	AllocationMb *int64 `json:"allocation_mb,omitempty"`

	// Minimum allocated storage.
	MinimumMb *int64 `json:"minimum_mb,omitempty"`

	// Maximum allocated storage.
	MaximumMb *int64 `json:"maximum_mb,omitempty"`

	// Step size storage can be adjusted.
	StepSizeMb *int64 `json:"step_size_mb,omitempty"`

	// Is this group's storage adjustable?.
	IsAdjustable *bool `json:"is_adjustable,omitempty"`

	// Is this group's storage optional?.
	IsOptional *bool `json:"is_optional,omitempty"`

	// Can this group's storage scale down?.
	CanScaleDown *bool `json:"can_scale_down,omitempty"`
}

// UnmarshalGroupDisk unmarshals an instance of GroupDisk from the specified map of raw messages.
func UnmarshalGroupDisk(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupDisk)
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allocation_mb", &obj.AllocationMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_mb", &obj.MinimumMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_mb", &obj.MaximumMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "step_size_mb", &obj.StepSizeMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_adjustable", &obj.IsAdjustable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_optional", &obj.IsOptional)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "can_scale_down", &obj.CanScaleDown)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupMembers : GroupMembers struct
type GroupMembers struct {
	// Units used for scaling number of members.
	Units *string `json:"units,omitempty"`

	// Allocated number of members.
	AllocationCount *int64 `json:"allocation_count,omitempty"`

	// Minimum number of members.
	MinimumCount *int64 `json:"minimum_count,omitempty"`

	// Maximum number of members.
	MaximumCount *int64 `json:"maximum_count,omitempty"`

	// Step size for number of members.
	StepSizeCount *int64 `json:"step_size_count,omitempty"`

	// Is this deployment's number of members adjustable?.
	IsAdjustable *bool `json:"is_adjustable,omitempty"`

	// Is this deployments's number of members optional?.
	IsOptional *bool `json:"is_optional,omitempty"`

	// Can this deployment's number of members scale down?.
	CanScaleDown *bool `json:"can_scale_down,omitempty"`
}

// UnmarshalGroupMembers unmarshals an instance of GroupMembers from the specified map of raw messages.
func UnmarshalGroupMembers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupMembers)
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allocation_count", &obj.AllocationCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_count", &obj.MinimumCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_count", &obj.MaximumCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "step_size_count", &obj.StepSizeCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_adjustable", &obj.IsAdjustable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_optional", &obj.IsOptional)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "can_scale_down", &obj.CanScaleDown)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupMemory : GroupMemory struct
type GroupMemory struct {
	// Units used for scaling memory.
	Units *string `json:"units,omitempty"`

	// Allocated memory in MB.
	AllocationMb *int64 `json:"allocation_mb,omitempty"`

	// Minimum memory in MB.
	MinimumMb *int64 `json:"minimum_mb,omitempty"`

	// Maximum memory in MB.
	MaximumMb *int64 `json:"maximum_mb,omitempty"`

	// Step size memory can be adjusted by in MB.
	StepSizeMb *int64 `json:"step_size_mb,omitempty"`

	// Is this group's memory adjustable?.
	IsAdjustable *bool `json:"is_adjustable,omitempty"`

	// Is this group's memory optional?.
	IsOptional *bool `json:"is_optional,omitempty"`

	// Can this group's memory scale down?.
	CanScaleDown *bool `json:"can_scale_down,omitempty"`
}

// UnmarshalGroupMemory unmarshals an instance of GroupMemory from the specified map of raw messages.
func UnmarshalGroupMemory(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupMemory)
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allocation_mb", &obj.AllocationMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_mb", &obj.MinimumMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_mb", &obj.MaximumMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "step_size_mb", &obj.StepSizeMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_adjustable", &obj.IsAdjustable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_optional", &obj.IsOptional)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "can_scale_down", &obj.CanScaleDown)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupScaling : GroupScaling struct
type GroupScaling struct {
	Members *GroupScalingMembers `json:"members,omitempty"`

	Memory *GroupScalingMemory `json:"memory,omitempty"`

	CPU *GroupScalingCPU `json:"cpu,omitempty"`

	Disk *GroupScalingDisk `json:"disk,omitempty"`
}

// UnmarshalGroupScaling unmarshals an instance of GroupScaling from the specified map of raw messages.
func UnmarshalGroupScaling(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupScaling)
	err = core.UnmarshalModel(m, "members", &obj.Members, UnmarshalGroupScalingMembers)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "memory", &obj.Memory, UnmarshalGroupScalingMemory)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cpu", &obj.CPU, UnmarshalGroupScalingCPU)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disk", &obj.Disk, UnmarshalGroupScalingDisk)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupScalingCPU : GroupScalingCPU struct
type GroupScalingCPU struct {
	// Number of allocated CPUs.
	AllocationCount *int64 `json:"allocation_count,omitempty"`
}

// UnmarshalGroupScalingCPU unmarshals an instance of GroupScalingCPU from the specified map of raw messages.
func UnmarshalGroupScalingCPU(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupScalingCPU)
	err = core.UnmarshalPrimitive(m, "allocation_count", &obj.AllocationCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupScalingDisk : GroupScalingDisk struct
type GroupScalingDisk struct {
	// Allocated storage in MB.
	AllocationMb *int64 `json:"allocation_mb,omitempty"`
}

// UnmarshalGroupScalingDisk unmarshals an instance of GroupScalingDisk from the specified map of raw messages.
func UnmarshalGroupScalingDisk(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupScalingDisk)
	err = core.UnmarshalPrimitive(m, "allocation_mb", &obj.AllocationMb)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupScalingMembers : GroupScalingMembers struct
type GroupScalingMembers struct {
	// Allocated number of members.
	AllocationCount *int64 `json:"allocation_count,omitempty"`
}

// UnmarshalGroupScalingMembers unmarshals an instance of GroupScalingMembers from the specified map of raw messages.
func UnmarshalGroupScalingMembers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupScalingMembers)
	err = core.UnmarshalPrimitive(m, "allocation_count", &obj.AllocationCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupScalingMemory : GroupScalingMemory struct
type GroupScalingMemory struct {
	// Allocated memory in MB.
	AllocationMb *int64 `json:"allocation_mb,omitempty"`
}

// UnmarshalGroupScalingMemory unmarshals an instance of GroupScalingMemory from the specified map of raw messages.
func UnmarshalGroupScalingMemory(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupScalingMemory)
	err = core.UnmarshalPrimitive(m, "allocation_mb", &obj.AllocationMb)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KillConnectionsOptions : The KillConnections options.
type KillConnectionsOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewKillConnectionsOptions : Instantiate KillConnectionsOptions
func (*CloudDatabasesV5) NewKillConnectionsOptions(id string) *KillConnectionsOptions {
	return &KillConnectionsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *KillConnectionsOptions) SetID(id string) *KillConnectionsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *KillConnectionsOptions) SetHeaders(param map[string]string) *KillConnectionsOptions {
	options.Headers = param
	return options
}

// KillConnectionsResponse : KillConnectionsResponse struct
type KillConnectionsResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalKillConnectionsResponse unmarshals an instance of KillConnectionsResponse from the specified map of raw messages.
func UnmarshalKillConnectionsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KillConnectionsResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListDeployablesOptions : The ListDeployables options.
type ListDeployablesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDeployablesOptions : Instantiate ListDeployablesOptions
func (*CloudDatabasesV5) NewListDeployablesOptions() *ListDeployablesOptions {
	return &ListDeployablesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListDeployablesOptions) SetHeaders(param map[string]string) *ListDeployablesOptions {
	options.Headers = param
	return options
}

// ListDeployablesResponse : ListDeployablesResponse struct
type ListDeployablesResponse struct {
	Deployables []Deployables `json:"deployables,omitempty"`
}

// UnmarshalListDeployablesResponse unmarshals an instance of ListDeployablesResponse from the specified map of raw messages.
func UnmarshalListDeployablesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListDeployablesResponse)
	err = core.UnmarshalModel(m, "deployables", &obj.Deployables, UnmarshalDeployables)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListDeploymentBackupsOptions : The ListDeploymentBackups options.
type ListDeploymentBackupsOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDeploymentBackupsOptions : Instantiate ListDeploymentBackupsOptions
func (*CloudDatabasesV5) NewListDeploymentBackupsOptions(id string) *ListDeploymentBackupsOptions {
	return &ListDeploymentBackupsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ListDeploymentBackupsOptions) SetID(id string) *ListDeploymentBackupsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDeploymentBackupsOptions) SetHeaders(param map[string]string) *ListDeploymentBackupsOptions {
	options.Headers = param
	return options
}

// ListDeploymentScalingGroupsOptions : The ListDeploymentScalingGroups options.
type ListDeploymentScalingGroupsOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDeploymentScalingGroupsOptions : Instantiate ListDeploymentScalingGroupsOptions
func (*CloudDatabasesV5) NewListDeploymentScalingGroupsOptions(id string) *ListDeploymentScalingGroupsOptions {
	return &ListDeploymentScalingGroupsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ListDeploymentScalingGroupsOptions) SetID(id string) *ListDeploymentScalingGroupsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDeploymentScalingGroupsOptions) SetHeaders(param map[string]string) *ListDeploymentScalingGroupsOptions {
	options.Headers = param
	return options
}

// ListDeploymentScalingGroupsResponse : ListDeploymentScalingGroupsResponse struct
type ListDeploymentScalingGroupsResponse struct {
	Groups []Group `json:"groups,omitempty"`
}

// UnmarshalListDeploymentScalingGroupsResponse unmarshals an instance of ListDeploymentScalingGroupsResponse from the specified map of raw messages.
func UnmarshalListDeploymentScalingGroupsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListDeploymentScalingGroupsResponse)
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalGroup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListDeploymentTasksOptions : The ListDeploymentTasks options.
type ListDeploymentTasksOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDeploymentTasksOptions : Instantiate ListDeploymentTasksOptions
func (*CloudDatabasesV5) NewListDeploymentTasksOptions(id string) *ListDeploymentTasksOptions {
	return &ListDeploymentTasksOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ListDeploymentTasksOptions) SetID(id string) *ListDeploymentTasksOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDeploymentTasksOptions) SetHeaders(param map[string]string) *ListDeploymentTasksOptions {
	options.Headers = param
	return options
}

// ListRegionsOptions : The ListRegions options.
type ListRegionsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListRegionsOptions : Instantiate ListRegionsOptions
func (*CloudDatabasesV5) NewListRegionsOptions() *ListRegionsOptions {
	return &ListRegionsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListRegionsOptions) SetHeaders(param map[string]string) *ListRegionsOptions {
	options.Headers = param
	return options
}

// ListRegionsResponse : ListRegionsResponse struct
type ListRegionsResponse struct {
	// An array of region ids.
	Regions []string `json:"regions,omitempty"`
}

// UnmarshalListRegionsResponse unmarshals an instance of ListRegionsResponse from the specified map of raw messages.
func UnmarshalListRegionsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListRegionsResponse)
	err = core.UnmarshalPrimitive(m, "regions", &obj.Regions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListRemotesOptions : The ListRemotes options.
type ListRemotesOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListRemotesOptions : Instantiate ListRemotesOptions
func (*CloudDatabasesV5) NewListRemotesOptions(id string) *ListRemotesOptions {
	return &ListRemotesOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ListRemotesOptions) SetID(id string) *ListRemotesOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListRemotesOptions) SetHeaders(param map[string]string) *ListRemotesOptions {
	options.Headers = param
	return options
}

// ListRemotesResponse : ListRemotesResponse struct
type ListRemotesResponse struct {
	// Remotes.
	Remotes *Remotes `json:"remotes,omitempty"`
}

// UnmarshalListRemotesResponse unmarshals an instance of ListRemotesResponse from the specified map of raw messages.
func UnmarshalListRemotesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListRemotesResponse)
	err = core.UnmarshalModel(m, "remotes", &obj.Remotes, UnmarshalRemotes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MongoDbConnectionURI : MongoDbConnectionURI struct
type MongoDbConnectionURI struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// Scheme/protocol for URI connection.
	Scheme *string `json:"scheme,omitempty"`

	Hosts []ConnectionHost `json:"hosts,omitempty"`

	// Path for URI connection.
	Path *string `json:"path,omitempty"`

	// Query options to add to the URI connection.
	QueryOptions map[string]interface{} `json:"query_options,omitempty"`

	// Authentication data for Connection String.
	Authentication *ConnectionAuthentication `json:"authentication,omitempty"`

	Certificate *ConnectionCertificate `json:"certificate,omitempty"`

	// Indicates ssl is required for the connection.
	Ssl *bool `json:"ssl,omitempty"`

	// Indicates the address is accessible by browser.
	BrowserAccessible *bool `json:"browser_accessible,omitempty"`

	// Name of the database to use in the URI connection.
	Database *string `json:"database,omitempty"`

	// Name of the replica set to use in the URI connection.
	ReplicaSet *string `json:"replica_set,omitempty"`
}

// UnmarshalMongoDbConnectionURI unmarshals an instance of MongoDbConnectionURI from the specified map of raw messages.
func UnmarshalMongoDbConnectionURI(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MongoDbConnectionURI)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalConnectionHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_options", &obj.QueryOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalConnectionAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalConnectionCertificate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "browser_accessible", &obj.BrowserAccessible)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database", &obj.Database)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "replica_set", &obj.ReplicaSet)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MySQLConnectionURI : MySQLConnectionURI struct
type MySQLConnectionURI struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// Scheme/protocol for URI connection.
	Scheme *string `json:"scheme,omitempty"`

	Hosts []ConnectionHost `json:"hosts,omitempty"`

	// Path for URI connection.
	Path *string `json:"path,omitempty"`

	// Query options to add to the URI connection.
	QueryOptions map[string]interface{} `json:"query_options,omitempty"`

	// Authentication data for Connection String.
	Authentication *ConnectionAuthentication `json:"authentication,omitempty"`

	Certificate *ConnectionCertificate `json:"certificate,omitempty"`

	// Indicates ssl is required for the connection.
	Ssl *bool `json:"ssl,omitempty"`

	// Indicates the address is accessible by browser.
	BrowserAccessible *bool `json:"browser_accessible,omitempty"`

	// Name of the database to use in the URI connection.
	Database *string `json:"database,omitempty"`
}

// UnmarshalMySQLConnectionURI unmarshals an instance of MySQLConnectionURI from the specified map of raw messages.
func UnmarshalMySQLConnectionURI(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MySQLConnectionURI)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalConnectionHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_options", &obj.QueryOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalConnectionAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalConnectionCertificate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "browser_accessible", &obj.BrowserAccessible)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database", &obj.Database)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PointInTimeRecoveryData : PointInTimeRecoveryData struct
type PointInTimeRecoveryData struct {
	EarliestPointInTimeRecoveryTime *string `json:"earliest_point_in_time_recovery_time,omitempty"`
}

// UnmarshalPointInTimeRecoveryData unmarshals an instance of PointInTimeRecoveryData from the specified map of raw messages.
func UnmarshalPointInTimeRecoveryData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PointInTimeRecoveryData)
	err = core.UnmarshalPrimitive(m, "earliest_point_in_time_recovery_time", &obj.EarliestPointInTimeRecoveryTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostgreSQLConnectionURI : PostgreSQLConnectionURI struct
type PostgreSQLConnectionURI struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// Scheme/protocol for URI connection.
	Scheme *string `json:"scheme,omitempty"`

	Hosts []ConnectionHost `json:"hosts,omitempty"`

	// Path for URI connection.
	Path *string `json:"path,omitempty"`

	// Query options to add to the URI connection.
	QueryOptions map[string]interface{} `json:"query_options,omitempty"`

	// Authentication data for Connection String.
	Authentication *ConnectionAuthentication `json:"authentication,omitempty"`

	Certificate *ConnectionCertificate `json:"certificate,omitempty"`

	// Indicates ssl is required for the connection.
	Ssl *bool `json:"ssl,omitempty"`

	// Indicates the address is accessible by browser.
	BrowserAccessible *bool `json:"browser_accessible,omitempty"`

	// Name of the database to use in the URI connection.
	Database *string `json:"database,omitempty"`
}

// UnmarshalPostgreSQLConnectionURI unmarshals an instance of PostgreSQLConnectionURI from the specified map of raw messages.
func UnmarshalPostgreSQLConnectionURI(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PostgreSQLConnectionURI)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalConnectionHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_options", &obj.QueryOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalConnectionAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalConnectionCertificate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "browser_accessible", &obj.BrowserAccessible)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database", &obj.Database)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PromoteReadOnlyReplicaOptions : The PromoteReadOnlyReplica options.
type PromoteReadOnlyReplicaOptions struct {
	// Deployment ID of the read-only replica to promote.
	ID *string `json:"id" validate:"required,ne="`

	// Promotion and Upgrade options.
	Promotion map[string]interface{} `json:"promotion,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPromoteReadOnlyReplicaOptions : Instantiate PromoteReadOnlyReplicaOptions
func (*CloudDatabasesV5) NewPromoteReadOnlyReplicaOptions(id string) *PromoteReadOnlyReplicaOptions {
	return &PromoteReadOnlyReplicaOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *PromoteReadOnlyReplicaOptions) SetID(id string) *PromoteReadOnlyReplicaOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetPromotion : Allow user to set Promotion
func (_options *PromoteReadOnlyReplicaOptions) SetPromotion(promotion map[string]interface{}) *PromoteReadOnlyReplicaOptions {
	_options.Promotion = promotion
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PromoteReadOnlyReplicaOptions) SetHeaders(param map[string]string) *PromoteReadOnlyReplicaOptions {
	options.Headers = param
	return options
}

// PromoteReadOnlyReplicaResponse : PromoteReadOnlyReplicaResponse struct
type PromoteReadOnlyReplicaResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalPromoteReadOnlyReplicaResponse unmarshals an instance of PromoteReadOnlyReplicaResponse from the specified map of raw messages.
func UnmarshalPromoteReadOnlyReplicaResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PromoteReadOnlyReplicaResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RedisConnectionURI : RedisConnectionURI struct
type RedisConnectionURI struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// Scheme/protocol for URI connection.
	Scheme *string `json:"scheme,omitempty"`

	Hosts []ConnectionHost `json:"hosts,omitempty"`

	// Path for URI connection.
	Path *string `json:"path,omitempty"`

	// Query options to add to the URI connection.
	QueryOptions map[string]interface{} `json:"query_options,omitempty"`

	// Authentication data for Connection String.
	Authentication *ConnectionAuthentication `json:"authentication,omitempty"`

	Certificate *ConnectionCertificate `json:"certificate,omitempty"`

	// Indicates ssl is required for the connection.
	Ssl *bool `json:"ssl,omitempty"`

	// Indicates the address is accessible by browser.
	BrowserAccessible *bool `json:"browser_accessible,omitempty"`

	// Number of the database to use in the URI connection.
	Database *int64 `json:"database,omitempty"`
}

// UnmarshalRedisConnectionURI unmarshals an instance of RedisConnectionURI from the specified map of raw messages.
func UnmarshalRedisConnectionURI(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RedisConnectionURI)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalConnectionHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_options", &obj.QueryOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalConnectionAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalConnectionCertificate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "browser_accessible", &obj.BrowserAccessible)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database", &obj.Database)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Remotes : Remotes.
type Remotes struct {
	// Leader ID, if applicable.
	Leader *string `json:"leader,omitempty"`

	// Replica IDs, if applicable.
	Replicas []string `json:"replicas,omitempty"`
}

// UnmarshalRemotes unmarshals an instance of Remotes from the specified map of raw messages.
func UnmarshalRemotes(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Remotes)
	err = core.UnmarshalPrimitive(m, "leader", &obj.Leader)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "replicas", &obj.Replicas)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResyncReplicaOptions : The ResyncReplica options.
type ResyncReplicaOptions struct {
	// Deployment ID of the read-only replica.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewResyncReplicaOptions : Instantiate ResyncReplicaOptions
func (*CloudDatabasesV5) NewResyncReplicaOptions(id string) *ResyncReplicaOptions {
	return &ResyncReplicaOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ResyncReplicaOptions) SetID(id string) *ResyncReplicaOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ResyncReplicaOptions) SetHeaders(param map[string]string) *ResyncReplicaOptions {
	options.Headers = param
	return options
}

// ResyncReplicaResponse : ResyncReplicaResponse struct
type ResyncReplicaResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalResyncReplicaResponse unmarshals an instance of ResyncReplicaResponse from the specified map of raw messages.
func UnmarshalResyncReplicaResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResyncReplicaResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetAllowlistOptions : The SetAllowlist options.
type SetAllowlistOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	IPAddresses []AllowlistEntry `json:"ip_addresses,omitempty"`

	// Verify that the current allowlist matches a provided ETag value. Use in conjunction with the GET operation's ETag
	// header to ensure synchronicity between clients.
	IfMatch *string `json:"If-Match,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetAllowlistOptions : Instantiate SetAllowlistOptions
func (*CloudDatabasesV5) NewSetAllowlistOptions(id string) *SetAllowlistOptions {
	return &SetAllowlistOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *SetAllowlistOptions) SetID(id string) *SetAllowlistOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIPAddresses : Allow user to set IPAddresses
func (_options *SetAllowlistOptions) SetIPAddresses(ipAddresses []AllowlistEntry) *SetAllowlistOptions {
	_options.IPAddresses = ipAddresses
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *SetAllowlistOptions) SetIfMatch(ifMatch string) *SetAllowlistOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SetAllowlistOptions) SetHeaders(param map[string]string) *SetAllowlistOptions {
	options.Headers = param
	return options
}

// SetAllowlistResponse : SetAllowlistResponse struct
type SetAllowlistResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalSetAllowlistResponse unmarshals an instance of SetAllowlistResponse from the specified map of raw messages.
func UnmarshalSetAllowlistResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetAllowlistResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetAutoscalingConditionsOptions : The SetAutoscalingConditions options.
type SetAutoscalingConditionsOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Group ID.
	GroupID *string `json:"group_id" validate:"required,ne="`

	Autoscaling AutoscalingSetGroupAutoscalingIntf `json:"autoscaling" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetAutoscalingConditionsOptions : Instantiate SetAutoscalingConditionsOptions
func (*CloudDatabasesV5) NewSetAutoscalingConditionsOptions(id string, groupID string, autoscaling AutoscalingSetGroupAutoscalingIntf) *SetAutoscalingConditionsOptions {
	return &SetAutoscalingConditionsOptions{
		ID: core.StringPtr(id),
		GroupID: core.StringPtr(groupID),
		Autoscaling: autoscaling,
	}
}

// SetID : Allow user to set ID
func (_options *SetAutoscalingConditionsOptions) SetID(id string) *SetAutoscalingConditionsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetGroupID : Allow user to set GroupID
func (_options *SetAutoscalingConditionsOptions) SetGroupID(groupID string) *SetAutoscalingConditionsOptions {
	_options.GroupID = core.StringPtr(groupID)
	return _options
}

// SetAutoscaling : Allow user to set Autoscaling
func (_options *SetAutoscalingConditionsOptions) SetAutoscaling(autoscaling AutoscalingSetGroupAutoscalingIntf) *SetAutoscalingConditionsOptions {
	_options.Autoscaling = autoscaling
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SetAutoscalingConditionsOptions) SetHeaders(param map[string]string) *SetAutoscalingConditionsOptions {
	options.Headers = param
	return options
}

// SetAutoscalingConditionsResponse : SetAutoscalingConditionsResponse struct
type SetAutoscalingConditionsResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalSetAutoscalingConditionsResponse unmarshals an instance of SetAutoscalingConditionsResponse from the specified map of raw messages.
func UnmarshalSetAutoscalingConditionsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetAutoscalingConditionsResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetDeploymentScalingGroupOptions : The SetDeploymentScalingGroup options.
type SetDeploymentScalingGroupOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Group Id.
	GroupID *string `json:"group_id" validate:"required,ne="`

	Group *GroupScaling `json:"group,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetDeploymentScalingGroupOptions : Instantiate SetDeploymentScalingGroupOptions
func (*CloudDatabasesV5) NewSetDeploymentScalingGroupOptions(id string, groupID string) *SetDeploymentScalingGroupOptions {
	return &SetDeploymentScalingGroupOptions{
		ID: core.StringPtr(id),
		GroupID: core.StringPtr(groupID),
	}
}

// SetID : Allow user to set ID
func (_options *SetDeploymentScalingGroupOptions) SetID(id string) *SetDeploymentScalingGroupOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetGroupID : Allow user to set GroupID
func (_options *SetDeploymentScalingGroupOptions) SetGroupID(groupID string) *SetDeploymentScalingGroupOptions {
	_options.GroupID = core.StringPtr(groupID)
	return _options
}

// SetGroup : Allow user to set Group
func (_options *SetDeploymentScalingGroupOptions) SetGroup(group *GroupScaling) *SetDeploymentScalingGroupOptions {
	_options.Group = group
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SetDeploymentScalingGroupOptions) SetHeaders(param map[string]string) *SetDeploymentScalingGroupOptions {
	options.Headers = param
	return options
}

// SetDeploymentScalingGroupResponse : SetDeploymentScalingGroupResponse struct
type SetDeploymentScalingGroupResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalSetDeploymentScalingGroupResponse unmarshals an instance of SetDeploymentScalingGroupResponse from the specified map of raw messages.
func UnmarshalSetDeploymentScalingGroupResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetDeploymentScalingGroupResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StartOndemandBackupOptions : The StartOndemandBackup options.
type StartOndemandBackupOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewStartOndemandBackupOptions : Instantiate StartOndemandBackupOptions
func (*CloudDatabasesV5) NewStartOndemandBackupOptions(id string) *StartOndemandBackupOptions {
	return &StartOndemandBackupOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *StartOndemandBackupOptions) SetID(id string) *StartOndemandBackupOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *StartOndemandBackupOptions) SetHeaders(param map[string]string) *StartOndemandBackupOptions {
	options.Headers = param
	return options
}

// StartOndemandBackupResponse : StartOndemandBackupResponse struct
type StartOndemandBackupResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalStartOndemandBackupResponse unmarshals an instance of StartOndemandBackupResponse from the specified map of raw messages.
func UnmarshalStartOndemandBackupResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StartOndemandBackupResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Task : Task struct
type Task struct {
	// ID of the task.
	ID *string `json:"id,omitempty"`

	// Human-readable description of the task.
	Description *string `json:"description,omitempty"`

	// The status of the task.
	Status *string `json:"status,omitempty"`

	// ID of the deployment the task is being performed on.
	DeploymentID *string `json:"deployment_id,omitempty"`

	// Indicator as percentage of progress of the task.
	ProgressPercent *int64 `json:"progress_percent,omitempty"`

	// Date and time when the task was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`
}

// Constants associated with the Task.Status property.
// The status of the task.
const (
	TaskStatusCompletedConst = "completed"
	TaskStatusFailedConst = "failed"
	TaskStatusRunningConst = "running"
)

// UnmarshalTask unmarshals an instance of Task from the specified map of raw messages.
func UnmarshalTask(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Task)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployment_id", &obj.DeploymentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "progress_percent", &obj.ProgressPercent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Tasks : Tasks struct
type Tasks struct {
	Tasks []Task `json:"tasks,omitempty"`
}

// UnmarshalTasks unmarshals an instance of Tasks from the specified map of raw messages.
func UnmarshalTasks(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Tasks)
	err = core.UnmarshalModel(m, "tasks", &obj.Tasks, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateDatabaseConfigurationOptions : The UpdateDatabaseConfiguration options.
type UpdateDatabaseConfigurationOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	Configuration ConfigurationIntf `json:"configuration,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDatabaseConfigurationOptions : Instantiate UpdateDatabaseConfigurationOptions
func (*CloudDatabasesV5) NewUpdateDatabaseConfigurationOptions(id string) *UpdateDatabaseConfigurationOptions {
	return &UpdateDatabaseConfigurationOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *UpdateDatabaseConfigurationOptions) SetID(id string) *UpdateDatabaseConfigurationOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfiguration : Allow user to set Configuration
func (_options *UpdateDatabaseConfigurationOptions) SetConfiguration(configuration ConfigurationIntf) *UpdateDatabaseConfigurationOptions {
	_options.Configuration = configuration
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDatabaseConfigurationOptions) SetHeaders(param map[string]string) *UpdateDatabaseConfigurationOptions {
	options.Headers = param
	return options
}

// UpdateDatabaseConfigurationResponse : UpdateDatabaseConfigurationResponse struct
type UpdateDatabaseConfigurationResponse struct {
	Task *Task `json:"task,omitempty"`
}

// UnmarshalUpdateDatabaseConfigurationResponse unmarshals an instance of UpdateDatabaseConfigurationResponse from the specified map of raw messages.
func UnmarshalUpdateDatabaseConfigurationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateDatabaseConfigurationResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// User : User struct
type User struct {
	// Username for new user.
	Username *string `json:"username" validate:"required"`

	// Password for new user.
	Password *string `json:"password" validate:"required"`

	// Role for new user. Available for MongoDB Enterprise Ops Manager users ONLY.
	Role *string `json:"role,omitempty"`
}

// Constants associated with the User.Role property.
// Role for new user. Available for MongoDB Enterprise Ops Manager users ONLY.
const (
	UserRoleGroupDataAccessAdminConst = "group_data_access_admin"
	UserRoleGroupReadOnlyConst = "group_read_only"
)

// NewUser : Instantiate User (Generic Model Constructor)
func (*CloudDatabasesV5) NewUser(username string, password string) (_model *User, err error) {
	_model = &User{
		Username: core.StringPtr(username),
		Password: core.StringPtr(password),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalUser unmarshals an instance of User from the specified map of raw messages.
func UnmarshalUser(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(User)
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingSetGroupAutoscalingAutoscalingCPUGroup : AutoscalingSetGroupAutoscalingAutoscalingCPUGroup struct
// This model "extends" AutoscalingSetGroupAutoscaling
type AutoscalingSetGroupAutoscalingAutoscalingCPUGroup struct {
	CPU *AutoscalingCPUGroupCPU `json:"cpu,omitempty"`
}

func (*AutoscalingSetGroupAutoscalingAutoscalingCPUGroup) isaAutoscalingSetGroupAutoscaling() bool {
	return true
}

// UnmarshalAutoscalingSetGroupAutoscalingAutoscalingCPUGroup unmarshals an instance of AutoscalingSetGroupAutoscalingAutoscalingCPUGroup from the specified map of raw messages.
func UnmarshalAutoscalingSetGroupAutoscalingAutoscalingCPUGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingSetGroupAutoscalingAutoscalingCPUGroup)
	err = core.UnmarshalModel(m, "cpu", &obj.CPU, UnmarshalAutoscalingCPUGroupCPU)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingSetGroupAutoscalingAutoscalingDiskGroup : AutoscalingSetGroupAutoscalingAutoscalingDiskGroup struct
// This model "extends" AutoscalingSetGroupAutoscaling
type AutoscalingSetGroupAutoscalingAutoscalingDiskGroup struct {
	Disk *AutoscalingDiskGroupDisk `json:"disk,omitempty"`
}

func (*AutoscalingSetGroupAutoscalingAutoscalingDiskGroup) isaAutoscalingSetGroupAutoscaling() bool {
	return true
}

// UnmarshalAutoscalingSetGroupAutoscalingAutoscalingDiskGroup unmarshals an instance of AutoscalingSetGroupAutoscalingAutoscalingDiskGroup from the specified map of raw messages.
func UnmarshalAutoscalingSetGroupAutoscalingAutoscalingDiskGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingSetGroupAutoscalingAutoscalingDiskGroup)
	err = core.UnmarshalModel(m, "disk", &obj.Disk, UnmarshalAutoscalingDiskGroupDisk)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup : AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup struct
// This model "extends" AutoscalingSetGroupAutoscaling
type AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup struct {
	Memory *AutoscalingMemoryGroupMemory `json:"memory,omitempty"`
}

func (*AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup) isaAutoscalingSetGroupAutoscaling() bool {
	return true
}

// UnmarshalAutoscalingSetGroupAutoscalingAutoscalingMemoryGroup unmarshals an instance of AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup from the specified map of raw messages.
func UnmarshalAutoscalingSetGroupAutoscalingAutoscalingMemoryGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup)
	err = core.UnmarshalModel(m, "memory", &obj.Memory, UnmarshalAutoscalingMemoryGroupMemory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigurationMySQLConfiguration : MySQL Configuration.
// This model "extends" Configuration
type ConfigurationMySQLConfiguration struct {
	// Maximum age for a binlog in seconds. If a binlog is older, it's archived.
	MysqlMaxBinlogAgeSec *int64 `json:"mysql_max_binlog_age_sec,omitempty"`

	// Maximum number of allowed MySQL connections.
	MaxConnections *int64 `json:"max_connections,omitempty"`

	// Determines which authentication plugin the server assigns to new accounts created by CREATE USER and GRANT
	// statements that do not explicitly specify an authentication plugin.
	MysqlDefaultAuthenticationPlugin *string `json:"mysql_default_authentication_plugin,omitempty"`
}

// Constants associated with the ConfigurationMySQLConfiguration.MysqlDefaultAuthenticationPlugin property.
// Determines which authentication plugin the server assigns to new accounts created by CREATE USER and GRANT statements
// that do not explicitly specify an authentication plugin.
const (
	ConfigurationMySQLConfigurationMysqlDefaultAuthenticationPluginMysqlNativePasswordConst = "mysql_native_password"
	ConfigurationMySQLConfigurationMysqlDefaultAuthenticationPluginSha256PasswordConst = "sha256_password"
)

func (*ConfigurationMySQLConfiguration) isaConfiguration() bool {
	return true
}

// UnmarshalConfigurationMySQLConfiguration unmarshals an instance of ConfigurationMySQLConfiguration from the specified map of raw messages.
func UnmarshalConfigurationMySQLConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigurationMySQLConfiguration)
	err = core.UnmarshalPrimitive(m, "mysql_max_binlog_age_sec", &obj.MysqlMaxBinlogAgeSec)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_connections", &obj.MaxConnections)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "mysql_default_authentication_plugin", &obj.MysqlDefaultAuthenticationPlugin)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigurationPgConfiguration : PostgreSQL and EnterpriseDB Configuration.
// This model "extends" Configuration
type ConfigurationPgConfiguration struct {
	// Maximum connections allowed.
	MaxConnections *int64 `json:"max_connections,omitempty"`

	// Max number of transactions that can be in the "prepared" state simultaneously.
	MaxPreparedTransactions *int64 `json:"max_prepared_transactions,omitempty"`

	// Deadlock timeout in ms. The time to wait on a lock before checking for deadlock.  Also the duration where lock waits
	// will be logged.
	DeadlockTimeout *int64 `json:"deadlock_timeout,omitempty"`

	// Number of simultaneous requests that can be handled efficiently by the disk subsystem.
	EffectiveIoConcurrency *int64 `json:"effective_io_concurrency,omitempty"`

	// Maximum number of simultaneously defined replication slots.
	MaxReplicationSlots *int64 `json:"max_replication_slots,omitempty"`

	// Maximum number of simultaneously running WAL sender processes.
	MaxWalSenders *int64 `json:"max_wal_senders,omitempty"`

	// The number of 8kB shared memory buffers used by the server.  Set to 1/4 of memory.  Setting too high will cause
	// crashes or prevent the database from starting.
	SharedBuffers *int64 `json:"shared_buffers,omitempty"`

	// Sets the current transaction's synchronization level.  Off can result in data loss.  remote_write with enable
	// synchronous replication which will impact performance and availabilty.
	SynchronousCommit *string `json:"synchronous_commit,omitempty"`

	// WAL level.  Set to logical to use logical decoding or logical replication.
	WalLevel *string `json:"wal_level,omitempty"`

	// The number of seconds to wait before forces a switch to the next WAL file if a new file has not been started.
	ArchiveTimeout *int64 `json:"archive_timeout,omitempty"`

	// The minimum number of milliseconds for execution time above which statements will be logged.
	LogMinDurationStatement *int64 `json:"log_min_duration_statement,omitempty"`
}

// Constants associated with the ConfigurationPgConfiguration.SynchronousCommit property.
// Sets the current transaction's synchronization level.  Off can result in data loss.  remote_write with enable
// synchronous replication which will impact performance and availabilty.
const (
	ConfigurationPgConfigurationSynchronousCommitLocalConst = "local"
	ConfigurationPgConfigurationSynchronousCommitOffConst = "off"
)

// Constants associated with the ConfigurationPgConfiguration.WalLevel property.
// WAL level.  Set to logical to use logical decoding or logical replication.
const (
	ConfigurationPgConfigurationWalLevelHotStandbyConst = "hot_standby"
	ConfigurationPgConfigurationWalLevelLogicalConst = "logical"
)

func (*ConfigurationPgConfiguration) isaConfiguration() bool {
	return true
}

// UnmarshalConfigurationPgConfiguration unmarshals an instance of ConfigurationPgConfiguration from the specified map of raw messages.
func UnmarshalConfigurationPgConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigurationPgConfiguration)
	err = core.UnmarshalPrimitive(m, "max_connections", &obj.MaxConnections)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_prepared_transactions", &obj.MaxPreparedTransactions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deadlock_timeout", &obj.DeadlockTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "effective_io_concurrency", &obj.EffectiveIoConcurrency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_replication_slots", &obj.MaxReplicationSlots)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_wal_senders", &obj.MaxWalSenders)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "shared_buffers", &obj.SharedBuffers)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "synchronous_commit", &obj.SynchronousCommit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "wal_level", &obj.WalLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "archive_timeout", &obj.ArchiveTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "log_min_duration_statement", &obj.LogMinDurationStatement)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigurationRedisConfiguration : Redis Configuration.
// This model "extends" Configuration
type ConfigurationRedisConfiguration struct {
	// The maximum memory Redis should use, as bytes.
	MaxmemoryRedis *int64 `json:"maxmemory-redis,omitempty"`

	// The policy with which Redis evicts keys when maximum memory is reached.
	MaxmemoryPolicy *string `json:"maxmemory-policy,omitempty"`

	// If set to yes this will enable AOF persistence.
	Appendonly *string `json:"appendonly,omitempty"`

	// The maximum memory Redis should use, as bytes.
	MaxmemorySamples *int64 `json:"maxmemory-samples,omitempty"`

	// Whether or not to stop accepting writes when background persistence actions fail.
	StopWritesOnBgsaveError *string `json:"stop-writes-on-bgsave-error,omitempty"`
}

// Constants associated with the ConfigurationRedisConfiguration.MaxmemoryPolicy property.
// The policy with which Redis evicts keys when maximum memory is reached.
const (
	ConfigurationRedisConfigurationMaxmemoryPolicyAllkeysLruConst = "allkeys-lru"
	ConfigurationRedisConfigurationMaxmemoryPolicyAllkeysRandomConst = "allkeys-random"
	ConfigurationRedisConfigurationMaxmemoryPolicyNoevictionConst = "noeviction"
	ConfigurationRedisConfigurationMaxmemoryPolicyVolatileLruConst = "volatile-lru"
	ConfigurationRedisConfigurationMaxmemoryPolicyVolatileRandomConst = "volatile-random"
	ConfigurationRedisConfigurationMaxmemoryPolicyVolatileTTLConst = "volatile-ttl"
)

// Constants associated with the ConfigurationRedisConfiguration.Appendonly property.
// If set to yes this will enable AOF persistence.
const (
	ConfigurationRedisConfigurationAppendonlyNoConst = "no"
	ConfigurationRedisConfigurationAppendonlyYesConst = "yes"
)

// Constants associated with the ConfigurationRedisConfiguration.StopWritesOnBgsaveError property.
// Whether or not to stop accepting writes when background persistence actions fail.
const (
	ConfigurationRedisConfigurationStopWritesOnBgsaveErrorNoConst = "no"
	ConfigurationRedisConfigurationStopWritesOnBgsaveErrorYesConst = "yes"
)

func (*ConfigurationRedisConfiguration) isaConfiguration() bool {
	return true
}

// UnmarshalConfigurationRedisConfiguration unmarshals an instance of ConfigurationRedisConfiguration from the specified map of raw messages.
func UnmarshalConfigurationRedisConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigurationRedisConfiguration)
	err = core.UnmarshalPrimitive(m, "maxmemory-redis", &obj.MaxmemoryRedis)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxmemory-policy", &obj.MaxmemoryPolicy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "appendonly", &obj.Appendonly)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxmemory-samples", &obj.MaxmemorySamples)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "stop-writes-on-bgsave-error", &obj.StopWritesOnBgsaveError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionDataStaxConnection : DataStax Connection Strings.
// This model "extends" Connection
type ConnectionDataStaxConnection struct {
	Secure *DataStaxConnectionURI `json:"secure" validate:"required"`
}

func (*ConnectionDataStaxConnection) isaConnection() bool {
	return true
}

// UnmarshalConnectionDataStaxConnection unmarshals an instance of ConnectionDataStaxConnection from the specified map of raw messages.
func UnmarshalConnectionDataStaxConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionDataStaxConnection)
	err = core.UnmarshalModel(m, "secure", &obj.Secure, UnmarshalDataStaxConnectionURI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionElasticsearchConnection : Elasticsearch Connection Strings.
// This model "extends" Connection
type ConnectionElasticsearchConnection struct {
	HTTPS *ConnectionURI `json:"https" validate:"required"`

	// CLI Connection.
	Cli *ConnectionCli `json:"cli" validate:"required"`
}

func (*ConnectionElasticsearchConnection) isaConnection() bool {
	return true
}

// UnmarshalConnectionElasticsearchConnection unmarshals an instance of ConnectionElasticsearchConnection from the specified map of raw messages.
func UnmarshalConnectionElasticsearchConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionElasticsearchConnection)
	err = core.UnmarshalModel(m, "https", &obj.HTTPS, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCli)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionEnterpriseDbConnection : PostgreSQL and EnterpriseDB Connection Strings.
// This model "extends" Connection
type ConnectionEnterpriseDbConnection struct {
	Postgres *PostgreSQLConnectionURI `json:"postgres" validate:"required"`

	Emp *ConnectionURI `json:"emp" validate:"required"`

	// CLI Connection.
	Cli *ConnectionCli `json:"cli" validate:"required"`
}

func (*ConnectionEnterpriseDbConnection) isaConnection() bool {
	return true
}

// UnmarshalConnectionEnterpriseDbConnection unmarshals an instance of ConnectionEnterpriseDbConnection from the specified map of raw messages.
func UnmarshalConnectionEnterpriseDbConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionEnterpriseDbConnection)
	err = core.UnmarshalModel(m, "postgres", &obj.Postgres, UnmarshalPostgreSQLConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "emp", &obj.Emp, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCli)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionEtcdConnection : etcd3 Connection Strings.
// This model "extends" Connection
type ConnectionEtcdConnection struct {
	Grpc *ConnectionURI `json:"grpc" validate:"required"`

	// CLI Connection.
	Cli *ConnectionCli `json:"cli" validate:"required"`
}

func (*ConnectionEtcdConnection) isaConnection() bool {
	return true
}

// UnmarshalConnectionEtcdConnection unmarshals an instance of ConnectionEtcdConnection from the specified map of raw messages.
func UnmarshalConnectionEtcdConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionEtcdConnection)
	err = core.UnmarshalModel(m, "grpc", &obj.Grpc, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCli)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionMongoDbConnection : MongoDB Connection Strings.
// This model "extends" Connection
type ConnectionMongoDbConnection struct {
	Mongodb *MongoDbConnectionURI `json:"mongodb" validate:"required"`

	// CLI Connection.
	Cli *ConnectionCli `json:"cli" validate:"required"`
}

func (*ConnectionMongoDbConnection) isaConnection() bool {
	return true
}

// UnmarshalConnectionMongoDbConnection unmarshals an instance of ConnectionMongoDbConnection from the specified map of raw messages.
func UnmarshalConnectionMongoDbConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionMongoDbConnection)
	err = core.UnmarshalModel(m, "mongodb", &obj.Mongodb, UnmarshalMongoDbConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCli)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionMongoDbeeConnection : MongoDB Connection Strings.
// This model "extends" Connection
type ConnectionMongoDbeeConnection struct {
	Mongodb *MongoDbConnectionURI `json:"mongodb" validate:"required"`

	// CLI Connection.
	Cli *ConnectionCli `json:"cli" validate:"required"`

	BiConnector *ConnectionURI `json:"bi_connector,omitempty"`

	Analytics *ConnectionURI `json:"analytics,omitempty"`
}

func (*ConnectionMongoDbeeConnection) isaConnection() bool {
	return true
}

// UnmarshalConnectionMongoDbeeConnection unmarshals an instance of ConnectionMongoDbeeConnection from the specified map of raw messages.
func UnmarshalConnectionMongoDbeeConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionMongoDbeeConnection)
	err = core.UnmarshalModel(m, "mongodb", &obj.Mongodb, UnmarshalMongoDbConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCli)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "bi_connector", &obj.BiConnector, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "analytics", &obj.Analytics, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionMongoDbeeOpsManagerConnection : MongoDB Connection Strings.
// This model "extends" Connection
type ConnectionMongoDbeeOpsManagerConnection struct {
	OpsManager *ConnectionURI `json:"ops_manager,omitempty"`
}

func (*ConnectionMongoDbeeOpsManagerConnection) isaConnection() bool {
	return true
}

// UnmarshalConnectionMongoDbeeOpsManagerConnection unmarshals an instance of ConnectionMongoDbeeOpsManagerConnection from the specified map of raw messages.
func UnmarshalConnectionMongoDbeeOpsManagerConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionMongoDbeeOpsManagerConnection)
	err = core.UnmarshalModel(m, "ops_manager", &obj.OpsManager, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionMySQLConnection : MySQL Connection Strings.
// This model "extends" Connection
type ConnectionMySQLConnection struct {
	Mysql *MySQLConnectionURI `json:"mysql" validate:"required"`

	// CLI Connection.
	Cli *ConnectionCli `json:"cli" validate:"required"`
}

func (*ConnectionMySQLConnection) isaConnection() bool {
	return true
}

// UnmarshalConnectionMySQLConnection unmarshals an instance of ConnectionMySQLConnection from the specified map of raw messages.
func UnmarshalConnectionMySQLConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionMySQLConnection)
	err = core.UnmarshalModel(m, "mysql", &obj.Mysql, UnmarshalMySQLConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCli)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionPostgreSQLConnection : PostgreSQL and EnterpriseDB Connection Strings.
// This model "extends" Connection
type ConnectionPostgreSQLConnection struct {
	Postgres *PostgreSQLConnectionURI `json:"postgres" validate:"required"`

	// CLI Connection.
	Cli *ConnectionCli `json:"cli" validate:"required"`
}

func (*ConnectionPostgreSQLConnection) isaConnection() bool {
	return true
}

// UnmarshalConnectionPostgreSQLConnection unmarshals an instance of ConnectionPostgreSQLConnection from the specified map of raw messages.
func UnmarshalConnectionPostgreSQLConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionPostgreSQLConnection)
	err = core.UnmarshalModel(m, "postgres", &obj.Postgres, UnmarshalPostgreSQLConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCli)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionRabbitMqConnection : RabbitMQ Connection Strings.
// This model "extends" Connection
type ConnectionRabbitMqConnection struct {
	Amqps *ConnectionURI `json:"amqps" validate:"required"`

	Mqtts *ConnectionURI `json:"mqtts" validate:"required"`

	StompSsl *ConnectionURI `json:"stomp_ssl" validate:"required"`

	HTTPS *ConnectionURI `json:"https" validate:"required"`

	// CLI Connection.
	Cli *ConnectionCli `json:"cli" validate:"required"`
}

func (*ConnectionRabbitMqConnection) isaConnection() bool {
	return true
}

// UnmarshalConnectionRabbitMqConnection unmarshals an instance of ConnectionRabbitMqConnection from the specified map of raw messages.
func UnmarshalConnectionRabbitMqConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionRabbitMqConnection)
	err = core.UnmarshalModel(m, "amqps", &obj.Amqps, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "mqtts", &obj.Mqtts, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "stomp_ssl", &obj.StompSsl, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "https", &obj.HTTPS, UnmarshalConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCli)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionRedisConnection : Redis Connection Strings.
// This model "extends" Connection
type ConnectionRedisConnection struct {
	Rediss *RedisConnectionURI `json:"rediss" validate:"required"`

	// CLI Connection.
	Cli *ConnectionCli `json:"cli" validate:"required"`
}

func (*ConnectionRedisConnection) isaConnection() bool {
	return true
}

// UnmarshalConnectionRedisConnection unmarshals an instance of ConnectionRedisConnection from the specified map of raw messages.
func UnmarshalConnectionRedisConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionRedisConnection)
	err = core.UnmarshalModel(m, "rediss", &obj.Rediss, UnmarshalRedisConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCli)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
