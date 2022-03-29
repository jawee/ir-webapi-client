package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"time"

	"github.com/jawee/ir-webapi-client/client/response"
)

type Client struct {
	CookieJar *cookiejar.Jar
	loginTime time.Time
}

func New() *Client {

    jar, _ := cookiejar.New(nil)

    return &Client{
        CookieJar: jar,
    }
}

func (c *Client) Login(email, password string) error {
    uri := "https://members-ng.iracing.com/auth"

    client := &http.Client{
        Jar: c.CookieJar,
    }

    reqBody := bytes.NewBuffer([]byte(fmt.Sprintf("{\"email\": \"%s\", \"password\": \"%s\"}", email, password)))
    req, err := http.NewRequest("POST", uri, reqBody)

    if err != nil {
        return err
    }
    req.Header.Add("Content-Type", "application/json")

    resp, err := client.Do(req)

    if err != nil {
        log.Println(err)
        return err
    }

    defer resp.Body.Close()
    var respMap map[string]interface{}
    err = json.NewDecoder(resp.Body).Decode(&respMap)

    if err != nil {
        return err
    }

    if respMap["authcode"] == 0.0 {
        return fmt.Errorf("Login failed: %s\n", respMap["message"])
    }

    return nil
}

func printBody(body io.ReadCloser) {
    b, err := io.ReadAll(body)

    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println(string(b))
}

func printResp(resp *http.Response) {
    respDump, err := httputil.DumpResponse(resp, true)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("RESPONSE:\n%s", string(respDump))
}

func (c *Client) getLink(uri string) (*response.LinkResponse, error) {

    client := &http.Client{
        Jar: c.CookieJar,
    }

    req, err := http.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, err
    }

    resp, err := client.Do(req)

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    b, err := io.ReadAll(resp.Body)
    

    var linkResp response.LinkResponse
    err = json.Unmarshal(b, &linkResp)
    
    if err != nil {
        log.Fatalln(err)
    }

    return &linkResp, nil;
}

func (c *Client) fetchLink(uri string, resp interface{}) error {
    client := &http.Client{
        Jar: c.CookieJar,
    }

    req, err := http.NewRequest("GET", uri, nil)
    if err != nil {
        return nil
    }

    res, err := client.Do(req)
    if err != nil {
        return nil
    }

    defer res.Body.Close()

    b, err := io.ReadAll(res.Body)

    err = json.Unmarshal(b, resp)
    if err != nil {
        return nil
    }

    return nil
}


