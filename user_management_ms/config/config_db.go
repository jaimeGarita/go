package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "user_management"
)

var DB *sql.DB

func ConnectDB() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user, password, host, port, dbname)

	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("❌ Error conectando a la base de datos:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ No se pudo conectar a la base de datos:", err)
	}

	fmt.Println("✅ Conectado a la base de datos")
}
