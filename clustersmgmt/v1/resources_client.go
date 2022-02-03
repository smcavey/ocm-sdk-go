/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"bufio"
	"context"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// ResourcesClient is the client of the 'resources' resource.
//
// Manages a collection of resources for a cluster
type ResourcesClient struct {
	transport http.RoundTripper
	path      string
}

// NewResourcesClient creates a new client for the 'resources'
// resource using the given transport to send the requests and receive the
// responses.
func NewResourcesClient(transport http.RoundTripper, path string) *ResourcesClient {
	return &ResourcesClient{
		transport: transport,
		path:      path,
	}
}

// Get creates a request for the 'get' method.
//
// Retrieves a list of resources for a cluster in error state
func (c *ResourcesClient) Get() *ResourcesGetRequest {
	return &ResourcesGetRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// Live returns the target 'cluster_resources' resource.
//
// Retrieves a list of currently available resources for a cluster
func (c *ResourcesClient) Live() *ClusterResourcesClient {
	return NewClusterResourcesClient(
		c.transport,
		path.Join(c.path, "live"),
	)
}

// ResourcesPollRequest is the request for the Poll method.
type ResourcesPollRequest struct {
	request    *ResourcesGetRequest
	interval   time.Duration
	statuses   []int
	predicates []func(interface{}) bool
}

// Parameter adds a query parameter to all the requests that will be used to retrieve the object.
func (r *ResourcesPollRequest) Parameter(name string, value interface{}) *ResourcesPollRequest {
	r.request.Parameter(name, value)
	return r
}

// Header adds a request header to all the requests that will be used to retrieve the object.
func (r *ResourcesPollRequest) Header(name string, value interface{}) *ResourcesPollRequest {
	r.request.Header(name, value)
	return r
}

// Interval sets the polling interval. This parameter is mandatory and must be greater than zero.
func (r *ResourcesPollRequest) Interval(value time.Duration) *ResourcesPollRequest {
	r.interval = value
	return r
}

// Status set the expected status of the response. Multiple values can be set calling this method
// multiple times. The response will be considered successful if the status is any of those values.
func (r *ResourcesPollRequest) Status(value int) *ResourcesPollRequest {
	r.statuses = append(r.statuses, value)
	return r
}

// Predicate adds a predicate that the response should satisfy be considered successful. Multiple
// predicates can be set calling this method multiple times. The response will be considered successful
// if all the predicates are satisfied.
func (r *ResourcesPollRequest) Predicate(value func(*ResourcesGetResponse) bool) *ResourcesPollRequest {
	r.predicates = append(r.predicates, func(response interface{}) bool {
		return value(response.(*ResourcesGetResponse))
	})
	return r
}

// StartContext starts the polling loop. Responses will be considered successful if the status is one of
// the values specified with the Status method and if all the predicates specified with the Predicate
// method return nil.
//
// The context must have a timeout or deadline, otherwise this method will immediately return an error.
func (r *ResourcesPollRequest) StartContext(ctx context.Context) (response *ResourcesPollResponse, err error) {
	result, err := helpers.PollContext(ctx, r.interval, r.statuses, r.predicates, r.task)
	if result != nil {
		response = &ResourcesPollResponse{
			response: result.(*ResourcesGetResponse),
		}
	}
	return
}

// task adapts the types of the request/response types so that they can be used with the generic
// polling function from the helpers package.
func (r *ResourcesPollRequest) task(ctx context.Context) (status int, result interface{}, err error) {
	response, err := r.request.SendContext(ctx)
	if response != nil {
		status = response.Status()
		result = response
	}
	return
}

// ResourcesPollResponse is the response for the Poll method.
type ResourcesPollResponse struct {
	response *ResourcesGetResponse
}

// Status returns the response status code.
func (r *ResourcesPollResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.response.Status()
}

// Header returns header of the response.
func (r *ResourcesPollResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.response.Header()
}

// Error returns the response error.
func (r *ResourcesPollResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.response.Error()
}

// Body returns the value of the 'body' parameter.
//
// List of cluster resources
func (r *ResourcesPollResponse) Body() *ClusterResources {
	return r.response.Body()
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// List of cluster resources
func (r *ResourcesPollResponse) GetBody() (value *ClusterResources, ok bool) {
	return r.response.GetBody()
}

// Poll creates a request to repeatedly retrieve the object till the response has one of a given set
// of states and satisfies a set of predicates.
func (c *ResourcesClient) Poll() *ResourcesPollRequest {
	return &ResourcesPollRequest{
		request: c.Get(),
	}
}

// ResourcesGetRequest is the request for the 'get' method.
type ResourcesGetRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *ResourcesGetRequest) Parameter(name string, value interface{}) *ResourcesGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *ResourcesGetRequest) Header(name string, value interface{}) *ResourcesGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *ResourcesGetRequest) Send() (result *ResourcesGetResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *ResourcesGetRequest) SendContext(ctx context.Context) (result *ResourcesGetResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "GET",
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &ResourcesGetResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	reader := bufio.NewReader(response.Body)
	_, err = reader.Peek(1)
	if err == io.EOF {
		err = nil
		return
	}
	if result.status >= 400 {
		result.err, err = errors.UnmarshalErrorStatus(reader, result.status)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = readResourcesGetResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// ResourcesGetResponse is the response for the 'get' method.
type ResourcesGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *ClusterResources
}

// Status returns the response status code.
func (r *ResourcesGetResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *ResourcesGetResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *ResourcesGetResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
// List of cluster resources
func (r *ResourcesGetResponse) Body() *ClusterResources {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// List of cluster resources
func (r *ResourcesGetResponse) GetBody() (value *ClusterResources, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}
