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
	fmt.Println(newBill("akeem").format())
	fmt.Println(newBill("akeem").updateBill("iphone 11", 0.4))
	fmt.Println(newBill("kabiru").updateTips(30))
	fmt.Println(newBill("akeem").updateName("ayinde"))
	fmt.Println(newUser("akeem", "kabbydev@gmail", "kabby@12", map[string]float64{"number of course bought": 20}).userName())
	memo()
}
