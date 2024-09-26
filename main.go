package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"openapi-to-wit/openapi"
	"openapi-to-wit/utils"
	"openapi-to-wit/wit"

	"gopkg.in/yaml.v2"
)

const OPENAPI_SPEC_PATH = "api/schemas/paths/openapi.yaml"
const WIT_PATH = "out/world.wit"

func main() {
	schemaPath := OPENAPI_SPEC_PATH
	openAPISpec, err := openapi.LoadOpenAPISchema(schemaPath)
	if err != nil {
		log.Fatalf("Error loading OpenAPI spec: %v", err)
	}

	// Generate the WIT representation for the OpenAPI spec
	witString := wit.GenerateWITFromOpenAPI(openAPISpec)

	// Write the WIT output to file
	witFilePath := WIT_PATH
	err = utils.WriteWITToFile(witFilePath, witString)
	if err != nil {
		log.Fatalf("Error writing WIT to file: %v", err)
	}

	fmt.Println("WIT successfully generated and written to", witFilePath)
}

// Simulating OpenAPI schema and generating Go struct
type OpenAPISchema struct {
	Components struct {
		Schemas map[string]Schema `yaml:"schemas"`
	} `yaml:"components"`
}

type Schema struct {
	Properties map[string]Property `yaml:"properties"` // Map of properties and their types (simple example)
	Required   []string            `yaml:"required"`   // Required fields
}

type Property struct {
	Type   string `yaml:"type"`
	Format string `yaml:"format,omitempty"`
}

// LoadOpenAPISchema loads the OpenAPI schema from a file path
func LoadOpenAPISchema(filePath string) (*OpenAPISchema, error) {
	fileData, err := ioutil.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to read OpenAPI schema file: %w", err)
	}

	var schema OpenAPISchema
	err = yaml.Unmarshal(fileData, &schema)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal OpenAPI schema: %w", err)
	}

	return &schema, nil
}

// GenerateGoStruct generates Go struct code from the OpenAPI schema
func GenerateGoStruct(schema Schema, structName string) ComplexType {
	var fields []Parameter
	for name, property := range schema.Properties {
		var goType WITType
		switch property.Type {
		case "string":
			goType = WITTypeString
		case "integer":
			goType = WITTypeI32
		// Add more OpenAPI to Go type mappings as needed
		default:
			goType = WITTypeString // Default to string for this example
		}
		fields = append(fields, Parameter{Name: name, Type: goType})
	}

	return ComplexType{
		Name:     structName,
		IsStruct: true,
		Fields:   fields,
	}
}
