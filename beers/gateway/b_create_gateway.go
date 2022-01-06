package gateway

import (
	"github.com/TobiasSchnizel/Beer-API/beers/models"
	"github.com/TobiasSchnizel/Beer-API/internal/database"
)

type BeerCreateGateway interface {
	Create(cmd *models.CreateBeerCMD) (*models.Beer, error)
}
type BeerCreateGtw struct {
	BeerStorageGateway
}

func NewBeerCreateGateway(client *database.MySqlClient) BeerCreateGateway {
	return &BeerCreateGtw{&SBeerStorage{client}}
}
