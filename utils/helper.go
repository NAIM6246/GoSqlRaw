package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var csvWriter *CsvWriter

type CsvWriter struct {
	csvWriter *csv.Writer
}

func InitFileWriter() error {
	file, err := os.Create("response.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
		return err
	}
	w := csv.NewWriter(file)
	csvWriter = &CsvWriter{
		csvWriter: w,
	}
	return nil
}

func GetCSVFileWriter() *CsvWriter {
	return csvWriter
}

func (w *CsvWriter) Write(data interface{}, funcName, message string) {
	defer w.csvWriter.Flush()
	dataArray, err := convertStructToString(data, funcName)
	if err != nil {
		log.Println("failed to marshal data", err)
		return
	}
	fmt.Println(dataArray)
	if err := w.csvWriter.Write(dataArray); err != nil {
		log.Println("failed to write in csv", err)
		return
	}
	fmt.Println("written log in csv")
}

func convertStructToString(data interface{}, funcName string) ([]string, error) {
	tempData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	row := make([]string, 0)
	row = append(row, funcName)
	row = append(row, time.Now().String())
	row = append(row, string(tempData))
	return row, nil
}
