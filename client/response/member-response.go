package response

type MemberResponse struct {
	Success bool      `json:"success"`
	CustIds []int     `json:"cust_ids"`
	Members []Member `json:"members"`
}

