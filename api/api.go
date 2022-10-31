package api

import (
	"bom/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	URL          = "https://wwwapi.lcsc.com"
	PRODUCT_CODE = URL + "/v1/products/detail?product_code="
)

func FetchProductCode(code string) (model.ProductCodeResponse, error) {
	//We make HTTP request using the Get function
	path := PRODUCT_CODE + code
	resp, err := http.Get(path)
	if err != nil {
		fmt.Println("ooopsss an error occurred, please try again")
		log.Fatal("ooopsss an error occurred, please try again")
	}
	defer resp.Body.Close()
	//Create a variable of the same type as our model
	var response model.ProductCodeResponse
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}
	//Invoke the text output function & return it with nil as the error value
	return response, nil
}
