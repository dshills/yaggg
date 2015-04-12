package main

import (
	"fmt"
	"sort"
	"testing"
)

type testStruct struct {
	s string
	i int
}

func (t testStruct) String() string {
	return fmt.Sprintf("{s:%v i:%v}", t.s, t.i)
}

func (t testStruct) Cmp(b testStruct) int {
	if t.i == b.i {
		return 0
	}
	if t.i > b.i {
		return 1
	}
	return -1
}

func (t testStruct) Clone() testStruct {
	c := t
	return c
}

func (t testStruct) Clear() testStruct {
	return testStruct{}
}

// A tsSlice is a slice of testStructs
type tsSlice []testStruct

// Len will satisfy the sort.Interface
func (e tsSlice) Len() int { return len(e) }

// Less will satisfy the sort.Interface
func (e tsSlice) Less(i, j int) bool {
	if e[i].Cmp(e[j]) > 0 {
		return false
	}
	return true
}

// Swap will satisfy the sort.Interface
func (e tsSlice) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

// Sort will sort tsSlice using the .Cmp function
func (e tsSlice) Sort() { sort.Sort(e) }

// SSearch will return the index of the first occurence of n for a sorted tsSlice
func (e tsSlice) SSearch(n testStruct) int {
	return sort.Search(len(e), func(i int) bool { return e[i].Cmp(n) != -1 })
}

// Search will return the index of the first occurence of n in tsSlice
func (e tsSlice) Search(n testStruct) int {
	for i, v := range e {
		if v.Cmp(n) == 0 {
			return i
		}
	}
	return len(e)
}

// Filter will return a new tsSlice of testStruct that match by calling the function fx
func (e tsSlice) Filter(fx func(n testStruct) bool) tsSlice {
	f := make(tsSlice, 0, len(e))
	for _, v := range e {
		if fx(v) {
			f = append(f, v)
		}
	}
	return f
}

// FilterNot will return a new tsSlice of testStruct that do not match by calling the function fx
func (e tsSlice) FilterNot(fx func(n testStruct) bool) tsSlice {
	f := make(tsSlice, 0, len(e))
	for _, v := range e {
		if fx(v) == false {
			f = append(f, v)
		}
	}
	return f
}

// Map will run function fx on all items in tsSlice
func (e tsSlice) Map(fx func(n testStruct) testStruct) {
	for i := range e {
		e[i] = fx(e[i])
	}
}

// Cut will remove items i through j-1
// The Clear function wil be called to insure memory is properly handled
func (e *tsSlice) Cut(i, j int) {
	z := *e // copy slice header
	copy(z[i:], z[j:])
	for k, n := len(z)-j+i, len(z); k < n; k++ {
		z[k].Clear()
	}
	z = z[:len(z)-j+i]
	*e = z
}

// Delete will remove item i
// The Clear function wil be called on the removed item to insure memory is properly handled
func (e *tsSlice) Delete(i int) {
	z := *e // copy the slice header
	end := len(z) - 1
	e.Swap(i, end)
	copy(z[i:], z[i+1:])
	z[end] = z[end].Clear()
	z = z[:end]
	*e = z
}

// Insert will place a new item at position i
func (e *tsSlice) Insert(n testStruct, i int) {
	z := *e // copy the slice header
	z = append(z, testStruct{})
	copy(z[i+1:], z[i:])
	z[i] = n
	*e = z
}

// Append will add a new item at the end of tsSlice
func (e *tsSlice) Append(n testStruct) {
	*e = append(*e, n)
}

// Prepend will add a new item at the beginning of tsSlice
func (e *tsSlice) Prepend(n testStruct) {
	e.Insert(n, 0)
}

// Clonet will return a copy of tsSlice calling the Clone function on each item
func (e tsSlice) Clone() tsSlice {
	f := make(tsSlice, len(e))
	for i := range e {
		f[i] = e[i].Clone()
	}
	return f
}

func TestBasic(t *testing.T) {
	ts := make(tsSlice, 0, 5)
	ts.Append(testStruct{s: "e", i: 5})
	ts.Append(testStruct{s: "d", i: 4})
	ts.Append(testStruct{s: "c", i: 3})
	ts.Append(testStruct{s: "b", i: 2})
	ts.Append(testStruct{s: "a", i: 1})
	if len(ts) != 5 {
		t.Errorf("Append: expected len(ts) == 5 got %v\n", len(ts))
		t.Errorf("%v\n", ts)
	}

	cc := ts.Clone()
	if len(cc) != len(ts) {
		t.Errorf("Clone: expected len(cc) == len(ts) got %v\n", len(cc))
		t.Errorf("%v\n", cc)
	}
	cc.Cut(0, len(cc))
	if len(cc) > 0 {
		t.Errorf("Clone Cut: expected len(cc) == 0 got %v\n", len(cc))
		t.Errorf("%v\n", cc)
	}

	ts.Insert(testStruct{s: "f", i: 6}, 3)
	if len(ts) != 6 {
		t.Errorf("Insert: expected len(ts) == 6 got %v\n", len(ts))
		t.Errorf("%v\n", ts)
	}
	if ts[3].i != 6 {
		t.Errorf("Insert: expected i == 6 to be in position 3\n")
		t.Errorf("%v\n", ts)
	}

	ts.Delete(0)
	if len(ts) != 5 {
		t.Errorf("Delete: expected len(ts) == 5 got %v\n", len(ts))
		t.Errorf("%v\n", ts)
	}

	ts.Cut(0, 3)
	if len(ts) != 2 {
		t.Errorf("Cut: expected len(ts) == 2 got %v\n", len(ts))
		t.Errorf("%v\n", ts)
	}

	tt := tsSlice(make([]testStruct, 0, 5))
	tt.Append(testStruct{s: "e", i: 5})
	tt.Append(testStruct{s: "d", i: 4})
	tt.Append(testStruct{s: "c", i: 3})
	tt.Append(testStruct{s: "b", i: 2})
	tt.Append(testStruct{s: "a", i: 1})
}

func TestSort(t *testing.T) {
	ts := make(tsSlice, 0, 5)
	ts.Append(testStruct{s: "e", i: 5})
	ts.Append(testStruct{s: "d", i: 4})
	ts.Append(testStruct{s: "c", i: 3})
	ts.Append(testStruct{s: "b", i: 2})
	ts.Append(testStruct{s: "a", i: 1})

	ts.Sort()
	if ts[0].i != 1 {
		t.Errorf("Sort: expected first item to be i == 1 got %v\n", ts[0].i)
		t.Errorf("%v\n", ts)
	}
}

func TestSearch(t *testing.T) {
	ts := make(tsSlice, 0, 5)
	ts.Append(testStruct{s: "e", i: 5})
	ts.Append(testStruct{s: "d", i: 4})
	ts.Append(testStruct{s: "c", i: 3})
	ts.Append(testStruct{s: "b", i: 2})
	ts.Append(testStruct{s: "a", i: 1})

	x := testStruct{s: "b", i: 2}
	i := ts.Search(x)
	if i != 3 {
		t.Errorf("Search: expected 3 got %v\n", i)
		t.Errorf("%v\n", ts)
	}
}

func TestSSearch(t *testing.T) {
	ts := make(tsSlice, 0, 5)
	ts.Append(testStruct{s: "e", i: 5})
	ts.Append(testStruct{s: "d", i: 4})
	ts.Append(testStruct{s: "c", i: 3})
	ts.Append(testStruct{s: "b", i: 2})
	ts.Append(testStruct{s: "a", i: 1})

	ts.Sort()

	x := testStruct{s: "b", i: 2}
	i := ts.SSearch(x)
	if i != 1 {
		t.Errorf("SSearch: expected 1 got %v\n", i)
		t.Errorf("%v\n", ts)
	}
}
