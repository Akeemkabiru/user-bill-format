//package main
//
//import (
//	"fmt"
//	"math"
//	"strings"
//)
//
//func areaOfCircle(r float64) string {
//	area := fmt.Sprintf("0.1%f", math.Pi*r*r)
//	return area
//}
//
//func goodBye(n string) {
//	fmt.Printf("Hello %v \n", n)
//}
//func welcome(n []string, f func(string)) {
//	for _, value := range n {
//		fmt.Println(value)
//	}
//}
//
//func userName(n string) (string, string) {
//	strSlice := strings.Split(n, " ")
//	fmt.Println(strSlice)
//	var initials []string
//	for _, value := range strSlice {
//		initials = append(initials, value[:1])
//	}
//
//	return initials[0], initials[1]
//}
//func server() {
//	goodBye("Kabby")
//	welcome([]string{"akeem", "kabiru"}, goodBye)
//	areaOne := areaOfCircle(2)
//	areaTwo := areaOfCircle(3)
//	fmt.Println(areaTwo, areaOne)
//	return userName("Akeem Kabiru")
//}
//
////there should be just one main function in your main file which will be specified in each of other files

package main

import "fmt"

var courseMates = map[string]int{
	"Oluwatobi": 70,
	"Akeem":     70,
	"Marvelous": 70,
}

func ageCalculator(n int) int {

	fmt.Println(names)
	return 2024 - n
}
