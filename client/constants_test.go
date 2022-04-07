package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

type DivisionsRoundTripFunc func(req *http.Request) *http.Response

func (f DivisionsRoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewDivisionsTestClient(fn CarclassRoundTripFunc) *Client {
	jar, _ := cookiejar.New(nil)
	client := Client{
		Client: &http.Client{
			Transport: CarclassRoundTripFunc(fn),
		},
		CookieJar: jar,
	}
	return &client
}

func GetDivisionsSuccessFunc(req *http.Request) *http.Response {
	if req.URL.String() == "https://members-ng.iracing.com/data/constants/divisions" {
		return &http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(bytes.NewBufferString(`
            [
  {
    "label": "ALL",
    "value": -1
  },
  {
    "label": "Division 1",
    "value": 0
  },
  {
    "label": "Division 2",
    "value": 1
  },
  {
    "label": "Division 3",
    "value": 2
  },
  {
    "label": "Division 4",
    "value": 3
  },
  {
    "label": "Division 5",
    "value": 4
  },
  {
    "label": "Division 6",
    "value": 5
  },
  {
    "label": "Division 7",
    "value": 6
  },
  {
    "label": "Division 8",
    "value": 7
  },
  {
    "label": "Division 9",
    "value": 8
  },
  {
    "label": "Division 10",
    "value": 9
  },
  {
    "label": "Rookie",
    "value": 10
  }
]
`)),
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

func TestGetDivisions(t *testing.T) {
	client := NewDivisionsTestClient(GetDivisionsSuccessFunc)

	resp, err := client.GetDivisions()
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if resp == nil {
		t.Errorf("Response is nil")
	}

	if len(*resp) != 12 {
		t.Errorf("Response length is not 12, got %d\n", len(*resp))
	}
}
