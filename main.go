package main

import (
	"fmt"
	"os"
	"weekly/src/service"
)

func main() {
	for {
		fmt.Print("\x1bc")
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
			fmt.Print("\x1bc")
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
