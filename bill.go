package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

// user input
// create a reader with bufio package that read from standard input(stdin)

// Using helper function to for the logic of reading user input from the terminal

//func getInput(prompt string, r *bufio.Reader) (string, error) {
//	fmt.Print(prompt)
//	input, err = r.ReadString('\n')
//	return strings.TrimSpace(input), err
//}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("creat a new bill name:")
	//read the string that was input on the terminal and save in name while err is the error which is denoted with _ since it is not needed here
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	b := newBill(name)
	return b
}

func promptOptions() {
	fmt.Println("choose an option (a-add item) (s-save bill) (t-add tip)")
	reader := bufio.NewReader(os.Stdin)
	opt, _ := reader.ReadString('\n')
	switch strings.TrimSpace(opt) {
	case "a":
		fmt.Println("add")
	case "s":
		fmt.Println("save bill")
	case "t":
		fmt.Println("add tip")
	default:
		fmt.Println("that was not a valid option")
	}

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
	fmt.Println(createBill())
	promptOptions()
}
