package main

import (
	"fmt"
	"os"
	"weekly/src/service"
	"weekly/utils"
)

func main() {
	for {
		utils.ClearScreen()
		fmt.Println("=== Welcome To Pizza ===")
		fmt.Println(`
	Menu :
	1. Menu
	2. History Order
	
	0. Exit
		`)
		fmt.Print("Masukan menu:")
		var input string
		fmt.Scan(&input)
		switch input {
		case "1":
			service.OrderService()
		case "2":
			service.HistoryOrder()
		case "0":
			utils.ClearScreen()
			fmt.Println("Terima kasih! Datang Kembali.")
			os.Exit(0)
		default:
			fmt.Println("Invalid Input")
			fmt.Scanln()
			main()
			return
		}
	}
}
