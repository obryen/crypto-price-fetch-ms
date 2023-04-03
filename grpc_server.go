package main

import (
	"context"
	"net"

	"github.com/obryen/crypto-price-fetch/proto"
	"google.golang.org/grpc"
)

func MakeGRPCServerAndRun(listenAddr string, svc IPriceFetcherService) error {
	grpcPriceFetcher := NewGRPCPriceFetcher(svc)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts)
	proto.RegisterPriceFetcherServer(server, grpcPriceFetcher)
	return server.Serve(ln)
}

type GRPCPriceFetcherServer struct {
	svc IPriceFetcherService
	proto.UnimplementedPriceFetcherServer
}

func NewGRPCPriceFetcher(svc IPriceFetcherService) *GRPCPriceFetcherServer {
	return &GRPCPriceFetcherServer{
		svc: svc,
	}
}

func (server *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {
	price, err := server.svc.FetchPrice(ctx, req.Ticker)
	if err != nil {
		return nil, err
	}

	resp := &proto.PriceResponse{
		Ticker: req.Ticker,
		Price:  float32(price),
	}

	return resp, err
}
