package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

type Std [][]string

func main() {
	rows, err := parseCSV("resources/grades.csv")
	if err != nil {
		log.Fatal(err)
	}

	tw := tablewriter.NewWriter(os.Stdout)
	for i, s := range rows {
		if i == 0 {
			tw.SetHeader(s)
			continue
		}
		tw.Append(s)
	}
	tw.Render()
}

func parseCSV(path string) (Std, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	var st Std
	r := csv.NewReader(csvFile)
	rows, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		st = append(st, row)
	}
	return st, nil
}
