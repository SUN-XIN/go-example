// if you want to use Sum with float params
// need to compile with tags "mytag"
// go build -tags mytag .

// +build mytag

package main

func Sum(a, b float32) float32 {
	return a + b
}
