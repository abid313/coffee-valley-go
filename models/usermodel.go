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

// fungsi Register untuk menyimpan data user baru ke dalam database
func (u UserModel) Register(user *entities.User) error {
	_, err := u.db.Exec("INSERT INTO users VALUES (null, ?, ?, ?, ?)", user.NamaLengkap, user.Email, user.Username, user.Password)
	if err != nil {
		return err // Mengembalikan error jika terjadi kesalahan saat eksekusi query
	}
	return nil // Mengembalikan nil jika tidak ada error, artinya data berhasil disimpan
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

func (u *UserModel) UpdateDistributor(id string, name string, city string, region string, country string, phone string, email string) error {
	query := `UPDATE distributors SET Nama = ?, City = ?, Region = ?, Country = ?, Phone = ?, Email = ? WHERE Id = ?`
	_, err := u.db.Exec(query, name, city, region, country, phone, email, id)
	if err != nil {
		return fmt.Errorf("error updating distributor: %w", err)
	}
	return nil
}

func (u *UserModel) DeleteDistributor(id string) error {
	query := `DELETE FROM distributors WHERE Id = ?`
	_, err := u.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting distributor: %w", err)
	}
	return nil
}

func (u *UserModel) GetAllOrderStatus(orderStatus *entities.OrderStatus) ([]entities.OrderStatus, error) {
	row, err := u.db.Query("select * from orderstatus")

	if err != nil {
		return nil, err
	}

	defer row.Close()

	var orderStatusAll []entities.OrderStatus
	for row.Next() {
		err := row.Scan(&orderStatus.Id, &orderStatus.Bean, &orderStatus.Price, &orderStatus.Quantity, &orderStatus.Total, &orderStatus.Status)
		if err != nil {
			return nil, err
		}
		fmt.Println(orderStatus) // Menampilkan data dari database
		orderStatusAll = append(orderStatusAll, *orderStatus)
	}

	return orderStatusAll, nil
}

func (u *UserModel) AddOrderStatus(bean string, price float64, quantity int, total float64, status string) error {
	// Validate input parameters
	if bean == "" || price < 0 || quantity < 0 || total < 0 || status == "" {
		return fmt.Errorf("invalid input parameters")
	}

	query := `INSERT INTO orderstatus (Bean, Price, Quantity, Total, Status) VALUES (?, ?, ?, ?, ?)`
	_, err := u.db.Exec(query, bean, price, quantity, total, status)
	if err != nil {
		return fmt.Errorf("error adding order status: %w", err)
	}
	return nil
}

func (u *UserModel) DeleteOrderStatus(id string) error {
	query := `DELETE FROM orderstatus WHERE Id = ?`
	_, err := u.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting order status: %w", err)
	}
	return nil
}

func (u *UserModel) GetOrderStatusById(id string) (*entities.OrderStatus, error) {
	query := `SELECT * FROM orderstatus WHERE Id = ?`
	row := u.db.QueryRow(query, id)

	orderStatus := &entities.OrderStatus{}
	err := row.Scan(&orderStatus.Id, &orderStatus.Bean, &orderStatus.Price, &orderStatus.Quantity, &orderStatus.Total, &orderStatus.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no order status found with id %s", id)
		}
		return nil, fmt.Errorf("error fetching order status: %w", err)
	}

	return orderStatus, nil
}

func (u *UserModel) UpdateOrderStatus(id string, bean string, price float64, quantity int, total float64, status string) error {
	query := `UPDATE orderstatus SET Bean = ?, Price = ?, Quantity = ?, Total = ?, Status = ? WHERE Id = ?`
	_, err := u.db.Exec(query, bean, price, quantity, total, status, id)
	if err != nil {
		return fmt.Errorf("error updating order status: %w", err)
	}
	return nil
}

func (u *UserModel) MarkDoneOrderStatus(id string) error {
	query := `UPDATE orderstatus SET Status = 'Done' WHERE Id = ?`
	_, err := u.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error marking order status as done: %w", err)
	}
	return nil
}
