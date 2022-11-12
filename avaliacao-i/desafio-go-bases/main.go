package main

import (
	"fmt"

	"github.com/helenaperdigueiro/especializacao_backend_iii/tree/main/avaliacao-i/desafio-go-bases/internal/tickets"
)

func main() {
	fmt.Println("---")

	totalTicketsByDestination, messageDestination, err := tickets.GetAllTicketsByDestination("Brazil")
	if err != nil {
		panic("Something went wrong. Check parameters and try again.")
	}
	
	fmt.Printf("Total tickets returned: %d\nMessage returned: %s\n", totalTicketsByDestination, messageDestination)

	fmt.Println("---")

	totalTicketsByPeriod, messagePeriod, err := tickets.GetAllTicketsByPeriod("early morning")
	if err != nil {
		panic("Something went wrong. Check parameters and try again.")
	}
	fmt.Printf("Total tickets returned: %d\nMessage returned: %s\n", totalTicketsByPeriod, messagePeriod)

	fmt.Println("---")

	averageTicketsPerDestinations, messageAverage, err := tickets.GetAverageForTicketsPerDestinations()
	if err != nil {
		panic("Something went wrong. Check parameters and try again.")
	}
	fmt.Printf("Total tickets returned: %d\nMessage returned: %s\n", averageTicketsPerDestinations, messageAverage)

	fmt.Println("---")
}
