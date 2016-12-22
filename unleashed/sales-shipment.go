package unleashed

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
