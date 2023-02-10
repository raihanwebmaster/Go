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

	//Print 
	fmt.Print("Hello, ")
	fmt.Print("world! ")

	//new line
	fmt.Print("Hello, \n")
	fmt.Print("world! \n")

	//auto new line 
	fmt.Println("Hello Raihan")
	fmt.Println("GoogBye Raihan")

	age := 25
	name := "raihan"

	fmt.Println("my age is", age, "and my name is", name)

	//PrintF (formatted strings) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points! \n", 222.55)
	fmt.Printf("you scored %0.1f points! \n", 222.55)

	//sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)




}