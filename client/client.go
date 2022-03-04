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

type loginRequest struct {
    email string
    password string
}

func (c *Client) Login(email, password string) error {
    uri := "https://members-ng.iracing.com/auth"

    client := &http.Client{
        Jar: c.CookieJar,
    }

    loginRequest := loginRequest{
        email: email,
        password: password,
    }

    reqBody := bytes.NewBuffer([]byte(fmt.Sprintf("{\"email\": \"%s\", \"password\": \"%s\"}", loginRequest.email, loginRequest.password)))
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

    log.Printf("%v\n", respMap)

    if err != nil {
        log.Println(err)
        return err
    }

    if respMap["authcode"] == 0.0 {
        log.Printf("Login failed: %s\n", respMap["message"])
        return fmt.Errorf("Login failed: %s\n", respMap["message"])
    }

    log.Println("Login successful")
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


// /usr/bin/curl -b cookie-jar.txt -c cookie-jar.txt  https://members-ng.iracing.com/data/results/get?subsession_id=38280997
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
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))

    return nil;
}

// {
//   "link": "https://members-ng.iracing.com/data/member/get",
//   "parameters": {
//     "cust_ids": {
//       "type": "numbers",
//       "required": true,
//       "note": "?cust_ids=2,3,4"
//     },
//     "include_licenses": {
//       "type": "boolean"
//     }
//   }
// }
func (c *Client) GetMember(custId int) error {
    uri := "https://members-ng.iracing.com/data/member/get"

    client := &http.Client{
        Jar: c.CookieJar,
    }

    req, err := http.NewRequest("GET", uri, nil)

    if err != nil {
        return err
    }

    req.URL.RawQuery = fmt.Sprintf("cust_ids=%d", custId)

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


// {
//   "link": "https://members-ng.iracing.com/data/stats/member_recent_races",
//   "parameters": {
//     "cust_id": {
//       "type": "number",
//       "note": "Defaults to the authenticated member."
//     }
//   }
// }
func (c *Client) GetMemberRecentRaces(custId int) error {
    uri := "https://members-ng.iracing.com/data/stats/member_recent_races"

    client := &http.Client{
        Jar: c.CookieJar,
    }

    req, err := http.NewRequest("GET", uri, nil)

    if err != nil {
        return err
    }

    req.URL.RawQuery = fmt.Sprintf("cust_id=%d", custId)

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
