package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type row struct {
	time      string
	client_id int
	value     float32
}

type query struct {
	rows []row
}

func insertDB(conn *pgx.Conn) {

}

func queryClient(client_id int, limit int) {
	godotenv.Load(".env")
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), fmt.Sprintf("SELECT * FROM metrics WHERE client_id = %d LIMIT %d;", client_id, limit))
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		values, err := rows.Values()
    fmt.Println(values[0], " ", values[2])
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
	}

  conn.Close(context.Background())
}
