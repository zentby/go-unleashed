package unleashed

import (
	"fmt"
)

type CustomerService service

type CustomerList struct {
	Pagination *Pagination `json:"Pagination"`
	Items      []*Customer `json:"Items"`
}

type Address struct {
	AddressType   *string `json:"AddressType,omitempty"`
	AddressName   *string `json:"AddressName,omitempty"`
	StreetAddress *string `json:"StreetAddress,omitempty"`
	Suburb        *string `json:"Suburb,omitempty"`
	City          *string `json:"City,omitempty"`
	Region        *string `json:"Region,omitempty"`
	Country       *string `json:"Country,omitempty"`
	PostalCode    *string `json:"PostalCode,omitempty"`
}

type Currency struct {
	CurrencyCode   *string `json:"CurrencyCode,omitempty"`
	Description    *string `json:"Description,omitempty"`
	GUID           *string `json:"Guid,omitempty"`
	LastModifiedOn *string `json:"LastModifiedOn,omitempty"`
}

type SellPriceTierReference struct {
	Reference *string `json:"Reference,omitempty"`
}

type Customer struct {
	Addresses                        []*Address              `json:"Addresses,omitempty"`
	CustomerCode                     *string                 `json:"CustomerCode,omitempty"`
	CustomerName                     *string                 `json:"CustomerName,omitempty"`
	GSTVATNumber                     *string                 `json:"GSTVATNumber,omitempty"`
	BankName                         *string                 `json:"BankName,omitempty"`
	BankBranch                       *string                 `json:"BankBranch,omitempty"`
	BankAccount                      *string                 `json:"BankAccount,omitempty"`
	Website                          *string                 `json:"Website,omitempty"`
	PhoneNumber                      *string                 `json:"PhoneNumber,omitempty"`
	FaxNumber                        *string                 `json:"FaxNumber,omitempty"`
	MobileNumber                     *string                 `json:"MobileNumber,omitempty"`
	DDINumber                        *string                 `json:"DDINumber,omitempty"`
	TollFreeNumber                   *string                 `json:"TollFreeNumber,omitempty"`
	Email                            *string                 `json:"Email,omitempty"`
	EmailCC                          *string                 `json:"EmailCC,omitempty"`
	Currency                         *Currency               `json:"Currency,omitempty"`
	Notes                            *string                 `json:"Notes,omitempty"`
	Taxable                          *bool                   `json:"Taxable,omitempty"`
	XeroContactID                    *string                 `json:"XeroContactId,omitempty"`
	SalesPerson                      *SalesPerson            `json:"SalesPerson,omitempty"`
	DiscountRate                     *float64                `json:"DiscountRate,omitempty"`
	PrintPackingSlipInsteadOfInvoice *bool                   `json:"PrintPackingSlipInsteadOfInvoice,omitempty"`
	PrintInvoice                     *bool                   `json:"PrintInvoice,omitempty"`
	StopCredit                       *bool                   `json:"StopCredit,omitempty"`
	Obsolete                         *bool                   `json:"Obsolete,omitempty"`
	XeroSalesAccount                 *string                 `json:"XeroSalesAccount,omitempty"`
	XeroCostOfGoodsAccount           *string                 `json:"XeroCostOfGoodsAccount,omitempty"`
	SellPriceTier                    *string                 `json:"SellPriceTier,omitempty"`
	SellPriceTierReference           *SellPriceTierReference `json:"SellPriceTierReference,omitempty"`
	CustomerType                     *string                 `json:"CustomerType,omitempty"`
	PaymentTerm                      *string                 `json:"PaymentTerm,omitempty"`
	ContactFirstName                 *string                 `json:"ContactFirstName,omitempty"`
	ContactLastName                  *string                 `json:"ContactLastName,omitempty"`
	SourceID                         *string                 `json:"SourceId,omitempty"`
	CreatedBy                        *string                 `json:"CreatedBy,omitempty"`
	GUID                             *string                 `json:"Guid,omitempty"`
	LastModifiedOn                   *string                 `json:"LastModifiedOn,omitempty"`
}

func (i Customer) String() string {
	return Stringify(i)
}

func (s *CustomerService) List(opt *PageOptions, query *map[string]string) (*CustomerList, *Response, error) {
	u := "customers"
	customers := &CustomerList{}
	resp, err := s.client.GetRequestData(u, opt, query, customers)
	if err != nil {
		return nil, resp, err
	}

	return customers, resp, err
}

func (s *CustomerService) GetCustomerBy(id string) (*Customer, *Response, error) {
	u := "customer/" + id
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	customer := &Customer{}
	resp, err := s.client.Do(req, customer)
	if err != nil {
		return nil, resp, err
	}

	return customer, resp, err
}

func (s *CustomerService) CreateCustomer(customer *Customer) (*Customer, *Response, error) {
	u := fmt.Sprintf("customers/%v", *customer.GUID)
	req, err := s.client.NewRequest("POST", u, customer)
	if err != nil {
		return nil, nil, err
	}
	c := new(Customer)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}
