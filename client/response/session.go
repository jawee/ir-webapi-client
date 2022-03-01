package response

import "time"

type Session struct {
	SubsessionID          int               `json:"subsession_id"`
	SeasonID              int               `json:"season_id"`
	SeasonName            string            `json:"season_name"`
	SeasonShortName       string            `json:"season_short_name"`
	SeasonYear            int               `json:"season_year"`
	SeasonQuarter         int               `json:"season_quarter"`
	SeriesID              int               `json:"series_id"`
	SeriesName            string            `json:"series_name"`
	SeriesShortName       string            `json:"series_short_name"`
	SeriesLogo            string            `json:"series_logo"`
	RaceWeekNum           int               `json:"race_week_num"`
	SessionID             int               `json:"session_id"`
	LicenseCategoryID     int               `json:"license_category_id"`
	LicenseCategory       string            `json:"license_category"`
	PrivateSessionID      int               `json:"private_session_id"`
	StartTime             time.Time         `json:"start_time"`
	EndTime               time.Time         `json:"end_time"`
	NumLapsForQualAverage int               `json:"num_laps_for_qual_average"`
	NumLapsForSoloAverage int               `json:"num_laps_for_solo_average"`
	CornersPerLap         int               `json:"corners_per_lap"`
	CautionType           int               `json:"caution_type"`
	EventType             int               `json:"event_type"`
	EventTypeName         string            `json:"event_type_name"`
	DriverChanges         bool              `json:"driver_changes"`
	MinTeamDrivers        int               `json:"min_team_drivers"`
	MaxTeamDrivers        int               `json:"max_team_drivers"`
	DriverChangeRule      int               `json:"driver_change_rule"`
	DriverChangeParam1    int               `json:"driver_change_param1"`
	DriverChangeParam2    int               `json:"driver_change_param2"`
	MaxWeeks              int               `json:"max_weeks"`
	PointsType            string            `json:"points_type"`
	EventStrengthOfField  int               `json:"event_strength_of_field"`
	EventAverageLap       int               `json:"event_average_lap"`
	EventLapsComplete     int               `json:"event_laps_complete"`
	NumCautions           int               `json:"num_cautions"`
	NumCautionLaps        int               `json:"num_caution_laps"`
	NumLeadChanges        int               `json:"num_lead_changes"`
	OfficialSession       bool              `json:"official_session"`
	HeatInfoID            int               `json:"heat_info_id"`
	SpecialEventType      int               `json:"special_event_type"`
	DamageModel           int               `json:"damage_model"`
	CanProtest            bool              `json:"can_protest"`
	CooldownMinutes       int               `json:"cooldown_minutes"`
	LimitMinutes          int               `json:"limit_minutes"`
	Track                 Track             `json:"track"`
	Weather               Weather           `json:"weather"`
	TrackState            TrackState        `json:"track_state"`
	SessionResults        []SessionResult   `json:"session_results"`
	RaceSummary           RaceSummary       `json:"race_summary"`
	CarClasses            []CarClass        `json:"car_classes"`
	AllowedLicenses       []AllowedLicense  `json:"allowed_licenses"`
	ResultsRestricted     bool              `json:"results_restricted"`
}
