package main

import (
	"fmt"
	"math"
	"strconv"

	"rsc.io/quote"
)

// No constants in structs
type Vertex3D struct {
	X int
	Y int
	Z int
}

// functions

func sqrtAndCast(x float64) int {
	if x < 0 {
		return -1
	}
	return int(math.Sqrt(x))
}

func main() {
	fmt.Println(quote.Go())

	// This is a comment!

	/*
		Multiline comments!!
	*/

	// Verbose type decleration
	var sum1 int = 2

	// Implicit type decleration
	sum2 := 3
	fmt.Println(sum1 + sum2)

	// loops
	// While loop
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
		j := 3
		j++
	}

	// Array
	var name [7]string

	// For loop
	for j := 0; j < 7; j++ {
		name[j] = strconv.Itoa(j)
	}

	// Priniting arrays? Yes!
	fmt.Println(name)

	myVert := Vertex3D{1, 2, 3}

	fmt.Println(myVert)

	// Reassignment of structs work
	myVert.X = 12

	fmt.Println(myVert)

	// user input
	var test string
	fmt.Scan(&test)
	fmt.Println(test)

	// Pointers
	fmt.Println(&test)

	pTest := &test

	fmt.Println(*pTest)

	// Control flow
	if sqrtAndCast(-5) == -1 {
		fmt.Println("Learn about i silly :P")
	}

	fmt.Println(sqrtAndCast(15))

}
