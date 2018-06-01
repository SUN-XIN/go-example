package main

import "fmt"

func main() {
	fmt.Println("copy list by l1:=12")
	l1 := []int{1, 2, 3, 4, 5}
	l2 := l1
	l2[0] = 99
	fmt.Println("list 1:", l1)
	fmt.Println("list 2:", l2)
	fmt.Println()

	fmt.Println("copy list by copy()")
	l3 := make([]int, len(l1))
	copy(l3, l1)
	l3[0] = -1
	fmt.Println("list 1:", l1)
	fmt.Println("list 3:", l3)
}
