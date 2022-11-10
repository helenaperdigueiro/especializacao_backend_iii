package tickets

import (
	"fmt"
	"os"
	// "strconv"
	"strings"
)

type Ticket struct {
	id int
	name string
	email string
	destination string
	departure string
	price float32
}

// exemplo 1
func GetTotalTickets(destination string) (int, string, error) {
	res, err := os.ReadFile("./tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}

	tickets := strings.Split(string(res), "\n")

	totalTickets := 0

	for i := 0; i < len(tickets); i++ {
		attributes := strings.Split(tickets[i], ",")

		if attributes[3] == destination {
			totalTickets++
		}

	}
	message := fmt.Sprintf("Total tickets for %s: %d", destination, totalTickets)

	return totalTickets, message, nil
}

// exemplo 2
// func GetMornings(time string) (int, error) {}

// exemplo 3
// func AverageDestination(destination string, total int) (int, error) {}
