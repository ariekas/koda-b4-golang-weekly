package main

import (
	"fmt"
	"os"
	"weekly/internal/model"
	"weekly/internal/service"
)

func main() {
	defer func() {
		r := recover()

		if r != nil {
			fmt.Println("Error :", r)
			os.Exit(1)
		}
	}()

	db := &model.Db{
		Transactions: []model.Transaction{},
		Orders:       []model.Order{},
	}
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
			service.OrderService(db)
		case "2":
			service.HistoryOrder(db)
		// case "3":
		// 	service.Opsion()
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
