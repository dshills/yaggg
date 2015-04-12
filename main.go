package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var out, gen, con, val, key, pkg string
	flag.StringVar(&out, "output", "", "Output file name. Defaults to stdout")
	flag.StringVar(&gen, "generate", "slice", "What to generate [slice, map]")
	flag.StringVar(&con, "container", "", "Name of the new container")
	flag.StringVar(&val, "value", "", "Name of the type stored in in the container")
	flag.StringVar(&key, "key", "", "Name of the type used for indexing a map")
	flag.StringVar(&pkg, "package", "main", "Package name defaults to main")

	flag.Parse()

	if gen == "" || val == "" || con == "" || pkg == "" {
		fmt.Printf("\n\tgenerate, container, package and value are required\n\n")
		flag.PrintDefaults()
		return
	}

	if gen == "map" && key == "" {
		fmt.Printf("\n\tkey type is required for generating map functions\n\n")
		flag.PrintDefaults()
		return
	}

	f := os.Stdout

	if out != "" {
		file, err := os.Create(out)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		f = file
	}

	if gen == "slice" {
		st := SliceType{Slice: con, Type: val, Package: pkg}
		st.Generate(f)
	}

	if gen == "map" {
		mt := MapType{Package: pkg, Map: con, Value: val, Key: key}
		mt.Generate(f)
	}

}
