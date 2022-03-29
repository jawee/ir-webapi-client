package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetSession(subsessionId int) error {
    uri := fmt.Sprintf("https://members-ng.iracing.com/data/results/get?subsession_id=%d", subsessionId)

    client := &http.Client{
        Jar: c.CookieJar,
    }

    req, err := http.NewRequest("GET", uri, nil)

    if err != nil {
        return err
    }

    resp, err := client.Do(req)

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
