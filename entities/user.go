package entities

type User struct {
	Id          int64
	NamaLengkap string
	Email       string
	Username    string
	Password    string
}

type Catalog struct {
	Bean        string
	Description string
	Price       float32
}

type Distributor struct {
	Id      int64
	Nama    string
	City    string
	Region  string
	Country string
	Phone   string
	Email   string
}
