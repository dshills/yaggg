package main

//go:generate gogengen --container myCont --value myStruct > myStruct.tmp

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var out, gen, con, val, key string
	flag.StringVar(&out, "output", "", "Output file name. Defaults to stdout")
	flag.StringVar(&gen, "generate", "slice", "What to generate [slice, map]")
	flag.StringVar(&con, "container", "", "Name of the new container")
	flag.StringVar(&val, "value", "", "Name of the type stored in in the container")
	flag.StringVar(&key, "key", "", "Name of the type used for indexing a map")

	flag.Parse()

	if gen == "" || val == "" || con == "" {
		fmt.Printf("\n\tgenerate, container and value are required\n\n")
		flag.PrintDefaults()
		return
	}

	if gen == "map" && key == "" {
		fmt.Printf("\n\tkey type is required for generating map functions\n\n")
		flag.PrintDefaults()
		return
	}

	if gen == "slice" {
		st := SliceType{Slice: con, Type: val}
		st.Generate(os.Stdout)
	}

}
