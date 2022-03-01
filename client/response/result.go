package response

import "time"

type Result struct {
	CustID                int       `json:"cust_id"`
	DisplayName           string    `json:"display_name"`
	FinishPosition        int       `json:"finish_position"`
	FinishPositionInClass int       `json:"finish_position_in_class"`
	LapsLead              int       `json:"laps_lead"`
	LapsComplete          int       `json:"laps_complete"`
	OptLapsComplete       int       `json:"opt_laps_complete"`
	Interval              int       `json:"interval"`
	ClassInterval         int       `json:"class_interval"`
	AverageLap            int       `json:"average_lap"`
	BestLapNum            int       `json:"best_lap_num"`
	BestLapTime           int       `json:"best_lap_time"`
	BestNlapsNum          int       `json:"best_nlaps_num"`
	BestNlapsTime         int       `json:"best_nlaps_time"`
	BestQualLapAt         time.Time `json:"best_qual_lap_at"`
	BestQualLapNum        int       `json:"best_qual_lap_num"`
	BestQualLapTime       int       `json:"best_qual_lap_time"`
	ReasonOutID           int       `json:"reason_out_id"`
	ReasonOut             string    `json:"reason_out"`
	ChampPoints           int       `json:"champ_points"`
	DropRace              bool      `json:"drop_race"`
	ClubPoints            int       `json:"club_points"`
	Position              int       `json:"position"`
	QualLapTime           int       `json:"qual_lap_time"`
	StartingPosition      int       `json:"starting_position"`
	CarClassID            int       `json:"car_class_id"`
	ClubID                int       `json:"club_id"`
	ClubName              string    `json:"club_name"`
	ClubShortname         string    `json:"club_shortname"`
	Division              int       `json:"division"`
	DivisionName          string    `json:"division_name"`
	OldLicenseLevel       int       `json:"old_license_level"`
	OldSubLevel           int       `json:"old_sub_level"`
	OldCpi                int       `json:"old_cpi"`
	OldiRating            int       `json:"oldi_rating"`
	OldTtrating           int       `json:"old_ttrating"`
	NewLicenseLevel       int       `json:"new_license_level"`
	NewSubLevel           int       `json:"new_sub_level"`
	NewCpi                int       `json:"new_cpi"`
	NewiRating            int       `json:"newi_rating"`
	NewTtrating           int       `json:"new_ttrating"`
	Multiplier            int       `json:"multiplier"`
	LicenseChangeOval     int       `json:"license_change_oval"`
	LicenseChangeRoad     int       `json:"license_change_road"`
	Incidents             int       `json:"incidents"`
	MaxPctFuelFill        int       `json:"max_pct_fuel_fill"`
	WeightPenaltyKg       int       `json:"weight_penalty_kg"`
	LeaguePoints          int       `json:"league_points"`
	LeagueAggPoints       int       `json:"league_agg_points"`
	CarID                 int       `json:"car_id"`
	AggregateChampPoints  int       `json:"aggregate_champ_points"`
	Livery                Livery    `json:"livery"`
	Suit                  Suit      `json:"suit"`
	Helmet                Helmet    `json:"helmet"`
	Watched               bool      `json:"watched"`
	Friend                bool      `json:"friend"`
	Ai                    bool      `json:"ai"`
}
