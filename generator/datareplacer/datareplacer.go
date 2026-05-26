package datareplacer

import (
	"reflect"

	"github.com/stripe/stripe-mock/spec"
)

// DataReplacer takes a generated response and replaces values in it that share
// a name and type of parameters that were sent in with the request, as
// determined by the associated OpenAPI schema and the types of incoming
// values.
//
// This is designed to have the effect of making returned fixtures more
// realistic while also staying a simple heuristic that doesn't require very
// much maintenance.
type DataReplacer struct {
	Definitions map[string]*spec.Schema
	Schema      *spec.Schema
}

// ReplaceData projects data from the incoming request into response data as
// appropriate.
func (r *DataReplacer) ReplaceData(requestData map[string]interface{}, responseData map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Identical to the above except that we pass a schema as argument so that we
// can easily have the relevant one during recursion.
func (r *DataReplacer) replaceDataInternal(requestData map[string]interface{}, responseData map[string]interface{}, schema *spec.Schema) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Recursively call in to replace data, but only if the key is
// in both maps.
//
// A fairly obvious improvement here is if a key is in the
// request but not present in the response, then check the
// canonical schema to see if it's there. It might be an
// optional field that doesn't appear in the fixture, and if it
// was given to us with the request, we probably want to
// include it.

// In the non-map case, just set the respons key's value to
// what was in the request, but only if both values are the
// same type (this is to prevent problems where a field is set
// as an ID, but the response field is the hydrated object of
// that).
//
// While this will largely be "good enough", there's some
// obvious cases that aren't going to be handled correctly like
// index-based array updates (e.g.,
// `additional_owners[1][name]=...`). I'll have to iron out
// that rough edges later on.

func (r *DataReplacer) isSameType(schema *spec.Schema, requestValue interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

// Reflect in Go has the concept of a "zero Value" (not be confused with a
// type's zero value with a lowercase "v") and asking for Type on one will
// panic. I'm not exactly sure under what conditions these are generated,
// but they are occasionally, so here we hedge against them.
//
// https://github.com/stripe/stripe-mock/issues/75

// In the case of `anyOf`, allow replacement if any of the schema branches apply.

// Incoming value is not an array

// Allow the replacement if completely empty. In practice, this
// should never happen because you can't send an empty array via
// form data, but we'll cover the case anyway.

// Allow the replacement if the first item in the incoming slice is
// compatible with the array's `items` schema.

// Don't try to replace objects for now, the likelihood is that they're
// not compatible between request and response anyway.

// Unreachable because of `default` above

func (r *DataReplacer) maybeDereference(schema *spec.Schema, context string) (*spec.Schema, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// definitionFromJSONPointer extracts the name of a JSON schema definition from
// a JSON pointer, so "#/components/schemas/charge" would become just "charge".
// This is a simplified workaround to avoid bringing in JSON schema
// infrastructure because we can guarantee that the spec we're producing will
// take a certain shape. If this gets too hacky, it will be better to put a more
// legitimate JSON schema parser in place.
func definitionFromJSONPointer(pointer string) string { _ = "STUB: not implemented"; return "" }

func isIntegerKind(kind reflect.Kind) bool { _ = "STUB: not implemented"; return false }
