package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetData() ([]MenuItem, error){
	resp, err := http.Get("https://raw.githubusercontent.com/ariekas/koda-b4-golang-weekly-data/refs/heads/main/dataProduct.json")

	if err != nil {
		fmt.Println("Error: Failed to Fecth data")
	}

	body, err := io.ReadAll(
		resp.Body,
	)

	var menus []MenuItem

	if err != nil {
		fmt.Println("Failed to raid body")
	}

	err = json.Unmarshal(body, &menus)
	if err != nil {
		fmt.Println("Failed to converd data")
	}

	return menus, nil
}
type MenuItem struct{
	ID int
	Name string
	Price float64
	Category string
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


