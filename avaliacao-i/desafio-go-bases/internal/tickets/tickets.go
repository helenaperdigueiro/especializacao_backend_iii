package tickets

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	id int
	name string
	email string
	destination string
	departure time.Time
	price float32
}

// exemplo 1
func GetAllTicketsByDestination(destination string) (int, string, error) {
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
func GetAllTicketsByPeriod(period string) (int, error) {
	res, err := os.ReadFile("./tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}

	tickets := strings.Split(string(res), "\n")
	var minimum int
	var maximum int

	switch period {
	case "early morning":
		minimum = 0
		maximum = 6
	case "morning":
		minimum = 7
		maximum = 12
	case "afternoon":
		minimum = 13
		maximum = 19
	case "evening":
		minimum = 20
		maximum = 23
	}

	totalTickets := 0

	for i := 0; i < len(tickets); i++ {
		attributes := strings.Split(tickets[i], ",")
		time, err := strconv.Atoi(strings.Split(attributes[4])[0])
		if err != nil {
			panic("Could not get departure time.")
		}
		if time >= minimum || time <= maximum {
			totalTickets++;
		}
	}

	return totalTickets, nil

}

// exemplo 3
// func AverageDestination(destination string, total int) (int, error) {}
