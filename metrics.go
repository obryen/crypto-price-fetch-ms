package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricService{
		next: next,
	}
}

func (m *metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("storing and pushing to 3rd party ")
	// store metric , push to 3rd party monitering like prometheus

	return m.next.FetchPrice(ctx, ticker)
}
