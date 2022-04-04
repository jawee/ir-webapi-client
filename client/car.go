package client

import (
	"github.com/jawee/ir-webapi-client/client/response"
)

// https://members-ng.iracing.com/data/car/assets

func (c *Client) GetCarAssets() (*map[string] response.CarAssets, error) {
    uri := "https://members-ng.iracing.com/data/car/assets"
    var carAssetsResponse map[string] response.CarAssets
    err := c.get(uri, &carAssetsResponse)

    if err != nil {
        return nil, err
    }

    return &carAssetsResponse, nil
}

