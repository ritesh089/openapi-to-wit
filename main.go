package main

import (
	"log"
)

const OPENAPI_SPEC_PATH = "api/schemas/paths/openapi.yaml"
const WIT_PATH = "wit/world.wit"

func main() {
	// Load the OpenAPI schema from file path "api/schemas/paths/openapi.yaml"
	schemaPath := OPENAPI_SPEC_PATH
	openAPISchema, err := LoadOpenAPISchema(schemaPath)
	if err != nil {
		log.Fatalf("Error loading OpenAPI schema: %v", err)
	}

	// Extract the "Person" schema from the OpenAPI spec
	personSchema, exists := openAPISchema.Components.Schemas["Person"]
	if !exists {
		log.Fatalf("Person schema not found in OpenAPI components")
	}

	// Generate Go struct from OpenAPI schema for "Person"
	personStruct := GenerateGoStruct(personSchema, "Person")

	// Store complex types in a map for easy lookup
	complexTypes := map[string]ComplexType{
		"Person": personStruct,
	}

	// Define a WIT module for the OpenAPI-defined Person operations
	exampleModule := WITModule{
		ModuleName: "person_module",
		Functions: []FunctionSignature{
			{
				Name: "createPerson",
				Parameters: []Parameter{
					{Name: "name", Type: WITTypeString},
					{Name: "age", Type: WITTypeI32},
				},
				ReturnType: WITTypeStruct,
				ReturnRef:  &personStruct, // Return complex Person struct
			},
			{
				Name:       "getPerson",
				Parameters: []Parameter{}, // No input parameters
				ReturnType: WITTypeStruct,
				ReturnRef:  &personStruct, // Return complex Person struct
			},
		},
	}

	// Generate the WIT representation
	wit := GenerateWIT(exampleModule, complexTypes)

	// Write the WIT output to file "wit/world.wit"
	witFilePath := WIT_PATH
	err = WriteWITToFile(witFilePath, wit)
	if err != nil {
		log.Fatalf("Error writing WIT to file: %v", err)
	}
}
