package gateway

import (
	"github.com/TobiasSchnizel/Beer-API/gadgets/beers/models"
	"github.com/TobiasSchnizel/Beer-API/internal/database"
	"github.com/TobiasSchnizel/Beer-API/internal/logs"
)

type BeerStorageGateway interface {
	Add(cmd *models.CreateBeerCMD) (*models.Beer, error)
	//FindById(id int) (*models.Beer, error)
}
type BeerStorage struct {
	*database.MySqlClient
}

func (s *BeerStorage) BeerCreateGateway(cmd *models.CreateBeerCMD) (*models.Beer, error) {
	tx, err := s.MySqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
		return nil, err
	}

	res, err := tx.Exec(`insert into beer (name, price, country)
	values (?, ?, ?)`, cmd.Name, cmd.Price, cmd.Country)

	if err != nil {
		logs.Log().Error("cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()
	// rowsAf, err := res.RowsAffected()

	if err != nil {
		logs.Log().Error("cannot fetch last id")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &models.Beer{
		Id:      id,
		Name:    cmd.Name,
		Price:   cmd.Price,
		Country: cmd.Country,
	}, nil
}
