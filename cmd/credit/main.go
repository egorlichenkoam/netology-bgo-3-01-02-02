package main

import (
	"01/pkg/credit"
	"fmt"
)

/**
Домашняя работа bgo-3-02-01

https://https://github.com/netology-code/bgo-homeworks/tree/master/02_func
*/
func main() {
	monthlyPayment, overpayment, totalPayment := credit.Calculate(1_000_000_00, 36, 20)
	fmt.Printf("Monthly payment : %d\nOverpayment : %d\nTotal payment : %d\n", monthlyPayment, overpayment, totalPayment)
}
