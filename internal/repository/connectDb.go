package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnectDb() *pgx.Conn {
	godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")

	conn,err := pgx.Connect(context.Background(), dbURL)

	if err != nil {
		fmt.Printf("Error: Cant connect to database %v", err)
	}
	
	return conn

}