package model

type ProductPrice struct {
	Ladder         float64 `json:"ladder"`
	UsdPrice       float64 `json:"usdPrice"`
	CurrencyPrice  float64 `json:"currencyPrice"`
	CurrencySymbol string  `json:"currencySymbol"`
	DiscountRate   string  `json:"discountRate"`
}

type ProductCodeResponse struct {
	StockNumber      uint64         `json:"stockNumber"`
	Description      string         `json:"productIntroEn"`
	Package          string         `json:"encapStandard"`
	DistributorPart  string         `json:"productCode"`
	Manufacturer     string         `json:"brandNameEn"`
	ManufacturerPart string         `json:"productModel"`
	ProductPriceList []ProductPrice `json:"productPriceList"`
}
