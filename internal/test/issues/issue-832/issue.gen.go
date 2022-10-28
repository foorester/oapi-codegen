// Package issue_832 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package issue_832

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Defines values for Document_Status.
const (
	Four  Document_Status = "four"
	One   Document_Status = "one"
	Three Document_Status = "three"
	Two   Document_Status = "two"
)

// Document defines model for Document.
type Document struct {
	Name            *string          `json:"name,omitempty"`
	Document_Status *Document_Status `json:"status,omitempty"`
}

// Document_Status defines model for Document.status.
type Document_Status string

// DocumentStatus defines model for DocumentStatus.
type DocumentStatus struct {
	Value *string `json:"value,omitempty"`
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7ySz0/rMAzH/5XK7x27tq/vljMIIYQ47AgIhdRbM9o4Styxaer/jpxuReOHxIlLkyb2",
	"9/tx7AMY6j05dBxBHSCaFnudthdkhh4dy94H8hjYYrpxukdZee8RFEQO1q1hzCGy5iGFoBt6UPdADiEH",
	"fiX5tgHlb0VDgMf8Q3oOu8WaFpP2bP60nCTHcY6n5w0aFrtT0HK2Pefc6m74CvSzlhxZtyIJbjCaYD1b",
	"cqDgVr9gFoeAGbeas4BmCNFuMROFmOmAWatd02GTTebd/sFJsZY7ccCd7n0nZW8xxEmzKqrinxRAHp32",
	"FhT8L6qihhy85jaxl6dEdYA1piaIuhas6wYUXE73V8iQQ8DoycWp7LqqZDHk+Ng+7X1nTcotN1EYTp2W",
	"3d+AK1Dwp3wfhfI4B+U8BOmJzp/m7kZOx3xmrX8AW/8G7Tw03zGP41sAAAD//zZo9q75AgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
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

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
