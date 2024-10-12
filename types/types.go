package types

import (
	"reflect"
)

//region Public

// Return instances of basic types.
func BasicInstances() []any {
	return typeInstancesByTag("basic")
}

// Return instances of comparable types.
func ComparableInstances() []any {
	return typeInstancesByTag("comparable")
}

// Return instances of element types.
func ElementInstances() []any {
	return typeInstancesByTag("element")
}

// Return instances of key types, identical to comparable types.
func KeyInstances() []any {
	return typeInstancesByTag("key")
}

// Return instances of integer types.
func IntegerInstances() []any {
	return typeInstancesByTag("integer")
}

// Return instances of ordered types.
func OrderedInstances() []any {
	return typeInstancesByTag("ordered")
}

// Return instances of numeric types.
func NumbericInstances() []any {
	return typeInstancesByTag("numeric")
}

// Return instances of slice types.
func SliceInstances() []any {
	elementInstances := ElementInstances()
	var sliceInstances []any

	for _, v := range elementInstances {
		sliceInstances = append(
			sliceInstances,
			reflect.MakeSlice(
				reflect.SliceOf(
					reflect.TypeOf(v),
				),
				1,
				1,
			).Interface(),
		)
	}

	return sliceInstances
}

// Return instances of map types.
func MapInstances() []any {
	keyInstances := KeyInstances()
	elementInstances := ElementInstances()
	var mapInstances []any

	for _, k := range keyInstances {
		for _, e := range elementInstances {
			mapInstances = append(mapInstances,
				reflect.MakeMap(
					reflect.MapOf(
						reflect.TypeOf(k),
						reflect.TypeOf(e),
					),
				).Interface(),
			)
		}
	}

	return mapInstances
}

//endregion

//region Private

// Return type instances for their tag.
func typeInstancesByTag(tag string) []any {
	var filtered []any

	for _, v := range typeInstances {
		for _, vTag := range v.Tags {
			if vTag == tag {
				filtered = append(filtered, v.Value)
			}
		}
	}

	return filtered
}

//endregion
