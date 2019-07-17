package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{3, 1, 4, 1}
	fmt.Println("original:", values)
	fmt.Println("is sorted:", sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	fmt.Println("sort int:", values)                      // "[1 1 3 4]"
	fmt.Println("is sorted:", sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println("reverse int:", values)                   // "[4 3 1 1]"
	fmt.Println("is sorted:", sort.IntsAreSorted(values)) // "false"
}
