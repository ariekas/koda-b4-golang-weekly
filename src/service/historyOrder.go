package service

import (
	"fmt"
	"weekly/data"
	"weekly/utils"
)

func HistoryOrder(){
	utils.ClearScreen()
	fmt.Println("=== History Order ===")
	if len(data.Transactions) == 0 {
		fmt.Println("Belum ada transaksi.")
		fmt.Scanln()
		return
	}

	for i, transaction := range data.Transactions {
		fmt.Printf("%d %s - Tanggal: %s \n Nama: %s | Total: Rp %.0f \n\n",
	i+1, transaction.OrderID,transaction.DateOrder.Format("02-Jan-2006 15:04"),  transaction.Custemer, transaction.Total)
	}
	fmt.Println("\nTekan ENTER untuk kembali...")
	fmt.Scanln()
	return
}