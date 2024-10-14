package types

import (
	"reflect"
)

//region Public

// Return all kinds.
func AllKinds() []string {
	return transform(
		typeInstances,
		func(instance TypeInstance) string {
			return kindString(instance.Value)
		})
}

// Return all type instances, excluding the kinds given.
func AllTypeInstances() []any {
	return transform(typeInstances,
		func(instance TypeInstance) any { return instance.Value })
}

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

// Return a filtered list of kinds inclusively by the tags given,
//
//	then exclusively by specific kinds given.
func Kinds(includeTags []string, exclude []string) []string {
	var kinds []string

	// Include kinds with the tags given. Else, include all.
	if isEmpty(includeTags) {
		kinds = AllKinds()
	} else {
		for _, includeTag := range includeTags {
			kinds = transform(
				filter(typeInstances, func(instance TypeInstance) bool {
					return has(instance.Tags, includeTag)
				}),
				func(instance TypeInstance) string {
					return kindString(instance.Value)
				},
			)
		}
	}

	// Exclude kinds given
	if isNotEmpty(exclude) {
		kinds = filter(kinds, func(kind string) bool {
			return !has(exclude, kind)
		})
	}

	return kinds
}

// Return a filtered list of instances inclusively by the tags given,
//
//	then exclusively by specific kinds given.
func Instances(includeTags []string, excludeKinds []string) []any {
	var instances []any

	// Include kinds with the tags given. Else, include all.
	if isEmpty(includeTags) {
		instances = AllTypeInstances()
	} else {
		for _, includeTag := range includeTags {
			instances = transform(
				filter(typeInstances, func(instance TypeInstance) bool {
					return has(instance.Tags, includeTag)
				}),
				func(instance TypeInstance) any {
					return instance.Value
				},
			)
		}
	}

	// Exclude kinds given
	if isNotEmpty(excludeKinds) {
		instances = filter(instances, func(instance any) bool {
			return !has(excludeKinds, kindString(instance))
		})
	}

	return instances
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
			reflect.
				MakeSlice(
					reflect.SliceOf(
						reflect.TypeOf(v),
					),
					0,
					0).
				Interface(),
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

// Return a slice of elements that match the criteria given.
func filter[F any](filterables []F, criteria func(F) bool) []F {
	var filtered []F

	for _, filterable := range filterables {
		if criteria(filterable) {
			filtered = append(filtered, filterable)
		}
	}

	return filtered
}

// Return true if the slice given is empty. Else, false.
func isEmpty(values []string) bool {
	return len(values) == 0
}

// Return true if the slice given is not empty. Else, false.
func isNotEmpty(values []string) bool {
	return len(values) > 0
}

// Apply the transformer function given to each value in the slice given.
func transform[V any, T any](values []V, transformer func(value V) T) []T {
	transformed := make([]T, 0, len(values))

	for _, value := range values {
		transformed = append(transformed, transformer(value))
	}

	return transformed
}

func has(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

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
