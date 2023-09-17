package pandabase

type BaseError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type Client struct {
	Secret  string
	Version string
}

type GetUserResponse struct {
	ID        int      `json:"id"`
	ImageHash string   `json:"image_hash"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Country   string   `json:"country"`
	Balance   float64  `json:"balance"`
	Metadata  []string `json:"metadata"`
	Verified  bool     `json:"verified"`
	Limit     int      `json:"limit"`
	Role      string   `json:"role"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type GetCustomersResponse struct {
	Data []struct{} `json:"data"`
}

type GetCustomerResponse struct {
	ID         int     `json:"id"`
	Email      string  `json:"email"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Phone      *int    `json:"phone,omitempty"`
	Currency   *string `json:"currency,omitempty"`
	IsVerified bool    `json:"is_verified"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type CreateCustomerRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     *int   `json:"phone,omitempty"`
	Currency  string `json:"currency,omitempty"`
}

type UpdateCustomerRequest struct {
	ID   int `json:"id"`
	Data struct {
		Email     *string `json:"email,omitempty"`
		FirstName *string `json:"first_name,omitempty"`
		LastName  *string `json:"last_name,omitempty"`
		Phone     *int    `json:"phone,omitempty"`
		Currency  *string `json:"currency,omitempty"`
	} `json:"data"`
}

type GetProductResponse struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	PriceUSD    float64  `json:"price_usd"`
	Handle      string   `json:"handle"`
	Images      []string `json:"images"`
	InStock     bool     `json:"in_stock"`
	Tags        []string `json:"tags"`
	Discount    string   `json:"discount"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type CreateProductRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	PriceUSD    float64  `json:"price_usd"`
	Handle      string   `json:"handle"`
	Images      []string `json:"images"`
	InStock     bool     `json:"in_stock"`
	Tags        []string `json:"tags"`
}

type UpdateProductRequest struct {
	ID   int `json:"id"`
	Data struct {
		Title       *string   `json:"title,omitempty"`
		Description *string   `json:"description,omitempty"`
		PriceUSD    *float64  `json:"price_usd,omitempty"`
		Handle      *string   `json:"handle,omitempty"`
		Images      *[]string `json:"images,omitempty"`
		InStock     *bool     `json:"in_stock,omitempty"`
		Tags        *[]string `json:"tags,omitempty"`
	} `json:"data"`
}

type GetShopResponse struct {
	ID                   int      `json:"id"`
	ImageHash            string   `json:"image_hash"`
	Title                string   `json:"title"`
	Description          string   `json:"description"`
	MetaData             []string `json:"meta_data"`
	CSS                  *string  `json:"css,omitempty"`
	Category             []string `json:"category"`
	Plan                 string   `json:"plan"`
	Domain               bool     `json:"domain"`
	Handle               string   `json:"handle"`
	HasStorefront        bool     `json:"has_storefront"`
	ForceSSL             bool     `json:"force_ssl"`
	CheckoutAPISupported bool     `json:"checkout_api_supported"`
	Timezone             string   `json:"timezone"`
	CreatedAt            string   `json:"created_at"`
}

type CreateTransactionRequest struct {
	ProductID  int    `json:"product_id"`
	CustomerID int    `json:"customer_id"`
	Provider   string `json:"provider"`
}

type GetCreatedTransactionResponse struct {
	ID               int     `json:"id"`
	IntentID         string  `json:"intent_id"`
	Status           string  `json:"status"`
	Amount           float64 `json:"amount"`
	OrderInformation struct {
		OrderID     int    `json:"order_id"`
		CustomerID  int    `json:"customer_id"`
		ProductID   int    `json:"product_id"`
		OrderStatus string `json:"order_status"`
	} `json:"order_information"`
	Provider    string `json:"provider"`
	CheckoutURL string `json:"checkout_url"`
}

type GetTransactionResponse struct {
	ID               int     `json:"id"`
	IntentID         string  `json:"intent_id"`
	Status           string  `json:"status"`
	Amount           float64 `json:"amount"`
	OrderInformation struct {
		OrderID     int    `json:"order_id"`
		CustomerID  int    `json:"customer_id"`
		ProductID   int    `json:"product_id"`
		OrderStatus string `json:"order_status"`
	} `json:"order_information"`
	Provider string `json:"provider"`
}
