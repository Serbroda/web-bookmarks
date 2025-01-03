// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// BookmarkDto defines model for BookmarkDto.
type BookmarkDto struct {
	CollectionId Id      `json:"collectionId"`
	Description  *string `json:"description,omitempty"`
	Id           Id      `json:"id"`
	Title        *string `json:"title,omitempty"`
	Url          string  `json:"url"`
}

// BookmarkDtoList defines model for BookmarkDtoList.
type BookmarkDtoList = []BookmarkDto

// CollectionDto defines model for CollectionDto.
type CollectionDto struct {
	Description *string `json:"description,omitempty"`
	Id          Id      `json:"id"`
	Name        string  `json:"name"`
	ParentId    *Id     `json:"parentId,omitempty"`
	SpaceId     Id      `json:"spaceId"`
}

// CollectionDtoList defines model for CollectionDtoList.
type CollectionDtoList = []CollectionDto

// CreateBookmarkDto defines model for CreateBookmarkDto.
type CreateBookmarkDto struct {
	CollectionId Id      `json:"collectionId"`
	Description  *string `json:"description,omitempty"`
	Title        *string `json:"title,omitempty"`
	Url          string  `json:"url"`
}

// CreateCollectionDto defines model for CreateCollectionDto.
type CreateCollectionDto struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
	ParentId    *Id     `json:"parentId,omitempty"`
	SpaceId     Id      `json:"spaceId"`
}

// CreateSpaceDto defines model for CreateSpaceDto.
type CreateSpaceDto struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
}

// Error defines model for Error.
type Error struct {
	Message *string `json:"message,omitempty"`
}

// Id defines model for Id.
type Id = string

// SpaceDto defines model for SpaceDto.
type SpaceDto struct {
	Description *string `json:"description,omitempty"`
	Id          Id      `json:"id"`
	Name        string  `json:"name"`
}

// SpaceDtoList defines model for SpaceDtoList.
type SpaceDtoList = []SpaceDto

// UpdateBookmarkDto defines model for UpdateBookmarkDto.
type UpdateBookmarkDto struct {
	CollectionId Id      `json:"collectionId"`
	Description  *string `json:"description,omitempty"`
	Title        *string `json:"title,omitempty"`
	Url          string  `json:"url"`
}

// UpdateCollectionDto defines model for UpdateCollectionDto.
type UpdateCollectionDto struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
	ParentId    *Id     `json:"parentId,omitempty"`
	SpaceId     Id      `json:"spaceId"`
}

// UpdateSpaceDto defines model for UpdateSpaceDto.
type UpdateSpaceDto struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
}

// BadRequest defines model for BadRequest.
type BadRequest = Error

// NotFound defines model for NotFound.
type NotFound = Error

// CreateBookmarkJSONRequestBody defines body for CreateBookmark for application/json ContentType.
type CreateBookmarkJSONRequestBody = CreateBookmarkDto

// UpdateBookmarkJSONRequestBody defines body for UpdateBookmark for application/json ContentType.
type UpdateBookmarkJSONRequestBody = UpdateBookmarkDto

// CreateCollectionJSONRequestBody defines body for CreateCollection for application/json ContentType.
type CreateCollectionJSONRequestBody = CreateCollectionDto

// UpdateCollectionJSONRequestBody defines body for UpdateCollection for application/json ContentType.
type UpdateCollectionJSONRequestBody = UpdateCollectionDto

// CreateSpaceJSONRequestBody defines body for CreateSpace for application/json ContentType.
type CreateSpaceJSONRequestBody = CreateSpaceDto

// UpdateSpaceJSONRequestBody defines body for UpdateSpace for application/json ContentType.
type UpdateSpaceJSONRequestBody = UpdateSpaceDto

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List all bookmarks
	// (GET /bookmarks)
	ListBookmarks(ctx echo.Context) error
	// Create a new bookmark
	// (POST /bookmarks)
	CreateBookmark(ctx echo.Context) error
	// Deletes an existing bookmark
	// (DELETE /bookmarks/{bookmarkId})
	DeleteBookmark(ctx echo.Context, bookmarkId Id) error
	// Updates an existing bookmark
	// (PUT /bookmarks/{bookmarkId})
	UpdateBookmark(ctx echo.Context, bookmarkId Id) error
	// List all collections
	// (GET /collections)
	ListCollections(ctx echo.Context) error
	// Create a new collection
	// (POST /collections)
	CreateCollection(ctx echo.Context) error
	// Delete an existing collection
	// (DELETE /collections/{collectionId})
	DeleteCollection(ctx echo.Context, collectionId Id) error
	// Update an existing collection
	// (PUT /collections/{collectionId})
	UpdateCollection(ctx echo.Context, collectionId Id) error
	// List all spaces
	// (GET /spaces)
	ListSpaces(ctx echo.Context) error
	// Create a new space
	// (POST /spaces)
	CreateSpace(ctx echo.Context) error
	// Delete an existing space
	// (DELETE /spaces/{spaceId})
	DeleteSpace(ctx echo.Context, spaceId Id) error
	// Update an existing space
	// (PUT /spaces/{spaceId})
	UpdateSpace(ctx echo.Context, spaceId Id) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListBookmarks converts echo context to params.
func (w *ServerInterfaceWrapper) ListBookmarks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ListBookmarks(ctx)
	return err
}

// CreateBookmark converts echo context to params.
func (w *ServerInterfaceWrapper) CreateBookmark(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateBookmark(ctx)
	return err
}

// DeleteBookmark converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteBookmark(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "bookmarkId" -------------
	var bookmarkId Id

	err = runtime.BindStyledParameterWithLocation("simple", false, "bookmarkId", runtime.ParamLocationPath, ctx.Param("bookmarkId"), &bookmarkId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter bookmarkId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteBookmark(ctx, bookmarkId)
	return err
}

// UpdateBookmark converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateBookmark(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "bookmarkId" -------------
	var bookmarkId Id

	err = runtime.BindStyledParameterWithLocation("simple", false, "bookmarkId", runtime.ParamLocationPath, ctx.Param("bookmarkId"), &bookmarkId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter bookmarkId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateBookmark(ctx, bookmarkId)
	return err
}

// ListCollections converts echo context to params.
func (w *ServerInterfaceWrapper) ListCollections(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ListCollections(ctx)
	return err
}

// CreateCollection converts echo context to params.
func (w *ServerInterfaceWrapper) CreateCollection(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateCollection(ctx)
	return err
}

// DeleteCollection converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteCollection(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "collectionId" -------------
	var collectionId Id

	err = runtime.BindStyledParameterWithLocation("simple", false, "collectionId", runtime.ParamLocationPath, ctx.Param("collectionId"), &collectionId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter collectionId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteCollection(ctx, collectionId)
	return err
}

// UpdateCollection converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateCollection(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "collectionId" -------------
	var collectionId Id

	err = runtime.BindStyledParameterWithLocation("simple", false, "collectionId", runtime.ParamLocationPath, ctx.Param("collectionId"), &collectionId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter collectionId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateCollection(ctx, collectionId)
	return err
}

// ListSpaces converts echo context to params.
func (w *ServerInterfaceWrapper) ListSpaces(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ListSpaces(ctx)
	return err
}

// CreateSpace converts echo context to params.
func (w *ServerInterfaceWrapper) CreateSpace(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateSpace(ctx)
	return err
}

// DeleteSpace converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteSpace(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "spaceId" -------------
	var spaceId Id

	err = runtime.BindStyledParameterWithLocation("simple", false, "spaceId", runtime.ParamLocationPath, ctx.Param("spaceId"), &spaceId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter spaceId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteSpace(ctx, spaceId)
	return err
}

// UpdateSpace converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateSpace(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "spaceId" -------------
	var spaceId Id

	err = runtime.BindStyledParameterWithLocation("simple", false, "spaceId", runtime.ParamLocationPath, ctx.Param("spaceId"), &spaceId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter spaceId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateSpace(ctx, spaceId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/bookmarks", wrapper.ListBookmarks)
	router.POST(baseURL+"/bookmarks", wrapper.CreateBookmark)
	router.DELETE(baseURL+"/bookmarks/:bookmarkId", wrapper.DeleteBookmark)
	router.PUT(baseURL+"/bookmarks/:bookmarkId", wrapper.UpdateBookmark)
	router.GET(baseURL+"/collections", wrapper.ListCollections)
	router.POST(baseURL+"/collections", wrapper.CreateCollection)
	router.DELETE(baseURL+"/collections/:collectionId", wrapper.DeleteCollection)
	router.PUT(baseURL+"/collections/:collectionId", wrapper.UpdateCollection)
	router.GET(baseURL+"/spaces", wrapper.ListSpaces)
	router.POST(baseURL+"/spaces", wrapper.CreateSpace)
	router.DELETE(baseURL+"/spaces/:spaceId", wrapper.DeleteSpace)
	router.PUT(baseURL+"/spaces/:spaceId", wrapper.UpdateSpace)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xZS2/jNhD+KwLboxApzZ5820cXCFq0RdNFD0UOtDSxuSuRXJLK1jX03wuSetGiLXlt",
	"yUGSmyWR8/i+mfkoeYsSlnNGgSqJFlskQHJGJZiLdzj9E74WIJW+ShhVQM1PzHlGEqwIo9Fnyai+J5M1",
	"5Fj/+lHAA1qgH6LWdGSfyuhnIZhAZVmGKAWZCMK1EbTQvoLaWRmij0wsSZqCMe2ubB+VIfqNqY+soOn0",
	"Af61hkBySMgDgTQQIFkhEgi+YRlQpoIHE0UZok8UF2rNBPkP0n7wzlPtpPJr4GbsS47Flw+K6UsuGAeh",
	"iOUiYVkGibZymw7lcGsCcRxvkdpwQAsklSB0pZ+TkXYUURl4LRQi89wvQyTga0GEzv8f7SZ0o7cb78N6",
	"I1t+hsSw3kHgV2KrjijI5VCgXeTKxi4WAm/09fvGuxfacwFFce7HiWMBVI3lTXKcwLjFPqjr7VU8Ppgd",
	"PI4C2kXSB7UArGDmSj6xQkcWp03txFq6eI2MKQ+T6J1eeM4cdwLZ695O4J7XHKTEqz2Ge0YsNr3Qvj+p",
	"E4eAr1P3IlBHeVRrNql5uvITT59rV9rUXkBX2kQv1JV6GaEPxm3FayPWwd+wfMt58PaPWxSiRxDSHnWu",
	"r+KrWHtmHCjmBC3QzVV8dYM0rGptgo6WlRFztQJT7jorXFci0k3wrlkVuifUn+L4bCe/3cOH5wz4+y86",
	"nzfx9T5bTXCRe9DTm26GN3WOtrqcijzHYlNhEOAsC5YdIBReSU1ZfQ/d64Jl0oOhq8rIUg4a13RzNvz6",
	"0l+61aVEAWWPwOspCPSRZ+OryIiHyei8/VyEdBtvgAMK3xri/byXYaeTom398zYt7WTIQEG/Kj6Y+52q",
	"4FjgHBQIbX+LiEZNt2o9lhaotYx2mQ1HsmSm372/j13C7ookASlnQ1/veDO8o3nldOmyaMoA0wD+JVIR",
	"uhpgLUS88DSrK9bz03L+ydA/foyaDPFck8EptKc5GU6rTUvAMbWpJ0p7Ajuszu876yYksf/W+hQVOnHA",
	"qJFt7w6qdJvmpDq98yY/r1J7nD8jrU66DHorYKe7om33ZWeEZjs1MiwPO69Sr7rd121nNA4TeFC6L8vO",
	"VPL9HQMjnm9gvBAJP7JM9ZwxnxYOC/idXTIhf84Xraco27KGoMbR3BgUa5PXpDrdftabV6Jdv89InWVF",
	"2S7Rba9E2+pz3AglrgtgeMy3n/he9XdAf/dRdFB1L8HEVFp7XNPHszT9S1TYvbNCbwTxWBeZ+UMDRZiT",
	"6PEalffl/wEAAP//wEvRhVchAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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