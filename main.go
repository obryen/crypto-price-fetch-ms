package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/obryen/crypto-price-fetch/client"
)

func main() {
	client := client.NewClient("http://localhost:3000")
	price, err := client.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("price", price)

	return

	listenAddr := flag.String("listenaddr", ":3000", "listen port for the app")
	flag.Parse()
	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
