package main

import (
	"fmt"
	"sort"
	"strings"
)
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

	//The standard Library strings Package
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting,"hello"))
	fmt.Println(strings.Contains(greeting,"hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	fmt.Println("original string value =", greeting)


	//sort package
	agess := []int{45,20,49, 40, 88, 90}
	sort.Ints(agess)
	fmt.Println(agess)

	index := sort.SearchInts(agess, 40)
	fmt.Println(index)

	testNames := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(testNames)
	fmt.Println(testNames)

	fmt.Println(sort.SearchStrings(testNames,"bowser"))


	//Loop
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	for i := 0; i < 5; i++{
		fmt.Println("value of i is:", i)
	}

	loopName := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	for i := 0; i < len(loopName); i++{
		fmt.Println(loopName[i])
	}

	for index, value := range loopName{
		fmt.Printf("the value at index %v is %v \n",index, value)
	}

	for _, value := range loopName{
		fmt.Printf("the value  is %v \n",value)
	}

	// Boolean & conditionals
	booleanAge := 45
	fmt.Println(booleanAge <= 50)
	fmt.Println(booleanAge >= 50)
	fmt.Println(booleanAge == 45)
	fmt.Println(booleanAge != 50)

	if booleanAge < 30 {
		fmt.Println("age is less than 30")
	} else if booleanAge < 40 {
		fmt.Println("age is less than 40")
	}else {
		fmt.Println("age is not less than 45")
	}

	for index, value := range names {
		if index == 1{
			fmt.Println("continuing at pos ", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at post", index)
		}
		fmt.Printf("the value at pos %v ia %v \n", index, value)
	}
}