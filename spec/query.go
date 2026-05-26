package spec

// BuildQuerySchema builds a JSON schema that will be used to validate query
// parameters on the incoming request. Unlike request bodies, OpenAPI puts
// query parameters in a different, non-JSON schema part of an operation.
func BuildQuerySchema(operation *Operation) *Schema { _ = "STUB: not implemented"; return nil }
