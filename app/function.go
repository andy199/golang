package main

import (
	"fmt"
)


func calculateBill(price, no int) int{
	var totaPrice = price * no
	return totaPrice
}

func main() {
	price, no := 9, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)
}