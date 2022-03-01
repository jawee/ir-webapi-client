package response

type Livery struct {
	CarID        int         `json:"car_id"`
	Pattern      int         `json:"pattern"`
	Color1       string      `json:"color1"`
	Color2       string      `json:"color2"`
	Color3       string      `json:"color3"`
	NumberFont   int         `json:"number_font"`
	NumberColor1 string      `json:"number_color1"`
	NumberColor2 string      `json:"number_color2"`
	NumberColor3 string      `json:"number_color3"`
	NumberSlant  int         `json:"number_slant"`
	Sponsor1     int         `json:"sponsor1"`
	Sponsor2     int         `json:"sponsor2"`
	CarNumber    string      `json:"car_number"`
	WheelColor   interface{} `json:"wheel_color"`
	RimType      int         `json:"rim_type"`
}
