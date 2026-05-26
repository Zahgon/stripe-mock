package coercer

import (
	"regexp"

	"github.com/stripe/stripe-mock/spec"
)

// CoerceParams coerces the types of certain parameters according to typing
// information from their corresponding JSON schema. This is useful because an
// input format like form-encoding doesn't support anything but strings, and
// we'd like to work with a slightly wider variety of types like booleans and
// integers.
func CoerceParams(schema *spec.Schema, data map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// coerceSubSchema coerces a sub-param value according to an arbitrary JSON sub-schema.
// It is named with "sub-schema" because it can be array, primitive, or object, unlike the main
// `CoerceParams` only expecting object type with properties. This is also used in coercing each
// sub-schema of anyOf or array.
func coerceSubSchema(val interface{}, subSchema *spec.Schema) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil

	// Non-object schemas are anyOf, array, and primitive schemas.
	// Implicitly treats actual object schema with empty properties as non-object.
}

// `object` schema with properties

// unwrapping sub-schemas, and coerce contents in the map

//
// ---
//

// Various identifiers for types in JSON schema.
const (
	arrayType   = "array"
	booleanType = "boolean"
	integerType = "integer"
	numberType  = "number"
	objectType  = "object"
	stringType  = "string"
)

// maxSliceSize defines a somewhat arbitrary maximum size on an incoming
// integer-indexed map that we're willing to parse so that we don't run out of
// memory trying to allocate a slice.
const maxSliceSize = 1000

// numberPattern simply checks to see if an input string looks like a number.
var numberPattern = regexp.MustCompile(`\A\d+\z`)

// coercePrimitiveType tries to coerce a primitive type (e.g. bool, int, etc.)
// from the given generic interface{} value. On success it returns a coerced
// value with a boolean true. On failure (say the value wasn't a type that
// could be coerced) it returns nil and a boolean false.
func coercePrimitiveType(val interface{}, primitiveType string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// coerceNonObjectSchema tries to coerce a non-object schema given generic interface{} value.
//
// It's similar to coercePrimitiveType above (and indeed calls into it), but
// also handles array and anyOf schema (supporting a number of different primitive types)
func coerceNonObjectSchema(val interface{}, schema *spec.Schema) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// assuming enum value isn't numeric string. when given anyOf schema with enum and
// number, the numeric string won't falsely be taken as enum and miss its coercion

// underspecified array of primitive

// isSchemaPrimitiveType checks whether the given schema is a coercable
// primitive type (as opposed to an object or array).
//
// The conditional ladder in this function should be *identical* to the one in
// coercePrimitiveType (i.e., if support is added for a new type, it needs to
// be added in both places).
func isSchemaPrimitiveType(schema *spec.Schema) bool { _ = "STUB: not implemented"; return false }

// parseIntegerIndexedMap tries to parse a map that has all integer-indexed
// keys (e.g. { "0": ..., "1": "...", "2": "..." }) as a slice. We only try to
// do this when we know that the target schema requires an array.
func parseIntegerIndexedMap(valMap map[string]interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Already checked error above
