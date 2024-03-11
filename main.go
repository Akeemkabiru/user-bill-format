package main

import "fmt"

func main() {
	var nameOne string = "Akeem"
	var nameTwo = "Kabiru"
	fmt.Println(nameOne, nameTwo)
	nameOne = "Ezekiel"
	fmt.Println(nameOne)
	nameThree := "sunday"
	fmt.Println(nameThree)
	var ageOne int = 20
	fmt.Println(ageOne)
	bioData := 30
	fmt.Println(bioData)
	akeemAge := 50
	var akeemJuniorAge int = 20
	fmt.Println(akeemAge, akeemJuniorAge)
	var floatValue float32 = 32.3
	fmt.Println(floatValue)
	var myName string = "Akeem Kabiru Sunday"
	fmt.Print(myName, "\n")
	fmt.Println(myName)
	fmt.Print(myName, "\n")
	fmt.Print(myName)
	fmt.Println("my name is", myName, "i am", ageOne, "years old")
	fmt.Printf("my name is %v, I am %v years old and a biochemist", myName, ageOne)

}

//the shorthand for declaring variable cannot be used outside a function but var can be used to declare variable outside a function
//%_ can be used for format specifier
