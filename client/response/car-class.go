package response

type CarClass struct {
	CarClassID    int           `json:"car_class_id"`
	CarsInClass   []CarsInClass `json:"cars_in_class"`
	Name          string        `json:"name"`
	RelativeSpeed int           `json:"relative_speed"`
	ShortName     string        `json:"short_name"`
}
