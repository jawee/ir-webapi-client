package response

type MemberSummary struct {
	ThisYear ThisYear `json:"this_year"`
	CustID   int      `json:"cust_id"`
}
