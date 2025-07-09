package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/joho/godotenv"
)

// DBConn adalah fungsi untuk menghubungkan aplikasi dengan database MySQL
func DBConn() (db *sql.DB, err error) {
	// Load .env file
	// err = godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Gagal load file .env")
	// 	return nil, err
	// }

	// Ambil variabel dari environment
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Format DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Buka koneksi database
	db, err = sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}

	// Test koneksi
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
