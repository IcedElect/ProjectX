package main

import (
	"HOTA/internal/handlers"
	"HOTA/internal/printer"
	"fmt"
)

func Menu() int {
	fmt.Println()
	printer.PrintMeny()
	fmt.Println()

	var action int
	fmt.Scan(&action)
	fmt.Println()
	return action
}

func main() {
	actions := []func(){
		handlers.CreateUser,
		handlers.GetUserByNik,
		handlers.GetUserByID,
		handlers.SerheUserByStack,
		handlers.ListUser,
		handlers.UpdateUser,
		handlers.DeleteIserID,
		handlers.StatUser,
	}

	for {
		action := Menu()

		actions[action-1]()

	}

}
