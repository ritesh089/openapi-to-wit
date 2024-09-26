package wit

import (
	"fmt"
	"openapi-to-wit/models"
	"strings"
)

// GenerateWITFromOpenAPI generates WIT types and functions for any OpenAPI spec
func GenerateWITFromOpenAPI(spec *models.OpenAPISpec) string {
	var witBuilder strings.Builder

	// First, generate types for all schemas outside the world block
	witBuilder.WriteString(GenerateTypesForSchemas(spec.Components.Schemas))

	// Then, generate the world definition with exported functions inside it
	witBuilder.WriteString(fmt.Sprintf("world %s {\n", strings.ToLower(spec.Info.Title)))

	// Generate functions for each path in the OpenAPI spec
	witBuilder.WriteString(GenerateFunctionsForPaths(spec.Paths))

	// Close the world definition
	witBuilder.WriteString("}\n")

	return witBuilder.String()
}

// GenerateTypesForSchemas generates WIT record types for each schema in the OpenAPI spec
// These types will be placed outside the world block
func GenerateTypesForSchemas(schemas map[string]models.Schema) string {
	var witBuilder strings.Builder

	for name, schema := range schemas {
		witBuilder.WriteString(fmt.Sprintf("type %s = record {\n", name))
		for propertyName, property := range schema.Properties {
			witBuilder.WriteString(fmt.Sprintf("    %s: %s,\n", propertyName, MapOpenAPITypeToWIT(property.Type, property.Format)))
		}
		witBuilder.WriteString("}\n\n")
	}

	return witBuilder.String()
}

// GenerateFunctionsForPaths generates WIT functions for each OpenAPI path
// These functions will be placed inside the world block with "export" prepended
func GenerateFunctionsForPaths(paths map[string]models.PathItem) string {
	var witBuilder strings.Builder

	for path, item := range paths {
		for method, operation := range item.Operations {
			// Convert method and path to a function name
			functionName := fmt.Sprintf("%s%s", strings.Title(strings.ToLower(method)), pathToFunctionName(path))

			// Start function signature with "export"
			witBuilder.WriteString(fmt.Sprintf("    export %s: func(", functionName))

			// Handle request body parameters, if any
			if operation.RequestBody != nil && operation.RequestBody.Content != nil {
				for _, content := range operation.RequestBody.Content {
					for name, schema := range content.Schema.Properties {
						witBuilder.WriteString(fmt.Sprintf("%s: %s, ", name, MapOpenAPITypeToWIT(schema.Type, schema.Format)))
					}
				}
			}

			// Trim trailing comma if needed
			witBuilder.WriteString(")")

			// Handle response types, using WIT types
			if response, ok := operation.Responses["200"]; ok && response.Content != nil {
				for _, content := range response.Content {
					// Replace OpenAPI $ref with just the schema name
					refType := strings.TrimPrefix(content.Schema.Ref, "#/components/schemas/")
					witBuilder.WriteString(fmt.Sprintf(" -> %s", refType))
				}
			}

			// End function signature
			witBuilder.WriteString("\n")
		}
	}

	return witBuilder.String()
}

// Helper function to map OpenAPI types to WIT types
func MapOpenAPITypeToWIT(openAPIType, format string) string {
	switch openAPIType {
	case "string":
		return "string"
	case "integer":
		if format == "int32" {
			return "i32"
		} else if format == "int64" {
			return "i64"
		}
	case "boolean":
		return "bool"
	case "array":
		return "list"
	}
	return "string" // default to string if unknown
}

// Helper function to convert OpenAPI paths to valid WIT function names
func pathToFunctionName(path string) string {
	// Replace slashes with underscores and remove leading/trailing slashes
	return strings.ReplaceAll(strings.Trim(path, "/"), "/", "_")
}
