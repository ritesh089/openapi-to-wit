package openapi

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"openapi-to-wit/models"

	"gopkg.in/yaml.v2"
)

// LoadOpenAPISchema loads the OpenAPI spec from a YAML file
func LoadOpenAPISchema(filePath string) (*models.OpenAPISpec, error) {
	fileData, err := ioutil.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to read OpenAPI spec file: %w", err)
	}

	var spec models.OpenAPISpec
	err = yaml.Unmarshal(fileData, &spec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal OpenAPI spec: %w", err)
	}

	return &spec, nil
}
