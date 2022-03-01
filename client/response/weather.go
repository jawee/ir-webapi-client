package response

import "time"

type Weather struct {
	Type                    int       `json:"type"`
	TempUnits               int       `json:"temp_units"`
	TempValue               int       `json:"temp_value"`
	RelHumidity             int       `json:"rel_humidity"`
	Fog                     int       `json:"fog"`
	WindDir                 int       `json:"wind_dir"`
	WindUnits               int       `json:"wind_units"`
	WindValue               int       `json:"wind_value"`
	Skies                   int       `json:"skies"`
	WeatherVarInitial       int       `json:"weather_var_initial"`
	WeatherVarOngoing       int       `json:"weather_var_ongoing"`
	TimeOfDay               int       `json:"time_of_day"`
	SimulatedStartUtcTime   time.Time `json:"simulated_start_utc_time"`
	SimulatedStartUtcOffset int       `json:"simulated_start_utc_offset"`
}
