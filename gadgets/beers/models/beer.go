package models

// model struct for Beer
type Beer struct {
	Id				int64 // ID autoincremental in SQL
	Name			string
	Price			float32
	Country			string
}

type CreateBeerCMD struct {
	Name			string	`json:"name"`
	Price			float32 `json:"price"`
	Country			string	`json:"country"`
}