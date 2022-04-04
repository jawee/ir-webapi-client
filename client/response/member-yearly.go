package response 

type MemberYearly struct {
	Stats  []Stats `json:"stats"`
	CustID int     `json:"cust_id"`
}
