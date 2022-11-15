package main

import (
	"fmt"

	"github.com/helenaperdigueiro/especializacao_backend_iii/tree/main/avaliacao-i/desafio-go-bases/internal/tickets"
)

func main() {
	fmt.Println("---")

	totalTicketsByDestination, messageDestination, ticketsFromDestination, err := tickets.GetAllTicketsByDestination("Brazil")
	if err != nil {
		panic("Something went wrong. Check parameters and try again.")
	}
	
	fmt.Printf("Total tickets returned: %d\nMessage returned: %s\nTickets returned: %s\n", totalTicketsByDestination, messageDestination, ticketsFromDestination)

	fmt.Println("---")

	totalTicketsByPeriod, messagePeriod, ticketsFromPeriod, err := tickets.GetAllTicketsByPeriod("early morning")
	if err != nil {
		panic("Something went wrong. Check parameters and try again.")
	}

	fmt.Printf("Total tickets returned: %d\nMessage returned: %s\nTickets returned: %s\n", totalTicketsByPeriod, messagePeriod, ticketsFromPeriod)

	fmt.Println("---")

	averageTicketsPerDestinations, messageAverage, err := tickets.GetAverageForTicketsPerDestinations()
	if err != nil {
		panic("Something went wrong. Check parameters and try again.")
	}

	fmt.Printf("Total tickets returned: %d\nMessage returned: %s\n", averageTicketsPerDestinations, messageAverage)

	fmt.Println("---")

	totalDistinctDestinations, distinctDestinations := tickets.GetDistinctDestinations()

	fmt.Printf("Total distinct destinations: %d\nDistinct destinations: %s\n", totalDistinctDestinations, distinctDestinations)

	fmt.Println("---")

}
