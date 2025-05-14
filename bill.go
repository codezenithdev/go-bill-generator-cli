package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tips  float64
}

// make a new bill

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tips:  0,
	}

	return b
}

func (b *bill) format() string {
	fs := "Bill breakdown: \n"

	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tips)

	// print total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tips)

	return fs
}

// update tip

func (b *bill) updateTip(tip float64) {
	b.tips = tip
}

// add item to the bill

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save the bill

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved !!!")
}
