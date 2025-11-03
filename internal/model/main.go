package model

import (
	"fmt"
	"time"
)

type MenuItem struct {
	Id          int
	Name        string
	Price       float64
	Stock       int
	Description string
	Created_at  time.Time
	Updated_At  time.Time
}

// PrintProduct implements ShowData.
func (m MenuItem) PrintProduct(i int) {
	fmt.Printf("%d. %s - Rp %.0f \n", i+1, m.Name, m.Price)
}

type Order struct {
	Item     MenuItem
	Quantity int
}

// PrintProduct implements ShowData.
func (o Order) PrintProduct(i int) {
	fmt.Printf("%d.\nProduct: %s\nPrice: Rp %.0f\nQuantity: %d\n\n",
		i+1, o.Item.Name, o.Item.Price, o.Quantity)
}

type Transaction struct {
	OrderID   string
	Custemer  string
	Order     []Order
	Total     float64
	DateOrder time.Time
}

type Db struct {
	Transactions []Transaction
	Orders       []Order
}

type ShowData interface {
	PrintProduct(i int)
}

func Print(s ShowData, i int) {
	s.PrintProduct(i)
}


