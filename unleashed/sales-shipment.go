package unleashed

import (
	"fmt"
)

type SalesShipmentService service

type SalesShipment struct {
	OrderNumber        *string              `json:"OrderNumber"`
	ShipmentNumber     *string              `json:"ShipmentNumber"`
	ShipmentStatus     *string              `json:"ShipmentStatus"`
	DispatchDate       *string              `json:"DispatchDate"`
	TrackingNumber     *string              `json:"TrackingNumber"`
	ShippingCompany    *string              `json:"ShippingCompany"`
	Comments           *string              `json:"Comments"`
	CreatedOn          *string              `json:"CreatedOn"`
	CreatedBy          *string              `json:"CreatedBy"`
	LastModifiedBy     *string              `json:"LastModifiedBy"`
	SalesShipmentLines []*SalesShipmentLine `json:"SalesShipmentLines"`
	GUID               *string              `json:"Guid"`
	LastModifiedOn     *string              `json:"LastModifiedOn"`
}

type SalesShipmentLine struct {
	LineNumber           *int         `json:"LineNumber"`
	Product              *ProductKey  `json:"Product"`
	ShipmentQty          *int         `json:"ShipmentQty"`
	UnitCost             *float64     `json:"UnitCost"`
	CreatedOn            *string      `json:"CreatedOn"`
	CreatedBy            *string      `json:"CreatedBy"`
	LastModifiedBy       *string      `json:"LastModifiedBy"`
	SalesOrderLineNumber *int         `json:"SalesOrderLineNumber"`
	SalesOrderLineID     *int         `json:"SalesOrderLineId"`
	SerialNumbers        []*SerialKey `json:"SerialNumbers"`
	BatchNumbers         []*BatchKey  `json:"BatchNumbers"`
	GUID                 *string      `json:"Guid"`
	LastModifiedOn       *string      `json:"LastModifiedOn"`
}

type SalesShipmentList struct {
	Pagination *Pagination      `json:"Pagination"`
	Items      []*SalesShipment `json:"Items"`
}

func (i SalesShipment) String() string {
	return Stringify(i)
}

func (s *SalesShipmentService) List(opt *PageOptions, query *map[string]string) (*SalesShipmentList, *Response, error) {
	u := "salesshipments"

	salesshipments := &SalesShipmentList{}
	resp, err := s.client.GetRequestData(u, opt, query, salesshipments)
	if err != nil {
		return nil, resp, err
	}

	return salesshipments, resp, err
}

func (s *SalesShipmentService) GetSalesShipmentBy(id string) (*SalesShipment, *Response, error) {
	u := "salesshipments/" + id

	salesShipment := &SalesShipment{}
	resp, err := s.client.GetRequestData(u, nil, nil, salesShipment)
	if err != nil {
		return nil, resp, err
	}

	return salesShipment, resp, err
}

func (s *SalesShipmentService) CreateSalesShipment(salesShipment *SalesShipment) (*SalesShipment, *Response, error) {
	u := fmt.Sprintf("salesshipments/%v", *salesShipment.GUID)
	req, err := s.client.NewRequest("POST", u, salesShipment)
	if err != nil {
		return nil, nil, err
	}
	c := new(SalesShipment)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

func (s *SalesShipmentService) UpdateSalesShipment(salesShipment *SalesShipment) (*SalesShipment, *Response, error) {
	u := fmt.Sprintf("salesshipments/%v", *salesShipment.GUID)
	req, err := s.client.NewRequest("PUT", u, salesShipment)
	if err != nil {
		return nil, nil, err
	}
	c := new(SalesShipment)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

func (s *SalesShipmentService) DeleteSalesShipment(guid string) (*SalesShipment, *Response, error) {
	u := fmt.Sprintf("salesshipments/%v", guid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}
	c := new(SalesShipment)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}
