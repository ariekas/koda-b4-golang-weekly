package repository

import (
	"context"
	"fmt"
	"weekly/internal/model"
	"github.com/jackc/pgx/v5"
)


func GetData() []model.MenuItem{
	conn := ConnectDb()

	rows, err := conn.Query(context.Background(), 
`SELECT id, name, price, stock, description,  created_at, updated_at
FROM product
ORDER BY id
`)
if err != nil {
	fmt.Println("Error: Failed to query products")
}

defer rows.Close()

products, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.MenuItem])

	if err != nil {
		fmt.Println("Error: Failed to map product")
	}

	return products



}