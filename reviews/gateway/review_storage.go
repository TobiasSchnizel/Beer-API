package gateway

import (
	"context"
	"github.com/TobiasSchnizel/Beer-API/intenal/database"
	"github.com/TobiasSchnizel/Beer-API/intenal/logs"
	"github.com/TobiasSchnizel/Beer-API/reviews/models"
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
			{"beerId", cmd.BeerId},
		})

	if err != nil {
		logs.Log().Error("cannot insert review")
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID)

	return id.String(), nil
}