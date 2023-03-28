package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/obryen/crypto-price-fetch/types"
)

type Client struct {
	endpoint string
}

func NewClient(endPoint string) *Client {
	return &Client{
		endpoint: endPoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceRes, error) {
	req, err := http.NewRequest("get", c.endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("no status OK code %s", httpErr["error"])
	}

	priceRes := new(types.PriceRes)
	if err3 := json.NewDecoder(resp.Body).Decode(priceRes); err3 != nil {
		return nil, err
	}
	return priceRes, nil
}
