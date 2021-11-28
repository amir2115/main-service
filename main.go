package main

import (
	"MainService/dao"
	"encoding/csv"
	"io"
	"os"
)

func main() {
	//config.ConnectDatabase()
	//r := routes.SetupRouter()
	//
	//r.Run("0.0.0.0:8000")
	readCSVFile("assets/vgsales.csv")
}

func readCSVFile(filePath string) (gameSalesHistories []dao.GameSalesHistory) {
	isFirstRow := true
	headerMap := make(map[string]int)
	f, _ := os.Open(filePath)
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		//checkError("Some other error occurred", err)
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
