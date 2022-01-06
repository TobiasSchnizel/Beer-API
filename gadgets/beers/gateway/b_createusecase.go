package gateway

import (
	"github.com/TobiasSchnizel/Beer-API/gadgets/beers/models"
)

func (s *BeerCreateGtw) create(cmd *models.CreateBeerCMD) (*models.Beer, error) {
	return s.create(cmd)
}