package unleashed

type WarehouseService service

type Warehouse struct {
	WarehouseCode  *string `json:"WarehouseCode,omitempty"`
	WarehouseName  *string `json:"WarehouseName,omitempty"`
	IsDefault      *bool   `json:"IsDefault,omitempty"`
	StreetNo       *string `json:"StreetNo,omitempty"`
	AddressLine1   *string `json:"AddressLine1,omitempty"`
	AddressLine2   *string `json:"AddressLine2,omitempty"`
	City           *string `json:"City,omitempty"`
	Region         *string `json:"Region,omitempty"`
	Country        *string `json:"Country,omitempty"`
	PostCode       *string `json:"PostCode,omitempty"`
	PhoneNumber    *string `json:"PhoneNumber,omitempty"`
	FaxNumber      *string `json:"FaxNumber,omitempty"`
	MobileNumber   *string `json:"MobileNumber,omitempty"`
	DDINumber      *string `json:"DDINumber,omitempty"`
	ContactName    *string `json:"ContactName,omitempty"`
	Obsolete       *bool   `json:"Obsolete,omitempty"`
	GUID           *string `json:"Guid,omitempty"`
	LastModifiedOn *string `json:"LastModifiedOn,omitempty"`
}
