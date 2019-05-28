// Package schemas provides primitives to interact the openapi HTTP API.
//
// This is an autogenerated file, any edits which you make here will be lost!
package schemas

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strings"
)

// AnyType1 defines component schema for AnyType1.
type AnyType1 interface{}

// AnyType2 defines component schema for AnyType2.
type AnyType2 interface{}

// CustomStringType defines component schema for CustomStringType.
type CustomStringType string

// GenericObject defines component schema for GenericObject.
type GenericObject map[string]interface{}

// Issue9Params defines parameters for Issue9.
type Issue9Params struct {
	Foo string `json:"foo"`
}

// Issue9RequestBody defines body for Issue9 for application/json ContentType.
type Issue9RequestBody interface{}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example.
	Server string

	// HTTP client with any customized settings, such as certificate chains.
	Client http.Client

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor func(req *http.Request, ctx context.Context) error
}

// The interface specification for the client above.
type ClientInterface interface {
	// Issue9 request with JSON body
	Issue9(ctx context.Context, params *Issue9Params, body *Issue9RequestBody) (*http.Response, error)
}

// Issue9 request with JSON body
func (c *Client) Issue9(ctx context.Context, params *Issue9Params, body *Issue9RequestBody) (*http.Response, error) {
	req, err := NewIssue9Request(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewIssue9Request generates requests for Issue9 with JSON body
func NewIssue9Request(server string, params *Issue9Params, body *Issue9RequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(buf)
	}
	return NewIssue9RequestWithBody(server, params, "application/json", bodyReader)
}

// NewIssue9RequestWithBody generates requests for Issue9 with non-JSON body
func NewIssue9RequestWithBody(server string, params *Issue9Params, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/issues/9", server)

	var queryStrings []string

	var queryParam0 string

	queryParam0, err = runtime.StyleParam("form", true, "foo", params.Foo)
	if err != nil {
		return nil, err
	}

	queryStrings = append(queryStrings, queryParam0)

	if len(queryStrings) != 0 {
		queryUrl += "?" + strings.Join(queryStrings, "&")
	}

	req, err := http.NewRequest("GET", queryUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	//  (GET /issues/9)
	Issue9(ctx echo.Context, params Issue9Params) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Issue9 converts echo context to params.
func (w *ServerInterfaceWrapper) Issue9(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params Issue9Params
	// ------------- Required query parameter "foo" -------------
	if paramValue := ctx.QueryParam("foo"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument foo is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "foo", ctx.QueryParams(), &params.Foo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter foo: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Issue9(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router runtime.EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/issues/9", wrapper.Issue9)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/2xSsW7cMAz9FYGz4cu1U7y1GYpMKZBsvQw8iT4rsCWFpK41DP97IfmMNmg3SnyPfE9P",
	"C9g4pRgoqEC3gNiBJqzllzC/zImO0C1rs58+lY4jseyT+higg5fBi5Eh5tGZMxkMxgcl7tHSssLawEMW",
	"jdOzsg+XMqOM6CNPqNCBrU1oQGsHpMIK7RsFYm+fzm9ktXBuiLhdrOvagA99/I8iEjUWhcT0kc0V2ccs",
	"xovkepWDM/FKbNRP1JrvI6GQQecMGt25hXoKGGZzzhfT+1/k2lMoQr2OtG95Jr4SQwNXYtm2H9u79q4Y",
	"iIkCJg8dfG7v2iM0kFCH+raHTcvhvhwupP96eBg9BTUJGScxxbrxwUZmsjrOpR6zI1cNMr3nIuan18Gc",
	"o5sNBncKlUtKLJvwmIixjH900MFjUXBfRe0w6H4s4Mv290w8QwMBp2K1jxEaKGs8k4NOOVNz+yt/JbNn",
	"t75uYBL9Gt1cEDYGpVB9Ykqjt1XI4U2K2eXPqJrqx5d4qgWO1dkHGT2OQmulSM1hc5B5hA4G1dQdDrcQ",
	"SqytI0oTphY9rK/r7wAAAP//9Cn8pfsCAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

