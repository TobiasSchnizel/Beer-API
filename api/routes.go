package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/TobiasSchnizel/Beer-API/gadgets/beers/web"
)

func Routes(bh *CreateBeerHandler ) *chi.Mux {
	mux := chi.NewMux()
	// globals
	mux.Use(
		middleware.Logger, // log every http request
		middleware.Recoverer, // recover recover if panic occurs
	)

	mux.Get("/beers", beersHandler)
	mux.Post("/beers", saveBeerHandler)
	mux.Get("/beer/{beerID}", findBeerHandler)
	mux.Get("/beer/{beerID}/boxprice", priceBoxBeerHandler)

	return mux
}

func beersHandler(w http.ResponseWriter, r *http.Request) {

}

func findBeerHandler(w http.ResponseWriter, r *http.Request){

}
func priceBoxBeerHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("done-by", "tobias")

	res:= map[string]interface{}{"message": "hello, World"}
	_ = json.NewEncoder(w).Encode(res)
}