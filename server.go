package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	var firstName string = "akeem"
	fmt.Println(firstName)
	secondName := "Kabiru"
	fmt.Println(secondName)
	var ageOne int = 20
	ageTwo := 30
	fmt.Println(ageTwo, ageOne)
	//string formatter
	var thirdName string = "Sunday"
	fmt.Printf("my third name is %v \n", thirdName) //template literal approximate in golang, %_ is a specifier ,while _ can be T for determine type,_ can be v which is used to format string by joining it with value, %q is is used to quote the value
	fmt.Print("Kabiru")
	//Array and slices: array has specific length and cannot be changed but slices can be changed and no specific length

	var arrOne [4]float64 = [4]float64{1.1, 1.2, 1.3, 1.4}
	fmt.Println(arrOne)
	arrTwo := [3]string{"akeem", "kabiru", "sunday"}
	fmt.Println(arrTwo)
	arrThree := []string{"easyyy"}
	arr := append(arrThree, "ezekiel")
	fmt.Println(arr)
	fmt.Println(len(arr))
	strValue := fmt.Sprintf("Hello %v", firstName)
	fmt.Println(strValue)

}
