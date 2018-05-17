package main

import "fmt"

type App interface {
	Add() interface{}
}

type MyInt struct {
	V1 int
	V2 int
}

func (mi MyInt) Add() interface{} {
	return mi.V1 + mi.V2
}

type MyFloat struct {
	V1 float32
	V2 float32
}

func (mf MyFloat) Add() interface{} {
	return mf.V1 + mf.V2
}

func main() {
	var ap App
	m1 := MyInt{
		V1: 1,
		V2: 99,
	}

	m2 := MyFloat{
		V1: 1.11,
		V2: 99.9,
	}

	ap = m1
	fmt.Printf("1 Add: %+v \n", ap.Add())

	ap = m2
	fmt.Printf("2 Add: %+v \n", ap.Add())
}
