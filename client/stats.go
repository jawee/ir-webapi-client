package client

import (
	"fmt"

	"github.com/jawee/ir-webapi-client/client/response"
)

func (c *Client) GetMemberRecentRaces(custId int) (*response.RecentRacesResponse, error) {
	uri := fmt.Sprintf("https://members-ng.iracing.com/data/stats/member_recent_races?cust_id=%d", custId)

	var memberRecentRacesResp response.RecentRacesResponse
    err := c.get(uri, &memberRecentRacesResp)

	if err != nil {
        return nil, err
	}

	return &memberRecentRacesResp, nil
}
