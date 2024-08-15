package domain

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Address struct {
	Jalan  string `json:"jalan"`
	Alamat string `json:"alamat"`
}
