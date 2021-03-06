package gateway

import (
	"github.com/TobiasSchnizel/Beer-API/internal/database"
	"github.com/TobiasSchnizel/Beer-API/reviews/models"
)

type ReviewGateway interface {
	AddReview(cmd *models.CreateReviewCMD) (string, error)
}

type ReviewGtw struct {
	ReviewStorage
}

func NewReviewGateway(mongoClient *database.Mongo) ReviewGateway {
	return &ReviewGtw{&ReviewStg{mongoClient}}
}
