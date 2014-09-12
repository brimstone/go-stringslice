package stringslice

import "fmt"

// Define a type named "StringSlice" as a slice of strings

type StringSlice []string

// This is mainly for use with the flags package
func (i *StringSlice) String() string {
	return fmt.Sprint(*i)
}

// This method is also required by the flags package
func (i *StringSlice) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// This compares one object of StringSlice type with another
func (p *StringSlice) Compare(a *StringSlice) bool {
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
func New(a string) *StringSlice {
	foo := new(StringSlice)
	foo.Set(a)
	return foo
}
