package response

import "time"

type Car struct {
	AiEnabled               bool                 `json:"ai_enabled"`
	AllowNumberColors       bool                 `json:"allow_number_colors"`
	AllowNumberFont         bool                 `json:"allow_number_font"`
	AllowSponsor1           bool                 `json:"allow_sponsor1"`
	AllowSponsor2           bool                 `json:"allow_sponsor2"`
	AllowWheelColor         bool                 `json:"allow_wheel_color"`
	AwardExempt             bool                 `json:"award_exempt"`
	CarDirpath              string               `json:"car_dirpath"`
	CarID                   int                  `json:"car_id"`
	CarName                 string               `json:"car_name"`
	CarNameAbbreviated      string               `json:"car_name_abbreviated"`
	CarTypes                []CarTypes           `json:"car_types"`
	CarWeight               int                  `json:"car_weight"`
	Categories              []string             `json:"categories"`
	Created                 time.Time            `json:"created"`
	FreeWithSubscription    bool                 `json:"free_with_subscription"`
	HasHeadlights           bool                 `json:"has_headlights"`
	HasMultipleDryTireTypes bool                 `json:"has_multiple_dry_tire_types"`
	Hp                      int                  `json:"hp"`
	MaxPowerAdjustPct       int                  `json:"max_power_adjust_pct"`
	MaxWeightPenaltyKg      int                  `json:"max_weight_penalty_kg"`
	MinPowerAdjustPct       int                  `json:"min_power_adjust_pct"`
	PackageID               int                  `json:"package_id"`
	PaintRules              map[string]PaintRule `json:"paint_rules"`
	Patterns                int                  `json:"patterns"`
	Price                   float64              `json:"price"`
	Retired                 bool                 `json:"retired"`
	SearchFilters           string               `json:"search_filters"`
	Sku                     int                  `json:"sku"`
}
