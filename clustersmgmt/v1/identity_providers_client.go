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
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// IdentityProvidersClient is the client of the 'identity_providers' resource.
//
// Manages the collection of identity providers of a cluster.
type IdentityProvidersClient struct {
	transport http.RoundTripper
	path      string
}

// NewIdentityProvidersClient creates a new client for the 'identity_providers'
// resource using the given transport to send the requests and receive the
// responses.
func NewIdentityProvidersClient(transport http.RoundTripper, path string) *IdentityProvidersClient {
	return &IdentityProvidersClient{
		transport: transport,
		path:      path,
	}
}

// Add creates a request for the 'add' method.
//
// Adds a new identity provider to the cluster.
func (c *IdentityProvidersClient) Add() *IdentityProvidersAddRequest {
	return &IdentityProvidersAddRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// List creates a request for the 'list' method.
//
// Retrieves the list of identity providers.
func (c *IdentityProvidersClient) List() *IdentityProvidersListRequest {
	return &IdentityProvidersListRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// IdentityProvider returns the target 'identity_provider' resource for the given identifier.
//
// Reference to the service that manages an specific identity provider.
func (c *IdentityProvidersClient) IdentityProvider(id string) *IdentityProviderClient {
	return NewIdentityProviderClient(
		c.transport,
		path.Join(c.path, id),
	)
}

// IdentityProvidersAddRequest is the request for the 'add' method.
type IdentityProvidersAddRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	body      *IdentityProvider
}

// Parameter adds a query parameter.
func (r *IdentityProvidersAddRequest) Parameter(name string, value interface{}) *IdentityProvidersAddRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *IdentityProvidersAddRequest) Header(name string, value interface{}) *IdentityProvidersAddRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
// Description of the cluster.
func (r *IdentityProvidersAddRequest) Body(value *IdentityProvider) *IdentityProvidersAddRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *IdentityProvidersAddRequest) Send() (result *IdentityProvidersAddResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *IdentityProvidersAddRequest) SendContext(ctx context.Context) (result *IdentityProvidersAddResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := &bytes.Buffer{}
	err = writeIdentityProvidersAddRequest(r, buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "POST",
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &IdentityProvidersAddResponse{}
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
	err = readIdentityProvidersAddResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// IdentityProvidersAddResponse is the response for the 'add' method.
type IdentityProvidersAddResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *IdentityProvider
}

// Status returns the response status code.
func (r *IdentityProvidersAddResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *IdentityProvidersAddResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *IdentityProvidersAddResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
// Description of the cluster.
func (r *IdentityProvidersAddResponse) Body() *IdentityProvider {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// Description of the cluster.
func (r *IdentityProvidersAddResponse) GetBody() (value *IdentityProvider, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// IdentityProvidersListRequest is the request for the 'list' method.
type IdentityProvidersListRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	page      *int
	size      *int
}

// Parameter adds a query parameter.
func (r *IdentityProvidersListRequest) Parameter(name string, value interface{}) *IdentityProvidersListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *IdentityProvidersListRequest) Header(name string, value interface{}) *IdentityProvidersListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *IdentityProvidersListRequest) Page(value int) *IdentityProvidersListRequest {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Number of items contained in the returned page.
func (r *IdentityProvidersListRequest) Size(value int) *IdentityProvidersListRequest {
	r.size = &value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *IdentityProvidersListRequest) Send() (result *IdentityProvidersListResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *IdentityProvidersListRequest) SendContext(ctx context.Context) (result *IdentityProvidersListResponse, err error) {
	query := helpers.CopyQuery(r.query)
	if r.page != nil {
		helpers.AddValue(&query, "page", *r.page)
	}
	if r.size != nil {
		helpers.AddValue(&query, "size", *r.size)
	}
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
	result = &IdentityProvidersListResponse{}
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
	err = readIdentityProvidersListResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// IdentityProvidersListResponse is the response for the 'list' method.
type IdentityProvidersListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	items  *IdentityProviderList
	page   *int
	size   *int
	total  *int
}

// Status returns the response status code.
func (r *IdentityProvidersListResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *IdentityProvidersListResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *IdentityProvidersListResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of identity providers.
func (r *IdentityProvidersListResponse) Items() *IdentityProviderList {
	if r == nil {
		return nil
	}
	return r.items
}

// GetItems returns the value of the 'items' parameter and
// a flag indicating if the parameter has a value.
//
// Retrieved list of identity providers.
func (r *IdentityProvidersListResponse) GetItems() (value *IdentityProviderList, ok bool) {
	ok = r != nil && r.items != nil
	if ok {
		value = r.items
	}
	return
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *IdentityProvidersListResponse) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the requested page, where one corresponds to the first page.
func (r *IdentityProvidersListResponse) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Number of items contained in the returned page.
func (r *IdentityProvidersListResponse) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Number of items contained in the returned page.
func (r *IdentityProvidersListResponse) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection.
func (r *IdentityProvidersListResponse) Total() int {
	if r != nil && r.total != nil {
		return *r.total
	}
	return 0
}

// GetTotal returns the value of the 'total' parameter and
// a flag indicating if the parameter has a value.
//
// Total number of items of the collection.
func (r *IdentityProvidersListResponse) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}
