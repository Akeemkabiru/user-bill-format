package main

import (
	"fmt"
	"math"
	"strings"
)

func areaOfCircle(r float64) string {
	area := fmt.Sprintf("0.1%f", math.Pi*r*r)
	return area
}

func goodBye(n string) {
	fmt.Printf("Hello %v \n", n)
}
func welcome(n []string, f func(string)) {
	for _, value := range n {
		fmt.Println(value)
	}
}

func userName(n string) (string, string) {
	strSlice := strings.Split(n, " ")
	fmt.Println(strSlice)
	var initials []string
	for _, value := range strSlice {
		initials = append(initials, value[:1])
	}

	return initials[0], initials[1]
}
func main() {
	goodBye("Kabby")
	welcome([]string{"akeem", "kabiru"}, goodBye)
	areaOne := areaOfCircle(2)
	areaTwo := areaOfCircle(3)
	fmt.Println(areaTwo, areaOne)
	fmt.Println(userName("Akeem Kabiru"))

}
