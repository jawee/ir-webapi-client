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

func (c *Client) GetMemberSummary(custId int) (*response.MemberSummary, error) {
    uri := fmt.Sprintf("https://members-ng.iracing.com/data/stats/member_summary?cust_id=%d", custId)

    var memberSummaryResp response.MemberSummary
    err := c.get(uri, &memberSummaryResp)

    if err != nil {
        return nil, err
    }

    return &memberSummaryResp, nil
}

func (c *Client) GetMemberYearly(custId int) (*response.MemberYearly, error) {
    uri := fmt.Sprintf("https://members-ng.iracing.com/data/stats/member_yearly?cust_id=%d", custId)
    var memberYearlyResp response.MemberYearly
    err := c.get(uri, &memberYearlyResp)

    if err != nil {
        return nil, err
    }

    return &memberYearlyResp, nil
}
