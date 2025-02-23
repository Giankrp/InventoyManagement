package models

type ProductResponse struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Barcode     string  `json:"barcode"`
}

type PaginatedResponse struct {
	Data       []ProductResponse `json:"data"`
	Page       int             `json:"page"`
	Limit      int             `json:"limit"`
	TotalItems int64           `json:"total_items"`
	TotalPages int             `json:"total_pages"`
}
