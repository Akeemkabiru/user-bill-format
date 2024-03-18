package main

import (
	"fmt"
	"strings"
)

type user struct {
	name     string
	email    string
	password string
	products map[string]float64
}

func newUser(name string, email string, password string, products map[string]float64) user {
	firstUser := user{
		name:     name,
		email:    email,
		password: password,
		products: products,
	}
	return firstUser
}

func (firstUser user) userName() string {
	firstLetter := firstUser.name[:2]
	return strings.ToUpper(firstLetter)
}

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

func (b bill) updateName(n string) string {
	b.name = n
	return n
}

// slice
var scores []int = []int{1, 2, 3, 4}

// map
var users map[string]string = map[string]string{"Akeem": "Man"}

//receiver function that update the bill and change the tip

// Pointer and Memory
func memo() {
	age := 20
	m := &age
	fmt.Println(m)  //ampersand is used to get location of a value in memory
	fmt.Println(*m) //asteriks is used to find value in a particular memory location
}
