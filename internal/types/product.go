package types

type Product struct {
	ID             int     `json:"product_id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	StockQuantity  int     `json:"stock_quantity"`
	Category       int     `json:"category"`
	ImageURL       string  `json:"image_url"`
	UpdatedAt      string  `json:"updated_at"`
	CreatedAt      string  `json:"created_at"`
	CategoryString string
	Sum            int
	ViewCount      int
	PurchaseCount  int
	LastPurchaseAt string
	LastViewAt     string
}

type CreateProductRequest struct {
	Name          string  `form:"name"`
	Description   string  `form:"description"`
	Price         float64 `form:"price"`
	StockQuantity int     `form:"stock_quantity"`
	Category      int     `form:"category"`
	ImageURL      string
}

func NewProduct(cpr *CreateProductRequest) *Product {
	return &Product{
		Name:          cpr.Name,
		Description:   cpr.Description,
		Price:         cpr.Price,
		StockQuantity: cpr.StockQuantity,
		Category:      cpr.Category,
		ImageURL:      cpr.ImageURL,
	}
}
