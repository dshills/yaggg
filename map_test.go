package main

import "testing"

type mytype struct{ i int }
type mykey string
type mymap map[mykey]mytype

func (m mytype) Clone() mytype {
	return mytype{m.i}
}

func (m mytype) Cmp(b mytype) int {
	if m.i == b.i {
		return 0
	}
	if m.i > b.i {
		return 1
	}
	return -1
}

func (m mymap) Keys() []mykey {
	keys := make([]mykey, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m mymap) Values() []mytype {
	values := make([]mytype, 0, len(m))
	for _, v := range m {
		values = append(values, v.Clone())
	}
	return values
}

func (m mymap) Union(b mymap) mymap {
	z := make(mymap)
	for k, v := range b {
		z[k] = v.Clone()
	}
	for k, v := range m {
		z[k] = v.Clone()
	}
	return z
}

func (m mymap) Intersect(b mymap) mymap {
	z := make(mymap)
	sm := m
	lg := b
	if len(m) > len(b) {
		sm = b
		lg = m
	}

	for k := range sm {
		if _, ok := lg[k]; ok == true {
			z[k] = m[k].Clone()
		}
	}
	return z
}

func (m mymap) Diff(b mymap) mymap {
	z := make(mymap)
	for k, v := range m {
		if _, ok := b[k]; ok == false {
			z[k] = v.Clone()
		}
	}
	return z
}

func (m mymap) Equal(b mymap) bool {
	if len(m) != len(b) {
		return false
	}
	for k, v := range m {
		if _, ok := b[k]; ok == true && b[k].Cmp(v) == 0 {
			continue
		}
		return false
	}
	return true
}

func (m mymap) Clone() mymap {
	z := make(mymap)
	for k, v := range m {
		z[k] = v.Clone()
	}
	return z
}

func TestMap(t *testing.T) {
	mm := make(mymap)
	mm["one"] = mytype{1}
	mm["two"] = mytype{2}
	mm["three"] = mytype{3}
	mm["four"] = mytype{4}
	mm["five"] = mytype{5}
	mm["six"] = mytype{6}

	k := mm.Keys()
	if len(k) != len(mm) {
		t.Errorf("Keys: expected len(k) == len(mm) got %v\n", len(k))
	}
	v := mm.Values()
	if len(v) != len(mm) {
		t.Errorf("Values: expected len(v) == len(mm) got %v\n", len(v))
	}
}

func TestUnion(t *testing.T) {
	m := make(mymap)
	m["one"] = mytype{1}
	m["two"] = mytype{2}
	m["three"] = mytype{3}

	mm := make(mymap)
	mm["four"] = mytype{4}
	mm["five"] = mytype{5}
	mm["six"] = mytype{6}

	z := m.Union(mm)
	if len(z) != (len(m) + len(mm)) {
		t.Errorf("Union: wrong len(z) got %v\n", len(z))
		t.Errorf("%v\n", z)
	}
}

func TestInterset(t *testing.T) {
	m := make(mymap)
	m["one"] = mytype{1}
	m["two"] = mytype{2}
	m["three"] = mytype{3}

	mm := make(mymap)
	mm["two"] = mytype{2}
	mm["three"] = mytype{3}
	mm["four"] = mytype{4}
	mm["five"] = mytype{5}
	mm["six"] = mytype{6}

	z := m.Intersect(mm)
	if len(z) != 2 {
		t.Errorf("Intersect: wrong len(z) expected 2 got %v\n", len(z))
		t.Errorf("%v\n", z)
	}
}

func TestDiff(t *testing.T) {
	m := make(mymap)
	m["one"] = mytype{1}
	m["two"] = mytype{2}
	m["three"] = mytype{3}

	mm := make(mymap)
	mm["two"] = mytype{2}
	mm["three"] = mytype{3}
	mm["four"] = mytype{4}
	mm["five"] = mytype{5}
	mm["six"] = mytype{6}

	z := m.Diff(mm)
	if len(z) != 1 {
		t.Errorf("Diff: wrong len(z) expected 1 got %v\n", len(z))
		t.Errorf("%v\n", z)
	}

	z = mm.Diff(m)
	if len(z) != 3 {
		t.Errorf("Diff: wrong len(z) expected 1 got %v\n", len(z))
		t.Errorf("%v\n", z)
	}
}

func TestClone(t *testing.T) {
	m := make(mymap)
	m["one"] = mytype{1}
	m["two"] = mytype{2}
	m["three"] = mytype{3}
	m["four"] = mytype{4}
	m["five"] = mytype{5}
	m["six"] = mytype{6}

	z := m.Clone()
	if len(z) != len(m) {
		t.Errorf("Clone: wrong len(z) expected 6 got %v\n", len(z))
		t.Errorf("%v\n", z)
	}
	if &z == &m {
		t.Errorf("Clone address should not match got %v\n", &z)
	}
}

func TestEqual(t *testing.T) {
	m := make(mymap)
	m["one"] = mytype{1}
	m["two"] = mytype{2}
	m["three"] = mytype{3}

	mm := make(mymap)
	mm["one"] = mytype{1}
	mm["two"] = mytype{2}

	if m.Equal(mm) == true {
		t.Errorf("Equal: expected not equal got equal\n")
		t.Errorf("Equal: m: %v mm:%v\n", m, mm)
	}

	mm["three"] = mytype{3}
	if m.Equal(mm) == false {
		t.Errorf("Equal: expected equal got not equal\n")
		t.Errorf("Equal: m: %v mm:%v\n", m, mm)
	}

}
