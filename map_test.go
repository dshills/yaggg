package main

import "testing"

type mymap map[string]int

func (m mymap) Keys() []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m mymap) Values() []int {
	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func TestMap(t *testing.T) {
	mm := make(mymap)
	mm["one"] = 1
	mm["two"] = 2
	mm["three"] = 3
	mm["four"] = 4
	mm["five"] = 5
	mm["six"] = 6

	k := mm.Keys()
	if len(k) != len(mm) {
		t.Errorf("Keys: expected len(k) == len(mm) got %v\n", len(k))
	}
	v := mm.Values()
	if len(v) != len(mm) {
		t.Errorf("Values: expected len(v) == len(mm) got %v\n", len(v))
	}
}
