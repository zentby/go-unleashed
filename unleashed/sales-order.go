package unleashed

import (
	"fmt"
)

type SalesOrderService service

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

func (i SalesOrder) String() string {
	return Stringify(i)
}

func (s *SalesOrderService) List(opt *PageOptions, query *map[string]string) (*SalesOrderList, *Response, error) {
	u := "salesorders"

	salesOrders := &SalesOrderList{}
	resp, err := s.client.GetRequestData(u, opt, query, salesOrders)
	if err != nil {
		return nil, resp, err
	}

	return salesOrders, resp, err
}

func (s *SalesOrderService) GetSalesOrderBy(id string) (*SalesOrder, *Response, error) {
	u := "salesorders/" + id
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	salesorders := &SalesOrder{}
	resp, err := s.client.Do(req, salesorders)
	if err != nil {
		return nil, resp, err
	}

	return salesorders, resp, err
}

func (s *SalesOrderService) CreateSalesOrder(salesorder *SalesOrder) (*SalesOrder, *Response, error) {
	u := fmt.Sprintf("salesorders/%v", *salesorder.GUID)
	req, err := s.client.NewRequest("POST", u, salesorder)
	if err != nil {
		return nil, nil, err
	}
	c := new(SalesOrder)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

func (s *SalesOrderService) UpdateSalesOrder(salesorder *SalesOrder) (*SalesOrder, *Response, error) {
	u := fmt.Sprintf("salesorders/%v", *salesorder.GUID)
	req, err := s.client.NewRequest("PUT", u, salesorder)
	if err != nil {
		return nil, nil, err
	}
	c := new(SalesOrder)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

func (s *SalesOrderService) CompleteSalesOrder(guid string) (*SalesOrder, *Response, error) {
	u := fmt.Sprintf("salesorders/%v/complete", guid)
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}
	c := new(SalesOrder)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}
