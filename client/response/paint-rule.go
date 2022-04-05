package response

type PaintRule struct {
	PaintCarAvailable bool   `json:"PaintCarAvailable"`
	Color1            string `json:"Color1"`
	Color2            string `json:"Color2"`
	Color3            string `json:"Color3"`
	Sponsor1Available bool   `json:"Sponsor1Available"`
	Sponsor2Available bool   `json:"Sponsor2Available"`
	Sponsor1          string `json:"Sponsor1"`
	Sponsor2          string `json:"Sponsor2"`
	RulesExplanation  string `json:"RulesExplanation"`
}
