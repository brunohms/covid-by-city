package main

import (
	"covid-crawler/Helper"
	"fmt"
	"github.com/tealeg/xlsx/v3"
)

func ProcessCity(city Helper.City) {
	fmt.Println("Processing:", city.Name())

	//downloadLinksFile := "/tmp/download_links"
	//
	//if _, err := os.Stat("downloads/"); os.IsNotExist(err) {
	//	os.Mkdir("downloads/", os.ModePerm)
	//}
	//
	//if _, err := os.Stat("downloads/" + city.Name()); os.IsNotExist(err) {
	//	os.Mkdir("downloads/" + city.Name(), os.ModePerm)
	//}
	//city.FetchDownloadLinks(downloadLinksFile)
	//city.DownloadFiles(downloadLinksFile)
	covidData := city.ProcessFiles()
	if covidData == nil {
		fmt.Println("Error processing files from city:", city.Name())
		return
	}
}

func cellVisitor(c *xlsx.Cell) error {
	value, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Cell value:", value)
	}
	return err
}

func rowVisitor(r *xlsx.Row) error {
	return r.ForEachCell(cellVisitor)
}

func rowStuff() {
	filename := "/home/brunohs/Downloads/xlsx_laboratorios.xlsx"
	wb, err := xlsx.OpenFile(filename)
	if err != nil {
		panic(err)
	}

	for i, sh := range wb.Sheets {
		fmt.Println("Max row is", sh.MaxRow, i, sh.Name)
		row, err := sh.Row(0)
		if err != nil {
			panic(err)
		}
		rowVisitor(row)
		//err = sh.ForEachRow(rowVisitor)
		//if err != nil {
		//	panic(err)
		//}
	}
}

func main() {
	fmt.Println("== xlsx package tutorial ==")
	rowStuff()

	//var cities []Helper.City
	//
	//cities = append(cities, Cities.OuroPreto{})
	//
	//for _, city := range cities {
	//	ProcessCity(city)
	//}
}
