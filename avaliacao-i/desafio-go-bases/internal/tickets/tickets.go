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
	departure string
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

const (
	earlyMorning string = "early morning"
	morning string = "morning"
	afternoon string = "afternoon"
	evening string = "evening"
)

// exemplo 2
func GetAllTicketsByPeriod(period string) (int, string, error) {
	var minimum int
	var maximum int

	switch period {
	case earlyMorning:
		minimum = 0
		maximum = 6
	case morning:
		minimum = 7
		maximum = 12
	case afternoon:
		minimum = 13
		maximum = 19
	case evening:
		minimum = 20
		maximum = 23
	default:
		return 0, "", fmt.Errorf("Invalid period: %s. Valid options: %s, %s, %s, %s.", period, earlyMorning, morning, afternoon, evening)
	}

	res, err := os.ReadFile("./tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}

	tickets := strings.Split(string(res), "\n")

	totalTickets := 0

	for i := 0; i < len(tickets); i++ {
		attributes := strings.Split(tickets[i], ",")
		time, err := strconv.Atoi(strings.Split(attributes[4], ":")[0])
		if err != nil {
			panic("Could not get departure time.")
		}
		if time >= minimum || time <= maximum {
			totalTickets++;
		}
	}

	message := fmt.Sprintf("Total tickets for period %s: %d", period, totalTickets)

	return totalTickets, message, nil

}

// exemplo 3
// func AverageDestination(destination string, total int) (int, error) {}
