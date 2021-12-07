package dal

import (
	"MainService/dao"
	"encoding/csv"
	"io"
	"os"
)

func ReadCSVFile(filePath string) (gameSalesHistories []dao.GameSalesHistory) {
	isFirstRow := true
	headerMap := make(map[string]int)
	f, _ := os.Open(filePath)
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if isFirstRow {
			isFirstRow = false

			for i, v := range record {
				headerMap[v] = i
			}

			continue
		}
		gameSalesHistories = append(gameSalesHistories, dao.GameSalesHistory{
			Publisher: record[headerMap["Year"]],
		})
	}
	return
}
