// stackoverflow
// https://stackoverflow.com/questions/38645175/why-have-arrays-in-go
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [5]int{1, 2, 30, 4, 5}
	s := []int{1, 2, 30, 4, 5}

	fmt.Printf("s[0] pointer: %p (before)\n", &s[0])
	fmt.Printf("s[0] value: %d (before) \n", s[0])
	sliceFunc(s)
	fmt.Printf("s[0] pointer: %p (after)\n", &s[0])
	fmt.Printf("s[0] value: %d (after)\n", s[0])

	fmt.Println()

	fmt.Printf("a[0] pointer: %p (before)\n", &a[0])
	fmt.Printf("a[0] value: %d (before) \n", a[0])
	arrayFunc(a)
	fmt.Printf("a[0] pointer: %p (after)\n", &a[0])
	fmt.Printf("a[0] value: %d (after)\n", a[0])

	fmt.Println()

	fmt.Printf("before sort slice: %v \n", s)
	sort.Ints(s)
	fmt.Printf("after sort slice: %v \n", s)

	fmt.Println()

	// CAN NOT sort array!
	// cannot use a (type [5]int) as type []int in argument to sort.Ints
	//sort.Ints(a)

	appendTest := make([]int, 0, 3*3)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			appendSlice(appendTest, i*j)
		}
	}
	fmt.Printf("append slice end, got: %v \n", appendTest)

	appendTest = appendTest[:0]
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			appendTest = appendSliceBis(appendTest, i*j)
		}
	}
	fmt.Printf("append slice end, got: %v \n", appendTest)
}

// slice pass as pointer
func sliceFunc(s []int) {
	s[0] = 99

	fmt.Printf("s[0] pointer in func: %p \n", &s[0])
	fmt.Printf("s[0] value: %d in func \n", s[0])
}

// slice pass as value -> copy
func arrayFunc(a [5]int) {
	a[0] = 99

	fmt.Printf("a[0] pointer in func: %p \n", &a[0])
	fmt.Printf("a[0] value: %d in func \n", a[0])
}

func appendSlice(s []int, v int) {
	s = append(s, v)
	fmt.Printf("append slice in func: %v \t\t %p\n", s, &s[0])
}

func appendSliceBis(s []int, v int) []int {
	s = append(s, v)
	fmt.Printf("append slice in func: %v \t\t %p\n", s, &s[0])

	return s
}
