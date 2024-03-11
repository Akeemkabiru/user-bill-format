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
	strOne := fmt.Sprintf("Hello %v", nameTwo)
	fmt.Println(strOne)

	//Array and slices

	var ages [3]int = [3]int{20, 25, 30}
	fmt.Println(ages)
	names := [3]int{1, 2, 3}
	fmt.Println(names[1])
	//slices (uses arrays under the hood)
	var scores [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(scores, len(scores))
	//both array and slice are mutable the diff btw the two is that array has specific number while slice do not have

	arrOne := [3]string{"ayo", "ade", "bola"}
	fmt.Println(arrOne)
	sliceOne := []int{1, 2, 3}
	fmt.Println(sliceOne)
	arrOne[2] = "middey"
	fmt.Println(arrOne)
	sliceOne[1] = 3
	fmt.Println(sliceOne)
	//both are mutable , and the diff is that we can append slice
	//Append do not mutate a slice but instead create another slices containing the update version

	sliceOne = append(sliceOne, 4)
	fmt.Println(sliceOne)
	//to update a slice you just simply update the whole slice which means we are basically creating new slice and not the old one

	studentNames := []string{"sunday", "oluwatobi", "marvelous"}
	fmt.Println(studentNames)
	studentNames = append(studentNames, "John")
	fmt.Println(studentNames)

	//slice range: it is used to take part of an array and add create a variable with it
	rangeOne := sliceOne[1:3] //inclusive of the first number but not the last number
	fmt.Println(rangeOne)
	rangeTwo := sliceOne[2:3]
	rangeThree := sliceOne[1:] //two inclusive and the rest upward
	rangeFour := sliceOne[:3]
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)
	fmt.Println(rangeFour)
	sliceTwo := append(sliceOne, 5) //update slice
	fmt.Println(sliceTwo)
	//length of array cannot change , length of slices can change. You use slices most time and array is rarely used
}

//the shorthand for declaring variable cannot be used outside a function but var can be used to declare variable outside a function
//%_ can be used for format specifier
