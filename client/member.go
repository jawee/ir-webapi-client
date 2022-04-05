package client

import (
	"fmt"

	"github.com/jawee/ir-webapi-client/client/response"
)

func (c *Client) GetMember(custId int) (*response.MemberResponse, error) {
	uri := fmt.Sprintf("https://members-ng.iracing.com/data/member/get?cust_ids=%d", custId)

	var memberResponse response.MemberResponse
	err := c.get(uri, &memberResponse)
	if err != nil {
		return nil, err
	}

	return &memberResponse, nil
}
