package main

import (
	"fmt"

	"github.com/helenaperdigueiro/especializacao_backend_iii/tree/main/avaliacao-i/desafio-go-bases/internal/tickets"
)

func main() {
	totalTickets, message, err := tickets.GetTotalTickets("Brazil")
	if err != nil {
		panic("Something went wrong. Check parameters and try again.")
	}
	
	fmt.Printf("Total tickets returned: %d\nMessage returned: %s", totalTickets, message)

}
