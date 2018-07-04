package main

import "log"

func main() {
	// Test with map
	mapTest := make(map[int]bool)
	mapTest[0] = true
	mapTest[1] = true
	mapTest[2] = true
	updateMap(mapTest)
	if mapTest[0] {
		log.Printf("Pass a map into function -> no changed")
	} else {
		log.Printf("Pass a map into function -> changed")
	}

	// Test with slice
	sliceTest := []int{1, 2, 3}
	updateSlice(sliceTest)
	//log.Println(sliceTest)
	if sliceTest[0] > 0 {
		log.Printf("Pass a slice into function -> no changed")
	} else {
		log.Printf("Pass a slice into function -> changed")
	}

	// Test with pointer slice
	v1, v2, v3 := 1, 2, 3
	slicePointerTest := []*int{&v1, &v2, &v3}
	updatePointerSlice(slicePointerTest)
	if *slicePointerTest[0] > 0 {
		log.Printf("Pass a pointer slice into function -> no changed")
	} else {
		log.Printf("Pass a pointer slice into function -> changed")
	}

	// Test with array
	arrayTest := [3]int{1, 2, 3}
	updateArray(arrayTest)
	if arrayTest[0] > 0 {
		log.Printf("Pass a array into function -> no changed")
	} else {
		log.Printf("Pass a array into function -> changed")
	}

	// Test with pointer array
	arrayPointerTest := [3]*int{&v1, &v2, &v3}
	updatePointerArray(arrayPointerTest)
	if *arrayPointerTest[0] > 0 {
		log.Printf("Pass a pointer array into function -> no changed")
	} else {
		log.Printf("Pass a pointer array into function -> changed")
	}

	// Test with pointer
	pointerTest := [3]int{1, 2, 3}
	updatePointer(&pointerTest)
	if pointerTest[0] > 0 {
		log.Printf("Pass a pointer into function -> no changed")
	} else {
		log.Printf("Pass a pointer into function -> changed")
	}

	// Test with copy pointer in function
	val := &Val{
		Vint:    1,
		Vstring: "a",
	}
	updateWithCopy(val, -1, "zzz")
	if val.Vint > 0 {
		log.Printf("Pass a pointer into function and then copy in function -> no changed")
	} else {
		log.Printf("Pass a pointer into function and then copy in function -> changed")
	}
}

func updateMap(m map[int]bool) {
	m[0] = false
}

func updateSlice(s []int) {
	s[0] = -1
}

func updatePointerSlice(s []*int) {
	newVal := -1
	s[0] = &newVal
}

func updateArray(a [3]int) {
	a[0] = -1
}

func updatePointerArray(a [3]*int) {
	newVal := -1
	a[0] = &newVal
}

func updatePointer(a *[3]int) {
	a[0] = -1
}

func updateWithCopy(v *Val, newInt int, newString string) {
	v = &Val{
		Vint:    newInt,
		Vstring: newString,
	}
}

type Val struct {
	Vint    int
	Vstring string
}
