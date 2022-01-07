package models

import "errors"


const maxLengthString = 150
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

func (cmd *CreateBeerCMD) validate() error {
	if cmd.Price < 1 {
		return errors.New("price is lower")
	}

	if len(cmd.Name) > maxLengthString || len(cmd.Country) > maxLengthString {
		return errors.New("characters must be less than 150 chars")
	}
	return nil

}