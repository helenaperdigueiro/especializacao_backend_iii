package tickets

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id string
	Name string
	Email string
	Destination string
	Departure string
	Price string
}

func GetAllTicketsByDestination(destination string) (int, string, []Ticket, error) {
	res, err := os.ReadFile("./tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}

	tickets := strings.Split(string(res), "\n")

	var ticketsFromDestination []Ticket

	for i := 0; i < len(tickets); i++ {
		attributes := strings.Split(tickets[i], ",")

		if attributes[3] == destination {
			ticket := Ticket{
				Id: attributes[0],
				Name: attributes[1],
				Email: attributes[2],
				Destination: attributes[3],
				Departure: attributes[4],
				Price: attributes[5],
			}
			ticketsFromDestination = append(ticketsFromDestination, ticket)
		}

	}

	totalTickets := len(ticketsFromDestination)
	message := fmt.Sprintf("Total tickets for %s: %d", destination, totalTickets)

	return totalTickets, message, ticketsFromDestination, nil
}

const (
	earlyMorning string = "early morning"
	morning string = "morning"
	afternoon string = "afternoon"
	evening string = "evening"
)

func getRangeForPeriod(period string) (int, int, error) {
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
		return 0, 0, fmt.Errorf("Invalid period: %s. Valid options: %s, %s, %s, %s.", period, earlyMorning, morning, afternoon, evening)
	}
	return minimum, maximum, nil
}

func GetAllTicketsByPeriod(period string) (int, string, []Ticket, error) {
	minimum, maximum, err := getRangeForPeriod(period)
	if err != nil {
		panic(err)
	}

	res, err := os.ReadFile("./tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}

	tickets := strings.Split(string(res), "\n")
	
	var ticketsFromPeriod []Ticket

	for i := 0; i < len(tickets); i++ {
		attributes := strings.Split(tickets[i], ",")
		time, err := strconv.Atoi(strings.Split(attributes[4], ":")[0])
		if err != nil {
			panic("Could not get departure time.")
		}
		if time >= minimum && time <= maximum {
			ticket := Ticket{
				Id: attributes[0],
				Name: attributes[1],
				Email: attributes[2],
				Destination: attributes[3],
				Departure: attributes[4],
				Price: attributes[5],
			}
			ticketsFromPeriod = append(ticketsFromPeriod, ticket)
		}
	}

	totalTickets := len(ticketsFromPeriod)
	message := fmt.Sprintf("Total tickets for period %s: %d", period, totalTickets)

	return totalTickets, message, ticketsFromPeriod, nil
}

func containsInArray(arrayOfElements []string, element string) bool {
	for _, elementInArray := range arrayOfElements {
		if elementInArray == element {
			return true
		}
	}
	return false
}

func GetDistinctDestinations() (int, []string) {
	res, err := os.ReadFile("./tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}

	tickets := strings.Split(string(res), "\n")
	distinctDestinations := []string{}

	for i := 0; i < len(tickets); i++ {
		attributes := strings.Split(tickets[i], ",")
		destination := attributes[3]

		if !containsInArray(distinctDestinations, destination) {
			distinctDestinations = append(distinctDestinations, destination)
		}
	}

	totalDistinctDestinations := len(distinctDestinations)

	return totalDistinctDestinations, distinctDestinations
}

func GetAverageForTicketsPerDestinations() (int, string, error) {
	res, err := os.ReadFile("./tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}

	tickets := strings.Split(string(res), "\n")

	totalTickets := len(tickets)
	totalDistinctDestinations, _ := GetDistinctDestinations()
	var averageTicketsPerDestinations int
	if totalDistinctDestinations != 0 {
		averageTicketsPerDestinations = totalTickets / totalDistinctDestinations
	}
	
	message := fmt.Sprintf("Average tickets per destinations: %d", averageTicketsPerDestinations)

	return averageTicketsPerDestinations, message, nil
}
