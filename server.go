package main

import (
	"fmt"
)

func main() {
	studentsNames := []string{"d", "a", "c"}
	studentsScore := []int{2, 3, 1, 4, 5}
	//these loops do not mutate the slices cos we still have same values in the slice
	for i := 0; i < len(studentsScore); i++ {
		fmt.Println(studentsScore[i] + 2)
	}
	for index, value := range studentsNames {
		fmt.Println(index, value)
	}
	fmt.Println(studentsScore)
	//fmt.Println(append(studentsNames, "d"))
	//sort.Strings(studentsNames)
	//
	//sort.Ints(studentsScore)
	//fmt.Println(sort.SearchInts(studentsScore, 1))
	//fmt.Println(studentsNames, studentsScore)

}
