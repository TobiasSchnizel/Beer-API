package gateway

import (
	"github.com/TobiasSchnizel/Beer-API/beers/models"
)

func (s *BeerCreateGtw) Create(cmd *models.CreateBeerCMD) (*models.Beer, error) {
	return s.create(cmd)
}