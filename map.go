// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.
package main

import (
	"io"
	"text/template"
)

// A MapType describes the types for code generation
type MapType struct {
	Package string // package name for generated code
	Map     string // user type map name
	Key     string // name of the type used for the keys
	Value   string // name of the type used for the values
}

// Generate will write the map functions to w
func (st MapType) Generate(w io.Writer) (err error) {
	t := template.Must(template.New("Map").Parse(mapTmpl))
	err = t.Execute(w, st)
	return
}

const mapTmpl = `package {{.Package}}

type {{.Map}} map[{{.Key}}]{{.Value}}

func (m {{.Map}}) Keys() []{{.Key}} {
	keys := make([]{{.Key}}, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m {{.Map}}) Values() []{{.Value}} {
	values := make([]{{.Value}}, 0, len(m))
	for _, v := range m {
		values = append(values, v.Clone())
	}
	return values
}

func (m {{.Map}}) Union(b {{.Map}}) {{.Map}} {
	z := make({{.Map}})
	for k, v := range b {
		z[k] = v.Clone()
	}
	for k, v := range m {
		z[k] = v.Clone()
	}
	return z
}

func (m {{.Map}}) Intersect(b {{.Map}}) {{.Map}} {
	z := make({{.Map}})
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

func (m {{.Map}}) Diff(b {{.Map}}) {{.Map}} {
	z := make({{.Map}})
	for k, v := range m {
		if _, ok := b[k]; ok == false {
			z[k] = v.Clone()
		}
	}
	return z
}

func (m {{.Map}}) Equal(b {{.Map}}) bool {
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

func (m {{.Map}}) Clone() {{.Map}} {
	z := make({{.Map}})
	for k, v := range m {
		z[k] = v.Clone()
	}
	return z
}`
