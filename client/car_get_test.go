package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

type GetCarRoundTripFunc func(req *http.Request) *http.Response

func (f GetCarRoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewGetCarTestClient(fn GetCarRoundTripFunc) *Client {
	jar, _ := cookiejar.New(nil)
	client := Client{
		Client: &http.Client{
			Transport: GetCarRoundTripFunc(fn),
		},
		CookieJar: jar,
	}
	return &client
}

func GetCarSuccessFunc(req *http.Request) *http.Response {
	// Test request parameters
	if req.URL.String() == "https://members-ng.iracing.com/data/car/get" {
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
			Body: ioutil.NopCloser(bytes.NewBufferString(`[
    {
        "ai_enabled": false,
        "allow_number_colors": true,
        "allow_number_font": true,
        "allow_sponsor1": true,
        "allow_sponsor2": true,
        "allow_wheel_color": false,
        "award_exempt": true,
        "car_dirpath": "williamsfw31",
        "car_id": 33,
        "car_make": "Williams",
        "car_model": "FW31",
        "car_name": "Williams-Toyota FW31",
        "car_name_abbreviated": "FW31",
        "car_types": [
        {
            "car_type": "f1"
        },
        {
            "car_type": "formula1"
        },
        {
            "car_type": "formulaone"
        },
        {
            "car_type": "openwheel"
        },
        {
            "car_type": "road"
        }
        ],
        "car_weight": 1331,
        "categories": [
        "road"
        ],
        "created": "2010-08-05T10:10:05Z",
        "free_with_subscription": false,
        "has_headlights": false,
        "has_multiple_dry_tire_types": false,
        "hp": 750,
        "max_power_adjust_pct": 0,
        "max_weight_penalty_kg": 250,
        "min_power_adjust_pct": -5,
        "package_id": 99,
        "paint_rules": {
        "RestrictCustomPaint": true
        },
        "patterns": 2,
        "price": 11.95,
        "retired": false,
        "search_filters": "road,openwheel,f1,formula1,formulaone",
        "site_url": "http://www.attwilliams.com/",
        "sku": 10222
    },
    {
        "ai_enabled": false,
        "allow_number_colors": true,
        "allow_number_font": true,
        "allow_sponsor1": true,
        "allow_sponsor2": true,
        "allow_wheel_color": true,
        "award_exempt": false,
        "car_dirpath": "skmodified",
        "car_id": 2,
        "car_name": "Modified - SK",
        "car_name_abbreviated": "SK",
        "car_types": [
          {
            "car_type": "mod"
          },
          {
            "car_type": "nascar"
          },
          {
            "car_type": "oval"
          },
          {
            "car_type": "sk mod"
          }
        ],
        "car_weight": 2600,
        "categories": [
          "oval"
        ],
        "created": "2006-05-03T21:15:00Z",
        "free_with_subscription": false,
        "has_headlights": false,
        "has_multiple_dry_tire_types": false,
        "hp": 370,
        "max_power_adjust_pct": 0,
        "max_weight_penalty_kg": 250,
        "min_power_adjust_pct": -5,
        "package_id": 16,
        "paint_rules": {
          "1": {
            "PaintCarAvailable": false,
            "Color1": "FF0000",
            "Color2": "00FF00",
            "Color3": "0000FF",
            "Sponsor1Available": false,
            "Sponsor2Available": false,
            "Sponsor1": "0",
            "Sponsor2": "0",
            "RulesExplanation": "The currently selected car pattern doesn't allow painting the car or the car's numbers."
          }
        },
        "patterns": 25,
        "price": 11.95,
        "retired": false,
        "search_filters": "oval,nascar,mod, sk mod",
        "sku": 10010
    },
    {
        "ai_enabled": true,
        "allow_number_colors": true,
        "allow_number_font": true,
        "allow_sponsor1": true,
        "allow_sponsor2": true,
        "allow_wheel_color": true,
        "award_exempt": false,
        "car_dirpath": "trucks\\tundra2015",
        "car_id": 62,
        "car_make": "Toyota",
        "car_model": "Tundra",
        "car_name": "[Retired] NASCAR Gander Outdoors Toyota Tundra",
        "car_name_abbreviated": "TT",
        "car_types": [
          {
            "car_type": "nascar"
          },
          {
            "car_type": "oval"
          },
          {
            "car_type": "truck"
          }
        ],
        "car_weight": 3400,
        "categories": [
          "oval"
        ],
        "created": "2015-02-25T13:17:52Z",
        "free_with_subscription": false,
        "has_headlights": false,
        "has_multiple_dry_tire_types": false,
        "hp": 650,
        "max_power_adjust_pct": 0,
        "max_weight_penalty_kg": 250,
        "min_power_adjust_pct": -5,
        "package_id": 175,
        "paint_rules": {
          "1": {
            "PaintCarAvailable": false,
            "Color1": "FF0000",
            "Color2": "00FF00",
            "Color3": "0000FF",
            "Sponsor1Available": false,
            "Sponsor2Available": false,
            "Sponsor1": "0",
            "Sponsor2": "0",
            "RulesExplanation": "The currently selected car pattern doesn't allow painting the car or the car's numbers."
          },
          "2": {
            "PaintCarAvailable": false,
            "Color1": "FF0000",
            "Color2": "00FF00",
            "Color3": "0000FF",
            "Sponsor1Available": false,
            "Sponsor2Available": false,
            "Sponsor1": "0",
            "Sponsor2": "0",
            "RulesExplanation": "The currently selected car pattern doesn't allow painting the car or the car's numbers."
          },
          "3": {
            "PaintCarAvailable": false,
            "Color1": "FF0000",
            "Color2": "00FF00",
            "Color3": "0000FF",
            "Sponsor1Available": false,
            "Sponsor2Available": false,
            "Sponsor1": "0",
            "Sponsor2": "0",
            "RulesExplanation": "The currently selected car pattern doesn't allow painting the car or the car's numbers."
          }
        },
        "patterns": 27,
        "price": 0,
        "retired": true,
        "search_filters": "oval,nascar,truck",
        "site_url": "http://www.iracing.com",
        "sku": 0
      }
]`)),
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

func TestGetCar(t *testing.T) {
	client := NewGetCarTestClient(GetCarSuccessFunc)

	resp, err := client.GetCars()
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if resp == nil {
		t.Errorf("Response is nil")
	}

    if len(resp.Cars) != 3 {
        t.Errorf("Expected 3 cars, got %d", len(resp.Cars))
    }

}
