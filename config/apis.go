package config

/*
	Package golang_openapi provides type definitions for (un)marshaling OpenAPI (formerly Swagger) specifications.
*/

import (
	"github.com/satori/go.uuid"
	"time"
	// "github.com/roscopecoltran/krakend/logging"
)

type APIConfig struct {
	UUID        uuid.UUID     `gorm:"-" storm:"uuid" storm:"uuid,omitempty" json:"uuid,omitempty" yaml:"uuid,omitempty" toml:"uuid,omitempty"`
	Disabled    bool          `gorm:"column:disabled" default:"false" json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty"`
	Specs       []SpecsConfig `gorm:"column:specs" json:"specs,omitempty" yaml:"specs,omitempty" toml:"specs,omitempty"`
	Credentials string        `gorm:"column:credentials" json:"credentials,omitempty" yaml:"credentials,omitempty" toml:"credentials,omitempty"`
}

type OpenAPIConfig struct {
	// SwaggerVersion, as of this writing, is usually "2.0"
	SwaggerVersion string `json:"swagger,omitempty" yaml:"swagger,omitempty" mapstructure:"swagger" toml:"swagger,omitempty"`
	// Info is a section defining basic information about the API
	Info ApiInfo `json:"info,omitempty" yaml:"info,omitempty" mapstructure:"info" toml:"info,omitempty"`
	// Host is the FQDN of the server for API requests
	Host string `json:"host,omitempty" yaml:"host,omitempty" mapstructure:"host" toml:"host,omitempty"`
	// Schemes is a list of accepted schemes for API requests (http, https)
	Schemes []string `json:"schemes" yaml:"schemes" mapstructure:"schemes" toml:"schemes"`
	// BasePath is a path to append to the Host in all API requests, if any
	BasePath string `json:"basePath,omitempty" yaml:"basePath,omitempty" mapstructure:"basePath" toml:"basePath,omitempty"`
	// Produces is a list of generated mime types for responses
	Produces []string `json:"produces" yaml:"produces" mapstructure:"produces" toml:"produces"`
	// Paths is a map of string paths with Path specifications for API requests
	Paths map[string]map[string]ApiRequest `json:"paths" yaml:"paths" mapstructure:"paths" toml:"paths"`
	// Definitions define type definitions for use in requests and responses
	Definitions map[string]ApiDefinition `json:"definitions" yaml:"definitions" mapstructure:"definitions" toml:"definitions"`
	// Security indicates the required security definition for operations performed
	Security []map[string][]string `json:"security" yaml:"security" mapstructure:"security" toml:"security"`
	// SecurityDefinitions specifies security requirements
	SecurityDefinitions map[string]ApiSecurityDefinition `json:"securityDefinitions" yaml:"securityDefinitions" mapstructure:"securityDefinitions" toml:"securityDefinitions"`
	// number of concurrent calls this endpoint must send to the backends
	ConcurrentCalls int `default:"1" json:"concurrent_calls,omitempty" yaml:"concurrent_calls,omitempty" mapstructure:"concurrent_calls" toml:"concurrent_calls,omitempty"`
	// timeout of this endpoint
	Timeout time.Duration `default:"2000" json:"timeout,omitempty" yaml:"timeout" mapstructure:"timeout" toml:"timeout,omitempty"`
}

// Info is an OpenAPI section for basic information about the API
type ApiInfo struct {
	// Title is the short name of the API
	Title string `json:"title,omitempty" yaml:"title,omitempty" mapstructure:"title" toml:"title,omitempty"`
	// Description is a description of the API
	Description string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description" toml:"description,omitempty"`
	// Version is the API's current version
	Version string `json:"version,omitempty" yaml:"version,omitempty" mapstructure:"version" toml:"version,omitempty"`
}

// Request is a specification of request parameters and documentation for an HTTP verb and its request
type ApiRequest struct {
	// Summary is a title for the HTTP request
	Summary string `json:"summary,omitempty" yaml:"summary,omitempty" mapstructure:"summary" toml:"summary,omitempty"`
	// Description is a description of the HTTP request
	Description string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description" toml:"description,omitempty"`
	// Parameters is a list of parameters for the HTTP request
	Parameters []ApiParameter `json:"parameters" yaml:"parameters" mapstructure:"parameters" toml:"parameters"`
	// Tags is an array of tags for the request
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty" mapstructure:"tags" toml:"tags,omitempty"`
	// Responses houses possible responses to a request
	Responses map[string]ApiResponse `json:"responses" yaml:"responses" mapstructure:"responses" toml:"responses"`
}

// Parameter is an HTTP request parameter
type ApiParameter struct {
	// Name is a name for the parameter
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name" toml:"name,omitempty"`
	// In defines where the parameter is specified. One of "query", "header", "path", "formData", or "body"
	In string `json:"in" yaml:"in" mapstructure:"in" toml:"in"`
	// Description is a text description of the parameter
	Description string `json:"description" yaml:"description" mapstructure:"description" toml:"description"`
	// Required indicates whether the parameter is required or not
	Required bool `json:"required,omitempty" yaml:"required,omitempty" mapstructure:"required,omitempty" toml:"required,omitempty"`
	// Type defines the type of data being inputted and is not needed if Query=="body". The value MUST
	// be one of "string", "number", "integer", "boolean", "array" or "file". If type is "file", the consumes
	// MUST be either "multipart/form-data", " application/x-www-form-urlencoded" or both and the parameter
	// MUST be in "formData".
	Type string `json:"type" yaml:"type" mapstructure:"type" toml:"type"`
	// Format further specifies the format of the data in the parameter and is not needed if Query=="body".
	// See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#dataTypeFormat
	Format string `json:"format" yaml:"format" mapstructure:"format" toml:"format"`
	// Schema declares the schema for body parameters
	Schema *ApiSchema `json:"schema" yaml:"schema" mapstructure:"schema" toml:"schema"`
}

// Response is an individual response for an HTTP request
type ApiResponse struct {
	// Description is a short description of the response
	Description string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description" toml:"description,omitempty"`
	// Schema is the Schema for the returned response
	Schema ApiSchema `json:"schema" yaml:"schema" mapstructure:"schema" toml:"schema"`
}

// Schema is a response schema definition
type ApiSchema struct {
	// Type is the type of response returned
	Type string `json:"type" yaml:"type" mapstructure:"type" toml:"type"`
	// Items is an item type reference for the response
	Items *ApiItemRef `json:"items" yaml:"items" mapstructure:"items" toml:"items"`
	// Ref is a reference (if the type is not an array)
	Ref string `json:"$ref" yaml:"$ref" mapstructure:"$ref" toml:"$ref"`
}

// ItemRef is a reference to the type of item used in a request or a response
type ApiItemRef struct {
	// Ref indicates a reference to a definition (if any)
	Ref string `json:"$ref" yaml:"$ref" mapstructure:"$ref" toml:"$ref"`
	// Type indicates the type of items for an array
	Type string `json:"type" yaml:"type" mapstructure:"type" toml:"type"`
}

// Definition is a type definition for an HTTP request or response
type ApiDefinition struct {
	// Type is the type of data being used (usually "object", "string", "integer", "array")
	Type string `json:"type" yaml:"type" mapstructure:"type" toml:"type"`
	// Properties is a list of properties of an object
	Properties map[string]ApiProperty `json:"properties" yaml:"properties" mapstructure:"properties" toml:"properties"`
	// Required is an array of strings representing the names of the required parameters
	Required []string `json:"required" yaml:"required" mapstructure:"required" toml:"required"`
}

// Property is a definition of a property of a Definition
type ApiProperty struct {
	// Type is a data type. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types
	Type string `json:"type" yaml:"type" mapstructure:"type" toml:"type"`
	// Format is the data type format for an Property
	Format string `json:"format" yaml:"format" mapstructure:"format" toml:"format"`
	// Description is a description of the property
	Description string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description" toml:"description,omitempty"`
	// Items is an option ItemRef for the items in the Property if the Property references another item type
	Items *ApiItemRef `json:"items" yaml:"items" mapstructure:"items" toml:"items"`
	// Ref is a reference for the type of object if the Type=="object"
	Ref string `json:"$ref" yaml:"$ref" mapstructure:"$ref" toml:"$ref"`
}

type ApiSecurityDefinition struct {
	// Type determines the type of security used
	Type string `json:"type" yaml:"type" mapstructure:"type" toml:"type"`
	// Name is the name of the parameter used for security enforcement
	Name string `json:"name" yaml:"name" mapstructure:"name" toml:"name"`
	// In indicates one of "query" or "header" depending on where the security is verfied in the request
	In string `json:"in" yaml:"in" mapstructure:"in" toml:"in"`
}

// Definer is an interface for objects that can define Definition specifications
type ApiDefiner interface {
	OpenAPIDefinitions() map[string]ApiDefinition
}
