package main

import "fmt"

type bill struct {
	name string
	bill map[string]float64
	tip  float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		bill: map[string]float64{"macBookPro": 3.5, "monitor": 0.1},
		tip:  0,
	}
	return b
}

// receiver function
func (b bill) format() string {
	var fs string
	var total float64 = 0
	for key, value := range b.bill {
		fs += fmt.Sprintf("%v...%-25v \n", key, value)
		total += value
	}
	fs += fmt.Sprintf("total...%-25v \n", total)
	return fs
}
