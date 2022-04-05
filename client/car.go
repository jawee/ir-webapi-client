package client

import (
	"github.com/jawee/ir-webapi-client/client/response"
)

// https://members-ng.iracing.com/data/car/assets

// func (c *Client) GetCarAssets() (*map[string] response.CarAssets, error) {
func (c *Client) GetCarAssets() (*response.CarAssetsResponse, error) {
    uri := "https://members-ng.iracing.com/data/car/assets"
    var carAssets map[string] response.CarAssets
    err := c.get(uri, &carAssets)

    if err != nil {
        return nil, err
    }

    resp := response.CarAssetsResponse{
        CarAssets: carAssets,
    }

    return &resp, nil
}

