package service

import (
	"fmt"
	"math/rand"
	"time"
	"weekly/data"
)

// Struct yang berisi var yang tipe data nya slice []data.Transaction/Order yang di dapat dari package data
type Db struct {
	Transactions []data.Transaction
	Orders       []data.Order
}

type ShowData interface{
	PrintProduct(i int)
}

func Print (s ShowData, i int){
	s.PrintProduct(i)
}

// Method
func (d *Db) OrderService() {
	var choise, quantity int
	var menus, error = data.GetData()

	if error != nil {
		fmt.Println("Error: Failed to getting data product")
	}

	for {
		fmt.Print("\x1bc")
		fmt.Println("=== Menu ===")
		for i, menu := range menus {
			Print(menu, i)
		}

		fmt.Println("\n \n99. Lanjut")
		fmt.Println("0. Kembali")

		fmt.Print("Masukan No Menu yang di pilih! ")
		fmt.Scan(&choise)
		if choise == 99 {
			break
		}
		if choise == 0 {
			fmt.Println("Kembali ke menu utama...")
			return
		}

		if choise < 1 || choise > len(menus) {
			fmt.Println("Nomor menu tidak tersedia!")
			fmt.Scanln()
			continue
		}

		fmt.Print("Masukan jumlah yang di beli! ")
		fmt.Scan(&quantity)

		if quantity > 50 {
			panic(fmt.Sprintf("Jumlah yang dimasukan melebihi batas, %d", quantity))
		}

		selecteMenu := menus[choise-1]

		order := data.Order{
			Item:     selecteMenu,
			Quantity: quantity,
		}

		d.Orders = append(d.Orders, order)

		fmt.Println(selecteMenu.Name+" Dibeli sebanyak ", quantity)
		fmt.Scanln()
	}
	d.CheckoutService()
}

// Method
func (d *Db) CheckoutService() {
	var choiseMenu int
	var quantity, nomerOrder int
	var custumer string

	fmt.Print("\x1bc")
	for {
		fmt.Print("\x1bc")

		fmt.Println("=== Detail Pesanan ===")
		for i, menu := range d.Orders {Print(menu, i)}
		fmt.Println(`
1. Menghapus pesanan
2. Edit Pesanan
3. Pesan
	
0. Kembali

Pilih Menu !
		`)
		fmt.Scan(&choiseMenu)

		if choiseMenu == 0 {

			fmt.Print("\x1bc")

			fmt.Println("Kembali ke menu home")

			return
		}
		switch choiseMenu {
		case 1:
			fmt.Println("Masukan Nomer pesanan: ")
			fmt.Scan(&nomerOrder)

			if nomerOrder < 1 || nomerOrder > len(d.Orders) {
				fmt.Println("Nomor pesanan tidak valid.")
				fmt.Scanln()
				continue
			}

			d.Orders = append(d.Orders[:nomerOrder-1], d.Orders[nomerOrder:]...)
			fmt.Println("Pesanan berhasil di hapus !")
		case 2:
			fmt.Print("Masukan nomor pesanan : ")
			fmt.Scan(&nomerOrder)

			if nomerOrder < 1 || nomerOrder > len(d.Orders) {
				fmt.Println("Nomor pesanan tidak valid.")
				fmt.Scanln()
				continue
			}

			fmt.Print("Masukan jumlah baru untuk pesanan ini: ")
			fmt.Scan(&quantity)

			d.Orders[nomerOrder-1].Quantity = quantity
			fmt.Println("Quantity berhasil diperbarui!")
		case 3:

			fmt.Print("\x1bc")

			fmt.Print("Masukan nama pelanggan : ")
			fmt.Scan(&custumer)
			var total float64
			for _, o := range d.Orders {
				total += o.Item.Price * float64(o.Quantity)
			}

			dateStr := time.Now().Format("020106")
			randomNum := rand.Intn(900) + 100
			orderID := fmt.Sprintf("Ord-%s-%d-%s", dateStr, randomNum, custumer)

			transaction := data.Transaction{
				OrderID:   orderID,
				Custemer:  custumer,
				Order:     d.Orders,
				Total:     total,
				DateOrder: time.Now(),
			}

			d.Transactions = append(d.Transactions, transaction)
			d.Orders = nil

			fmt.Print("\x1bc")

			fmt.Println("\n=== Transaksi Berhasil! ===")
			fmt.Printf("Order ID: %s\nNama: %s\nTotal: Rp %.0f\nTanggal: %s\n",
				transaction.OrderID, transaction.Custemer, transaction.Total,
				transaction.DateOrder.Format("02-Jan-2006 15:04"))
			fmt.Println("============================")
			fmt.Println("Tekan ENTER untuk kembali ke menu utama...")
			fmt.Scanln()
			return
		default:
			fmt.Println("Invalid Input")
			fmt.Scanln()
			return
		}
	}
}
