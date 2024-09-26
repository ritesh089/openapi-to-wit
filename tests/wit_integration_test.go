package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"openapi-to-wit/openapi"
	"openapi-to-wit/wit"
)

// TODO: Integration tests are failing because of order of functions in wit
// Helper function to load file contents
func loadFileContents(filePath string) (string, error) {
	bytes, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Test the simple OpenAPI spec to WIT generation
func TestSimpleOpenAPISpecToWIT(t *testing.T) {
	// Load the OpenAPI spec from the test file
	openAPIPath := "data/specs/openapi1.yaml"
	openAPISpec, err := openapi.LoadOpenAPISchema(openAPIPath)
	if err != nil {
		t.Fatalf("Failed to load OpenAPI spec: %v", err)
	}

	// Generate the WIT from the OpenAPI spec
	generatedWIT := wit.GenerateWITFromOpenAPI(openAPISpec)

	// Load the expected WIT output
	expectedWITPath := "data/wit/wit_expected_1.txt"
	expectedWIT, err := loadFileContents(expectedWITPath)
	if err != nil {
		t.Fatalf("Failed to load expected WIT output: %v", err)
	}

	// Compare the generated WIT with the expected WIT output
	trimmedGeneratedWIT := strings.TrimSpace(generatedWIT)
	trimmedExpectedWIT := strings.TrimSpace(expectedWIT)

	if trimmedGeneratedWIT != trimmedExpectedWIT {
		t.Errorf("Generated WIT does not match expected WIT:\nGenerated:\n%s\nExpected:\n%s", trimmedGeneratedWIT, trimmedExpectedWIT)
	}
}

// // Add more tests for complex OpenAPI specs as needed
// func TestComplexOpenAPISpecToWIT(t *testing.T) {
// 	// Similar steps for a complex OpenAPI spec
// 	openAPIPath := "../api/schemas/test_complex_openapi.yaml"
// 	openAPISpec, err := openapi.LoadOpenAPISchema(openAPIPath)
// 	if err != nil {
// 		t.Fatalf("Failed to load OpenAPI spec: %v", err)
// 	}

// 	// Generate the WIT from the OpenAPI spec
// 	generatedWIT := wit.GenerateWITFromOpenAPI(openAPISpec)

// 	// Load the expected WIT output
// 	expectedWITPath := "../api/schemas/expected_wit_complex.txt"
// 	expectedWIT, err := loadFileContents(expectedWITPath)
// 	if err != nil {
// 		t.Fatalf("Failed to load expected WIT output: %v", err)
// 	}

// 	// Compare the generated WIT with the expected WIT output
// 	trimmedGeneratedWIT := strings.TrimSpace(generatedWIT)
// 	trimmedExpectedWIT := strings.TrimSpace(expectedWIT)

// 	if trimmedGeneratedWIT != trimmedExpectedWIT {
// 		t.Errorf("Generated WIT does not match expected WIT:\nGenerated:\n%s\nExpected:\n%s", trimmedGeneratedWIT, trimmedExpectedWIT)
// 	}
// }
