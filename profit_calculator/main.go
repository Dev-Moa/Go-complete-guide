package main

import "fmt"

func main() {
	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Println("Profit calculator")

	fmt.Print("revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("expense: ")
	fmt.Scan(&expense)

	fmt.Print("taxRate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expense

	profit := ebt * (1 * taxRate / 100)

	ratio := ebt / profit

	fmt.Println("EBT", ebt)
	fmt.Println("profit", profit)
	fmt.Println("ratio", ratio)
}
