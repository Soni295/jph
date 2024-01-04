package model

type UserId uint64

type User struct {
	Id       UserId `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`

	Address Address `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	ZipCode int16  `json:"zipcode"`

	Geo Geo `json:"geo"`
}

type Geo struct {
	Latitude float64 `json:"lat"`
	Longitud float64 `json:"lon"`
}