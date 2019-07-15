// version 1.0 2018-08-10
package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Printf("%q\n", m.Get("lang")) // "en"
	fmt.Printf("%q\n", m.Get("q"))    // ""
	fmt.Printf("%q\n", m.Get("item")) // "1"     (first value)
	fmt.Printf("%q\n", m["item"])     // "[1 2]" (direct map access)

	// url.Values is a map type, any updates and deletions that url.Values.Add makes
	// to the map elements are visible t the caller.  However, as with ordinary
	// functions, any changes a method makes to the reference itself, like setting
	// it to nil or making it refer to a different map data structure, will not be
	// reflected in the caller.
	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // panic: assignment to entry in nil map
}
