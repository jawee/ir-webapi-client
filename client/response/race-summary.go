package response

type RaceSummary struct {
	SubsessionID         int    `json:"subsession_id"`
	AverageLap           int    `json:"average_lap"`
	LapsComplete         int    `json:"laps_complete"`
	NumCautions          int    `json:"num_cautions"`
	NumCautionLaps       int    `json:"num_caution_laps"`
	NumLeadChanges       int    `json:"num_lead_changes"`
	FieldStrength        int    `json:"field_strength"`
	NumOptLaps           int    `json:"num_opt_laps"`
	HasOptPath           bool   `json:"has_opt_path"`
	SpecialEventType     int    `json:"special_event_type"`
	SpecialEventTypeText string `json:"special_event_type_text"`
}
