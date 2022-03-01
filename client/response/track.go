package response

type Track struct {
	TrackID    int    `json:"track_id"`
	TrackName  string `json:"track_name"`
	ConfigName string `json:"config_name"`
	CategoryID int    `json:"category_id"`
	Category   string `json:"category"`
}
