package types

import (
	"reflect"
	"testing"
)

var this = Types{}

func TestBasicInstances(t *testing.T) {
	expected := this.BasicKinds()
	actual := instancesToKindStrings(this.BasicInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestBasicKinds(t *testing.T) {
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
	actual := this.BasicKinds()

	assertManyEquals(t, expected, actual)
}

func TestComparableInstances(t *testing.T) {
	expected := this.ComparableKinds()
	actual := instancesToKindStrings(this.ComparableInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestComparableKinds(t *testing.T) {
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
	actual := this.ComparableKinds()

	assertManyEquals(t, expected, actual)
}

func TestElementInstances(t *testing.T) {
	expected := this.ElementKinds()
	actual := instancesToKindStrings(this.ElementInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestElementKinds(t *testing.T) {
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
	actual := this.ElementKinds()

	assertManyEquals(t, expected, actual)
}

func TestKeyInstances(t *testing.T) {
	expected := this.KeyKinds()
	actual := instancesToKindStrings(this.KeyInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestKeyKinds(t *testing.T) {
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
	actual := this.KeyKinds()

	assertManyEquals(t, expected, actual)
}

func TestIntegerInstances(t *testing.T) {
	expected := this.IntegerKinds()
	actual := instancesToKindStrings(this.IntegerInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestIntegerKinds(t *testing.T) {
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
	actual := this.IntegerKinds()

	assertManyEquals(t, expected, actual)
}

func TestOrderedInstances(t *testing.T) {
	expected := this.OrderedKinds()
	actual := instancesToKindStrings(this.OrderedInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestOrderedKinds(t *testing.T) {
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
	actual := this.OrderedKinds()

	assertManyEquals(t, expected, actual)
}

func TestNumbericInstances(t *testing.T) {
	expected := this.NumbericKinds()
	actual := instancesToKindStrings(this.NumbericInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestNumbericKinds(t *testing.T) {
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
	actual := this.NumbericKinds()

	assertManyEquals(t, expected, actual)
}

func TestSliceInstances(t *testing.T) {
	expected := instancesToKindStrings(this.ElementInstances())
	actual := slicesToElementKindStrings(this.SliceInstances())

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestMapInstances(t *testing.T) {
	keyKindStrings := instancesToKindStrings(this.KeyInstances())
	elementKindStrings := instancesToKindStrings(this.ElementInstances())
	var expected []string
	actual := mapsToKeyElementKindStrings(this.MapInstances())

	// Populate expected
	for _, keyKindString := range keyKindStrings {
		for _, elementKindString := range elementKindStrings {
			expected = append(expected, keyKindString+"|"+elementKindString)
		}
	}

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestAllTypeInstances(t *testing.T) {
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
		"func",
		"ptr",
		"<nil>",
	}
	filteredInstances := this.Instances([]string{}, []string{"slice", "map"})
	var actual []string

	// Populate actual
	for _, instance := range filteredInstances {
		actual = append(actual, kindString(instance))
	}

	assertManyEquals(t, expected, actual)
	assertEquals(t, len(expected), len(actual))
}

func TestKinds(t *testing.T) {
	expected := []string{
		"int",
		"int8",
		"int32",
		"int64",
		"uint8",
		"uint16",
		"uint64",
	}
	actual := this.Kinds(
		[]string{"integer"},
		[]string{"int16", "uint", "uint32"},
	)

	assertEquals(t, len(expected), len(actual))
	assertManyEquals(t, expected, actual)
}

func TestAllKinds(t *testing.T) {
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
		"<nil>",
	}
	actual := this.AllKinds()

	assertEquals(t, len(expected), len(actual))
	assertManyEquals(t, expected, actual)
}

//region Duplicate code to prevent circular dependencies

// Assert that the expected and actual slices given are equal.
func assertManyEquals[T comparable](t *testing.T, expected []T, actual []T) {
	for i := 0; i < len(expected); i++ {
		assertEquals(t, expected[i], actual[i])
	}
}

// Assert that the expected and actual values given are equal.
func assertEquals[T comparable](t *testing.T, expected T, actual T) {
	if actual != expected {
		t.Errorf("\nExpected:\t%#v\nActual:\t\t%#v\n", expected, actual)
	}
}

// Convert slice instances to element kind strings.
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

// Convert map instances to key|value kind strings.
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
