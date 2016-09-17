package unleashed

type ProductService service

type ProductList struct {
	Pagination Pagination
	Items      []Product
}

type SellPriceTier struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type UnitOfMeasure struct {
	Guid string `json:"Guid"`
	Name string `json:"Name"`
}

type SupplierWithProductCode struct {
	Guid                       string  `json:"Guid"`
	SupplierCode               string  `json:"SupplierCode"`
	SupplierName               string  `json:"SupplierName"`
	SupplierProductCode        string  `json:"SupplierProductCode"`
	SupplierProductDescription string  `json:"SupplierProductDescription"`
	SupplierProductPrice       float64 `json:"SupplierProductPrice"`
}

type Product struct {
	ProductCode            string                  `json:"ProductCode"`
	ProductDescription     string                  `json:"ProductDescription"`
	Barcode                string                  `json:"Barcode"`
	PackSize               float64                 `json:"PackSize"`
	Width                  float64                 `json:"Width"`
	Height                 float64                 `json:"Height"`
	Depth                  float64                 `json:"Depth"`
	Weight                 float64                 `json:"Weight"`
	MinStockAlertLevel     float64                 `json:"MinStockAlertLevel"`
	MaxStockAlertLevel     float64                 `json:"MaxStockAlertLevel"`
	ReOrderPoint           float64                 `json:"ReOrderPoint"`
	UnitOfMeasure          UnitOfMeasure           `json:"UnitOfMeasure"`
	NeverDiminishing       bool                    `json:"NeverDiminishing"`
	LastCost               float64                 `json:"LastCost"`
	DefaultPurchasePrice   float64                 `json:"DefaultPurchasePrice"`
	DefaultSellPrice       float64                 `json:"DefaultSellPrice"`
	AverageLandPrice       float64                 `json:"AverageLandPrice"`
	Obsolete               bool                    `json:"Obsolete"`
	Notes                  string                  `json:"Notes"`
	SellPriceTier1         SellPriceTier           `json:"SellPriceTier1"`
	SellPriceTier2         SellPriceTier           `json:"SellPriceTier2"`
	SellPriceTier3         SellPriceTier           `json:"SellPriceTier3"`
	SellPriceTier4         SellPriceTier           `json:"SellPriceTier4"`
	SellPriceTier5         SellPriceTier           `json:"SellPriceTier5"`
	SellPriceTier6         SellPriceTier           `json:"SellPriceTier6"`
	SellPriceTier7         SellPriceTier           `json:"SellPriceTier7"`
	SellPriceTier8         SellPriceTier           `json:"SellPriceTier8"`
	SellPriceTier9         SellPriceTier           `json:"SellPriceTier9"`
	SellPriceTier10        SellPriceTier           `json:"SellPriceTier10"`
	XeroTaxCode            string                  `json:"XeroTaxCode"`
	XeroTaxRate            float64                 `json:"XeroTaxRate"`
	TaxablePurchase        bool                    `json:"TaxablePurchase"`
	TaxableSales           bool                    `json:"TaxableSales"`
	XeroSalesTaxCode       string                  `json:"XeroSalesTaxCode"`
	XeroSalesTaxRate       float64                 `json:"XeroSalesTaxRate"`
	IsComponent            bool                    `json:"IsComponent"`
	IsAssembledProduct     bool                    `json:"IsAssembledProduct"`
	ProductGroup           string                  `json:"ProductGroup"`
	XeroSalesAccount       string                  `json:"XeroSalesAccount"`
	XeroCostOfGoodsAccount string                  `json:"XeroCostOfGoodsAccount"`
	BinLocation            string                  `json:"BinLocation"`
	Supplier               SupplierWithProductCode `json:"Supplier"`
	SourceID               string                  `json:"SourceId"`
	CreatedBy              string                  `json:"CreatedBy"`
	SourceVariantParentID  string                  `json:"SourceVariantParentId"`
	GUID                   string                  `json:"Guid"`
	LastModifiedOn         string                  `json:"LastModifiedOn"`
}

func (i Product) String() string {
	return Stringify(i)
}

func (s *ProductService) List(opt *PageOptions) (*ProductList, *Response, error) {
	u := "products"
	if opt != nil {
		t, err := addOptions(u, *opt)
		if err != nil {
			return nil, nil, err
		}
		u = t
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove custom Accept header when this API fully launches.
	req.Header.Set("Accept", jsontype)

	products := &ProductList{}
	resp, err := s.client.Do(req, products)
	if err != nil {
		return nil, resp, err
	}

	return products, resp, err
}
