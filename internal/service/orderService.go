package service

import (
	"fmt"
	"math/rand/v2"
	"time"
	"weekly/internal/model"
	"weekly/internal/repository"
)

func OrderService(db *model.Db) {
	var choise, quantity int
	var menus = repository.GetData()

	for {
		fmt.Print("\x1bc")
		fmt.Println("=== Menu ===")
		for i, menu := range menus {
			model.Print(menu, i)
		}

		fmt.Println("\n \n98. Cari produk")
		fmt.Println("99. Order")
		fmt.Println("0. Kembali")

		fmt.Print("Masukan No Menu yang di pilih! ")
		fmt.Scan(&choise)
		if choise == 99 {
			break
		}

		if choise == 98 {
			var keyword string
			fmt.Print("Masukkan kata kunci pencarian: ")
			fmt.Scan(&keyword)

			results := repository.SearchMenu(menus, keyword)
			if len(results) == 0 {
				fmt.Println("Tidak ada hasil ditemukan.")
			} else {
				fmt.Println("=== Hasil Pencarian ===")
				for i, menu := range results {
					model.Print(menu, i)
				}
			}
			fmt.Println("\nTekan ENTER untuk melanjutkan...")
			fmt.Scanln()
			fmt.Scanln()
			continue
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
			fmt.Printf("Jumlah yang dimasukan melebihi batas maksimal 50, %d\n", quantity)
			fmt.Println("Tekan ENTER untuk melanjutkan...")
			fmt.Scanln()
			fmt.Scanln()
			continue
		}

		if quantity < 1 {
			fmt.Printf("Jumlah tidak boleh kurang dari 1, %d\n", quantity)
			fmt.Println("Tekan ENTER untuk melanjutkan...")
			fmt.Scanln()
			fmt.Scanln()
			continue
		}

		selecteMenu := menus[choise-1]

		// Validasi stok
		if quantity > selecteMenu.Stock {
			fmt.Printf("Stok tidak cukup! Stok tersedia: %d\n", selecteMenu.Stock)
			fmt.Println("Tekan ENTER untuk melanjutkan...")
			fmt.Scanln()
			fmt.Scanln()
			continue
		}

		order := model.Order{
			Item:     selecteMenu,
			Quantity: quantity,
		}

		db.Orders = append(db.Orders, order)

		fmt.Println(selecteMenu.Name+" Dibeli sebanyak ", quantity)
		fmt.Println("Tekan ENTER untuk melanjutkan...")
		fmt.Scanln()
		fmt.Scanln()

	}
	CheckoutService(db)
}

func CheckoutService(db *model.Db) {
	var choiseMenu int
	var quantity, nomerOrder int
	var custumer string

	for {
		fmt.Print("\x1bc")

		fmt.Println("=== Detail Pesanan ===")
		if len(db.Orders) == 0 {
			fmt.Println("Tidak ada data order")
		} else {
			var totalAll float64
			for i, menu := range db.Orders {
				model.Print(menu, i)
				totalAll += menu.Item.Price * float64(menu.Quantity)
			}
			fmt.Printf("Total: Rp %.0f\n", totalAll)
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
			time.Sleep(1 * time.Second)
			return
		}

		switch choiseMenu {
		case 1:
			fmt.Print("Masukan Nomer pesanan: ")
			fmt.Scan(&nomerOrder)

			if nomerOrder < 1 || nomerOrder > len(db.Orders) {
				fmt.Println("Nomor pesanan tidak valid.")
				fmt.Scanln()
				fmt.Scanln()
				continue
			}

			db.Orders = append(db.Orders[:nomerOrder-1], db.Orders[nomerOrder:]...)
			fmt.Println("Pesanan berhasil di hapus !")
			fmt.Println("Tekan ENTER untuk melanjutkan...")
			fmt.Scanln()
			fmt.Scanln()

		case 2:
			fmt.Print("Masukan nomor pesanan : ")
			fmt.Scan(&nomerOrder)

			if nomerOrder < 1 || nomerOrder > len(db.Orders) {
				fmt.Println("Nomor pesanan tidak valid.")
				fmt.Scanln()
				fmt.Scanln()
				continue
			}

			fmt.Print("Masukan jumlah baru untuk pesanan ini: ")
			fmt.Scan(&quantity)

			// Validasi stok
			if quantity > db.Orders[nomerOrder-1].Item.Stock {
				fmt.Printf("Stok tidak cukup! Stok tersedia: %d\n", db.Orders[nomerOrder-1].Item.Stock)
				fmt.Scanln()
				fmt.Scanln()
				continue
			}

			if quantity < 1 {
				fmt.Println("Jumlah tidak boleh kurang dari 1")
				fmt.Scanln()
				fmt.Scanln()
				continue
			}

			db.Orders[nomerOrder-1].Quantity = quantity
			fmt.Println("Quantity berhasil diperbarui!")
			fmt.Println("Tekan ENTER untuk melanjutkan...")
			fmt.Scanln()
			fmt.Scanln()

		case 3:
			if len(db.Orders) == 0 {
				fmt.Println("Tidak ada pesanan untuk diproses!")
				fmt.Println("Tekan ENTER untuk melanjutkan...")
				fmt.Scanln()
				fmt.Scanln()
				continue
			}

			fmt.Print("\x1bc")
			fmt.Print("Masukan nama pelanggan : ")
			fmt.Scan(&custumer)

			var total float64
			for _, o := range db.Orders {
				total += o.Item.Price * float64(o.Quantity)
			}

			dateStr := time.Now().Format("020106")
			randomNum := rand.IntN(900) + 100
			orderID := fmt.Sprintf("Ord-%s-%d-%s", dateStr, randomNum, custumer)

			transaction := model.Transaction{
				OrderID:   orderID,
				Custemer:  custumer,
				Order:     db.Orders,
				Total:     total,
				DateOrder: time.Now(),
			}

			// Simpan transaksi ke database
			err := repository.CreateOrder(transaction)
			if err != nil {
				fmt.Printf("Error menyimpan transaksi: %v\n", err)
				fmt.Println("Tekan ENTER untuk kembali...")
				fmt.Scanln()
				fmt.Scanln()
				return
			}

			// Simpan ke slice lokal juga (opsional)
			db.Transactions = append(db.Transactions, transaction)
			db.Orders = nil

			fmt.Print("\x1bc")
			fmt.Println("\n=== Transaksi Berhasil! ===")
			fmt.Printf("Order ID: %s\nNama: %s\nTotal: Rp %.0f\nTanggal: %s\n",
				transaction.OrderID, transaction.Custemer, transaction.Total,
				transaction.DateOrder.Format("02-Jan-2006 15:04"))
			fmt.Println("Transaksi berhasil disimpan ke database!")
			fmt.Println("============================")
			fmt.Println("Tekan ENTER untuk kembali ke menu utama...")
			fmt.Scanln()
			fmt.Scanln()
			return
		default:
			fmt.Println("Invalid Input")
			fmt.Scanln()
			fmt.Scanln()
		}
	}
}

func HistoryOrder(db *model.Db) {
	var choise, choiseMenu int

	fmt.Print("\x1bc")
	fmt.Println("=== History Order ===")
	if len(db.Transactions) == 0 {
		fmt.Println("Belum ada transaksi.")
		fmt.Scanln()
		return
	}

	for i, transaction := range db.Transactions {
		fmt.Printf("%d %s - Tanggal: %s \n Nama: %s | Total: Rp %.0f \n\n",
			i+1, transaction.OrderID, transaction.DateOrder.Format("02-Jan-2006 15:04"), transaction.Custemer, transaction.Total)
	}
	fmt.Println("1. Cetak Struk")
	fmt.Println("0. Kembali")

	fmt.Print("\nPilih Menu: ")
	fmt.Scan(&choiseMenu)
	if choiseMenu == 1 {
		fmt.Print("\nPilih nomor transaksi untuk cetak struk: ")

		fmt.Scan(&choise)

		if choise == 0 {
			return
		}

		if choise < 1 || choise > len(db.Transactions) {
			fmt.Println("Nomor transaksi tidak valid.")
			fmt.Println("Tekan ENTER untuk melanjutkan...")
			fmt.Scanln()
		}

		selectedTransaction := db.Transactions[choise-1]
		fmt.Print("\x1bc")
		PrintStruk(selectedTransaction)

		fmt.Println("\nTekan ENTER untuk kembali ke menu utama...")
		fmt.Scanln()
		return
	}

}

