package main

import (
	"context"
	"fmt"
	"time"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetching(ctx, ticker)

}

var priceMock = map[string]float64{
	"BTC": 20000.0,
	"ETH": 200.1,
	"RPX": 100.5,
}

func MockPriceFetching(ctx context.Context, ticker string) (float64, error) {
	time.Sleep(100 * time.Millisecond)

	price, ok := priceMock[ticker]

	if !ok {
		return price, fmt.Errorf("the given ticker (%s) does not exist", ticker)
	}

	return price, nil
}
