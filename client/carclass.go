package client

import (
	"github.com/jawee/ir-webapi-client/client/response"
)

func (c *Client) GetCarClass() (*response.CarClassResponse, error) {
	uri := "https://members-ng.iracing.com/data/carclass/get"
	var carClassResp response.CarClassResponse 
	err := c.get(uri, &carClassResp)

	if err != nil {
		return nil, err
	}

	return &carClassResp, nil
}
