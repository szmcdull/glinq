package gmap

import (
	"fmt"
	"testing"
)

func TestShallowCopy(t *testing.T) {
	l := map[int]int{1: 2, 2: 4, 3: 6}
	result := ShallowCopy(l)
	if result[1] != 2 ||
		result[2] != 4 ||
		result[3] != 6 {

		t.Fail()
	}
}

func ExampleSortedKeys() {
	m := map[int]string{}
	m[1] = "John"
	m[5] = "Marry"
	m[9] = "Rose"

	// get the sorted keys
	keys := SortedKeys(m)

	// print the sorted keys
	fmt.Println(keys)

	// Output: [1 5 9]
}

func TestKeyValueOrder(t *testing.T) {
	m := map[int]string{}
	m[1] = "John"
	m[5] = "Marry"
	m[9] = "Rose"
	keys := Keys(m)
	values := Values(m)
	for i := range keys {
		if m[keys[i]] != values[i] {
			t.Fail()
		}
	}
}
