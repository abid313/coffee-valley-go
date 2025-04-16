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

type OrderStatus struct {
	Id       int64
	Bean     string
	Price    float64
	Quantity int64
	Total    float64
	Status   string
}
