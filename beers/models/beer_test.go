package models

import "testing"

func NewBeer(name string, price float32, country string) *CreateBeerCMD {
	return &CreateBeerCMD{
		Name:   name,
		Price: 	price,
		Country: country,
	}
}
func Test_withCorrectParams(t *testing.T) {
	r := NewBeer("Paceña", 1.5, "Amazing")

	err := r.validate()

	if err != nil {
		t.Error("the validation did not past")
		t.Fail()
	}
}