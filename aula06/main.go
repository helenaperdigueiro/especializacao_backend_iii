package main

import "fmt"

type CreditCard struct {
}

type WireTransfer struct {
}

type Cash struct {
}

type PaymentMethod interface {
	Pay(purchase *Purchase)
}

func (method *CreditCard) Pay(price float64) float64 {
	return 1.1 * price
}

func (method *WireTransfer) Pay(price float64) float64 {
	return price
}

func (method *Cash) Pay(price float64) float64 {
	return 1.05 * price
}

type Purchase struct {
	price float64
	creditCard CreditCard
	wireTransfer WireTransfer
	cash Cash
}

func (purchase *Purchase) setPrice(price float64)  {
	purchase.price = price
}

func (purchase *Purchase) processPayment(price float64)  {
	purchase.price = price
}

func main() {

}