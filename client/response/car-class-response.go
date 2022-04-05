package response

type CarClassResponse []struct {
	CarClassID    int         `json:"car_class_id"`
	CarsInClass   []CarsClass `json:"cars_in_class"`
	CustID        int         `json:"cust_id"`
	Name          string      `json:"name"`
	RelativeSpeed int         `json:"relative_speed"`
	ShortName     string      `json:"short_name"`
}
type CarsClass struct {
	CarDirpath string `json:"car_dirpath"`
	CarID      int    `json:"car_id"`
	Retired    bool   `json:"retired"`
}
