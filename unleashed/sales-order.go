package unleashed

type SalesOrder struct {
	SalesOrderLines           []*SalesOrderLine `json:"SalesOrderLines"`
	OrderNumber               *string           `json:"OrderNumber"`
	OrderDate                 *string           `json:"OrderDate"`
	RequiredDate              *string           `json:"RequiredDate"`
	OrderStatus               *string           `json:"OrderStatus"`
	Customer                  *CustomerKey      `json:"Customer"`
	CustomerRef               *string           `json:"CustomerRef"`
	Comments                  *string           `json:"Comments"`
	Warehouse                 *Warehouse        `json:"Warehouse"`
	ReceivedDate              *string           `json:"ReceivedDate"`
	DeliveryName              *string           `json:"DeliveryName"`
	DeliveryStreetAddress     *string           `json:"DeliveryStreetAddress"`
	DeliverySuburb            *string           `json:"DeliverySuburb"`
	DeliveryCity              *string           `json:"DeliveryCity"`
	DeliveryRegion            *string           `json:"DeliveryRegion"`
	DeliveryCountry           *string           `json:"DeliveryCountry"`
	DeliveryPostCode          *string           `json:"DeliveryPostCode"`
	Currency                  *Currenry         `json:"Currency"`
	ExchangeRate              *float64          `json:"ExchangeRate"`
	DiscountRate              *float64          `json:"DiscountRate"`
	Tax                       *Tax              `json:"Tax"`
	TaxRate                   *float64          `json:"TaxRate"`
	XeroTaxCode               *string           `json:"XeroTaxCode"`
	SubTotal                  *float64          `json:"SubTotal"`
	TaxTotal                  *float64          `json:"TaxTotal"`
	Total                     *float64          `json:"Total"`
	TotalVolume               *float64          `json:"TotalVolume"`
	TotalWeight               *float64          `json:"TotalWeight"`
	BCSubTotal                *float64          `json:"BCSubTotal"`
	BCTaxTotal                *float64          `json:"BCTaxTotal"`
	BCTotal                   *float64          `json:"BCTotal"`
	PaymentDueDate            *string           `json:"PaymentDueDate"`
	AllocateProduct           *bool             `json:"AllocateProduct"`
	SalesOrderGroup           *string           `json:"SalesOrderGroup"`
	DeliveryMethod            *string           `json:"DeliveryMethod"`
	SalesPerson               *SalesPerson      `json:"SalesPerson"`
	SendAccountingJournalOnly *bool             `json:"SendAccountingJournalOnly"`
	SourceID                  *string           `json:"SourceId"`
	CreatedBy                 *string           `json:"CreatedBy"`
	GUID                      *string           `json:"Guid"`
	LastModifiedOn            *string           `json:"LastModifiedOn"`
}

type SalesOrderLine struct {
	LineNumber                     *int         `json:"LineNumber"`
	LineType                       *string      `json:"LineType"`
	Product                        *ProductKey  `json:"Product"`
	DueDate                        *string      `json:"DueDate"`
	OrderQuantity                  *float64     `json:"OrderQuantity"`
	UnitPrice                      *float64     `json:"UnitPrice"`
	DiscountRate                   *float64     `json:"DiscountRate"`
	LineTotal                      *float64     `json:"LineTotal"`
	Volume                         *float64     `json:"Volume"`
	Weight                         *float64     `json:"Weight"`
	Comments                       *string      `json:"Comments"`
	AverageLandedPriceAtTimeOfSale *float64     `json:"AverageLandedPriceAtTimeOfSale"`
	TaxRate                        *float64     `json:"TaxRate"`
	LineTax                        *float64     `json:"LineTax"`
	XeroTaxCode                    *string      `json:"XeroTaxCode"`
	BCUnitPrice                    *float64     `json:"BCUnitPrice"`
	BCLineTotal                    *float64     `json:"BCLineTotal"`
	BCLineTax                      *float64     `json:"BCLineTax"`
	LineTaxCode                    *string      `json:"LineTaxCode"`
	XeroSalesAccount               *string      `json:"XeroSalesAccount"`
	SerialNumbers                  []*SerialKey `json:"SerialNumbers"`
	BatchNumbers                   []*BatchKey  `json:"BatchNumbers"`
	GUID                           *string      `json:"Guid"`
	LastModifiedOn                 *string      `json:"LastModifiedOn"`
}

type SerialKey struct {
	Identifier *string `json:"Identifier"`
}

type BatchKey struct {
	Number *string `json:"Number"`
}

type ProductKey struct {
	GUID               *string `json:"Guid"`
	ProductCode        *string `json:"ProductCode"`
	ProductDescription *string `json:"ProductDescription"`
}

type CustomerKey struct {
	CustomerCode   *string `json:"CustomerCode"`
	CustomerName   *string `json:"CustomerName"`
	CurrencyID     *int    `json:"CurrencyId"`
	GUID           *string `json:"Guid"`
	LastModifiedOn *string `json:"LastModifiedOn"`
}

type SalesOrderList struct {
	Pagination *Pagination   `json:"Pagination"`
	Items      []*SalesOrder `json:"Items"`
}
