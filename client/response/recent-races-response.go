package response

type RecentRacesResponse struct {
	Races  []Race `json:"races"`
	CustID int     `json:"cust_id"`
}
