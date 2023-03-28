package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/obryen/crypto-price-fetch/types"
)

type APIfunc func(context.Context, http.ResponseWriter, *http.Request) error

type JSONAPIServer struct {
	listAdd string
	svc     PriceFetcher
}

func NewJSONAPIServer(listAdd string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listAdd: listAdd,
		svc:     svc,
	}
}

func (j *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHttpHandlerFunc(j.handleFetchPrice))

	http.ListenAndServe(j.listAdd, nil)

}

func makeHttpHandlerFunc(apiFn APIfunc) http.HandlerFunc {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (j *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := j.svc.FetchPrice(ctx, ticker)

	if err != nil {
		return err
	}

	priceResp := types.PriceRes{
		Price:  price,
		Ticker: ticker,
	}

	return writeJSON(w, http.StatusAccepted, priceResp)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
