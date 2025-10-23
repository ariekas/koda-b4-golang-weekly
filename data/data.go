package data

import "time"

type MenuItem struct{
	ID int
	Name string
	Price float64
	Category string
}

// Tempat untuk menyimpan menu 	
var Menus = []MenuItem{
	{ID: 1, Name: "Nasi", Price: 5000, Category: "Makanan"},
	{ID: 2, Name: "Teh", Price: 2000, Category: "Minuman"},
	{ID: 3, Name: "Sate", Price: 10000, Category: "Makanan"},
}

// Menyimpan data pesanan user yang berisi item pesanan user dan juga input quantity
type Order struct{
	Item MenuItem
	Quantity int
}

// Data penyimpanan untuk semua transaksi
type Transaction struct{
	OrderID string
	Custemer string
	Order []Order
	Total float64
	DateOrder time.Time
}

// Slice global yang menyimpan semua transaksi yang pernah di lakukan
var Transactions []Transaction
