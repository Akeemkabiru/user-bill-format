package main

import "fmt"

type bill struct {
	name  string
	bills map[string]float64
	tip   float64
}

// creating new bill for user
func newBill(name string) bill {
	b := bill{
		name:  name,
		bills: map[string]float64{"macbook": 20, "phone": 30, "workspace": 100},
		tip:   0,
	}
	return b
}

// Receiver function to format the bill
func (b bill) formatBills() string {
	var totalBills float64
	var formattedBills string
	for key, value := range b.bills {
		formattedBills += fmt.Sprintf("%-12v...%v\n", key, value)
		//total bills
		totalBills = totalBills + value
	}
	return fmt.Sprintf("%-10v\ntotal bills...%v", formattedBills, totalBills)
}

// Receiver function that update bill and tip
func (b bill) updateTip(billValue float64) float64 {
	b.tip = billValue
	return b.tip
}

// Receiver function that update the bill
func (b bill) addBill(product string, price float64) map[string]float64 {
	b.bills = map[string]float64{product: price}
	return b.bills
}

// main
func main() {
	bill := newBill("Akeem")
	fmt.Println(bill)
	formatted := newBill("kabiru").formatBills()
	fmt.Println(formatted)
	newTip := newBill("akeem").updateTip(30)
	fmt.Println(newTip)
	fmt.Println(newBill("Dorcas"))
}
