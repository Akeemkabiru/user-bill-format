package main

import (
	"fmt"
	"math"
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

func main() {
	goodBye("Kabby")
	welcome([]string{"akeem", "kabiru"}, goodBye)
	areaOne := areaOfCircle(2)
	areaTwo := areaOfCircle(3)
	fmt.Println(areaTwo, areaOne)

}
