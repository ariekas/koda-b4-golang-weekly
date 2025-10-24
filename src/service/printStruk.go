package service

import (
	"fmt"
	"weekly/data"
)

func PrintStruk(transaction data.Transaction) {
	fmt.Println("========== STRUK PEMBELIAN ==========")
	fmt.Printf("Order ID : %s\nNama     : %s\nTanggal  : %s\n",
		transaction.OrderID, transaction.Custemer, transaction.DateOrder.Format("02-Jan-2006 15:04"))
	fmt.Println("-------------------------------------")
	for _, o := range transaction.Order {
		fmt.Printf("%-20s x%d  Rp %.0f\n", o.Item.Name, o.Quantity, o.Item.Price*float64(o.Quantity))
	}
	fmt.Println("-------------------------------------")
	fmt.Printf("TOTAL : Rp %.0f\n", transaction.Total)
	fmt.Println("=====================================")
}