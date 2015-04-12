# yaggg [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/dshills/yaggg)
>Yet Another Go Generic Generator

## Overview
There are a few options for statically generating generics for Go. I mention them in the Alternatives section. However, I was interested in the new go generate in combination with the template library. It currently only supports slices and maps and only using the built in templates that I wrote. Future versions will have support for external templates.

Feel free to fork or pull.

Vary much a work in progress but functional and useful.

## Features
- uses Go templates to generate output
- useful running with go generate
- No interface{} nonsense, brute force, type specific

## Slice
- Swap // swap values in a slice
- Sort // sort a slice using a function
- Search // search for a value using a function
- SSearch // search a sorted list
- Filter // return a new slice based on a function
- FilterNot // return a new slice based on a function
- Map // run a function across all elements
- Cut // remove a range of values
- Delete // remove a value
- Insert // insert a new item in the middle of a slice
- Append // append a value to the end of a slice (same as append)
- Prepend // add a value to the beginning of a slice
- Clone // create a new slice based on the old slice using a defined function
- Extract // create a slice from a key value within value

## Map
- Keys // returns a slice of the keys
- Values // returns a slice of the values
- Union // creates a new map from two maps
- Intersect // creates a map of the intersection of two maps
- Diff // creates a map of the unique keys in a map
- Equal // compares two maps for value equality
- Clone // creates a new map using a defined function

## Instalation
	go get github.com/dshills/yaggg

## Usage
User Type must implement 3 functions Cmp, Clone, Clear in the following forms:

func (a MyType)Cmp(b MyType) int
Cmp must return 0 for equality, 1 when a > b, -1 when a < b

func (a MyType)Clone() MyType
Clone must return a deep copy of MyType

func (a MyType)Clear() MyType
Clear must return a valid empty MyType

	yaggg --output <outfile> --generate <map | slice> --container <conName> --value <mytype> --key <mapKey>

example:

	yagg --output mytypemap.go --generate map --container mytypemap --value mytype --key string

Using Go Generate

	//go:generate yagg --output mytypemap.go --generate map --container mytypemap --value mytype --key string

## Sample
```Go
package main

//go:generate yaggg --output boguses.go --container boguses --value bogus
//go:generate yaggg --output bogusmap.go --container bogusmap --value bogus --key string --generate map

type bogus struct {
	ID    int
	Bool  bool
	Int   int
	Float float64
	Slice []string
}

func (b bogus) Cmp(j bogus) int {
	if b.ID == j.ID {
		return 0
	}
	if b.ID > j.ID {
		return 1
	}
	return -1
}

func (b bogus) Clone() bogus {
	c := b
	return c
}

func (b bogus) Clear() bogus {
	return bogus{}
}

func main() {

}
```

## To Do
Almost everything

## Alternatives
* [gengen](https://github.com/joeshaw/gengen)
* [gen](http://clipperhouse.github.io/gen/)

## License
Copyright 2015 Davin Hills. All rights reserved.
MIT license. License details can be found in the LICENSE file.
