package response

import "time"

type Race struct {
	SeasonID           int          `json:"season_id"`
	SeriesID           int          `json:"series_id"`
	SeriesName         string       `json:"series_name"`
	CarID              int          `json:"car_id"`
	CarClassID         int          `json:"car_class_id"`
	Livery             Livery       `json:"livery"`
	LicenseLevel       int          `json:"license_level"`
	SessionStartTime   time.Time    `json:"session_start_time"`
	WinnerGroupID      int          `json:"winner_group_id"`
	WinnerName         string       `json:"winner_name"`
	WinnerHelmet       WinnerHelmet `json:"winner_helmet"`
	WinnerLicenseLevel int          `json:"winner_license_level"`
	StartPosition      int          `json:"start_position"`
	FinishPosition     int          `json:"finish_position"`
	QualifyingTime     int          `json:"qualifying_time"`
	Laps               int          `json:"laps"`
	LapsLed            int          `json:"laps_led"`
	Incidents          int          `json:"incidents"`
	ClubPoints         int          `json:"club_points"`
	Points             int          `json:"points"`
	StrengthOfField    int          `json:"strength_of_field"`
	SubsessionID       int          `json:"subsession_id"`
	OldSubLevel        int          `json:"old_sub_level"`
	NewSubLevel        int          `json:"new_sub_level"`
	OldiRating         int          `json:"oldi_rating"`
	NewiRating         int          `json:"newi_rating"`
	Track              Track        `json:"track"`
}
