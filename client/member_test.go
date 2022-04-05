package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *Client {
	jar, _ := cookiejar.New(nil)
	client := Client{
		Client: &http.Client{
			Transport: RoundTripFunc(fn),
		},
		CookieJar: jar,
	}
	return &client
}

func GetMemberSuccessFunc(req *http.Request) *http.Response {
	// Test request parameters
	if req.URL.String() == "https://members-ng.iracing.com/data/member/get?cust_ids=123" {
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
                    "link": "https://some-url.com"
                }`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	}

	if req.URL.String() == "https://some-url.com" {
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
                  "success": true,
                  "cust_ids": [
                    123
                  ],
                  "members": [
                    {
                      "cust_id": 123,
                      "display_name": "Some Name",
                      "helmet": {
                        "pattern": 0,
                        "color1": "060eb1",
                        "color2": "fff505",
                        "color3": "0000d6",
                        "face_type": 0,
                        "helmet_type": 0
                      },
                      "last_login": "2022-03-26T11:34:24.065609Z",
                      "member_since": "2022-03-29",
                      "club_id": 43,
                      "club_name": "Scandinavia",
                      "ai": false
                    }
                  ]
                }`)),
			Header: make(http.Header),
		}
	}

	return &http.Response{
		StatusCode: 401,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
              "error": "Unauthorized"
            }`)),
		Header: make(http.Header),
	}
}

func TestGetMemberUnauthorized(t *testing.T) {
	client := NewTestClient(GetMemberSuccessFunc)
	resp, err := client.GetMember(124)
	if err == nil {
		t.Errorf("Error should be nil")
	}

	if resp != nil {
		t.Errorf("Response is not nil")
	}
}

func TestGetMember(t *testing.T) {
	client := NewTestClient(GetMemberSuccessFunc)

	resp, err := client.GetMember(123)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if len(resp.Members) == 0 {
		t.Errorf("Error: %s", "Members is empty")
	}

	if len(resp.CustIds) == 0 {
		t.Errorf("Error: %s", "cust_ids is empty")
	}

	for _, member := range resp.Members {
		if member.CustID != 123 {
			t.Errorf("Error: %s", "CustID is not 123")
		}
	}
}
