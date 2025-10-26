package main

import (
	"fmt"
	"os"
	"weekly/data"
	"weekly/src/service"
)

func main() {
	defer func() {
		r := recover()

		if r != nil {
			fmt.Println("Error :", r)
			os.Exit(1)
		}
	}()

	// Var untuk mengambil struct Db dari package service
	var d *service.Db = &service.Db{
		Transactions: []data.Transaction{},
		Orders: []data.Order{},
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
			d.OrderService()
		case "2":
			d.HistoryOrder()
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
