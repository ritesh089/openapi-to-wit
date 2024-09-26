package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

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
