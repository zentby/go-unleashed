package unleashed

import (
	"fmt"
)

type SalesPersonService service

type SalesPerson struct {
	FullName       *string `json:"FullName"`
	Email          *string `json:"Email"`
	Obsolete       *bool   `json:"Obsolete"`
	GUID           *string `json:"Guid"`
	LastModifiedOn *string `json:"LastModifiedOn"`
}

type SalesPersonList struct {
	Items []*SalesPerson `json:"Items"`
}

func (i SalesPerson) String() string {
	return Stringify(i)
}

func (s *SalesPersonService) List(opt *PageOptions, query *map[string]string) (*SalesPersonList, *Response, error) {
	u := "salespersons"

	salespersons := &SalesPersonList{}
	resp, err := s.client.GetRequestData(u, opt, query, salespersons)
	if err != nil {
		return nil, resp, err
	}

	return salespersons, resp, err
}

func (s *SalesPersonService) GetSalesPersonBy(id string) (*SalesPerson, *Response, error) {
	u := "salespersons/" + id

	salesPerson := &SalesPerson{}
	resp, err := s.client.GetRequestData(u, nil, nil, salesPerson)
	if err != nil {
		return nil, resp, err
	}

	return salesPerson, resp, err
}

func (s *SalesPersonService) CreateSalesPerson(salesPerson *SalesPerson) (*SalesPerson, *Response, error) {
	u := fmt.Sprintf("salespersons/%v", *salesPerson.GUID)
	req, err := s.client.NewRequest("POST", u, salesPerson)
	if err != nil {
		return nil, nil, err
	}
	c := new(SalesPerson)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

func (s *SalesPersonService) UpdateSalesPerson(salesPerson *SalesPerson) (*SalesPerson, *Response, error) {
	u := fmt.Sprintf("salespersons/%v", *salesPerson.GUID)
	req, err := s.client.NewRequest("POST", u, salesPerson)
	if err != nil {
		return nil, nil, err
	}
	c := new(SalesPerson)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}
