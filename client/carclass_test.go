package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

type CarclassRoundTripFunc func(req *http.Request) *http.Response

func (f CarclassRoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewCarclassTestClient(fn CarclassRoundTripFunc) *Client {
	jar, _ := cookiejar.New(nil)
	client := Client{
		Client: &http.Client{
			Transport: CarclassRoundTripFunc(fn),
		},
		CookieJar: jar,
	}
	return &client
}

func GetCarclassSuccessFunc(req *http.Request) *http.Response {
	if req.URL.String() == "https://members-ng.iracing.com/data/carclass/get" {
		return &http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
                    "link": "https://some-url.com"
                }`)),
			Header: make(http.Header),
		}
	}

	if req.URL.String() == "https://some-url.com" {
		return &http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(bytes.NewBufferString(`
            [
    {
      "car_class_id": 1,
      "cars_in_class": [
        {
          "car_dirpath": "rt2000",
          "car_id": 1,
          "retired": false
        }
      ],
      "cust_id": 0,
      "name": "Skip Barber Race Series",
      "relative_speed": 40,
      "short_name": "SBRS"
    },
    {
        "car_class_id": 2041,
        "cars_in_class": [
          {
            "car_dirpath": "ferrarievogt3",
            "car_id": 144,
            "retired": false
          },
          {
            "car_dirpath": "mercedesamggt3",
            "car_id": 72,
            "retired": false
          }
        ],
        "cust_id": 0,
        "name": "GT3 Class 2",
        "relative_speed": 52,
        "short_name": "GT3 Class 2"
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

func TestGetCarclass(t *testing.T) {
	client := NewCarclassTestClient(GetCarclassSuccessFunc)

	resp, err := client.GetCarClass()
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if resp == nil {
		t.Errorf("Response is nil")
	}

    if len(*resp) != 2 {
        t.Errorf("Response length is not 2")
    }
}
