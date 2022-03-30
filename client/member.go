package client

import (
	"fmt"

	"github.com/jawee/ir-webapi-client/client/response"
)

func (c *Client) GetMember(custId int) (*response.MemberResponse, error) {
	uri := fmt.Sprintf("https://members-ng.iracing.com/data/member/get?cust_ids=%d", custId)

    linkResp, err := c.getLink(uri)
	if err != nil {
		return nil, err
	}

    var memberResponse response.MemberResponse
    err = c.fetchLink(linkResp.Link, &memberResponse)

    if err != nil {
        return nil, err
    }

    return &memberResponse, nil
}

func (c *Client) GetMemberRecentRaces(custId int) (*response.RecentRacesResponse, error) {
	uri := fmt.Sprintf("https://members-ng.iracing.com/data/stats/member_recent_races?cust_id=%d", custId)

	linkResp, err := c.getLink(uri)
	if err != nil {
        return nil, err
	}

	var memberRecentRacesResp response.RecentRacesResponse
	err = c.fetchLink(linkResp.Link, &memberRecentRacesResp)

	if err != nil {
        return nil, err
	}

	return &memberRecentRacesResp, nil
}
