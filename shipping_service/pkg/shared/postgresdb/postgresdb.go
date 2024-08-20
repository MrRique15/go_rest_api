package postgresdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectPostgresDB(connectionString string, dbname string) *sql.DB  {
    db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS shipping (
		_id SERIAL PRIMARY KEY,
        shipping_type VARCHAR(255),
        order_id VARCHAR(255),
        shipping_price NUMERIC
    );`

	_, err = db.Exec(createTableQuery)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgresDB: ", dbname)

	return db
}