package types

import (
	"reflect"
	"testing"
)

func TestBasicInstances(t *testing.T) {
	expected := []string{
		"bool",
		"string",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",
		"float32",
		"float64",
		"complex64",
		"complex128",
	}
	actual := instancesToKindStrings(BasicInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestComparableInstances(t *testing.T) {
	expected := []string{
		"bool",
		"string",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",
		"float32",
		"float64",
		"complex64",
		"complex128",
		"array",
		"struct",
		"chan",
		"ptr",
	}
	actual := instancesToKindStrings(ComparableInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestElementInstances(t *testing.T) {
	expected := []string{
		"bool",
		"string",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",
		"float32",
		"float64",
		"complex64",
		"complex128",
		"slice",
		"array",
		"map",
		"struct",
		"chan",
		"func",
		"ptr",
	}
	actual := instancesToKindStrings(ElementInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestKeyInstances(t *testing.T) {
	expected := []string{
		"bool",
		"string",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",
		"float32",
		"float64",
		"complex64",
		"complex128",
		"array",
		"struct",
		"chan",
		"ptr",
	}
	actual := instancesToKindStrings(KeyInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TesIntegerInstances(t *testing.T) {
	expected := []string{
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
	}
	actual := instancesToKindStrings(IntegerInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestOrderedInstances(t *testing.T) {
	expected := []string{
		"bool",
		"string",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",
		"float32",
		"float64",
	}
	actual := instancesToKindStrings(OrderedInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestNumbericInstances(t *testing.T) {
	expected := []string{
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"float32",
		"float64",
		"complex64",
		"complex128",
	}
	actual := instancesToKindStrings(NumbericInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestSliceInstances(t *testing.T) {
	expected := instancesToKindStrings(ElementInstances())
	actual := slicesToElementKindStrings(SliceInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestMapInstances(t *testing.T) {
	keyKindStrings := instancesToKindStrings(KeyInstances())
	elementKindStrings := instancesToKindStrings(ElementInstances())
	var expected []string
	actual := mapsToKeyElementKindStrings(MapInstances())

	// Populate expected
	for _, keyKindString := range keyKindStrings {
		for _, elementKindString := range elementKindStrings {
			expected = append(expected, keyKindString+"|"+elementKindString)
		}
	}

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

//region Duplicate code to prevent circular dependencies

func assertManyEquals[T comparable](t *testing.T, expected []T, actual []T) {
	for i := 0; i < len(expected); i++ {
		assertEquals(t, expected[i], actual[i])
	}
}

func assertEquals[T comparable](t *testing.T, expected T, actual T) {
	if actual != expected {
		t.Errorf("\nExpected:\t%#v\nActual:\t\t%#v\n", expected, actual)
	}
}

func instancesToKindStrings(instances []any) []string {
	var kindStrings []string

	for _, instance := range instances {
		kindStrings = append(kindStrings, kindString(instance))
	}

	return kindStrings
}

// Return the kind of the given type as a string.
func kindString(v any) string {
	if v == nil {
		return "<nil>"
	} else {
		return reflect.TypeOf(v).Kind().String()
	}
}

func slicesToElementKindStrings(slices []any) []string {
	var elementKindStrings []string

	for _, slice := range slices {
		elementKindStrings = append(
			elementKindStrings,
			reflect.
				TypeOf(slice).
				Elem().
				Kind().
				String(),
		)
	}

	return elementKindStrings
}

func mapsToKeyElementKindStrings(maps []any) []string {
	var keyElementKindStrings []string

	for _, m := range maps {
		keyElementKindStrings = append(
			keyElementKindStrings,
			reflect.TypeOf(m).Key().Kind().String()+"|"+
				reflect.TypeOf(m).Elem().Kind().String(),
		)
	}

	return keyElementKindStrings
}

//endregion
