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
	t := template.Must(template.New("Slice").Parse(sliceTmpl))
	err = t.Execute(w, st)
	return
}

const mapTmpl = `package {{.Package}}

func (m {{.Map}}) Keys() []{{.Key}} {
	keys := make([]{{.Key}}, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m {{.Map}}) Values() []{{.Value}} {
	values := make([]{{.Value}}, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
`
