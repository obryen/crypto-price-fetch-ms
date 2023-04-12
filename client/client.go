package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/obryen/crypto-price-fetch/proto"
	"github.com/obryen/crypto-price-fetch/types"
	"google.golang.org/grpc"
)

func NewGRPCClient(remoteAddress string) (proto.PriceFetcherClient, error) {
	conn, err := grpc.Dial(remoteAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := proto.NewPriceFetcherClient(conn)

	return c, nil
}

type Client struct {
	endpoint string
}

func NewClient(endPoint string) *Client {
	return &Client{
		endpoint: endPoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceRes, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)
	fmt.Println("endpoint val \n", endpoint)
	req, err := http.NewRequest("get", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		httpErr := make(map[string]interface{})
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service responded with none status OK code: %s", httpErr["error"])
	}

	priceRes := new(types.PriceRes)
	if err3 := json.NewDecoder(resp.Body).Decode(priceRes); err3 != nil {
		return nil, err
	}
	return priceRes, nil
}
