package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

// Converts OpenAPI types to WIT types
func openAPITypeToWITType(openAPIType string) string {
	switch openAPIType {
	case "string":
		return "string"
	case "integer":
		return "u32"
	case "boolean":
		return "bool"
	case "array":
		return "list"
	default:
		return "string"
	}
}

// Generate WIT file content from OpenAPI specification
func generateWITFromOpenAPI(api *spec.Swagger) (string, error) {
	wit := "world api {\n"

	// Loop through each path in the OpenAPI spec
	for path, pathItem := range api.Paths.Paths {
		for method, operation := range getOperations(pathItem) {
			// Sanitize and format the path
			witFunctionName := sanitizePath(path)

			// Create WIT function signature
			wit += fmt.Sprintf("  import %s_%s: func(", method, witFunctionName)

			// Add parameters to WIT function
			for i, param := range operation.Parameters {
				witType := openAPITypeToWITType(param.Type)
				if i > 0 {
					wit += ", "
				}
				wit += fmt.Sprintf("%s: %s", param.Name, witType)
			}

			// Add return type (for now, assuming 200 response is relevant)
			if len(operation.Responses.StatusCodeResponses) > 0 {
				wit += ") -> ("
				response := operation.Responses.StatusCodeResponses[200]
				if response.Schema != nil {
					witType := openAPITypeToWITType(response.Schema.Type[0])
					wit += fmt.Sprintf("response: %s", witType)
				}
				wit += ")"
			} else {
				wit += ")"
			}

			// Close function definition
			wit += "\n"
		}
	}

	// Close the world
	wit += "}\n"
	return wit, nil
}

// Helper function to sanitize path strings for WIT naming conventions
func sanitizePath(path string) string {
	// Remove slashes and replace with underscores for function naming
	return filepath.Base(path)
}

// Get all operations from a given path item
func getOperations(pathItem spec.PathItem) map[string]*spec.Operation {
	operations := map[string]*spec.Operation{}
	if pathItem.Get != nil {
		operations["get"] = pathItem.Get
	}
	if pathItem.Post != nil {
		operations["post"] = pathItem.Post
	}
	if pathItem.Put != nil {
		operations["put"] = pathItem.Put
	}
	if pathItem.Delete != nil {
		operations["delete"] = pathItem.Delete
	}
	return operations
}

func main() {
	// Define paths
	openAPISpecPath := "api/schemas/paths/openapi.yaml"
	witOutputPath := "wit/world.wit"

	// Load the OpenAPI specification
	swaggerDoc, err := loads.Spec(openAPISpecPath)
	if err != nil {
		log.Fatalf("Error loading OpenAPI spec: %v", err)
	}

	// Generate WIT from OpenAPI
	witContent, err := generateWITFromOpenAPI(swaggerDoc.Spec())
	if err != nil {
		log.Fatalf("Error generating WIT content: %v", err)
	}

	// Ensure the output directory exists
	err = os.MkdirAll(filepath.Dir(witOutputPath), os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating directory for WIT output: %v", err)
	}

	// Write the WIT content to the specified file
	err = ioutil.WriteFile(witOutputPath, []byte(witContent), 0644)
	if err != nil {
		log.Fatalf("Error writing WIT file: %v", err)
	}

	fmt.Printf("WIT file successfully generated at %s\n", witOutputPath)
}

