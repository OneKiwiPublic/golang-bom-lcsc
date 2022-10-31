package main

import (
	"fmt"
	"strconv"

	"bom/api"
	"bom/csv"
)

func main() {
	csvPath := "/home/vanson/working/kicad/radio-classd/assembly/bom_radio-classd.csv"
	boms, err := csv.OpenBOM(csvPath)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range boms {

		if v.Distributor == "LCSC" && v.DistributorPart != "" {
			fmt.Printf("Fetch LCSC Part: %s\n", v.DistributorPart)
			response, _ := api.FetchProductCode(v.DistributorPart)
			boms[i].Description = response.Description
			boms[i].DistributorPart = response.DistributorPart
			boms[i].Manufacturer = response.Manufacturer
			boms[i].ManufacturerPart = response.ManufacturerPart
			boms[i].Stock = strconv.FormatUint(response.StockNumber, 10)
			price := fmt.Sprintf("%f", response.ProductPriceList[0].UsdPrice)
			boms[i].UnitPrice = price
		}

	}
	csv.WriteBOM(boms)
	fmt.Println("done")
}
