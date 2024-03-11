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
	var ageOne int = 23
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
	//Println start from new line while Print and Printf do not
	fmt.Println(myName)
	fmt.Print(myName, "\n")
	fmt.Print(myName)
	fmt.Println("my name is", myName, "i am", ageOne, "years old")

	//format specifier for integer and string
	fmt.Printf("my name is %v, I am %v years old and a biochemist \n", myName, ageOne)
	fmt.Printf("Hello %v how was your day %v \n", nameTwo, nameOne)
	fmt.Printf("Hello %q \n", nameOne)
	fmt.Printf("%T", ageOne)

	//formatting float
	fmt.Printf("your score is %0.2f \n", 20.12345)

	//sprintf (save formatted string) this is used to save string that contain value in it

	str := fmt.Sprintf("Hello %v", nameThree)
	fmt.Println(str)
}

//the shorthand for declaring variable cannot be used outside a function but var can be used to declare variable outside a function
//%_ can be used for format specifier
