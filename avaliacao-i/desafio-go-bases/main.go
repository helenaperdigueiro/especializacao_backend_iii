package main

import (
	"fmt"

	"github.com/helenaperdigueiro/especializacao_backend_iii/tree/main/avaliacao-i/desafio-go-bases/internal/tickets"
)

func main() {
	totalTicketsByDestination, messageDestination, err := tickets.GetAllTicketsByDestination("Brazil")
	if err != nil {
		panic("Something went wrong. Check parameters and try again.")
	}
	
	fmt.Printf("Total tickets returned: %d\nMessage returned: %s", totalTicketsByDestination, messageDestination)

	totalTicketsByPeriod, messagePeriod, err := tickets.GetAllTicketsByPeriod("early morning")
	if err != nil {
		panic("Something went wrong. Check parameters and try again.")
	}
	fmt.Printf("Total tickets returned: %d\nMessage returned: %s", totalTicketsByPeriod, messagePeriod)

}
