package main

import "fmt"

func main() {
	total := sum([]int{1, 2, 3, 4, 5})
	fmt.Println(total)
	fmt.Println(userName("akeem", "kabiru"))
	fmt.Println(changeDepartment(userScore, 30))
	fmt.Printf("what is your name? %q \n", myNames)
	fmt.Println("hello", myNames)
	newName := "kabiru" //in another memory location
	m := &newName
	updateName(m) //a copy of the newName is created and stored in another memory location
	fmt.Println(&newName)
	fmt.Println(newBill("Ezekiel"))
}
