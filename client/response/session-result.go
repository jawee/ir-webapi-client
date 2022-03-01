package response

type SessionResult struct {
	SimsessionNumber   int       `json:"simsession_number"`
	SimsessionType     int       `json:"simsession_type"`
	SimsessionTypeName string    `json:"simsession_type_name"`
	SimsessionSubtype  int       `json:"simsession_subtype"`
	SimsessionName     string    `json:"simsession_name"`
	Results            []Result  `json:"results"`
}
