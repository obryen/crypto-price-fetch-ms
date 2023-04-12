package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/obryen/crypto-price-fetch/client"
	"github.com/obryen/crypto-price-fetch/proto"
)

func main() {
	var (
		jsonListenAddr = flag.String("listenaddr", ":3000", "json listn address for the app")
		grpcListenAddr = flag.String("listenaddr", ":4000", "grpc listn address for the app")
		svc            = loggingService{&priceFetcher{}}
		ctx            = context.Background()
	)

	flag.Parse()
	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		time.Sleep(3 * time.Second)
		resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v/n", resp)
	}()

	go MakeGRPCServerAndRun(*grpcListenAddr, svc)

	jsonServer := NewJSONAPIServer(*jsonListenAddr, svc)
	jsonServer.Run()

}
