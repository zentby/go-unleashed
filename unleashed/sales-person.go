package unleashed

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
