package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type row struct {
	created      time.Time
	client_id int32
	value     float64
}

type queryRows struct {
  conn *pgx.Conn
	rows []row
}

func (query *queryRows) connect() {
  godotenv.Load(".env")
  var err error
  query.conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
    os.Exit(1)

    defer query.conn.Close(context.Background())
  }
}

func (query *queryRows) queryClose() {
  query.queryEmpty() 
  
  err := query.conn.Close(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}

func (query queryRows) queryPrint() {
  for _, row := range query.rows {
    fmt.Println(row.created, row.client_id, row.value)
  }
}

func (query *queryRows) queryEmpty() {
  query.rows = query.rows[:0]
}

func (query *queryRows) queryClient(client_id int) {

	rows, err := query.conn.Query(context.Background(), fmt.Sprintf("SELECT * FROM metrics WHERE client_id = %d;", client_id))
	if err != nil {
		_, err2 := fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		if err2 != nil {
			return
		}
		os.Exit(1)
	}

	for rows.Next() {
		values, err := rows.Values()

    query.rows = append(query.rows, row{values[0].(time.Time), values[1].(int32), values[2].(float64)})
		
    if err != nil {
			log.Fatal("error while iterating dataset")
		}
	}
}

func (query *queryRows) queryClientTime(client_id int, startTime time.Time, endTime time.Time) {
  querystr := "SELECT * FROM metrics WHERE created BETWEEN $1 AND $2;"

  rows, err := query.conn.Query(context.Background(), querystr, startTime, endTime)
  if err != nil {
    fmt.Println(err) 
    _, err2 := fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
    if err2 != nil {
      return
    }
    os.Exit(1)
  }
  for rows.Next() {
    values, err := rows.Values()
    query.rows = append(query.rows, row{values[0].(time.Time), values[1].(int32), values[2].(float64)})
    if err != nil {
      log.Fatal("error while iterating dataset")
		}
	}

}

func (query *queryRows) insert(client_id int, value float64) {
	sqlinsert := fmt.Sprintf("INSERT INTO metrics(client_id, value)  VALUES  (%d, %f);", client_id, value)
  _, err := query.conn.Query(context.Background(), sqlinsert)
	if err != nil {
		_, err2 := fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		if err2 != nil {
			return
		}
		os.Exit(1)
	}
}
