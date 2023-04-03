package main

import (
	"flag"
)

func main() {
	// client := client.NewClient("http://localhost:3000/")
	// price, err := client.FetchPrice(context.Background(), "ETH")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print("price", price)

	// return

	jsonListenAddr := flag.String("listenaddr", ":3000", "json listn address for the app")
	grpcListenAddr := flag.String("listenaddr", ":3000", "grpc listn address for the app")
	flag.Parse()
	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	jsonServer := NewJSONAPIServer(*jsonListenAddr, svc)
	jsonServer.Run()

	go MakeGRPCServerAndRun(*grpcListenAddr, svc)
}
