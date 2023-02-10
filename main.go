package main

import "fmt"
func main() {
	fmt.Println("Hello, raihan")
	//strings
	var nameOne string = "raihan"
	var nameTwo = "uddin"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameour := "yoshi"

	fmt.Println(nameour)

	//ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne,ageTwo, ageThree)

	//bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint = 25

	fmt.Println(numOne,numTwo,numThree)

	//float
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 55050038038309430.7
	scoreThree := 1.5

	fmt.Println(scoreOne,scoreTwo,scoreThree)
}