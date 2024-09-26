package models

// OpenAPISpec represents the full OpenAPI specification
type OpenAPISpec struct {
	OpenAPI    string              `yaml:"openapi"`
	Info       Info                `yaml:"info"`
	Paths      map[string]PathItem `yaml:"paths"`
	Components Components          `yaml:"components"`
}

// Info represents metadata about the API
type Info struct {
	Title   string `yaml:"title"`
	Version string `yaml:"version"`
}

// PathItem represents the operations available on a single path
type PathItem struct {
	Operations map[string]Operation `yaml:",inline"` // Methods like GET, POST, etc.
}

// Operation represents an operation on a path (like GET, POST)
type Operation struct {
	Summary     string              `yaml:"summary"`
	OperationID string              `yaml:"operationId"`
	RequestBody *RequestBody        `yaml:"requestBody,omitempty"`
	Responses   map[string]Response `yaml:"responses"`
}

// RequestBody represents a request body in an operation
type RequestBody struct {
	Content map[string]MediaType `yaml:"content"`
}

// MediaType represents the media type of a request or response body
type MediaType struct {
	Schema Schema `yaml:"schema"`
}

// Response represents a response for an operation
type Response struct {
	Description string               `yaml:"description"`
	Content     map[string]MediaType `yaml:"content"`
}

// Components represent the reusable parts of an OpenAPI spec (e.g., schemas)
type Components struct {
	Schemas map[string]Schema `yaml:"schemas"`
}

// Schema represents the structure of data
type Schema struct {
	Type       string              `yaml:"type"`
	Format     string              `yaml:"format,omitempty"`
	Properties map[string]Property `yaml:"properties,omitempty"`
	Ref        string              `yaml:"$ref,omitempty"`
}

// Property represents a property in a schema
type Property struct {
	Type   string `yaml:"type"`
	Format string `yaml:"format,omitempty"`
}
