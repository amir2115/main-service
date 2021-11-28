package dao

type GameSalesHistory struct {
	Rank        uint    `json:"rank"`
	Name        string  `json:"name"`
	Platform    string  `json:"platform"`
	Year        int     `json:"year"`
	Genre       string  `json:"genre"`
	Publisher   string  `json:"publisher"`
	NASales     float32 `json:"na_sales"`
	EUSales     float32 `json:"eu_sales"`
	JPSales     float32 `json:"jp_sales"`
	OtherSales  float32 `json:"other_sales"`
	GlobalSales float32 `json:"global_sales"`
}
