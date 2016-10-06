package unleashed

type PurchaseOrderService service

type PurchaseOrderLine struct {
	GUID                string  `json:"Guid,omitempty"`
	LineNumber          int     `json:"LineNumber,omitempty"`
	Product             Product `json:"Product,omitempty"`
	DueDate             string  `json:"DueDate,omitempty"`
	OrderQuantity       float64 `json:"OrderQuantity,omitempty"`
	UnitPrice           float64 `json:"UnitPrice,omitempty"`
	LineTotal           float64 `json:"LineTotal,omitempty"`
	Volume              float64 `json:"Volume,omitempty"`
	Weight              float64 `json:"Weight,omitempty"`
	Comments            string  `json:"Comments,omitempty"`
	ReceiptQuantity     string  `json:"ReceiptQuantity,omitempty"`
	BCUnitPrice         float64 `json:"BCUnitPrice,omitempty"`
	BCSubTotal          float64 `json:"BCSubTotal,omitempty"`
	Tax                 Tax     `json:"Tax,omitempty"`
	LineTax             float64 `json:"LineTax,omitempty"`
	LastModifiedOn      string  `json:"LastModifiedOn,omitempty"`
	DiscountedUnitPrice float64 `json:"DiscountedUnitPrice,omitempty"`
	DiscountRate        float64 `json:"DiscountRate,omitempty"`
}

type Supplier struct {
	GUID         string `json:"Guid,omitempty"`
	SupplierCode string `json:"SupplierCode,omitempty"`
	SupplierName string `json:"SupplierName,omitempty"`
}

type PurchaseOrder struct {
	GUID                  string              `json:"Guid,omitempty"`
	OrderNumber           string              `json:"OrderNumber,omitempty"`
	OrderDate             string              `json:"OrderDate,omitempty"`
	RequiredDate          string              `json:"RequiredDate,omitempty"`
	CompletedDate         string              `json:"CompletedDate,omitempty"`
	Supplier              Supplier            `json:"Supplier,omitempty"`
	SupplierRef           string              `json:"SupplierRef,omitempty"`
	Comments              string              `json:"Comments,omitempty"`
	Printed               bool                `json:"Printed,omitempty"`
	OrderStatus           string              `json:"OrderStatus,omitempty"`
	ReceivedDate          string              `json:"ReceivedDate,omitempty"`
	DeliveryName          string              `json:"DeliveryName,omitempty"`
	DeliveryStreetAddress string              `json:"DeliveryStreetAddress,omitempty"`
	DeliverySuburb        string              `json:"DeliverySuburb,omitempty"`
	DeliveryRegion        string              `json:"DeliveryRegion,omitempty"`
	DeliveryCity          string              `json:"DeliveryCity,omitempty"`
	DeliveryCountry       string              `json:"DeliveryCountry,omitempty"`
	DeliveryPostCode      string              `json:"DeliveryPostCode,omitempty"`
	Currency              Currency            `json:"Currency,omitempty"`
	ExchangeRate          float64             `json:"ExchangeRate,omitempty"`
	Tax                   Tax                 `json:"Tax,omitempty"`
	TaxRate               float64             `json:"TaxRate,omitempty"`
	XeroTaxCode           string              `json:"XeroTaxCode,omitempty"`
	SubTotal              float64             `json:"SubTotal,omitempty"`
	TaxTotal              float64             `json:"TaxTotal,omitempty"`
	Total                 float64             `json:"Total,omitempty"`
	TotalVolume           float64             `json:"TotalVolume,omitempty"`
	TotalWeight           float64             `json:"TotalWeight,omitempty"`
	SupplierInvoiceDate   string              `json:"SupplierInvoiceDate,omitempty"`
	BCSubTotal            float64             `json:"BCSubTotal,omitempty"`
	BCTaxTotal            float64             `json:"BCTaxTotal,omitempty"`
	BCTotal               float64             `json:"BCTotal,omitempty"`
	PurchaseOrderLines    []PurchaseOrderLine `json:"PurchaseOrderLines,omitempty"`
	Warehouse             Warehouse           `json:"Warehouse,omitempty"`
	LastModifiedOn        string              `json:"LastModifiedOn,omitempty"`
	DiscountRate          float64             `json:"DiscountRate,omitempty"`
}
