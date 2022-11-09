package tickets

import {
	"fmt"
	"os"
}

type Ticket struct {
	id int
	name string
	email string
	destination string
	departure string
	price float32
}

// exemplo 1
func GetTotalTickets(destination string) (int, error) {
	res, err := os.ReadFile("./tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}
	fmt.Println(res)

}

// exemplo 2
func GetMornings(time string) (int, error) {}

// exemplo 3
func AverageDestination(destination string, total int) (int, error) {}
