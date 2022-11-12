package main

import (
	"fmt"

	"github.com/helenaperdigueiro/especializacao_backend_iii/tree/main/avaliacao-i/desafio-go-bases/internal/tickets"
)

func main() {
	totalTicketsByDestination, message, err := tickets.GetAllTicketsByDestination("Brazil")
	if err != nil {
		panic("Something went wrong. Check parameters and try again.")
	}
	
	fmt.Printf("Total tickets returned: %d\nMessage returned: %s", totalTicketsByDestination, message)

	totalTicketsByPeriod, err := tickets.GetAllTicketsByPeriod("x")
	fmt.Println(totalTicketsByPeriod)
	fmt.Println(err)
}
