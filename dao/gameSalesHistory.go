package dao

import "strconv"

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

func ConvertToGameSalesHistory(record []string, headerMap map[string]int) GameSalesHistory {
	rank, err := strconv.Atoi(record[headerMap["Rank"]])
	year, err := strconv.Atoi(record[headerMap["Year"]])
	NASales, err := strconv.ParseFloat(record[headerMap["NA_Sales"]], 64)
	EUSales, err := strconv.ParseFloat(record[headerMap["EU_Sales"]], 64)
	JPSales, err := strconv.ParseFloat(record[headerMap["JP_Sales"]], 64)
	OtherSales, err := strconv.ParseFloat(record[headerMap["Other_Sales"]], 64)
	GlobalSales, err := strconv.ParseFloat(record[headerMap["Global_Sales"]], 64)
	if err != nil {
		panic(err.Error())
	}
	return GameSalesHistory{
		Rank:        uint(rank),
		Name:        record[headerMap["Name"]],
		Platform:    record[headerMap["Platform"]],
		Year:        year,
		Genre:       record[headerMap["Genre"]],
		Publisher:   record[headerMap["Publisher"]],
		NASales:     float32(NASales),
		EUSales:     float32(EUSales),
		JPSales:     float32(JPSales),
		OtherSales:  float32(OtherSales),
		GlobalSales: float32(GlobalSales),
	}
}
