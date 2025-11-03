package service

import (
	"fmt"
	"weekly/internal/model"
)

type menu model.MenuItem
type order  model.Order

func (m menu) Print(i int){
	fmt.Printf("%d. %s - Rp %.0f \n", i+1, m.Name, m.Price)
}

func (o order) Print(i int){
	fmt.Printf("%d.\nProduct: %s\nPrice: Rp %.0f\nQuantity: %d\n\n",
		i+1, o.Item.Name, o.Item.Price, o.Quantity)
}