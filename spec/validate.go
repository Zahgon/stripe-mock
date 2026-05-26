package spec

import (
	"github.com/lestrrat-go/jsval"
)

// ComponentsForValidation is a collection of components for an OpenAPI
// specification that's been translated into equivalent JSON Schemas.
type ComponentsForValidation struct {
	root interface{}
}

// GetValidatorForOpenAPI3Schema gets a JSON Schema validator for a given
// OpenAPI specification and set of JSON Schema components.
func GetValidatorForOpenAPI3Schema(oaiSchema *Schema, components *ComponentsForValidation) (*jsval.JSVal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetComponentsForValidation translates a collection of components for an
// OpenAPI specification into equivalent JSON schemas.
//
// See also the comment on getJSONSchemaForOpenAPI3Schema.
func GetComponentsForValidation(components *Components) *ComponentsForValidation {
	_ = "STUB: not implemented"
	return nil
}

// Given an OpenAPI 3 schema represented as JSON, returns an equivalent JSON
// Schema represented as JSON. The important difference between OpenAPI 3
// schemas and JSON schemas is that OpenAPI 3 uses "nullable: true" to mark
// values that can be null, whereas JSON schemas represent "null" as a type just
// like "string".
//
// This converter only handles the options that are supported by the spec.Schema
// type, and it must be updated when new options are supported.
func getJSONSchemaForOpenAPI3Schema(oai *Schema) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Note that the major format that will be seen here, unix-time, will
// not be supported by the validator we're using -- we should probably
// see if we can support that properly.
