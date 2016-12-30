package unleashed

import "fmt"

type StockAdjustmentService service

type StockAdjustment struct {
	AdjustmentNumber     *string                `json:"AdjustmentNumber"`
	Warehouse            *Warehouse             `json:"Warehouse"`
	AdjustmentDate       *string                `json:"AdjustmentDate"`
	AdjustmentReason     *string                `json:"AdjustmentReason"`
	Status               *string                `json:"Status"`
	StockAdjustmentLines []*StockAdjustmentLine `json:"StockAdjustmentLines"`
	ConfirmedOn          *string                `json:"ConfirmedOn"`
	ConfirmedBy          *string                `json:"ConfirmedBy"`
	AccountCode          *string                `json:"AccountCode"`
	GUID                 *string                `json:"Guid"`
	LastModifiedOn       *string                `json:"LastModifiedOn"`
}

type StockAdjustmentLine struct {
	LineNumber     *int        `json:"LineNumber"`
	Product        *ProductKey `json:"Product"`
	NewQuantity    *int        `json:"NewQuantity"`
	NewActualValue *int        `json:"NewActualValue"`
	Comments       *string     `json:"Comments"`
	GUID           *string     `json:"Guid"`
	LastModifiedOn *string     `json:"LastModifiedOn"`
}

type StockAdjustmentList struct {
	Items []*StockAdjustment `json:"Items"`
}

func (i StockAdjustment) String() string {
	return Stringify(i)
}

func (s *StockAdjustmentService) List(opt *PageOptions, query *map[string]string) (*StockAdjustmentList, *Response, error) {
	u := "stockadjustments"

	stockadjustments := &StockAdjustmentList{}
	resp, err := s.client.GetRequestData(u, opt, query, stockadjustments)
	if err != nil {
		return nil, resp, err
	}

	return stockadjustments, resp, err
}

func (s *StockAdjustmentService) GetStockAdjustmentBy(id string) (*StockAdjustment, *Response, error) {
	u := "stockadjustments/" + id

	stockAdjustment := &StockAdjustment{}
	resp, err := s.client.GetRequestData(u, nil, nil, stockAdjustment)
	if err != nil {
		return nil, resp, err
	}

	return stockAdjustment, resp, err
}

func (s *StockAdjustmentService) CreateStockAdjustment(stockAdjustment *StockAdjustment) (*StockAdjustment, *Response, error) {
	u := fmt.Sprintf("stockadjustments/%v", *stockAdjustment.GUID)
	req, err := s.client.NewRequest("POST", u, stockAdjustment)
	if err != nil {
		return nil, nil, err
	}
	c := new(StockAdjustment)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

func (s *StockAdjustmentService) UpdateStockAdjustment(stockAdjustment *StockAdjustment) (*StockAdjustment, *Response, error) {
	u := fmt.Sprintf("stockadjustments/%v", *stockAdjustment.GUID)
	req, err := s.client.NewRequest("POST", u, stockAdjustment)
	if err != nil {
		return nil, nil, err
	}
	c := new(StockAdjustment)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}
