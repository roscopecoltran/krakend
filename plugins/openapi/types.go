/*
	Package golang_openapi provides type definitions for (un)marshaling OpenAPI (formerly Swagger) specifications.
*/
package openapi

import (
	"time"
	// "github.com/roscopecoltran/krakend/logging"
)

type OpenAPI struct {
	// SwaggerVersion, as of this writing, is usually "2.0"
	SwaggerVersion string `json:"swagger,omitempty" yaml:"swagger,omitempty" mapstructure:"swagger,omitempty"`
	// Info is a section defining basic information about the API
	Info Info `json:"info,omitempty" yaml:"info,omitempty" mapstructure:"info,omitempty"`
	// Host is the FQDN of the server for API requests
	Host string `json:"host,omitempty" yaml:"host,omitempty" mapstructure:"host,omitempty"`
	// Schemes is a list of accepted schemes for API requests (http, https)
	Schemes []string `json:"schemes,omitempty" yaml:"schemes,omitempty" mapstructure:"schemes,omitempty"`
	// BasePath is a path to append to the Host in all API requests, if any
	BasePath string `json:"basePath,omitempty" yaml:"basePath,omitempty" mapstructure:"basePath,omitempty"`
	// Produces is a list of generated mime types for responses
	Produces []string `json:"produces,omitempty" yaml:"produces,omitempty" mapstructure:"produces,omitempty"`
	// Paths is a map of string paths with Path specifications for API requests
	Paths map[string]map[string]Request `json:"paths,omitempty" yaml:"paths,omitempty" mapstructure:"paths,omitempty"`
	// Definitions define type definitions for use in requests and responses
	Definitions map[string]Definition `json:"definitions,omitempty" yaml:"definitions,omitempty" mapstructure:"definitions,omitempty"`
	// Security indicates the required security definition for operations performed
	Security []map[string][]string `json:"security,omitempty" yaml:"security,omitempty" mapstructure:"security,omitempty"`
	// SecurityDefinitions specifies security requirements
	SecurityDefinitions map[string]SecurityDefinition `json:"securityDefinitions,omitempty" yaml:"securityDefinitions,omitempty" mapstructure:"securityDefinitions,omitempty"`
	// number of concurrent calls this endpoint must send to the backends
	ConcurrentCalls int `json:"concurrent_calls" yaml:"concurrent_calls" mapstructure:"concurrent_calls"`
	// timeout of this endpoint
	Timeout time.Duration `json:"timeout" yaml:"timeout" mapstructure:"timeout"`
}

// Info is an OpenAPI section for basic information about the API
type Info struct {
	// Title is the short name of the API
	Title string `json:"title,omitempty" yaml:"title,omitempty" mapstructure:"title,omitempty"`
	// Description is a description of the API
	Description string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`
	// Version is the API's current version
	Version string `json:"version,omitempty" yaml:"version,omitempty" mapstructure:"version,omitempty"`
}

// Request is a specification of request parameters and documentation for an HTTP verb and its request
type Request struct {
	// Summary is a title for the HTTP request
	Summary string `json:"summary,omitempty" yaml:"summary,omitempty" mapstructure:"summary,omitempty"`
	// Description is a description of the HTTP request
	Description string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`
	// Parameters is a list of parameters for the HTTP request
	Parameters []Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty" mapstructure:"parameters,omitempty"`
	// Tags is an array of tags for the request
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Responses houses possible responses to a request
	Responses map[string]Response `json:"responses,omitempty" yaml:"responses,omitempty" mapstructure:"responses,omitempty"`
}

// Parameter is an HTTP request parameter
type Parameter struct {
	// Name is a name for the parameter
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`
	// In defines where the parameter is specified. One of "query", "header", "path", "formData", or "body"
	In string `json:"in,omitempty" yaml:"in,omitempty" mapstructure:"in,omitempty"`
	// Description is a text description of the parameter
	Description string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`
	// Required indicates whether the parameter is required or not
	Required bool `json:"required,omitempty" yaml:"required,omitempty" mapstructure:"required,omitempty"`
	// Type defines the type of data being inputted and is not needed if Query=="body". The value MUST
	// be one of "string", "number", "integer", "boolean", "array" or "file". If type is "file", the consumes
	// MUST be either "multipart/form-data", " application/x-www-form-urlencoded" or both and the parameter
	// MUST be in "formData".
	Type string `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	// Format further specifies the format of the data in the parameter and is not needed if Query=="body".
	// See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#dataTypeFormat
	Format string `json:"format,omitempty" yaml:"format,omitempty" mapstructure:"format,omitempty"`
	// Schema declares the schema for body parameters
	Schema *Schema `json:"schema,omitempty" yaml:"schema,omitempty" mapstructure:"schema,omitempty"`
}

// Response is an individual response for an HTTP request
type Response struct {
	// Description is a short description of the response
	Description string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`
	// Schema is the Schema for the returned response
	Schema Schema `json:"schema,omitempty" yaml:"schema,omitempty" mapstructure:"schema,omitempty"`
}

// Schema is a response schema definition
type Schema struct {
	// Type is the type of response returned
	Type string `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	// Items is an item type reference for the response
	Items *ItemRef `json:"items,omitempty" yaml:"items,omitempty" mapstructure:"items,omitempty"`
	// Ref is a reference (if the type is not an array)
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty" mapstructure:"$ref,omitempty"`
}

// ItemRef is a reference to the type of item used in a request or a response
type ItemRef struct {
	// Ref indicates a reference to a definition (if any)
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty" mapstructure:"$ref,omitempty"`
	// Type indicates the type of items for an array
	Type string `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
}

// Definition is a type definition for an HTTP request or response
type Definition struct {
	// Type is the type of data being used (usually "object", "string", "integer", "array")
	Type string `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	// Properties is a list of properties of an object
	Properties map[string]Property `json:"properties,omitempty" yaml:"properties,omitempty" mapstructure:"properties,omitempty"`
	// Required is an array of strings representing the names of the required parameters
	Required []string `json:"required,omitempty" yaml:"required,omitempty" mapstructure:"required,omitempty"`
}

// Property is a definition of a property of a Definition
type Property struct {
	// Type is a data type. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types
	Type string `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	// Format is the data type format for an Property
	Format string `json:"format,omitempty" yaml:"format,omitempty" mapstructure:"format,omitempty"`
	// Description is a description of the property
	Description string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`
	// Items is an option ItemRef for the items in the Property if the Property references another item type
	Items *ItemRef `json:"items,omitempty" yaml:"items,omitempty" mapstructure:"items,omitempty"`
	// Ref is a reference for the type of object if the Type=="object"
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty" mapstructure:"$ref,omitempty"`
}

type SecurityDefinition struct {
	// Type determines the type of security used
	Type string `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	// Name is the name of the parameter used for security enforcement
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`
	// In indicates one of "query" or "header" depending on where the security is verfied in the request
	In string `json:"in,omitempty" yaml:"in,omitempty" mapstructure:"in,omitempty"`
}

// Definer is an interface for objects that can define Definition specifications
type Definer interface {
	OpenAPIDefinitions() map[string]Definition
}
