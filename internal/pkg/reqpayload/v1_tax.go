package reqpayload

type CreateNewTax struct {
	Name    string `json:"name" form:"name" validate:"required" example:"Big Mac"`
	TaxCode int    `json:"tax_code" form:"tax_code" validate:"required" example:"1"`
	Price   int64  `json:"price" form:"price" validate:"required,min=0" example:"1000"`
}
