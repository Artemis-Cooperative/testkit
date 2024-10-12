package types

import (
	"reflect"
)

//region Public

// Return instances of basic types.
func BasicInstances() []any {
	return typeInstancesByTag("basic")
}

// Return kinds of basic types.
func BasicKinds() []string {
	return instancesToKindStrings(BasicInstances())
}

// Return instances of comparable types.
func ComparableInstances() []any {
	return typeInstancesByTag("comparable")
}

// Return kinds of comparable types.
func ComparableKinds() []string {
	return instancesToKindStrings(ComparableInstances())
}

// Return instances of element types.
func ElementInstances() []any {
	return typeInstancesByTag("element")
}

// Return kinds of element types.
func ElementKinds() []string {
	return instancesToKindStrings(ElementInstances())
}

// Return instances of key types, identical to comparable types.
func KeyInstances() []any {
	return typeInstancesByTag("key")
}

// Return kinds of key types.
func KeyKinds() []string {
	return instancesToKindStrings(KeyInstances())
}

// Return instances of integer types.
func IntegerInstances() []any {
	return typeInstancesByTag("integer")
}

// Return kinds of integer types.
func IntegerKinds() []string {
	return instancesToKindStrings(IntegerInstances())
}

// Return instances of ordered types.
func OrderedInstances() []any {
	return typeInstancesByTag("ordered")
}

// Return kinds of ordered types.
func OrderedKinds() []string {
	return instancesToKindStrings(OrderedInstances())
}

// Return instances of numeric types.
func NumbericInstances() []any {
	return typeInstancesByTag("numeric")
}

// Return kinds of numeric types.
func NumbericKinds() []string {
	return instancesToKindStrings(NumbericInstances())
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

// Convert the instances given to kind strings.
func instancesToKindStrings(instances []any) []string {
	var kindStrings []string

	for _, instance := range instances {
		kindStrings = append(kindStrings, kindString(instance))
	}

	return kindStrings
}

// Return the kind of the value given as a string.
func kindString(v any) string {
	if v == nil {
		return "<nil>"
	} else {
		return reflect.TypeOf(v).Kind().String()
	}
}

//endregion
