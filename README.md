# go-stringslice
A simple package to provide the flag package a way to accept a string slice as an option.

## Requirements
Nothing

## Installation
Just include it in the imports.

## Usage

```go
package main

import (
    "fmt"
    "flag"
    stringslice "github.com/brimstone/go-stringslice"
)


func main() {
    manyThings := new(stringslice.StringSlice)
    // parse our command line args
    flag.Var(manyThings, "thing", "Things, specify more than one")
    flag.Parse()
    things := []string(*manyThings)

    fmt.Println("One thing:", things[0]
}

```