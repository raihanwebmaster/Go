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

	//Print (formatted strings) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points! \n", 222.55)
	fmt.Printf("you scored %0.1f points! \n", 222.55)

	//sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)

	//Array
	var ages [3]int = [3]int{20, 25, 30}
	var agestwo = [3]int{20, 25, 30}
	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	fmt.Println(ages, len(ages))
	fmt.Println(agestwo, len(agestwo))
	fmt.Println(names, len(names))

	//Slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 35
	scores = append(scores,85)
	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)


}