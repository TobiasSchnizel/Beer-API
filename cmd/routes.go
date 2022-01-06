package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/TobiasSchnizel/Beer-API/beers/web"
	reviews "github.com/TobiasSchnizel/beer-API/reviews/web"
)

func Routes(
	beerhandler *web.CreateBeerHandler,
	reviewHandler *reviews.reviewHandler,
) *chi.Mux {
	mux := chi.NewMux()
	// globals
	mux.Use(
		middleware.Logger, // log http requests
		middleware.Recoverer, // recover in panic
	)

	mux.Get("/beers", beersHandler.SaveBeerHandler)
	mux.Post("/beers", SaveBeerHandler)
	mux.Get("/beer/{beerID}", findBeerHandler)
	mux.Get("/beer/{beerID}/boxprice", priceBoxBeerHandler)
	mux.Post("/reviews", reviewHandler.AddReviewHandler)

	return mux
}

func beersHandler(w http.ResponseWriter, r *http.Request) {

}

func findBeerHandler(w http.ResponseWriter, r *http.Request){

}
func priceBoxBeerHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("programer:", "Tobias")

	res:= map[string]interface{}{"message": "hello, World"}
	_ = json.NewEncoder(w).Encode(res)
}