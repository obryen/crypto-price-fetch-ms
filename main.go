package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	price, err := svc.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the price of %s is %f", "ETH", price)
}
