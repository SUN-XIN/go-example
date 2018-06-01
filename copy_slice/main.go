package main

import "fmt"

func main() {
	fmt.Println("copy list by l1:=12")
	l1 := []int{1, 2, 3, 4, 5}
	l2 := l1
	l2[0] = 99
	fmt.Printf("list 1: %p %v \n", &l1[0], l1)
	fmt.Printf("list 2: %p %v \n", &l2[0], l2)
	fmt.Println()

	fmt.Println("copy list by copy()")
	l3 := make([]int, len(l1))
	copy(l3, l1)
	l3[0] = -1
	fmt.Printf("list 1: %p %v \n", &l1[0], l1)
	fmt.Printf("list 3: %p %v \n", &l3[0], l3)
}
