package repository

import (
	"context"
	"fmt"
	"time"
	"weekly/internal/model"

	"github.com/jackc/pgx/v5"
)

func CreateOrder(t model.Transaction) error {
	conn := ConnectDb()

	defer conn.Close(context.Background())

	var transactionID int
	err := conn.QueryRow(context.Background(),
		`
INSERT INTO transactions (order_id, customer_name, total, order_date)
VALUES ($1, $2, $3, $4)
RETURNING id
`, t.OrderID, t.Custemer, t.Total, t.DateOrder).Scan(&transactionID)

	if err != nil {
		fmt.Printf("failed to insert transaction: %v", err)
	}

	for _, order := range t.Order {
		_, err = conn.Exec(context.Background(),
			`INSERT INTO order_items (transaction_id, product_id, product_name, price, quantity)
	VALUES ($1, $2, $3, $4, $5)`,
			transactionID, order.Item.Id, order.Item.Name, order.Item.Price, order.Quantity,
		)
		if err != nil {
			fmt.Printf("failed to insert order item: %v", err)
		}

		_, err = conn.Exec(context.Background(),
			`UPDATE products SET stock = stock - $1, updated_at = NOW() WHERE id = $2`, order.Quantity, order.Item.Id)

		if err != nil {
			fmt.Printf("failed to update product stock: %v", err)
		}
	}

	return nil

}

func GetOrderById(orderId string) (*model.Transaction, error) {
	conn := ConnectDb()

	defer conn.Close(context.Background())

	var transaction model.Transaction
	var transactionID int

	err := conn.QueryRow(context.Background(),
		`
SELECT id, order_id, customer_name, total, order_date 
FROM transactions
WHERE order_id = $1
`, orderId).Scan(&transactionID, &transaction.OrderID, &transaction.Custemer, &transaction.Total, &transaction.DateOrder)

	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Printf("transaction not found")
		}
		fmt.Printf("failed to query transaction: %v", err)
	}

	orderRows, err := conn.Query(context.Background(),
		`
SELECT oi.product_id, oi.product_name, oi.price, oi.quantity, p.stock, p.description, p.created_at, p.updated_at
FROM order_items oi
LEFT JOIN product p ON oi.product_id = p.id
WHERE oi.transaction_id = $1
`, transactionID)

	if err != nil {
		fmt.Printf("failed to query order items: %v", err)
	}
	defer orderRows.Close()

	var orders []model.Order

	for orderRows.Next() {
		var order model.Order
		var stock int
		var description string
		var createdAt, updatedAt time.Time

		err := orderRows.Scan(
			&order.Item.Id,
			&order.Item.Name,
			&order.Item.Price,
			&order.Quantity,
			&stock,
			&description,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			fmt.Printf("failed to scan order item: %v", err)
		}

		order.Item.Stock = stock
		order.Item.Description = description
		order.Item.Created_at = createdAt
		order.Item.Updated_At = updatedAt

		orders = append(orders, order)
	}

	transaction.Order = orders
	return &transaction, nil

}

func GetOrderHistory() ([]model.Transaction, error) {
	conn := ConnectDb()

	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		`SELECT id, order_id, customer_name, total, order_date
FROM transactions
ORDER BY order_date DESC`)

	if err != nil {
		fmt.Printf("failed to query transactions: %v", err)
	}

	defer rows.Close()

	var transactions []model.Transaction

	for rows.Next() {
		var transaction model.Transaction
		var transactionID int

		err = rows.Scan(
			&transactionID,
			&transaction.OrderID,
			&transaction.Custemer,
			&transaction.Total,
			&transaction.DateOrder,
		)

		if err != nil {
			fmt.Printf("failed to scan transaction: %v", err)
		}

		orderRows, err := conn.Query(context.Background(),
			`SELECT oi.product_id, oi.product_name, oi.price, oi.quantity, p.stock, p.description, p.created_at, p.updated_at
		 FROM order_items oi
		 LEFT JOIN product p ON oi.product_id = p.id
		 WHERE oi.transaction_id = $1`,
			transactionID,
		)

		if err != nil {
			 fmt.Printf("failed to query order items: %v", err)
		}

		var orders []model.Order
		for orderRows.Next() {
			var order model.Order
			var stock int
			var description string
			var createdAt, updatedAt time.Time

			err = orderRows.Scan(
				&order.Item.Id,
				&order.Item.Name,
				&order.Item.Price,
				&order.Quantity,
				&stock,
				&description,
				&createdAt,
				&updatedAt,
			)

			if err != nil {
				orderRows.Close()
				fmt.Printf("failed to scan order item: %v", err)
			}

			order.Item.Stock = stock
			order.Item.Description = description
			order.Item.Created_at = createdAt
			order.Item.Updated_At = updatedAt

			orders = append(orders, order)
		}
		orderRows.Close()

		transaction.Order = orders
		transactions = append(transactions, transaction)
	}
	return transactions, nil

}
