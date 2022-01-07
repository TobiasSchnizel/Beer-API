package main

import (
	"encoding/json"
	"net/http"

	"github.com/TobiasSchnizel/Beer-API/beers/web"
	reviews "github.com/TobiasSchnizel/Beer-API/reviews/web"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(beerHandler *web.CreateBeerHandler, reviewHandler *reviews.ReviewHandler) *chi.Mux {
	mux := chi.NewMux()
	// globals
	mux.Use(
		middleware.Logger,    // log http requests
		middleware.Recoverer, // recover in panic
	)

	mux.Post("/beers", beerHandler.SaveBeerHandler)
	mux.Get("/beers", nil)
	mux.Get("/beer/{beerID}", nil)
	mux.Get("/beer/{beerID}/boxprice", priceBoxBeerHandler)
	mux.Post("/reviews", reviewHandler.AddReviewHandler)

	return mux
}

func priceBoxBeerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("programer:", "Tobias")

	res := map[string]interface{}{"message": "hello, World"}
	_ = json.NewEncoder(w).Encode(res)
}
