package response

type Stats struct {
	CategoryID        int     `json:"category_id"`
	Category          string  `json:"category"`
	Starts            int     `json:"starts"`
	Wins              int     `json:"wins"`
	Top5              int     `json:"top5"`
	Poles             int     `json:"poles"`
	AvgStartPosition  int     `json:"avg_start_position"`
	AvgFinishPosition int     `json:"avg_finish_position"`
	Laps              int     `json:"laps"`
	LapsLed           int     `json:"laps_led"`
	AvgIncidents      float64 `json:"avg_incidents"`
	AvgPoints         int     `json:"avg_points"`
	WinPercentage     float64 `json:"win_percentage"`
	Top5Percentage    float64 `json:"top5_percentage"`
	LapsLedPercentage float64 `json:"laps_led_percentage"`
	TotalClubPoints   int     `json:"total_club_points"`
	Year              int     `json:"year"`
}
