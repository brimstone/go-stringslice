package stringslice

import "fmt"

// Define a type named "stringSlice" as a slice of strings

type stringSlice []string

// This is mainly for use with the flags package
func (i *stringSlice) String() string {
	return fmt.Sprint(*i)
}

// This method is also required by the flags package
func (i *stringSlice) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// This compares one object of stringSlice type with another
func (p *stringSlice) Compare(a *stringSlice) bool {
	b := []string(*a)
	// simple case
	if len(*p) != len(b) {
		return false
	}
	// deep inspection
	for i, x := range *p {
		if b[i] != x {
			return false
		}
	}
	return true
}

// Conventional shortcut
func New(a string) *stringSlice {
	foo := new(stringSlice)
	foo.Set(a)
	return foo
}
