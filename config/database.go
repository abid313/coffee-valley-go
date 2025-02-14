package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// DBConn adalah fungsi untuk menghubungkan aplikasi dengan database MySQL
// Fungsi ini akan membuka koneksi ke database dan mengembalikan objek db yang digunakan untuk berinteraksi dengan database.
func DBConn() (db *sql.DB, err error) {
	// Menentukan jenis driver database yang digunakan (MySQL)
	dbDriver := "mysql"

	// Menyediakan kredensial untuk mengakses database
	dbUser := "root"      // Username untuk mengakses database
	dbPass := "root"      // Password untuk mengakses database
	dbName := "webdataon" // Nama database yang akan digunakan

	// Membuka koneksi ke database MySQL dengan kredensial yang diberikan
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	// Mengembalikan objek db dan error (jika ada)
	return db, err
}
