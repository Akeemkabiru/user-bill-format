package main

import (
	"fmt"
	"strings"
)

func sum(n []int) int {
	var totalValue = 0
	for i := 0; i < len(n); i++ {
		totalValue = totalValue + n[i]
	}
	return totalValue
}

func userName(n string, m string) []string {
	nameSlice := []string{n, m}
	var initials []string
	for _, value := range nameSlice {
		initials = append(initials, strings.ToUpper(value[:1]))
	}
	return initials
}

var userScore = map[string]any{
	"name":       "Akeem",
	"age":        23,
	"department": "Biochemistry",
}

func changeDepartment(n map[string]any, m int) any {
	n["age"] = m
	return n
}

var myNames = "Akeem"

func main() {
	total := sum([]int{1, 2, 3, 4, 5})
	fmt.Println(total)
	fmt.Println(userName("akeem", "kabiru"))
	fmt.Println(changeDepartment(userScore, 30))
	fmt.Printf("what is your name? %q \n", myNames)
	fmt.Println("hello", myNames)
}
