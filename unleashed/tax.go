package unleashed

type TaxService service

type Tax struct {
	TaxCode            *string  `json:"TaxCode,omitempty"`
	Description        *string  `json:"Description,omitempty"`
	TaxRate            *float64 `json:"TaxRate,omitempty"`
	CanApplyToExpenses *bool    `json:"CanApplyToExpenses,omitempty"`
	CanApplyToRevenue  *bool    `json:"CanApplyToRevenue,omitempty"`
	Obsolete           *bool    `json:"Obsolete,omitempty"`
	GUID               *string  `json:"Guid,omitempty"`
	LastModifiedOn     *string  `json:"LastModifiedOn,omitempty"`
}
