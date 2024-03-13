package main

import (
	"fmt"
	"sort"
)

func main() {
	studentsNames := []string{"d", "a", "c"}
	studentsScore := []int{2, 3, 1, 4, 5}
	fmt.Println(append(studentsNames, "d"))
	sort.Strings(studentsNames)

	sort.Ints(studentsScore)
	fmt.Println(sort.SearchInts(studentsScore, 1))
	fmt.Println(studentsNames, studentsScore)
}
