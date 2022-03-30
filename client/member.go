package client

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jawee/ir-webapi-client/client/response"
)

func (c *Client) GetMember(custId int) error {
    uri := "https://members-ng.iracing.com/data/member/get"

    c.Client.Jar = c.CookieJar

    req, err := http.NewRequest("GET", uri, nil)

    if err != nil {
        return err
    }

    req.URL.RawQuery = fmt.Sprintf("cust_ids=%d", custId)

    resp, err := c.Client.Do(req)

    if err != nil {
        return err
    }

    defer resp.Body.Close()

    b, err := io.ReadAll(resp.Body)

    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println(string(b))

    return nil;
}

func (c *Client) GetMemberRecentRaces(custId int) (*response.RecentRacesResponse, error) {
    uri := fmt.Sprintf("https://members-ng.iracing.com/data/stats/member_recent_races?cust_id=%d", custId)

    linkResp, err := c.getLink(uri) 
    if err != nil {
        log.Fatalln(err)
    }

    var memberRecentRacesResp response.RecentRacesResponse
    err = c.fetchLink(linkResp.Link, &memberRecentRacesResp)

    if err != nil {
        log.Fatalln(err)
    }

    return &memberRecentRacesResp, nil;
}
