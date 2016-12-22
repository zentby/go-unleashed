package unleashed

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
