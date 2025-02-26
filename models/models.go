package models

type Product struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type Order struct {
	Id     string    `json:"id"`
	Status string    `json:"status"`
	Item   []Product `json:"item"`
}
