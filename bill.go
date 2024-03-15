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
		bill: map[string]float64{},
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

func (b bill) updateBill(product string, price float64) map[string]float64 {
	b.bill = map[string]float64{product: price}
	return b.bill
}

func (b bill) updateTips(newTip float64) float64 {
	b.tip = newTip
	return b.tip
}

//receiver function that update the bill and change the tip
