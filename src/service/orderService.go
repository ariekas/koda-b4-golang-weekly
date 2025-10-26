package service

import (
	"fmt"
	"math/rand"
	"time"
	"weekly/data"
)

func OrderService() {
	var choise, quantity int

	for {
		fmt.Print("\x1bc")
		fmt.Println("=== Menu ===")
		for i, menu := range data.Menus {
			fmt.Printf("%d. %s - Rp %.0f \n", i+1, menu.Name, menu.Price)
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

		if choise < 1 || choise > len(data.Menus) {
			fmt.Println("Nomor menu tidak tersedia!")
			fmt.Scanln()
			continue
		}

		fmt.Print("Masukan jumlah yang di beli! ")
		fmt.Scan(&quantity)

		selecteMenu := data.Menus[choise-1]

		order := data.Order{
			Item:     selecteMenu,
			Quantity: quantity,
		}

		data.Orders = append(data.Orders, order)

		fmt.Println(selecteMenu.Name+" Dibeli sebanyak ", quantity)
		fmt.Scanln()
	}
	CheckoutService()
}

func CheckoutService() {
	var choiseMenu int
	var quantity, nomerOrder int
	var custumer string

	fmt.Print("\x1bc")
	for {
		fmt.Print("\x1bc")

		fmt.Println("=== Detail Pesanan ===")
		for i, menu := range data.Orders {
			fmt.Printf("%d.\nProduct: %s\nPrice: Rp %.0f\nQuantity: %d\n\n",
				i+1, menu.Item.Name, menu.Item.Price, menu.Quantity)
		}
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

			if nomerOrder < 1 || nomerOrder > len(data.Orders) {
				fmt.Println("Nomor pesanan tidak valid.")
				fmt.Scanln()
				continue
			}

			data.Orders = append(data.Orders[:nomerOrder-1], data.Orders[nomerOrder:]...)
			fmt.Println("Pesanan berhasil di hapus !")
		case 2:
			fmt.Print("Masukan nomor pesanan : ")
			fmt.Scan(&nomerOrder)

			if nomerOrder < 1 || nomerOrder > len(data.Orders) {
				fmt.Println("Nomor pesanan tidak valid.")
				fmt.Scanln()
				continue
			}

			fmt.Print("Masukan jumlah baru untuk pesanan ini: ")
			fmt.Scan(&quantity)

			data.Orders[nomerOrder-1].Quantity = quantity
			fmt.Println("Quantity berhasil diperbarui!")
		case 3:

			fmt.Print("\x1bc")

			fmt.Print("Masukan nama pelanggan : ")
			fmt.Scan(&custumer)
			var total float64
			for _, o := range data.Orders {
				total += o.Item.Price * float64(o.Quantity)
			}

			dateStr := time.Now().Format("020106")
			randomNum := rand.Intn(900) + 100
			orderID := fmt.Sprintf("Ord-%s-%d-%s", dateStr, randomNum, custumer)

			transaction := data.Transaction{
				OrderID:   orderID,
				Custemer:  custumer,
				Order:     data.Orders,
				Total:     total,
				DateOrder: time.Now(),
			}

			data.Transactions = append(data.Transactions, transaction)
			data.Orders = nil

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
