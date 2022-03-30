package client

import (
	"fmt"

	"github.com/jawee/ir-webapi-client/client/response"
)

func (c *Client) GetSession(subsessionId int) (*response.Session, error) {
    uri := fmt.Sprintf("https://members-ng.iracing.com/data/results/get?subsession_id=%d", subsessionId)

    linkResp, err := c.getLink(uri)
	if err != nil {
		return nil, err
	}

    var sessionResp response.Session
    err = c.fetchLink(linkResp.Link, &sessionResp)

    if err != nil {
        return nil, err
    }

    return &sessionResp, nil
}
