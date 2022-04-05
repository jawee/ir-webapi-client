package client

import (
	"github.com/jawee/ir-webapi-client/client/response"
)

func (c *Client) GetCarAssets() (*response.CarAssetsResponse, error) {
	uri := "https://members-ng.iracing.com/data/car/assets"
	var carAssets map[string]response.CarAssets
	err := c.get(uri, &carAssets)

	if err != nil {
		return nil, err
	}

	resp := response.CarAssetsResponse{
		CarAssets: carAssets,
	}

	return &resp, nil
}

func (c *Client) GetCars() (*response.Car, error) {
	uri := "https://members-ng.iracing.com/data/car/get"
	var cars response.Car
	err := c.get(uri, &cars)

	if err != nil {
		return nil, err
	}

	return &cars, nil
}
