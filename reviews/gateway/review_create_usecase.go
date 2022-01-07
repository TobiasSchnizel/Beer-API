package gateway

import "github.com/TobiasSchnizel/Beer-API/reviews/models"

func (g *ReviewGtw) AddReview(cmd *models.CreateReviewCMD) (string, error) {
	return g.Add(cmd)
}
