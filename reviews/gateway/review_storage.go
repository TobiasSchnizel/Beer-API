package gateway

import (
	"context"
	"github.com/TobiasSchnizel/Beer-API/internal/database"
	"github.com/TobiasSchnizel/Beer-API/internal/logs"
	"github.com/TobiasSchnizel/Beer-API/reviews/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReviewStorage interface {
	Add(cmd *models.CreateReviewCMD) (string, error)
}

type ReviewStg struct {
	*database.Mongo
}

func (s *ReviewStg) Add(cmd *models.CreateReviewCMD) (string, error) {
	coll := s.Client.Database("reviewDB").Collection("reviews")

	res, err := coll.InsertOne(context.Background(),
		bson.D{
			{"comment", cmd.Comment},
			{"stars", cmd.Stars},
			{"createdAt", time.Now()},
			{"beerId", cmd.BeerID},
		})

	if err != nil {
		logs.Log().Error("cannot insert review")
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID)

	return id.String(), nil
}
