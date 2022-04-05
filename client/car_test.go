package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

type CarRoundTripFunc func(req *http.Request) *http.Response

func (f CarRoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewCarTestClient(fn CarRoundTripFunc) *Client {
	jar, _ := cookiejar.New(nil)
	client := Client{
		Client: &http.Client{
			Transport: CarRoundTripFunc(fn),
		},
		CookieJar: jar,
	}
	return &client
}

func GetCarAssetsSuccessFunc(req *http.Request) *http.Response {
	// Test request parameters
	if req.URL.String() == "https://members-ng.iracing.com/data/car/assets" {
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
    "1": {
      "car_id": 1,
      "car_rules": [],
      "detail_copy": "<b>Some stuff</b>\n\r\n",
      "detail_screen_shot_images": "formula_skipbarber/fsb2000_ss_1,formula_skipbarber/fsb2000_ss_2",
      "detail_techspecs_copy": "<center>\r\n<b>Constructor</b><br>\r\nSkip Barber\r\n<br>",
      "folder": "/img/cars/skipbarberformula2000",
      "gallery_images": "8",
      "gallery_prefix": null,
      "group_image": null,
      "group_name": null,
      "large_image": "skipbarberformula2000-large.jpg",
      "logo": "/img/logos/partners/skipbarber-logo.png",
      "small_image": "skipbarberformula2000-small.jpg",
      "sponsor_logo": null,
      "template_path": "car_templates/1_template_SBRS.zip"
    },
    "2": {
      "car_id": 2,
      "car_rules": [],
      "detail_copy": "SK Modified racing is one of the most popular.",
      "detail_screen_shot_images": "modified/sk_modified/shot_ss_1",
      "detail_techspecs_copy": "<center>\r\n<b>Constructor</b>",
      "folder": "/img/cars/modifiedsk",
      "gallery_images": "8",
      "gallery_prefix": null,
      "group_image": null,
      "group_name": null,
      "large_image": "modifiedsk-large.jpg",
      "logo": "/img/logos/brand/iracingnotext-logo.png",
      "small_image": "modifiedsk-small.jpg",
      "sponsor_logo": null,
      "template_path": "car_templates/2_template_SK.zip"
    }
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

func TestGetCarAssets(t *testing.T) {
	client := NewCarTestClient(GetCarAssetsSuccessFunc)

	resp, err := client.GetCarAssets()
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if resp == nil {
		t.Errorf("Response is nil")
	}

	if len(resp.CarAssets) != 2 {
		t.Errorf("Expected 2 cars, got %d", len(resp.CarAssets))
	}
}
