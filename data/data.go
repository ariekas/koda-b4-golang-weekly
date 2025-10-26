package data

import (
	"fmt"
	"time"
)

type MenuItem struct{
	ID int
	Name string
	Price float64
	Category string
}

var Menus = []MenuItem{
	{ID: 1, Name: "Nasi", Price: 5000, Category: "Makanan"},
	{ID: 2, Name: "Teh", Price: 2000, Category: "Minuman"},
	{ID: 3, Name: "Sate", Price: 10000, Category: "Makanan"},
}

type Order struct{
	Item MenuItem
	Quantity int
}

type Transaction struct{
	OrderID string
	Custemer string
	Order []Order
	Total float64
	DateOrder time.Time
}


func (m MenuItem) PrintProduct(i int){
	fmt.Printf("%d. %s - Rp %.0f \n", i+1, m.Name, m.Price)
}

func (o Order) PrintProduct(i int){
	fmt.Printf("%d.\nProduct: %s\nPrice: Rp %.0f\nQuantity: %d\n\n",
	i+1, o.Item.Name, o.Item.Price, o.Quantity)
}


