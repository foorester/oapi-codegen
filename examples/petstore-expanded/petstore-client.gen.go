// Package petstore provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package petstore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// Error defines model for Error.
type Error struct {

	// Error code
	Code int32 `json:"code"`

	// Error message
	Message string `json:"message"`
}

// NewPet defines model for NewPet.
type NewPet struct {

	// Name of the pet
	Name string `json:"name"`

	// Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// Pet defines model for Pet.
type Pet struct {
	// Embedded struct due to allOf(#/components/schemas/NewPet)
	NewPet `yaml:",inline"`
	// Embedded fields due to inline allOf schema

	// Unique id of the pet
	Id int64 `json:"id"`
}

// FindPetsParams defines parameters for FindPets.
type FindPetsParams struct {

	// tags to filter by
	Tags *[]string `json:"tags,omitempty"`

	// maximum number of results to return
	Limit *int32 `json:"limit,omitempty"`
}

// AddPetJSONBody defines parameters for AddPet.
type AddPetJSONBody NewPet

// AddPetJSONRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody AddPetJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// FindPets request
	FindPets(ctx context.Context, params *FindPetsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AddPet request  with any body
	AddPetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AddPet(ctx context.Context, body AddPetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeletePet request
	DeletePet(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindPetById request
	FindPetById(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) FindPets(ctx context.Context, params *FindPetsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFindPetsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AddPetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAddPetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AddPet(ctx context.Context, body AddPetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAddPetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeletePet(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeletePetRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) FindPetById(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFindPetByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewFindPetsRequest generates requests for FindPets
func NewFindPetsRequest(server string, params *FindPetsParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/pets")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if params.Tags != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "tags", *params.Tags); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Limit != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "limit", *params.Limit); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAddPetRequest calls the generic AddPet builder with application/json body
func NewAddPetRequest(server string, body AddPetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAddPetRequestWithBody(server, "application/json", bodyReader)
}

// NewAddPetRequestWithBody generates requests for AddPet with any type of body
func NewAddPetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/pets")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeletePetRequest generates requests for DeletePet
func NewDeletePetRequest(server string, id int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "id", id)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/pets/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewFindPetByIdRequest generates requests for FindPetById
func NewFindPetByIdRequest(server string, id int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "id", id)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/pets/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// FindPets request
	FindPetsWithResponse(ctx context.Context, params *FindPetsParams) (*FindPetsResponse, error)

	// AddPet request  with any body
	AddPetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*AddPetResponse, error)

	AddPetWithResponse(ctx context.Context, body AddPetJSONRequestBody) (*AddPetResponse, error)

	// DeletePet request
	DeletePetWithResponse(ctx context.Context, id int64) (*DeletePetResponse, error)

	// FindPetById request
	FindPetByIdWithResponse(ctx context.Context, id int64) (*FindPetByIdResponse, error)
}

type FindPetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r FindPetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FindPetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AddPetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r AddPetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AddPetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeletePetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r DeletePetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeletePetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type FindPetByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r FindPetByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FindPetByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// FindPetsWithResponse request returning *FindPetsResponse
func (c *ClientWithResponses) FindPetsWithResponse(ctx context.Context, params *FindPetsParams) (*FindPetsResponse, error) {
	rsp, err := c.FindPets(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseFindPetsResponse(rsp)
}

// AddPetWithBodyWithResponse request with arbitrary body returning *AddPetResponse
func (c *ClientWithResponses) AddPetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*AddPetResponse, error) {
	rsp, err := c.AddPetWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseAddPetResponse(rsp)
}

func (c *ClientWithResponses) AddPetWithResponse(ctx context.Context, body AddPetJSONRequestBody) (*AddPetResponse, error) {
	rsp, err := c.AddPet(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseAddPetResponse(rsp)
}

// DeletePetWithResponse request returning *DeletePetResponse
func (c *ClientWithResponses) DeletePetWithResponse(ctx context.Context, id int64) (*DeletePetResponse, error) {
	rsp, err := c.DeletePet(ctx, id)
	if err != nil {
		return nil, err
	}
	return ParseDeletePetResponse(rsp)
}

// FindPetByIdWithResponse request returning *FindPetByIdResponse
func (c *ClientWithResponses) FindPetByIdWithResponse(ctx context.Context, id int64) (*FindPetByIdResponse, error) {
	rsp, err := c.FindPetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return ParseFindPetByIdResponse(rsp)
}

// ParseFindPetsResponse parses an HTTP response from a FindPetsWithResponse call
func ParseFindPetsResponse(rsp *http.Response) (*FindPetsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &FindPetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseAddPetResponse parses an HTTP response from a AddPetWithResponse call
func ParseAddPetResponse(rsp *http.Response) (*AddPetResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &AddPetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseDeletePetResponse parses an HTTP response from a DeletePetWithResponse call
func ParseDeletePetResponse(rsp *http.Response) (*DeletePetResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &DeletePetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseFindPetByIdResponse parses an HTTP response from a FindPetByIdWithResponse call
func ParseFindPetByIdResponse(rsp *http.Response) (*FindPetByIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &FindPetByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
