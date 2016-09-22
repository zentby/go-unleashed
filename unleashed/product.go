package unleashed

import (
	"fmt"
)

type ProductService service

type ProductList struct {
	Pagination Pagination
	Items      []Product
}

type SellPriceTier struct {
	Name  string `json:"Name,omitempty"`
	Value string `json:"Value,omitempty"`
}

type UnitOfMeasure struct {
	Guid string `json:"Guid,omitempty"`
	Name string `json:"Name,omitempty"`
}

type SupplierWithProductCode struct {
	Guid                       string  `json:"Guid,omitempty"`
	SupplierCode               string  `json:"SupplierCode,omitempty"`
	SupplierName               string  `json:"SupplierName,omitempty"`
	SupplierProductCode        string  `json:"SupplierProductCode,omitempty"`
	SupplierProductDescription string  `json:"SupplierProductDescription,omitempty"`
	SupplierProductPrice       float64 `json:"SupplierProductPrice,omitempty"`
}

type Product struct {
	ProductCode            string                  `json:"ProductCode,omitempty"`
	ProductDescription     string                  `json:"ProductDescription,omitempty"`
	Barcode                string                  `json:"Barcode,omitempty"`
	PackSize               float64                 `json:"PackSize,omitempty"`
	Width                  float64                 `json:"Width,omitempty"`
	Height                 float64                 `json:"Height,omitempty"`
	Depth                  float64                 `json:"Depth,omitempty"`
	Weight                 float64                 `json:"Weight,omitempty"`
	MinStockAlertLevel     float64                 `json:"MinStockAlertLevel,omitempty"`
	MaxStockAlertLevel     float64                 `json:"MaxStockAlertLevel,omitempty"`
	ReOrderPoint           float64                 `json:"ReOrderPoint,omitempty"`
	UnitOfMeasure          UnitOfMeasure           `json:"UnitOfMeasure,omitempty"`
	NeverDiminishing       bool                    `json:"NeverDiminishing,omitempty"`
	LastCost               float64                 `json:"LastCost,omitempty"`
	DefaultPurchasePrice   float64                 `json:"DefaultPurchasePrice,omitempty"`
	DefaultSellPrice       float64                 `json:"DefaultSellPrice,omitempty"`
	AverageLandPrice       float64                 `json:"AverageLandPrice,omitempty"`
	Obsolete               bool                    `json:"Obsolete,omitempty"`
	Notes                  string                  `json:"Notes,omitempty"`
	SellPriceTier1         SellPriceTier           `json:"SellPriceTier1,omitempty"`
	SellPriceTier2         SellPriceTier           `json:"SellPriceTier2,omitempty"`
	SellPriceTier3         SellPriceTier           `json:"SellPriceTier3,omitempty"`
	SellPriceTier4         SellPriceTier           `json:"SellPriceTier4,omitempty"`
	SellPriceTier5         SellPriceTier           `json:"SellPriceTier5,omitempty"`
	SellPriceTier6         SellPriceTier           `json:"SellPriceTier6,omitempty"`
	SellPriceTier7         SellPriceTier           `json:"SellPriceTier7,omitempty"`
	SellPriceTier8         SellPriceTier           `json:"SellPriceTier8,omitempty"`
	SellPriceTier9         SellPriceTier           `json:"SellPriceTier9,omitempty"`
	SellPriceTier10        SellPriceTier           `json:"SellPriceTier10,omitempty"`
	XeroTaxCode            string                  `json:"XeroTaxCode,omitempty"`
	XeroTaxRate            float64                 `json:"XeroTaxRate,omitempty"`
	TaxablePurchase        bool                    `json:"TaxablePurchase,omitempty"`
	TaxableSales           bool                    `json:"TaxableSales,omitempty"`
	XeroSalesTaxCode       string                  `json:"XeroSalesTaxCode,omitempty"`
	XeroSalesTaxRate       float64                 `json:"XeroSalesTaxRate,omitempty"`
	IsComponent            bool                    `json:"IsComponent,omitempty"`
	IsAssembledProduct     bool                    `json:"IsAssembledProduct,omitempty"`
	ProductGroup           string                  `json:"ProductGroup,omitempty"`
	XeroSalesAccount       string                  `json:"XeroSalesAccount,omitempty"`
	XeroCostOfGoodsAccount string                  `json:"XeroCostOfGoodsAccount,omitempty"`
	BinLocation            string                  `json:"BinLocation,omitempty"`
	Supplier               SupplierWithProductCode `json:"Supplier,omitempty"`
	SourceID               string                  `json:"SourceId,omitempty"`
	CreatedBy              string                  `json:"CreatedBy,omitempty"`
	SourceVariantParentID  string                  `json:"SourceVariantParentId,omitempty"`
	GUID                   string                  `json:"Guid,omitempty"`
	LastModifiedOn         string                  `json:"LastModifiedOn,omitempty"`
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

func (s *ProductService) CreateProduct(product *Product) (*Product, *Response, error) {
	u := fmt.Sprintf("products/%v", product.GUID)
	req, err := s.client.NewRequest("POST", u, product)
	if err != nil {
		return nil, nil, err
	}
	c := new(Product)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}
