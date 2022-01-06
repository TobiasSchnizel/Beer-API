package web

import (
	"encoding/json"
	"github.com/TobiasSchnizel/Beer-API/beers/gateway"
	"github.com/TobiasSchnizel/Beer-API/beers/models"
	"github.com/TobiasSchnizel/Beer-API/internal/database"
	"net/http"
)

func (h *CreateBeerHandler) SaveBeerHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	cmd := parseRequest(r)
	res, err := h.Create(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg":"error in create beer"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

type CreateBeerHandler struct{
	gateway.BeerCreateGateway
}

func NewCreateBeerHandler(client *database.MySqlClient) *CreateBeerHandler {
	return &CreateBeerHandler{gateway.NewBeerCreateGateway(client)}
}
func parseRequest(r *http.Request) *models.CreateBeerCMD {
	body := r.Body
	defer body.Close()
	var cmd models.CreateBeerCMD

	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}