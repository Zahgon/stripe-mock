// Package nestedtypeassembler takes a collection of form.Pair tuples and uses
// them to construct more complex data types using "Rack-style" conventions for
// arrays and maps with a few small Stripe-specific tweaks.
//
// When processing a request, data should first be parsed from the query, form,
// or multipart form, handed off to this package for assembly, then passed on
// the coercer to coerce string to other expected types.
package nestedtypeassembler

import (
	"github.com/stripe/stripe-mock/param/form"
)

//
// Public functions
//

// AssembleParams takes a collection of form.Pair tuples and translates them to
// parameter map that includes complex data types like arrays and other nested
// maps.
func AssembleParams(form form.Values) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//fmt.Printf("new params = %+v\n", pairParams)

//fmt.Printf("merge result = %+v\n\n", params)

//
// Private constants
//

const (
	keyTypeArray = iota
	keyTypeMap
	keyTypeRaw
)

//
// Private types
//

// keyPart is an interface for a struct that represents a "part" of a parameter
// key. For example, we might say that in `obj[]` the `[]` is an array part.
type keyPart interface {
	KeyType() int
	Content() string
}

// keyType represents an array part of a parameter key, like the `[]` in
// `obj[][foo]`.
type keyArray struct {
}

func (k *keyArray) KeyType() int { _ = "STUB: not implemented"; return 0 }

func (k *keyArray) Content() string { _ = "STUB: not implemented"; return "" }

// keyMap represents a map part of a parameter key, like the `[foo]` in
// `obj[][foo]`.
type keyMap struct {
	content string
}

func (k *keyMap) KeyType() int { _ = "STUB: not implemented"; return 0 }

func (k *keyMap) Content() string {
	_ = "STUB: not implemented"

	// keyRaw represents the raw name of a parameter key, like the `obj` in
	// `obj[][foo]`.
	return ""
}

type keyRaw struct {
	content string
}

func (k *keyRaw) KeyType() int { _ = "STUB: not implemented"; return 0 }

func (k *keyRaw) Content() string {
	_ = "STUB: not implemented"

	// Private functions
	return ""
}

func parseKey(key string) []keyPart { _ = "STUB: not implemented"; return nil }

// Fall through to inMap

func buildParamStructure(key string, parts []keyPart, value string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func maybeCollapseArrays(arr1, arr2 []interface{}) bool {
	_ = "STUB: not implemented"
	// fmt.Printf("maybe collapse arrays %+v %+v\n", arr1, arr2)
	return false
}

// arr1's merge candidate is its last element only

// If any of the keys in arr2's map are already in arr1's map, then don't
// merge we consider this a new map.

func mergeMapsRecursive(map1, map2 map[string]interface{}) { _ = "STUB: not implemented"; return }

// If not an array or map, or we couldn't reconcile types between the
// two maps, simply set the key in map1 to the value from map2.
