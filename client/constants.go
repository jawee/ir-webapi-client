package client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/jawee/ir-webapi-client/client/response"
)

func (c *Client) GetDivisions() (*response.Divisions, error) {
	uri := "https://members-ng.iracing.com/data/constants/divisions"

	c.Client.Jar = c.CookieJar

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Invalid response code")
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	var divisionsResp response.Divisions
	err = json.Unmarshal(b, &divisionsResp)
	if err != nil {
		return nil, err
	}

	return &divisionsResp, nil
}
