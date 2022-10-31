package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/tushar2708/altcsv"
)

type BOM struct {
	Item             string
	Category         string
	Value            string
	Package          string
	Reference        string
	Description      string
	Assembly         string
	Distributor      string
	DistributorPart  string
	Manufacturer     string
	ManufacturerPart string
	Quantity         string
	Stock            string
	UnitPrice        string
}

const (
	ITEM              = 0
	CATEGORY          = 1
	VALUE             = 2
	REFERENCES        = 3
	PACKAGE           = 4
	DESCRIPTION       = 5
	ASSEMBLY          = 6
	DISTRIBUTOR       = 7
	DISTRIBUTOR_PART  = 8
	MANUFACTURER      = 9
	MANUFACTURER_PART = 10
	QUANTITY          = 11
)

func OpenBOM(fileName string) ([]BOM, error) {

	data := []BOM{}

	csvFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error open file", fileName)
		return nil, fmt.Errorf("Error open file BOM")
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	reader.Comma = ','
	reader.LazyQuotes = true

	csvLines, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}
	for i, items := range csvLines {
		if i > 0 {

			item := BOM{
				Item:             items[ITEM],
				Category:         items[CATEGORY],
				Value:            items[VALUE],
				Package:          items[PACKAGE],
				Reference:        items[REFERENCES],
				Description:      items[DESCRIPTION],
				Assembly:         items[ASSEMBLY],
				Distributor:      items[DISTRIBUTOR],
				DistributorPart:  items[DISTRIBUTOR_PART],
				Manufacturer:     items[MANUFACTURER],
				ManufacturerPart: items[MANUFACTURER_PART],
				Quantity:         items[QUANTITY],
			}
			data = append(data, item)
		}
	}
	return data, nil
}

type Employee struct {
	ID  string
	Age int
}

func WriteBOM(records []BOM) {
	headers := []string{"Item", "Category", "Value", "References",
		"Package", "Description", "Assembly", "Distributor",
		"Distributor Part#", "Manufacturer", "Manufacturer Part#",
		"Quantity", "Stock", "Unit Price"}
	fileWtr, _ := os.Create("bom.csv")
	csvWtr := altcsv.NewWriter(fileWtr)
	csvWtr.Quote = '"'      // use | as "quote"
	csvWtr.AllQuotes = true // surround each field with '|'
	csvWtr.Write(headers)
	for _, v := range records {
		row := []string{v.Item, v.Category,
			v.Value, v.Reference, v.Package, v.Description, v.Assembly,
			v.Distributor, v.DistributorPart, v.Manufacturer, v.ManufacturerPart,
			v.Quantity, v.Stock, v.UnitPrice,
		}
		if err := csvWtr.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	csvWtr.Flush()
	fileWtr.Close()
}
