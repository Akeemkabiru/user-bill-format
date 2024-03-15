package main

import (
	"fmt"
	"strings"
)

//Struct is a Non pointer wrapper that is used as blueprint in creating an object equivalent

//type bill struct {
//	name string
//	bill map[string]float64
//	tip  float64
//}

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

// look for the value at the memory loaction passed here and use it in the function
func updateName(x *string) {
	*x = "akeem"
	fmt.Println(&x)
}

// Pointer wrapper values : maps, functions, slices store both value and variable name in separate memory location
// Non Pointer Value : integer, float, string, struct , boolean stores their value and name in same memory block
// & ampersand  is used to locate memory address of a value while * is used as pointer to get value in a memory address
