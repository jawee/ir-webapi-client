package client

import (
	"fmt"

	"github.com/jawee/ir-webapi-client/client/response"
)

func (c *Client) GetSession(subsessionId int) (*response.Session, error) {
	uri := fmt.Sprintf("https://members-ng.iracing.com/data/results/get?subsession_id=%d", subsessionId)

	var sessionResp response.Session
	err := c.get(uri, &sessionResp)

	if err != nil {
		return nil, err
	}

	return &sessionResp, nil
}
