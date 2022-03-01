package response

import "time"

type Member struct {
	CustID      int       `json:"cust_id"`
	DisplayName string    `json:"display_name"`
	Helmet      Helmet    `json:"helmet"`
	LastLogin   time.Time `json:"last_login"`
	MemberSince string    `json:"member_since"`
	ClubID      int       `json:"club_id"`
	ClubName    string    `json:"club_name"`
	Ai          bool      `json:"ai"`
}
