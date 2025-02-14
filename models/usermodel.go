package models

import (
	"WebsiteDataOn/config"
	"WebsiteDataOn/entities"
	"database/sql"
	"fmt"
)

type UserModel struct {
	db *sql.DB // Mendeklarasikan field db yang akan digunakan untuk koneksi ke database
}

// Fungsi NewUserModel bertujuan untuk membuat objek UserModel baru dan menghubungkannya ke database
func NewUserModel() *UserModel {
	// Memanggil fungsi DBConn dari config untuk mendapatkan koneksi database
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}
	// Mengembalikan objek UserModel dengan koneksi database yang sudah terbentuk
	return &UserModel{
		db: conn,
	}
}

// Fungsi Where mencari data user berdasarkan field tertentu dan nilai field tersebut
func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {
	// Menjalankan query SQL untuk mengambil data user berdasarkan nama field dan nilainya
	// Limit 1 digunakan untuk mengambil hanya satu hasil pencarian
	row, err := u.db.Query("select id, nama, email, username, password from users where "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	// Jika query mengembalikan baris data, baca hasilnya dan masukkan ke dalam objek user
	// row.Next() akan mengembalikan true jika ada baris data untuk dibaca
	for row.Next() {
		// Memindahkan data dari row ke dalam field objek user
		row.Scan(&user.Id, &user.NamaLengkap, &user.Email, &user.Username, &user.Password)
	}

	// Kembalikan nil karena tidak ada error yang terjadi
	return nil
}

func (u UserModel) IndexHome(catalog *entities.Catalog) (*entities.Catalog, error) {
	row, err := u.db.Query("select * from beans where price >= 0.0 limit 1")

	if err != nil {
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		err := row.Scan(&catalog.Bean, &catalog.Description, &catalog.Price)
		if err != nil {
			return nil, err
		}
		fmt.Println(catalog) // Menampilkan data dari database
	}

	return catalog, nil
}

// Fungsi GetAllCatalog untuk mengambil semua data di table beans
func (u UserModel) GetAllCatalog(catalog *entities.Catalog) ([]entities.Catalog, error) {
	row, err := u.db.Query("select * from beans")

	if err != nil {
		return nil, err
	}

	defer row.Close()

	var catalogAll []entities.Catalog
	for row.Next() {
		err := row.Scan(&catalog.Bean, &catalog.Description, &catalog.Price)
		if err != nil {
			return nil, err
		}
		fmt.Println(catalog) // Menampilkan data dari database
		catalogAll = append(catalogAll, *catalog)
	}

	return catalogAll, nil
}

// Fungsi GetAllCatalog untuk mengambil semua data di table beans
func (u UserModel) GetAllDistributor(distributor *entities.Distributor) ([]entities.Distributor, error) {
	row, err := u.db.Query("select id,nama,city from distributors")

	if err != nil {
		return nil, err
	}

	defer row.Close()

	var distributorAll []entities.Distributor
	for row.Next() {
		err := row.Scan(&distributor.Id, &distributor.Nama, &distributor.City)
		if err != nil {
			return nil, err
		}
		fmt.Println(distributor) // Menampilkan data dari database
		distributorAll = append(distributorAll, *distributor)
	}

	return distributorAll, nil
}

func (u UserModel) AddDistributor(fieldName, fieldCity, fieldRegion, fieldCountry, fieldPhone, fieldEmail string) error {
	_, err := u.db.Exec("INSERT INTO distributors VALUES (null,?, ?, ?, ?, ?, ?)", fieldName, fieldCity, fieldRegion, fieldCountry, fieldPhone, fieldEmail)
	if err != nil {
		return err
	}

	return nil
}

func (u UserModel) EditDetailDistributor(fieldId string) (distributor *entities.Distributor, err error) {
	row, err := u.db.Query("SELECT * FROM distributors WHERE id = ?", fieldId)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	// Initialize the distributor to avoid nil pointer dereference
	distributor = &entities.Distributor{}

	for row.Next() {
		// Pass the fields directly into Scan without naming them
		err = row.Scan(&distributor.Id, &distributor.Nama, &distributor.City, &distributor.Region, &distributor.Country, &distributor.Phone, &distributor.Email)
		if err != nil {
			return nil, err
		}
	}

	// Check for any errors from iterating over the rows
	if err = row.Err(); err != nil {
		return nil, err
	}

	return distributor, nil
}
