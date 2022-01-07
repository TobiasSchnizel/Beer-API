package web

import (
	"encoding/json"
	"net/http"

	"github.com/TobiasSchnizel/Beer-API/internal/database"
	"github.com/TobiasSchnizel/Beer-API/reviews/gateway"
	"github.com/TobiasSchnizel/Beer-API/reviews/models"
)

func (h *ReviewHandler) AddReviewHandler(w http.ResponseWriter, r *http.Request) {
	cmd := params(r)
	res, err := h.AddReview(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Whrite([]byte(res))
}

func params(r *http.Request) *models.CreateReviewCMD {
	var cmd models.CreateReviewCMD

	_ = json.NewDecoder(r.Body).Decode(&cmd)

	return &cmd
}

type ReviewHandler struct {
	gateway.ReviewGateway
}

func NewReviewHandler(mongo *database.Mongo) *ReviewHandler {
	return &ReviewHandler{gateway.NewReviewGateway(mongo)}
}
