// Package components provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package components

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"net/http"
	"strings"
)

// SchemaObject defines component schema for SchemaObject.
type SchemaObject struct {
	FirstName string `json:"firstName"`
	Role      string `json:"role"`
}

// ParameterObject defines component parameter for "ParameterObject"
type ParameterObject string

// ResponseObject defines component response for ResponseObject.
type ResponseObject struct {
	Field SchemaObject `json:"Field"`
}

// RequestBody defines component requestBodies for RequestBody.
type RequestBody struct {
	Field SchemaObject `json:"Field"`
}

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
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router runtime.EchoRouter, si ServerInterface) {

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9ySzU7rQAyFXyXyvcsoacUuO1ggseBHtDvUxTBxwdVkZmq7FVWVd0czSVsofQFYxXbs",
	"75yjZA82dDF49CrQ7CEaNh0qcu6eDt3j6wqtppENXtHn0sToyBql4OuVBJ9mYt+xM5nEISIrYSbdEro2",
	"Ff8Zl9DAv/qkWw9HUs/yc9Tq+xIY1xtibKF5GQmLNFb80Do6Q2eSuosIDYgy+Tfo02qLYpli8ggNmOKY",
	"D0pI57DeIO+OYih6E9rR8/NxsPt1yQeGxODlEGZo/siXvC6EuuiwOIQswklsdJFA34z8yLIkFn0wHV7Q",
	"LIGDu/TiLE3eKr+gFtku+WVIx44seskcn4Xg/m6e6Eqa8DBH0WKGvM3/5BZZhoTTalJN0mKI6E0kaOCq",
	"mlRT6D8DAAD//9SMdbW0AwAA",
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

