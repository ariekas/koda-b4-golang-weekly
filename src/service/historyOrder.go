package service

import (
	"fmt"
)

func (d *Db) HistoryOrder() {
	var choise, choiseMenu int

	fmt.Print("\x1bc")
	fmt.Println("=== History Order ===")
	if len(d.Transactions) == 0 {
		fmt.Println("Belum ada transaksi.")
		fmt.Scanln()
		return
	}

	for i, transaction := range d.Transactions {
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

		if choise < 1 || choise > len(d.Transactions) {
			fmt.Println("Nomor transaksi tidak valid.")
			fmt.Println("Tekan ENTER untuk melanjutkan...")
			fmt.Scanln()
		}

		selectedTransaction := d.Transactions[choise-1]
		fmt.Print("\x1bc")
		PrintStruk(selectedTransaction)

		fmt.Println("\nTekan ENTER untuk kembali ke menu utama...")
		fmt.Scanln()
		return
	}

}
