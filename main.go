package main

import (
	"fmt"
	"weekly/data"
)

func main() {
	fmt.Println("=== Welcome To Pizza ===")
	fmt.Println(`
	Menu :
	1. Menu


	0. Exit
	`)
	fmt.Print("Masukan menu:")
	var input string
	fmt.Scan(&input)
	switch input {
	case "1":
		Buyying()
	case "0":
		fmt.Print("Exiting Program")
		return
	default:
		fmt.Println("Invalid Input")
	}
}


func Buyying(){
	var choise string
	fmt.Println("=== Menu ===")
	for i, menu := range data.Menus {
		fmt.Printf("%d. %s - Rp %.0f \n", i+1, menu.Name, menu.Price)
	}	

	fmt.Println("0. Kembali")
	fmt.Print("Masukan No Menu yang di pilih! ")
	fmt.Scan(&choise)
	// Input quantity
	// Showing data yang di pesan
	// Bisa ubah quantity dan delete
}



