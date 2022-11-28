package main

import (
	"fmt"
)

func main() {

	numbers := []int{25, 5, 8, 16, 50, 33}

	odd := make(chan int)
	even := make(chan int)

	go func() {
		for _,number := range numbers {
			if number%2 != 0 {
				odd <- number
			} else {
				even <- number
			}
		}
	}()

	//meensagem do channel odd
	go func() {
		for number := range odd {
			fmt.Println("number", number, "is odd")
		}
	}()

	//mensagem do channel even
	go func() {
		for number := range even {
			fmt.Println("number", number, "is even")
		}
	}()

	for {}
}
